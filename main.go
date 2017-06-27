package main

import (
	"fmt"
	"io"
	"net/http"
	"os/exec"
)

const LENGTH_LIMIT = 140
const COMMAND = "/run.sh"

// http://localhost:8080/sing/?words=horsemeat
func handleSing(w http.ResponseWriter, r *http.Request) {
	body := r.FormValue("words")
	if len(body) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if len(body) > LENGTH_LIMIT {
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, fmt.Sprintf("characters limited to %v", LENGTH_LIMIT))
		return
	}

	reader, err := audio(body)
	defer reader.Close()

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		io.WriteString(w, fmt.Sprintf("err %v", err))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("content-type", "audio/mpeg")
	if _, err := io.Copy(w, reader); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}

func audio(phrase string) (io.ReadCloser, error) {
	cmd := exec.Command(COMMAND)
	stdin, err := cmd.StdinPipe()
	if err != nil {
		return nil, err
	}
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return nil, err
	}

	cmd.Start()
	if _, err := io.WriteString(stdin, phrase); err != nil {
		return nil, err
	}
	return stdout, stdin.Close()
}

func main() {
	http.HandleFunc("/sing/", handleSing)
	http.ListenAndServe(":8080", nil)
}
