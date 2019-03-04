package main

import (
	"encoding/json"
	"fmt"

	"github.com/mitchellh/mapstructure"
)

type Message struct {
	Name string      `json: "name"`
	Data interface{} `json: "data"`
}

type Channel struct {
	Id   string `json: "id"`
	Name string `json: "name"`
}

func main() {
	recRawMsg := []byte(`{"name": "channel add", ` + `"data":{"name": "Hardware support"}}`)
	var recMessage Message
	err := json.Unmarshal(recRawMsg, &recMessage)
	if err != nil {
		fmt.Println(err)
		return
	}
	if recMessage.Name == "channel add" {
		channel, err := addChannel(recMessage.Data)
		var sendMessage Message
		sendMessage.Name = "channel add"
		sendMessage.Data = channel
		sendRawMsg, err := json.Marshal(sendMessage)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Print(string(sendRawMsg))
	}

}

func addChannel(data interface{}) (Channel, error) {
	var channel Channel
	err := mapstructure.Decode(data, &channel)
	if err != nil {
		return channel, err
	}
	channel.Id = "1"
	return channel, nil
}
