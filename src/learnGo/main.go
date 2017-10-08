package main

import (
	"fmt"
	"encoding/json"
	"math/rand"
	"time"
	"strconv"
)

type Book struct {
	Auth        string `json:"name"`
	PublishTime int `json:"time"`
	Test        []string
}

const (
	A = 4000+iota
	B
	C
)

func main() {
	b := &Book{"tom", 19901108, []string{"a", "b"}}
	d, err := json.MarshalIndent(b, "", "  ")
	if err != nil {
		return
	}
	fmt.Printf("%s\n", d)
	fmt.Println(stringAppend())
}

func stringAppend()string{
	rand.Seed(time.Now().Unix())
	i:=rand.Intn(10)
	return strconv.Itoa(i)
}
