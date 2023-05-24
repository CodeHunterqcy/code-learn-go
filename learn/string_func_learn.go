package learn

import (
	"fmt"
	"strconv"
	"strings"
)

func TestString() {
	// 1 字符串遍历处理中文，否则会乱码
	str1 := "123123中文测试3213"
	str2 := []rune(str1)
	for _, s := range str2 {
		fmt.Printf("%c\n", s)
	}

	// 2 字符串转整数
	n, err := strconv.Atoi("1234")
	if err != nil {
		fmt.Printf("err")
	} else {
		fmt.Println("结果", n)
	}

	// 3 整数转字符串
	str := strconv.Itoa(123456)
	fmt.Println(str)

	// 4 字符串转[]byte
	byte1 := []byte("hello")
	fmt.Printf("%v", byte1)

	// 5 []byte转字符串
	str = string(byte1)

	// 6 进制转换 base参数对应2，8，16进制
	strconv.FormatInt(12354, 2)

	// 7 查找子串是否在指定的字符串中,返回bool类型
	strings.Contains("hello", "llo")

	// 8 查找一个字符串有几个指定子串
	strings.Count("hellohello", "he")

	// 9 不区分大小的字符串比较（==是区分大小的）
	strings.EqualFold("abc", "Abc") // 返回true

	// 10 返回子串在字符串中第一次出现的位置
	index := strings.Index("hello", "lo")
	fmt.Println(index)

	// 11 返回子串在字符串中最后一次出现的位置
	lastIndex := strings.LastIndex("hellohello", "llo")
	fmt.Println(lastIndex)

	// 12 将指定子串替换为另一个子串,-1代表全部替换，如果想替换一个，该参数改为1即可
	testStr := "test test str replace"
	str3 := strings.Replace(testStr, "tset", "haha", -1)
	fmt.Println(str3)

	// 13 按照指定字符切割字符串
	strArr := strings.Split(testStr, " ")
	fmt.Println(strArr)

	// 14 字符串字母大小写转换
	strings.ToLower(testStr)
	strings.ToUpper(testStr)

	// 15 去掉字符串左右两边字符
	strings.TrimSpace("hello go go")
	strings.Trim("! hello !!!", "!")
	strings.TrimLeft("! hello !!!", "!")
	strings.TrimRight("! hello !!!", "!")

	// 16 判断字符串是否以指定字符串开头，返回bool类型
	res := strings.HasPrefix("test this func", "test")
	fmt.Println(res)

	// 17 判断字符串是否以指定字符串结束，返回bool类型
	res = strings.HasSuffix("test this func", "func")
	fmt.Println(res)

}
