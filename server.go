package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/go-martini/martini"
	"gopkg.in/yaml.v3"
)

// var tcpGroup map[string]string
var resultMap map[string]interface{}

func load() {
	// conf := new(module.Settings)
	yamlFile, err := ioutil.ReadFile("settings.yaml")
	log.Println("yamlFile:", string(yamlFile))
	if err != nil {
		log.Printf("yamlFile.Get err #%v ", err)
	}
	// err = yaml.Unmarshal(yamlFile, conf)
	err = yaml.Unmarshal(yamlFile, &resultMap)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}
	log.Println("conf", resultMap)
}

func main() {
	load()
	m := martini.Classic()
	m.Get("/healthz", healthz)
	m.Get("/port", tcpConnect)
	os.Setenv("PORT", "22615")
	m.RunOnAddr("0.0.0.0:22615")
}

func healthz(group string) string {
	return fmt.Sprintf("Hello %s", group)
}

func tcpConnect(res http.ResponseWriter, req *http.Request) {
	args := req.URL.Query()
	group := args.Get("group")
	fmt.Println("key", group)
	groupUrl := resultMap[group]
	fmt.Println("value", groupUrl)
	if groupUrl == "" {
		res.Write([]byte("group参数不正确"))
		return
	}

	url := groupUrl.(string)
	// dest, err := http.Get(url)
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println(err)
	}
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		res.Write([]byte("请求失败"))
		return
	}
	status := response.StatusCode
	var result []byte
	if status == http.StatusOK {
		result = []byte(fmt.Sprintf("%s前置服务状态正常", group))
	} else {
		result = []byte(fmt.Sprintf("%s前置服务状态异常", group))
	}

	// cmd := exec.Command("/bin/sh", "-c", `ls -l`)
	// stdout, _ := cmd.StdoutPipe()
	// outBytes, _ := ioutil.ReadAll(stdout)
	// out := string(outBytes)
	// var result []byte
	// if out == "200" {
	// 	result = []byte("请求成功")
	// } else {
	// 	result = []byte("请求失败")
	// }
	res.Write(result)
}
