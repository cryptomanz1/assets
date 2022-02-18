package main

import (
    "encoding/json"
    "fmt"
    "log"
    "os"
)

type Symbol struct {
    Symbol   string
    
}

func main() {
    // open the file pointer
    symbolFile, err := os.Open("info.json")
    if err != nil {
        log.Fatal(err)
    }
    defer symbolFile.Close()

    // create a new decoder
    var studentDecoder *json.Decoder = json.NewDecoder(symbolFile)
    if err != nil {
        log.Fatal(err)
    }

    // initialize the storage for the decoded data
    var studentList []Symbol
    
    // decode the data
    err = studentDecoder.Decode(&studentList)
    if err != nil {
        log.Fatal(err)
    }

    for i, student := range studentList {
        fmt.Println("Symbol", i+1)
        fmt.Println("Symbol:", student.Symbol)
    
    }
}