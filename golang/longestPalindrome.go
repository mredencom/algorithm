package main

import (
	"fmt"
	"strings"
)

//@todo 失败 (建议使用牛拉车算法)
//计算最长回文字符串
func longestPalindrome(s string) {
	r, a := []rune(s), []rune(s) //cbbd  [99 98 98 100]
	var res []int32              //存结果的
	fmt.Println(r, a)
	//字符串反转 cbbd 反转后 dbbc
	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	fmt.Println(string(r), string(a))
	//然后去同等下标的数据是否一样,一样就存在res,就是求交集
	for k, v := range r {
		if v == a[k] {
			res = append(res, v)
		}
	}
	fmt.Println(string(res))
	//return string(res)

}

//实现反转字符串 测试
func Reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}
//还没有实现
func main() {
	s := "google"
	//t := "$#"
	//for _, v := range s {
	//	_ = fmt.Sprintf("%s#%s", t, string(v))
	//}
	//fmt.Println(t)
	//a := Reverse(s)
	//fmt.Println(a)
	longestPalindrome(s)
	//longestPalindrome02(s)
}


//使用马拉车算法解决回文字符串
func longestPalindrome02(s string) {
	//首先判断这个字符串的长度是否大于2
	lens := len(s)
	//长度不够直接返回
	if lens < 2 {
		//return s
		fmt.Println("字符串长度少于2")
	}
	//newStr := []rune(s)
	var newStrArr []string
	//newStrArr := make([]string, 20)
	for _, v := range s {
		newStrArr = append(newStrArr, string(v))
	}
	//拼接构造字符
	resStr := fmt.Sprint("@#", strings.Join(newStrArr, "#"), "#")
	fmt.Println(resStr)
	var max int //最右回文位置
	var id int  //最右回文位置中心点
	var p []int //回文长度数组
	//var maxStr string
	var maxLen int
	for i := 1; i < lens; i++ {
		if maxLen > lens-i {
			continue
		}
		if max > i {
			z := 2*id - i
			p[i] = Minimum(max-i, p[z]).(int)
		} else {
			p[i] = 1
		}
	}
	ss := strings.Replace(resStr, "#", "", len(resStr))
	fmt.Println(ss)

}

//生成带#号的字符串
func ManacherString(s string) string {
	//newStr := []rune(s)
	var newStrArr []string
	//newStrArr := make([]string, 20)
	for _, v := range s {
		newStrArr = append(newStrArr, string(v))
	}
	//拼接构造字符
	return fmt.Sprint("#", strings.Join(newStrArr, "#"), "#")
}

//取一组数最小的数字
func Minimum(first interface{}, rest ...interface{}) interface{} {
	minimum := first
	for _, v := range rest {
		switch v.(type) {
		case int:
			if v := v.(int); v < minimum.(int) {
				minimum = v
			}
		case float64:
			if v := v.(float64); v < minimum.(float64) {
				minimum = v
			}
		case string:
			if v := v.(string); v < minimum.(string) {
				minimum = v
			}
		}
	}
	return minimum
}
