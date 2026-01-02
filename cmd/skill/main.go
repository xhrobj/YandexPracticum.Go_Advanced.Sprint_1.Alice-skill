package main

import "net/http"

func main() {
	if err := run(); err != nil {
		panic(err)
	}
}

func run() error {
	return http.ListenAndServe(":8080", http.HandlerFunc(webhook))
}

func webhook(rs http.ResponseWriter, rq *http.Request) {
	if rq.Method != http.MethodPost {
		rs.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	rs.Header().Set("Content-Type", "application/json")
	_, _ = rs.Write([]byte(`
		{
			"response": {
				"text": "Извините, я пока ничего не умею"
			},
			"version": "0.1"
		}`))
}
