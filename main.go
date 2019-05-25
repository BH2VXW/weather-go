package main

import "fmt"
import "./weatherserver"

func main() {
	fmt.Println("请输入要查询的城市或地区的名称，之后按回车键提交")
	var cityName string
LABEL1:
	fmt.Scanln(&cityName)
	result := weatherserver.GetWeather(cityName)
	fmt.Println(result)
	goto LABEL1
}
