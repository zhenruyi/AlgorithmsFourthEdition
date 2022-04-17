package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("d:/Documents/test.txt")
	if err != nil {
		fmt.Println("open error")
		return
	}
	defer file.Close()
	bytes := make([]byte, 3)
	n, _ := file.Read(bytes)
	fmt.Println(string(bytes), " ", n)
}
