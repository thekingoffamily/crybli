package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Response struct {
	Bpi struct {
		USD struct {
			RateFloat float64 `json:"rate_float"`
		} `json:"USD"`
	} `json:"bpi"`
}

func getBtcPriceHandler(w http.ResponseWriter, r *http.Request) {
	url := "https://api.coindesk.com/v1/bpi/currentprice.json"

	resp, err := http.Get(url)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var data Response
	err = json.Unmarshal(body, &data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	price := data.Bpi.USD.RateFloat

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.Write([]byte(fmt.Sprintf("%.2f", price)))
}

func main() {
	http.HandleFunc("/getBtcPrice", getBtcPriceHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}