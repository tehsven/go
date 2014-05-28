package main

import "fmt"
import "C"

func main() {
	hellogo();
}

//export hellogo
func hellogo() {
	fmt.Println("hello go")
}
