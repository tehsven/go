package main

import (
	"fmt"
	"bufio"
	"os"
	"log"
)

func main() {
	fmt.Println(os.Args)
	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
		return
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
		return
	}
}
