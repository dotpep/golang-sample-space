package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	resp, _ := http.Get("https://dotpep.xyz/")

	body, _ := io.ReadAll(resp.Body)
	fmt.Println(string(body))

	defer resp.Body.Close()
}
