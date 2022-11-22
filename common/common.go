package common

import (
	"strings"
)

// SplitContent :入力された文字を空白区切りで分割
// 第一引数 :配列の一番最初の文字
// 第二引数 :配列の一番最初の文字を除いた全ての文字
func SplitContent(content string) (string, []string) {
	result := strings.Fields(content)
	return result[0], result[1:]
}
