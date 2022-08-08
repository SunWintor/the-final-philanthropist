package common

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"github.com/SunWintor/tfp/configs"
)

func EncryptionPassword(password string) string {
	return MD5(password, ":", configs.GetConf().LoginSalt)
}

func MD5(str ...string) string {
	var bt bytes.Buffer
	for _, s := range str {
		bt.WriteString(s)
	}
	h := md5.New()
	h.Write([]byte(bt.String()))
	return hex.EncodeToString(h.Sum(nil))
}
