package main

import "fmt"

func main() {
	var challenge string

	fmt.Scanln(&challenge)
	fmt.Printf("Hello %v \n", challenge)
	fmt.Println(&challenge)
}
