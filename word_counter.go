package main

import (
    "fmt"
    "strings"
    "os"
    "io/ioutil"
)

var counter map[string] int

func WordCount(s string) map[string]int {
    a := strings.Fields(s)
    counter = make(map[string] int)

    for i := 0; i< len(a); i++ {
        if strings.Contains(a[i], "."){
            word_stripped := strings.Replace(a[i], ".", "", -1)
            counter[strings.ToLower(word_stripped)] += 1
        } else if strings.Contains(a[i], ","){
            word_stripped := strings.Replace(a[i], ",", "", -1)
            counter[strings.ToLower(word_stripped)] += 1
        } else {
            counter[strings.ToLower(a[i])] += 1
        }
    }

    return counter
}

func main() {
    file, e := ioutil.ReadFile("./words.txt")
    if e != nil {
        fmt.Printf("File error %v\n", e)
        os.Exit(1)
    }

    fmt.Println("Word Count")
    fmt.Println(WordCount(string(file)))
}