// 加密工具类

package util

import (
	"crypto/md5"
	"encoding/hex"
)

// 加密
func EncryptionMd5(s string) string {
	ctx := md5.New()
	ctx.Write([]byte(s))
	return hex.EncodeToString(ctx.Sum(nil))
}
