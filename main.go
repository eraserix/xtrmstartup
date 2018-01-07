package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		query, ok := r.URL.Query()["q"]
		if ok {
			response := handleQuery(query[0])
			w.Header().Set("Content-Type", "text/plain; charset=utf-8")
			w.WriteHeader(http.StatusOK)
			fmt.Fprintf(w, response)
		}
	})

	log.Fatal(http.ListenAndServe(":3002", nil))
}

func handleQuery(query string) string {
	splits := strings.SplitN(query, ":", 2)
	id := splits[0]
	question := strings.TrimSpace(splits[1])

	fmt.Printf("query [%v] : %v\n", id, question)
	answer := findTheAnswer(question)
	fmt.Printf("response [%v] : %v\n", id, answer)
	return answer
}

func answer(question string) string {
	return "go"
}
