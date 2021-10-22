package util

import (
	"crypto/md5"
	"fmt"
)

func HashPasswd(passwd string) string  {
	data := []byte(passwd)
	has := md5.Sum(data)
	md5Str := fmt.Sprintf("%x", has)
	return md5Str
}

