package go_grab_ip

//package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

//ip-api.com query
type IPData struct {
	Status      string  `json:"status"`
	Country     string  `json:"country"`
	CountryCode string  `json:"countryCode"`
	Region      string  `json:"region"`
	RegionName  string  `json:"region_name"`
	City        string  `json:"city"`
	Zip         string  `json:"zip"` //in some places it's a string
	Lat         float64 `json:"lat"`
	Lon         float64 `json:"lon"`
	Timezone    string  `json:"timezone"`
	Isp         string  `json:"isp"`
	Org         string  `json:"org"`
	Org2        string  `json:"as"`
	Ip          string  `json:"query"`
}

func GetIPData() (IPData, error) {
	resp, err := http.Get("http://ip-api.com/json/")
	if err != nil {
		return IPData{}, err //probably no internet
	}
	if resp.StatusCode != 200 {
		return IPData{}, fmt.Errorf("Invalid response code %v", resp.StatusCode)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	var data IPData
	if err = json.Unmarshal(body, &data); err != nil {
		return IPData{}, err
	}

	return data, nil
}

//
//func main() {
//	data, err := GetIPData()
//	if err != nil {
//		return
//	}
//	fmt.Println(data)
//}
