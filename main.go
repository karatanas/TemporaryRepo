package main

import (
  "fmt"
  "io/ioutil"
  "log"
  "net/http"
  "time"
)

var buffer string

//Initial run of the program
func init() {
  call, err := makeAPICall() //Call to API
  handleError(err)
  fmt.Printf("The price for a single dollar is: %s per BTC", call)
  insertNewLine()
  buffer = call
}

func main() {
  time.Sleep(time.Minute) //Wait one minute
  call, err := makeAPICall() //Call the API
  handleError(err)
  if buffer != call {
    insertNewLine()
    fmt.Printf("The price for a single dollar is: %s per BTC", call)
    insertNewLine()
  }
  buffer = call //Asing the buffer to the call
  main()
}

//makeAPICall function makes an API call to the function requesting the value of a single dollar per bitcoin
func makeAPICall() (string, error) {
  resp, err := http.Get("https://blockchain.info/tobtc?currency=USD&value=1") //Make a GET request to the API
  if err != nil {
    return "", err
  }
  currency, err := ioutil.ReadAll(resp.Body) //Read Response's body
  resp.Body.Close()                          //Close Response's body
  if err != nil {
    return "", err
  }
  return string(currency), nil //Convert the currency to a string and return
}

func handleError(err error) {
  if err != nil {
    log.Fatal(err)
  }
}
func insertNewLine() {
  fmt.Println("")
}