package main

import (
	"fmt"
	"net/http"
)

func main() {
	header := http.Header{}

	//using the header as slice

	header.Set("Auth-X", "abcdef1234")
	header.Add("Auth-X", "defhfhf")
	fmt.Println(header)

	//retrieving slice of values in headers

	resSlice := header["Auth-X"]
	fmt.Println(resSlice)

	// get the first value
	resFirst := header.Get("Auth-X")
	fmt.Println(resFirst)

	// replace all existing values with
	// this one
	header.Set("Auth-X", "newvalue")
	fmt.Println(header)

	// Remove header
	header.Del("Auth-X")
	fmt.Println(header)

}
