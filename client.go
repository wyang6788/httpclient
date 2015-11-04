package main

import (
	"bytes"
	"net/http"
	"fmt"
	"flag"
	"os"
	"io/ioutil"
)

func main(){
	urlPtr := flag.String("url", "", "the url")

	methodPtr := flag.String("method", "", "the method")

	dataPtr := flag.String("data", "", "the data")
	
	yearPtr := flag.String("year", "", "the year")

	flag.Parse()

	switch{
		case *methodPtr == "list":
			get(*urlPtr)
		case *methodPtr == "create":
			post(*urlPtr, *dataPtr)
		case *methodPtr == "update":
			update(*urlPtr)
		case *methodPtr == "remove":
			delete(*urlPtr, *yearPtr)
	}
}

func get(url string){
	req, err := http.NewRequest("GET", url, nil)
	req.Header.Set("Content-Type", "application/json")

  client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        os.Exit(1)
    }
    defer resp.Body.Close()

    body, _ := ioutil.ReadAll(resp.Body)
    fmt.Println(string(body))
}

func post(url string, data string){
  var jsonStr = []byte(data)
  req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
  req.Header.Set("Content-Type", "application/json")
 
  client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        os.Exit(1)
    }
    defer resp.Body.Close()

    body, _ := ioutil.ReadAll(resp.Body)
    fmt.Println(string(body))
}

func update(url string){
  req, err := http.NewRequest("PUT", url, nil)
  req.Header.Set("Content-Type", "application/json")
 
  client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        os.Exit(1)
    }
    defer resp.Body.Close()

    body, _ := ioutil.ReadAll(resp.Body)
    fmt.Println(string(body))
}

func delete(url string, year string){
  req, err := http.NewRequest("DELETE", url+"/"+year, nil)
  req.Header.Set("Content-Type", "application/json")

  client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        os.Exit(1)
    }
    defer resp.Body.Close()

    body, _ := ioutil.ReadAll(resp.Body)
    fmt.Println(string(body))
}
