package src

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

type gpt struct {
	API_URL string
}

type Response struct {
	Response string `json:"response"`
}

func NewGPT() *gpt {
	return &gpt{
		API_URL: os.Getenv("API_URL"),
	}
}

func (g *gpt) Requset(prompt string) string {
	var r Response
	data := fmt.Sprintf(`{"prompt":"%s"}`, prompt)
	resp, err := http.Post(g.API_URL, "application/json", strings.NewReader(data))

	if err != nil {
		log.Println(err)
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)

	if err != nil {
		log.Println(err)
	}

	json.Unmarshal(body, &r)

	return r.Response
}
