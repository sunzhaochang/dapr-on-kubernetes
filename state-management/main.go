package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	dapr "github.com/dapr/go-sdk/client"
)

const STATE_STORE_NAME = "statestore"
const ORDER_KEY = "orderId"

var client dapr.Client

func init() {

}

func main() {
	rand.Seed(time.Now().UnixMicro())

	var err error
	client, err = dapr.NewClient()
	if err != nil {
		panic(err)
	}
	defer client.Close()

	_ = client.SaveState(context.Background(), STATE_STORE_NAME, "orderType", []byte("test"), nil)

	for i := 0; ; i++ {
		switch i % 2 {
		case 0:
			invokeWithSDK()
		case 1:
			invokeWithAPI()
		}
		time.Sleep(2 * time.Second)
	}
}

func invokeWithSDK() {
	log.Printf("---invoke with sdk---")
	orderId := rand.Intn(1000-1) + 1
	if err := client.SaveState(context.Background(), STATE_STORE_NAME, ORDER_KEY, []byte(strconv.Itoa(orderId)), nil); err != nil {
		panic(err)
	}
	log.Printf("save state success, orderId: %d", orderId)

	result, err := client.GetState(context.Background(), STATE_STORE_NAME, ORDER_KEY, nil)
	if err != nil {
		panic(err)
	}
	log.Printf("get state key: %v, value: %v, etag: %v, metadata: %v", result.Key, string(result.Value), result.Etag, result.Metadata)

	if err = client.DeleteState(context.Background(), STATE_STORE_NAME, ORDER_KEY, nil); err != nil {
		panic(err)
	}
	log.Printf("delete state success")
}

func invokeWithAPI() {
	log.Printf("---invoke with api---")
	url := fmt.Sprintf("http://localhost:3500/v1.0/state/%s", STATE_STORE_NAME)

	orderId := rand.Intn(1000-1) + 1
	b, _ := json.Marshal([]map[string]string{
		{"key": ORDER_KEY, "value": strconv.Itoa(orderId)},
	})
	body := bytes.NewBuffer(b)

	resp, err := http.Post(url, "application/json", body)
	if err != nil {
		panic(err)
	}
	res, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	log.Printf("save state success, orderId: %v", orderId)

	resp, err = http.Get(fmt.Sprintf("%s/%s", url, ORDER_KEY))
	if err != nil {
		panic(err)
	}
	res, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	log.Printf("get state success, resp: %v", string(res))

	req, _ := http.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s", url, ORDER_KEY), nil)
	cli := http.Client{}
	resp, err = cli.Do(req)
	if err != nil {
		panic(err)
	}
	res, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	log.Printf("delete state success")
}
