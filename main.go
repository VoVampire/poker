package main

import (
	"fmt"
	"poker/service"
)

func main() {
	// todo 从JSON文件中获取牌组到数组中
	// todo 循环数组比较各组牌大小

	fmt.Println(service.Compare("Qs7d7hQcTs","3h6c5dAsQd"))
}