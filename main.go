package main

import (
	"EMQ/domain/services"
	"encoding/json"
)

func main() {
	topic := "test"
	action := "pub"
	num := 1
	message := map[string]interface{}{
		"testMessage": "I want you",
	}
	payload, _ := json.Marshal(message)
	qos := 0

	mQttServices := services.MQttServices{
		Topic: topic,
		Qos:   qos,
	}
	if action == "pub" {
		err := mQttServices.Publish(string(payload))
		if err != nil {
			panic(err)
		}
	} else {
		err := mQttServices.Subscribe(num)
		if err != nil {
			panic(err)
		}
	}

}
