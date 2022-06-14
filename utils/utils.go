package utils

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
)

const (
	contentType = "application/json"
)

// Config holds config file data
type Config struct {
	Url        string `json:"url"`
	PrevDays   int    `json:"prevdays"`
	SenderID   string `json:"sendermail"`
	RecieverID string `json:"recievermail"`
}

func GetConfigValues(filepath string) (config Config, err error) {
	var file *os.File
	file, err = os.Open(filepath)
	if err != nil {
		log.Println(err.Error())
		return
	}
	err = json.NewDecoder(file).Decode(&config)
	if err != nil {
		log.Println(err.Error())
		return
	}
	return
}

func MakeRequest(method, url string) (response *http.Response, err error) {
	client := &http.Client{}
	request, err := http.NewRequest(method, url, nil)
	if err != nil {
		log.Println(err.Error())
		return
	}
	request.Header.Add("Accept", contentType)
	request.Header.Add("Content-Type", contentType)
	return client.Do(request)
}
