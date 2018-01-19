package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	curr "github.com/vladimirvivien/learning-go/ch11/curr1"
)

var currencies = curr.Load("./data.csv")

func gui(resp http.ResponseWriter, req *http.Request) {
	file, err := os.Open("./currency.html")
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		fmt.Println(err)
		return
	}
	io.Copy(resp, file)
}

func currs(resp http.ResponseWriter, req *http.Request) {
	var currRequest curr.CurrencyRequest
	dec := json.NewDecoder(req.Body)
	if err := dec.Decode(&currRequest); err != nil {
		resp.WriteHeader(http.StatusBadRequest)
		fmt.Println(err)
		return
	}
	fmt.Println("Request:", currRequest)

	result := curr.Find(currencies, currRequest.Get)
	enc := json.NewEncoder(resp)
	if err := enc.Encode(&result); err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		fmt.Println(err)
		return
	}
	fmt.Println("Response:", result)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", gui)
	mux.HandleFunc("/currency", currs)

	if err := http.ListenAndServe(":4040", mux); err != nil {
		fmt.Println(err)
	}
}
