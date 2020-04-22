package main

import (
	"log"
	"net/http"
)

func main() {
	myfunction := func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`
		<html>
			<head>
				<title>チャット</title>
			</head>
			<body>
				<h1>チャットしましょう</h1>
			</body>
		</html>
		`))
	}
	http.HandleFunc("/", myfunction)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
