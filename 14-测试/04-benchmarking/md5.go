package benchmarking

import (
	"crypto/md5"
	"fmt"
)

func GetStringMd5(input string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(input)))
}
