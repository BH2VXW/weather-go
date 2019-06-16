package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"./telegram"
	"./weatherserver"
)

func main() {

	http.HandleFunc("/telegram/bot/", TelegramBot) //电报机器人回调
	http.HandleFunc("/weatherQuery/", WeatherQuery)
	http.ListenAndServe(":8888", nil)
	//	fmt.Println("请输入要查询的城市或地区的名称，之后按回车键提交")
	//	var cityName string
	//LABEL1:
	//	fmt.Scanln(&cityName)
	//	WeatherQuery(cityName)
	//	goto LABEL1
}

func TelegramBot(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)
	//body_str := string(body)
	//fmt.Println(body_str)
	telegram.Chat(body)
	message := r.FormValue("message")
	//检查是不是查天气
	if strings.IndexAny(message, "天气") >= 0 {
		//fmt.Fprintln(w, "查天气")
		result := weatherserver.GetWeatherByMessage(message)
		fmt.Fprintln(w, result)
	} else {
		fmt.Fprintln(w, "你说的“"+message+"”我听不懂")
	}
	//fmt.Fprintln(w, message)
}

func WeatherQuery(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintln(w, r.FormValue("city"))
	cityName := r.FormValue("city")
	result := weatherserver.GetWeather(cityName)
	fmt.Fprintln(w, result)
}
