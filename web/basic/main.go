package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "helloword")
}

func toGET() {
	resp, err := http.Get("https://www.bilibili.com/")
	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("err:", err)
	}
	fmt.Print(string(bytes))
	defer resp.Body.Close()

}

func toPost() {
	jsonBody := `{"name": "test","gender":"男"}`
	resp, err := http.Post("https://www.bilibili.com/", "application/json", strings.NewReader(jsonBody))
	if err != nil {
		fmt.Println("err:", err)
	}
	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("err:", err)
	}
	fmt.Print(string(bytes))

}
func main() {
	// server := &http.Server{Addr: ":80"}
	// http.HandleFunc("/", hello)
	// server.ListenAndServe()
	toGET()
}
