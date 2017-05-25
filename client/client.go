package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

func ubusCall(requestBody []byte) ([]byte, error) {
	body := bytes.NewReader(requestBody)
	req, err := http.NewRequest("POST", "http://127.0.0.1/ubus", body)
	if err != nil {
		return nil, err
	}

	dClient := http.DefaultClient
	resp, err := dClient.Do(req)
	if err != nil {
		return nil, err
	}
	return ioutil.ReadAll(resp.Body)
}

func main() {
	dd := websocket.DefaultDialer
	c, resp, err := dd.Dial("ws://192.168.56.1:8000/", nil)
	if resp == nil {
		fmt.Println(err)
		return
	}
	for {
		_, r, err := c.NextReader()
		if err != nil {
			fmt.Println(err)
			c.Close()
			break
		}
		mess, err := ioutil.ReadAll(r)
		if err != nil {
			fmt.Println(err)
			c.Close()
			break
		}
		fmt.Println("The Message is : ", string(mess))
		time.Sleep(time.Second)
		callResp, err := ubusCall(mess)
		if err != nil {
			fmt.Println(err)
			c.Close()
			break
		}
		err = c.WriteMessage(websocket.TextMessage, callResp)
		if err != nil {
			fmt.Println(err)
			c.Close()
			break
		}
	}
}
