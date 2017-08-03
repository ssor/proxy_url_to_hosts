package prequest

import (
	"fmt"

	"github.com/parnurzeal/gorequest"
)

func DoRequest(proxy, url string) error {
	request := gorequest.New().Proxy(proxy)
	resp, body, errs := request.Get(url).End()
	if errs != nil {
		fmt.Println("==================error:======================")
		fmt.Println(errs)
		return errs[0]
	}

	fmt.Println("status: ", resp.Status)
	fmt.Println("body: ", body)
	return nil
}
