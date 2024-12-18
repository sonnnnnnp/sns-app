package utils

import (
	"regexp"
)

// ポインタを string 型に変換
func Ptr(s string) *string { return &s }

// パスが除外リストにマッチするかをチェック
func IsExcludedPath(targetPath string, excludePaths []string) bool {
	for _, excludePath := range excludePaths {
		re := regexp.MustCompile(excludePath)
		if re.MatchString(targetPath) {
			return true
		}
	}
	return false
}
