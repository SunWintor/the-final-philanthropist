package common

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/SunWintor/tfp/configs"
	"strings"
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

func DisorderString(list []int64) string {
	var temp = make([]string, len(list))
	for k, v := range list {
		temp[k] = fmt.Sprintf("%d", v)
	}
	return strings.Join(temp, ",")
}
