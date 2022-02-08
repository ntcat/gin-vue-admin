package utils

import (
	"fmt"
	"strings"
	"testing"
)

func Test_fff(t *testing.T) {
	s := "root.cj3.ln2.01.status.input"

	//返回字符串str中的任何一个字符在字符串s中最后一次出现的位置。
	//如果找不到或str为空则返回-1 LastIndex
	index := strings.LastIndexAny(s, ".")
	fmt.Println(s[index+1:])
	fmt.Println(s[:index])

}
