package main

import (
	"fmt"
	"net/http"
)

func logginHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("login for:", r.URL.Query().Get("fusen"))
	//给前端的响应
	w.Write([]byte("你成功啦"))
}

func main() {

	http.HandleFunc("/user/login", logginHandler)
	http.ListenAndServe(":2280", nil)
}
