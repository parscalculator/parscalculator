package main

import (
    "fmt"
    "io/ioutil"
    "net/http"
)

func main() {
    req, err := http.NewRequest("GET", "https://www.google.com", nil)
    if err != nil {
        panic(err)
    }

    client := new(http.Client)
    response, err := client.Do(req)
    if err != nil {
        panic(err)
    }
    fmt.Println(ioutil.ReadAll(response.Body))
}