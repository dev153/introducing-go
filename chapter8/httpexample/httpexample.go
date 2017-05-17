package main

import (
	"io"
	"net/http"
)

func helloHandler(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "text/html")
	io.WriteString(res, `
	<DOCTYPE html>
	<html>
	<head>
	<title>Hello, World Page</title>
	</head>
	<body>
	Hello, World!
	</body>
	</html>
	`)
}

func main() {
	http.HandleFunc("/hello", helloHandler)
	http.ListenAndServe("127.0.0.1:9999", nil)
}
