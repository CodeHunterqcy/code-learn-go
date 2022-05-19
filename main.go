package main

import "fmt"

func main() {
	for k, v := range "hello world" {
		fmt.Printf("k = [%+v],v = [%+v]", k, v)
	}

}
