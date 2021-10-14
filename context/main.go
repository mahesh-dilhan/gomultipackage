package main

import (
	"context"
	"github.com/google/uuid"
	"net/http"
)

func main() {
	helloWorldHandler := http.HandlerFunc(HelloWorld)
	http.Handle("/welcome", inejctMsgID(helloWorldHandler))
	http.ListenAndServe(":8080", nil)
}

//HelloWorld hellow world handler
func HelloWorld(w http.ResponseWriter, r *http.Request) {
	msgID := ""
	if m := r.Context().Value("msgId"); m != nil {
		if value, ok := m.(string); ok {
			msgID = value
		}
	}
	w.Header().Add("msgId", msgID)
	w.Write([]byte("Hello, world"))
}
