package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os/exec"
	"strings"
)

const LENGTH_LIMIT = 140
const COMMAND = "/run.sh"

// http://localhost:8080/s?w=horsemeat
func handleSing(w http.ResponseWriter, r *http.Request) {
	body := r.FormValue("w")
	if len(body) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if len(body) > LENGTH_LIMIT {
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, fmt.Sprintf("characters limited to %v", LENGTH_LIMIT))
		return
	}

	resp, err := audio("raining", body)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("[app] err: ", err)
		return
	}

	w.Header().Set("content-type", "audio/mpeg")

	log.Println("[app:raining] phrase: ", body)

	if _, err := w.Write(resp); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("[app] err: ", err)
		return
	}
}

func handleIverson(w http.ResponseWriter, r *http.Request) {
	body := r.FormValue("w")
	if len(body) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if len(body) > LENGTH_LIMIT {
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, fmt.Sprintf("characters limited to %v", LENGTH_LIMIT))
		return
	}

	resp, err := audio("iverson", body)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("[app] err: ", err)
		return
	}

	w.Header().Set("content-type", "audio/mpeg")

	log.Println("[app:iverson] phrase: ", body)

	if _, err := w.Write(resp); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("[app] err: ", err)
		return
	}
}

func audio(theme string, phrase string) ([]byte, error) {
	cmd := exec.Command(COMMAND, "--", theme)
	cmd.Stdin = strings.NewReader(phrase)
	return cmd.Output()
}

func handlePing(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "pong")
	w.WriteHeader(http.StatusOK)
}

func main() {
	http.HandleFunc("/s", handleSing)
	http.HandleFunc("🎶", handleSing)
	http.HandleFunc("/raining", handleSing)
	http.HandleFunc("/ai", handleIverson)
	http.HandleFunc("/varz/ping", handlePing)
	http.ListenAndServe(":8080", nil)
}
