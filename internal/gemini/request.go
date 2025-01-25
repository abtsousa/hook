package gemini

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/abtsousa/hook/safe_error"
	"io"
	"net/http"
	"os"
	"runtime"
)

// Standard JSON output for Gemini requests.
type Request struct {
	Contents []struct {
		Parts []struct {
			Text string `json:"text"`
		} `json:"parts"`
	} `json:"contents"`
}

// Make a request to the Gemini API.
func makeRequest(text string) Request {

	const PROMPT = `You are an expert in generating terminal commands. Given a user's request describing what they want to do in the terminal, along with the following system information:
	
	Operating System: %s
	Architecture: %s
	Current Shell: %s
	
	Your task is to provide ONLY the most concise and correct terminal command that fulfills the user's request. Do NOT provide any explanation, commentary, or additional text, just the raw command.
	
	User Request: %s

	Command:
	`

	user_os := runtime.GOOS
	user_arch := runtime.GOARCH
	body := fmt.Sprintf(PROMPT, user_os, user_arch, os.Getenv("SHELL"), text)

	return Request{
		Contents: []struct {
			Parts []struct {
				Text string `json:"text"`
			} `json:"parts"`
		}{
			{
				Parts: []struct {
					Text string `json:"text"`
				}{
					{Text: body},
				},
			},
		},
	}
}

// Queries Gemini and outputs its response as a string.
func (c *Client) Query(query string) (string, error) {

	url := BASE_URL + API_KEY
	payload, err := json.Marshal(makeRequest(query))
	if err != nil {
		return "", safe_error.Return("Error marshalling request: %v", err)
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payload))
	if err != nil {
		return "", safe_error.Return("couldn't make request: %v", err)
	}

	rsp, err := c.client.Do(req)
	if err != nil {
		return "", safe_error.Return("couldn't get response: %v", err)
	}

	dat, err := io.ReadAll(rsp.Body)
	if err != nil {
		return "", safe_error.Return("couldn't parse received data: %v", err)
	}

	var t Response
	err = json.Unmarshal(dat, &t)
	if err != nil {
		return "", safe_error.Return("couldn't unmarshal received data: %v", err)
	}

	return t.Candidates[0].Content.Parts[0].Text, nil
}
