package shared_utils

import (
	"crypto/md5"
	"fmt"
)

func Md5Hash(text string) string {
	hash := md5.Sum([]byte(text))
	return fmt.Sprintf("%x", hash)
}
