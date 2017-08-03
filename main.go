package main

import (
	"flag"
	"fmt"

	"github.com/parnurzeal/gorequest"
	config "github.com/ssor/go_config"
)

var (
	// proxyHost = flag.String("proxy", "http://localhost:9900", "proxy host and port")
	// url       = flag.String("url", "", "url that will be requested")
	// urlHost   = flag.String("host", "127.0.0.1", "url's host")
	confFile = flag.String("conf", "conf/config.json", "config file")
)

func main() {
	flag.Parse()
	if flag.Parsed() == false {
		flag.PrintDefaults()
		return
	}

	conf, err := config.LoadConfig(*confFile)
	if err != nil {
		fmt.Println("==================error:======================")
		fmt.Println(err)
		return
	}

	proxy, err := validateProxyHost(conf.Get("proxy").(string))
	if err != nil {
		fmt.Println(err)
		return
	}
	completedURL, err := composeURL(conf.Get("host").(string), conf.Get("url").(string))
	if err != nil {
		fmt.Println(err)
		return
	}
	request := gorequest.New().Proxy(proxy)
	resp, body, errs := request.Get(completedURL).End()
	if errs != nil {
		fmt.Println("==================error:======================")
		fmt.Println(errs)
		return
	}

	fmt.Println("status: ", resp.Status)
	fmt.Println("body: ", body)
}

func validateProxyHost(host string) (string, error) {
	return host, nil
}

func composeURL(host, url string) (string, error) {
	return fmt.Sprintf("%s%s", host, url), nil
}
