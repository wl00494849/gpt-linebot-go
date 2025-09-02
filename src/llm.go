package src

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type gpt struct {
	PYTHON_GPT_URL string
}

type Response struct {
	Response string `json:"response"`
}

func NewGPT() *gpt {
	return &gpt{
		PYTHON_GPT_URL: os.Getenv("PYTHON_GPT_URL"),
	}
}

func (g *gpt) Requset(message string, userID string) string {
	var r Response
	data := struct {
		Message string `json:"message"`
		UserID  string `json:"userID"`
	}{Message: message, UserID: userID}

	b, _ := json.Marshal(data)

	req, err := http.NewRequest(
		http.MethodPost,
		g.PYTHON_GPT_URL+"/gpt/",
		bytes.NewReader(b),
	)

	if err != nil {
		fmt.Println("PostReqerr:", err)
		return ""
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("Decode error:", err)
		return ""
	}

	defer resp.Body.Close()

	json.NewDecoder(resp.Body).Decode(&r)

	return r.Response
}
