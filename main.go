package main

import (
	"flag"
	"fmt"

	config "github.com/ssor/go_config"
	"github.com/ssor/proxy_url_to_hosts/prequest"
)

var (
	confFile  = flag.String("conf", "conf/config.json", "config file")
	hostIndex = flag.Int("index", 0, "host index, if set 0 or no set, all host will be looped")
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
	proxy, url, hosts, err := getRequestParams(conf)
	if err != nil {
		fmt.Println("==================error:======================")
		fmt.Println(err)
		return
	}
	requestURLs := generateRequestURLs(hosts, url, *hostIndex)

	for _, url := range requestURLs {
		prequest.DoRequest(proxy, url)
	}

}

func getRequestParams(conf config.IConfigInfo) (proxy, url string, hosts []string, err error) {

	proxy, err = validateProxyHost(conf.Get("proxy").(string))
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("---- proxy: ", proxy)

	hosts = conf.Get("hosts").([]string)
	fmt.Println("---- hosts: ")
	for _, host := range hosts {
		fmt.Println("     * ", host)
	}

	url = conf.Get("url").(string)
	fmt.Println("---- url: ", url)
	return
}

func validateProxyHost(host string) (string, error) {
	return host, nil
}

func composeURL(host, url string) string {
	return fmt.Sprintf("%s%s", host, url)
}

func generateRequestURLs(hosts []string, url string, index int) []string {
	completedURLs := []string{}
	if index <= 0 {
		for _, host := range hosts {
			completedURLs = append(completedURLs, composeURL(host, url))
		}
	} else {
		completedURLs = append(completedURLs, composeURL(hosts[index-1], url))
	}
	return completedURLs
}
