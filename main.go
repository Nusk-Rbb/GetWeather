package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

type Response struct {
	Location []Location `json:"location"`
	Current  []Current  `json:"current"`
}

type Location struct {
	Name            string `json:"name"`
	Region          string `json:"region"`
	Country         string `json:"country"`
	Lat             int    `json:"lat"`
	Lon             int    `json:"lon"`
	Timezone        string `json:"tz_id"`
	LocaleTimeEpoch int    `json:"localetime_epoch"`
	LocaleTime      string `json:"localetime"`
}

type Current struct {
	LastUpdatedEpoch    int         `json:"last_updated_epoch"`
	LastUpdatedTime     string      `json:"last_updated"`
	TempCelsius         int         `json:"temp_c"`
	TempFahrenheit      int         `json:"temp_f"`
	IsDay               int         `json:"is_day"`
	Condition           []Condition `json:"condition"`
	Windmph             int         `json:"wind_mph"`
	Windkph             int         `json:"wind_kph"`
	WindDegree          int         `json:"wind_degree"`
	WindDir             string      `json:"wind_dir"`
	Pressuremb          int         `json:"pressure_mb"`
	Pressurein          int         `json:"pressure_in"`
	Precipmm            int         `json:"precip_mm"`
	Precipin            int         `json:"precip_in"`
	Humidity            int         `json:"humidity"`
	CloudPercentage     string      `json:"cloud"`
	FeelsLikeCelsius    int         `json:"feelslike_c"`
	FeelsLikeFahrenheit int         `json:"feelslike_f"`
	Visibilitykm        int         `json:"vis_km"`
	Visibilitymiles     int         `json:"vis_miles"`
	UV                  int         `json:"uv"`
	Gustmph             int         `json:"gust_mph"`
	Gustkph             int         `json:"gust_kph"`
}

type Condition struct {
	Text string `json:"text"`
	Icon string `json:"icon"`
	Code int    `json:"code"`
}

func main() {
	api_key, err := LoadEnv()
	if err != nil {
		fmt.Println(err)
	}
	locale := "niihama"
	resp, err := GetResponse(api_key, locale)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%+v\n", resp)
}

func LoadEnv() (string, error) {
	err := godotenv.Load("key.env")

	if err != nil {
		return "", err
	}

	message := os.Getenv("api_key")

	fmt.Println(message)

	return message, err
}

func GetResponse(api_key string, locale string) (*Response, error) {
	responses := Response{}
	api_url := "http://api.weatherapi.com/v1/current.json?"
	api_key = "key=" + api_key + "&"
	locale = "q=" + locale + "&"
	url := api_url + api_key + locale + "aqi=no"

	resp, err := http.Get(url)
	if err != nil {
		return &responses, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		fmt.Println("Status Code Error", resp.StatusCode)
		return &responses, nil
	}

	body, _ := io.ReadAll(resp.Body)

	fmt.Println(string(body))

	if err := json.Unmarshal(body, &responses); err != nil {
		return &responses, err
	}

	return &responses, nil
}
