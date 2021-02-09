package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Zip struct {
	Message string `json:"åmessage"`
	Results []struct {
		Address1 string `json:"address1"`
		Address2 string `json:"address2"`
		Kana1    string `json:"kana1"`
		Kana2    string `json:"kana2"`
		Kana3    string `json:"kana3"`
		Prefcode string `json:"prefcodde"`
		Zipcode  string `json:"zipcode"`
	}
	Status int `json:"status"`
}

func main() {
	zip, err := fetch()
	if err != nil {
		fmt.Println(err)
	}

	for _, i := range zip {
		fmt.Println(i.Results)
	}
}

func fetch() ([]Zip, error) {
	url := "https://zipcloud.ibsnet.co.jp/api/search?zipcode=0900802"
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	fmt.Println(res)

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var zip []Zip
	if err := json.Unmarshal(body, &zip); err != nil {
		return nil, err
	}

	return zip, nil
}
