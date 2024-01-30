package main

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

type Post struct {
	UserId int    `json:"userId"`
	Id     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

func main() {
	api_key := LoadEnv()
	locale := "niihama"
	GetResponse(api_key, locale)
}

func LoadEnv() string {
	err := godotenv.Load("key.env")

	if err != nil {
		fmt.Printf("Can't read env file : %v", err)
	}

	message := os.Getenv("api_key")

	return message
}

func GetResponse(api_key string, locale string) {
	api_url := "http://api.weatherapi.com/v1/current.json?"
	api_key = "key=" + api_key + "&"
	locale = "q=" + locale + "&"
	url := api_url + api_key + locale + "aqi=no"

	response, err := http.Get(url)
	if err != nil {
		fmt.Println("Response Error:", err)
		return
	}

	defer response.Body.Close()

	if response.StatusCode != 200 {
		fmt.Println("Status Code Error:", response.StatusCode)
		return
	}

	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))
}
