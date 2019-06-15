package weatherserver

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Result struct {
	Weatherinfo WeatherInfo `json:"weatherinfo"`
}

type WeatherInfo struct {
	City    string `json:"city"`
	CityID  string `json:"cityid"`
	Temp1   string `json:"temp1"`
	Temp2   string `json:"temp2"`
	Weather string `json:"weather"`
	Img1    string `json:"img1"`
	Img2    string `json:"img2"`
	Ptime   string `json:"ptime"`
}

//从内容中提取城市名称并查询天气
func GetWeatherByMessage(message string) string {
	//转换城市代码
	cityCode := cityCatch(message)
	result := ""
	if cityCode != "" {
		url := "http://www.weather.com.cn/data/cityinfo/" + cityCode + ".html"
		//fmt.Println(url)
		resp, err := http.Get(url)
		check(err)
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		check(err)
		dat := Result{}
		//fmt.Println(string(body))
		//jsonStr := `{"weatherinfo":{"city":"大连","cityid":"101070201","temp":"21.2","WD":"南风","WS":"小于3级","SD":"58%","AP":"999hPa","njd":"暂无实况","WSE":"<3","time":"17:50","sm":"2.8","isRadar":"1","Radar":"JC_RADAR_AZ9411_JB"}}`
		//jsonStr := `{"city":"大连","cityid":"101070201","temp":"21.2","WD":"南风","WS":"小于3级","SD":"58%","AP":"999hPa","njd":"暂无实况","WSE":"<3","time":"17:50","sm":"2.8","isRadar":"1","Radar":"JC_RADAR_AZ9411_JB"}`
		if err := json.Unmarshal(body, &dat); err != nil {
			fmt.Println(err)
			panic(err)
		}
		//fmt.Println(dat)
		result += "城市/地区：" + dat.Weatherinfo.City + "\r\n"
		result += "天气：" + dat.Weatherinfo.Weather + "\r\n"
		result += "气温：" + dat.Weatherinfo.Temp1 + "~" + dat.Weatherinfo.Temp2 + "\r\n"

		result += "---------------------"
	} else {
		result = "没有查询到该地区的天气情况"
	}

	return result
}

func GetWeather(cityName string) string {
	//转换城市代码
	cityCode := cityswitch(cityName)
	result := ""
	if cityCode != "" {
		url := "http://www.weather.com.cn/data/cityinfo/" + cityCode + ".html"
		//fmt.Println(url)
		resp, err := http.Get(url)
		check(err)
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		check(err)
		dat := Result{}
		//fmt.Println(string(body))
		//jsonStr := `{"weatherinfo":{"city":"大连","cityid":"101070201","temp":"21.2","WD":"南风","WS":"小于3级","SD":"58%","AP":"999hPa","njd":"暂无实况","WSE":"<3","time":"17:50","sm":"2.8","isRadar":"1","Radar":"JC_RADAR_AZ9411_JB"}}`
		//jsonStr := `{"city":"大连","cityid":"101070201","temp":"21.2","WD":"南风","WS":"小于3级","SD":"58%","AP":"999hPa","njd":"暂无实况","WSE":"<3","time":"17:50","sm":"2.8","isRadar":"1","Radar":"JC_RADAR_AZ9411_JB"}`
		if err := json.Unmarshal(body, &dat); err != nil {
			fmt.Println(err)
			panic(err)
		}
		//fmt.Println(dat)
		result += "城市/地区：" + dat.Weatherinfo.City + "\r\n"
		result += "天气：" + dat.Weatherinfo.Weather + "\r\n"
		result += "气温：" + dat.Weatherinfo.Temp1 + "~" + dat.Weatherinfo.Temp2 + "\r\n"

		result += "---------------------"
	} else {
		result = "没有查询到该地区的天气情况"
	}

	return result
}
