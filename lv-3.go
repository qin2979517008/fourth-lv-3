package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"strconv"
)

func main() {
	var name = [10162]string{}
	var j int
	for i := 2019210001; i <= 2019220162; i++ {
		a := strconv.Itoa(i)
		website := "http://jwzx.cqupt.edu.cn/kebiao/kb_stu.php?xh=" + a
		res, err := http.Get(website)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		// 读取资源数据 body: []byte
		body, err := ioutil.ReadAll(res.Body)
		// 关闭资源流
		res.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", website, err)
			os.Exit(1)
		}
		//写入文件
		ioutil.WriteFile("site.txt", body, 0644)
		file, _ := ioutil.ReadFile("site.txt")
		regall := `>>\d{10}([\s\S]*?)<`
		rega := regexp.MustCompile(regall)
		all := rega.FindAllStringSubmatch(string(file), 1)
		for _, v := range all {
			if v [1]!= "  "{
				fmt.Println(v[1])
				j = i - 2019210001
				name[j] = v[1]
			}
		}
	}
	res,max :=getMax(name)
	fmt.Println(res,max)
}
//getMax为找出数组中元素出现次数最多的元素以及次数
func getMax(name [10162]string) (string, int) {
	max := 0
	res := ""
	temp := make(map[string]int)
	for _, v := range name {
		if _, ok := temp[v]; ok&&v !=""{
			temp[v]++ // 如果map存在则增加1
		} else {
			temp[v] = 1
		}
		// 对最多次数的那个进行保存
		if max < temp[v] {
			max = temp[v]
			res = v
		}
	}
	return res, max
}
