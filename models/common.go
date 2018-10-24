package models

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"strings"
)

func init() {
	config()
}

func UnixConfigToMap(config string) map[string]map[string]string {
	var info_result = make(map[string]map[string]string)
	var inner_map = make(map[string]string)

	// fmt.Printf("%#v", info)
	info_arr := strings.Split(config, "\r\n")
	// log.Printf("%#v", info_arr)
	var title = ""
	for _, v := range info_arr {
		if v == "" {
			title = ""
			inner_map = make(map[string]string)
			continue
		}
		if title != "" {
			var x = strings.Split(v, ":")
			// println("find array")
			// log.Printf("%v", x)
			// log.Printf("%#v", v)

			tmp_map := make(map[string]string)
			tmp_map[x[0]] = x[1]
			info_result[title] = mapAdd(inner_map, tmp_map)
			// log.Printf("%#v", info_result)
		}
		if strings.Contains(v, "#") {
			// println("find title")
			title = strings.Replace(v, "# ", "", -1)
		}
		if v == "\r\n" {
			// println("find new line")
			title = ""
		}

	}
	return info_result

}

func mapAdd(current map[string]string, tmp map[string]string) map[string]string {
	for k, v := range tmp {
		current[k] = v
	}
	return current
}

func config() {
	data, err := ioutil.ReadFile("./config.json")
	if err != nil {
		log.Fatal("读取配置文件出错：", err)
		return
	}

	err = json.Unmarshal(data, &AppConfig)
	if err != nil {
		log.Fatal("解析配置文件出错：", err)
		return
	}
}
