package utils

import (
	"crypto/sha1"
	"fmt"
	"io"
	"sort"
)

// sha1签名
func Signature(params ...string) string {
	sort.Strings(params)
	s := sha1.New()
	for _, param := range params {
		_, _ = io.WriteString(s, param)
	}
	return fmt.Sprintf("%x", s.Sum(nil))
}
