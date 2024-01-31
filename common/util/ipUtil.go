// ip工具类
// xiaoRui
package util

import (
	"fmt"
	"net"
	"strings"

	"github.com/gogf/gf/encoding/gcharset"
	"github.com/gogf/gf/encoding/gjson"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/util/gconv"
)

func GetRealAddressByIP(ip string) string {
	toByteIp := ipToByte(ip)
	if isLocalIp(toByteIp) {
		return "服务器登录"
	}
	if isLANIp(toByteIp) {
		return "局域网"
	}
	return getLocation(ip)
}

func ipToByte(ipstr string) []byte {
	ips := strings.Split(ipstr, ".")
	ip := make([]byte, 0, len(ips))
	for _, s := range ips {
		u := gconv.Uint8(s)
		ip = append(ip, u)
	}
	return ip

}
func isLocalIp(IP net.IP) bool {
	if IP.IsLoopback() || IP.IsLinkLocalMulticast() || IP.IsLinkLocalUnicast() {
		return true
	}
	return false
}

func isLANIp(IP net.IP) bool {
	to4 := IP.To4()
	fmt.Println(to4)
	if ip4 := IP.To4(); ip4 != nil {
		switch true {
		case ip4[0] == 10:
			return true
		case ip4[0] == 172 && ip4[1] >= 16 && ip4[1] <= 31:
			return true
		case ip4[0] == 192 && ip4[1] == 168:
			return true
		default:
			return false
		}
	}
	return false
}
func getLocation(ip string) string {
	url := "https://whois.pconline.com.cn/ipJson.jsp?json=true&ip=" + ip
	bytes := ghttp.GetBytes(url)
	src := string(bytes)
	srcCharset := "GBK"
	tmp, _ := gcharset.ToUTF8(srcCharset, src)
	json, err := gjson.DecodeToJson(tmp)
	if err != nil {
		fmt.Println()
	}
	if json.GetInt("code") == 0 {
		addr := json.GetString("addr")
		return addr
	} else {
		return "未知地址"
	}
}
func GetLocalIP() (ip string, err error) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return
	}
	for _, addr := range addrs {
		ipAddr, ok := addr.(*net.IPNet)
		if !ok {
			continue
		}
		if ipAddr.IP.IsLoopback() {
			continue
		}
		if !ipAddr.IP.IsGlobalUnicast() {
			continue
		}
		return ipAddr.IP.String(), nil
	}
	return
}
