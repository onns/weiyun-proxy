package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
)

type GlobalConfig struct {
	Url    string `json:"url"`
	Port   string `json:"port"`
	Cookie string `json:"cookie"`
}

var OnnsGlobal GlobalConfig

func loadConfig() {
	filename := "config.json"
	if _, err := os.Stat(filename); err != nil {
		panic(err)
	}
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	json.Unmarshal(b, &OnnsGlobal)
}

// https://blog.csdn.net/Phoenix_smf/article/details/89278398
func init() {
	loadConfig()
}

func main() {
	remote, err := url.Parse(OnnsGlobal.Url)
	if err != nil {
		panic(err)
	}

	proxy := GoReverseProxy(&RProxy{
		remote: remote,
		cookie: OnnsGlobal.Cookie,
	})

	log.Println("当前代理地址： " + OnnsGlobal.Url + " 本地监听： http://127.0.0.1:" + OnnsGlobal.Port)

	serveErr := http.ListenAndServe(OnnsGlobal.Port, proxy)

	if serveErr != nil {
		panic(serveErr)
	}
}
