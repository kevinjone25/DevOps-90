package main

import (
	"encoding/hex"
	"fmt"
	"time"
)

func main() {
	s3cr3t := "4d475443437b5930755f6730745f6d337d"
	duration := time.Duration(100) * time.Second
	bs, err := hex.DecodeString(s3cr3t)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(bs))
	time.Sleep(duration)
}
