package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	dapr "github.com/dapr/go-sdk/client"
)

const (
	address = "http://localhost:3500"
	appID   = "http-server"
)

func main() {

	for i := 0; ; i++ {
		switch i % 3 {
		case 0:
			invokeA()
		case 1:
			invokeB()
		case 2:
			invokeC()
		}

		time.Sleep(1 * time.Second)
	}

}

// 不需要修改原有的URL, 只需要添加 dapr-app-id header即可
func invokeA() {
	client := &http.Client{}
	url := fmt.Sprintf("%s/echo?value=hello", address)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatalln(err)
	}
	req.Header.Add("dapr-app-id", appID)

	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("invokeA resp body: ", string(b))
}

//需要组装URL
func invokeB() {
	client := &http.Client{}
	url := fmt.Sprintf("%s/v1.0/invoke/%s/method/echo?value=hello", address, appID)
	req, err := http.NewRequest(http.MethodPost, url, nil)
	if err != nil {
		log.Fatalln(err)
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("invokeB resp body: ", string(b))
}

// 使用SDK, grpc -> http
func invokeC() {
	// 全局初始化一次, 程序退出时再调用close
	cli, err := dapr.NewClient()
	if err != nil {
		log.Fatalln(err)
	}
	//defer cli.Close()

	res, err := cli.InvokeMethod(context.TODO(), appID, "echo?value=hello", "get")
	log.Println("invokeC resp body: ", string(res))
}
