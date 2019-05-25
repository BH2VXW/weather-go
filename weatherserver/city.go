package weatherserver

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

//检查异常
func check(e error) {
	if e != nil {
		panic(e)
	}
}

//转换城市
func cityswitch(city string) string {
	file, err := getCurrentPath()
	check(err)
	jsonStr, err := ioutil.ReadFile(file + "data\\city.json")
	check(err)
	//解析JSON
	byt := []byte(jsonStr)
	var dat []map[string]string
	if err := json.Unmarshal(byt, &dat); err != nil {
		panic(err)
	}
	code := ""
	for _, item := range dat {
		for k, v := range item {
			if v == city {
				code = k
				break
			}
		}
	}
	return code
}

func getCurrentPath() (string, error) {
	file, err := exec.LookPath(os.Args[0])
	if err != nil {
		return "", err
	}
	path, err := filepath.Abs(file)
	if err != nil {
		return "", err
	}
	i := strings.LastIndex(path, "/")
	if i < 0 {
		i = strings.LastIndex(path, "\\")
	}
	if i < 0 {
		return "", errors.New(`error: Can't find "/" or "\".`)
	}
	return string(path[0 : i+1]), nil
}
