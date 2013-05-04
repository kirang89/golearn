package main

import (
    "fmt"
    "os"
    "json"
    "io/ioutil"
)

func main(){
    file, e := ioutil.ReadFile("./config.json")
    if e != nil {
        fmt.Printf("File error %v\n", e)
        os.Exit(1)
    }
    fmt.Printf("%s\n", string(file))

    var json_obj jsonobject
    json.Unmarshal(file, &json_obj)
    fmt.Printf("JSON STRUCT: %s\n", json_obj)
}
