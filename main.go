package main

import (
	"encoding/json"
	"net/http"
	"net/url"
	"strings"

	"github.com/gin-gonic/gin"
)

func extractOutputAndErrors(data map[string]interface{}) (string, string) {
	var output, errors string

	if events, ok := data["Events"].([]interface{}); ok {
		for _, event := range events {
			if eventMap, ok := event.(map[string]interface{}); ok && eventMap["Kind"] == "stdout" {
				output += eventMap["Message"].(string)
			}
		}
	}

	if err, ok := data["Errors"].(string); ok {
		errors = err
	}

	return output, errors
}

func sendPostRequest(content string) (map[string]interface{}, error) {
	client := &http.Client{}

	payload := url.Values{}
	payload.Set("version", "2")
	payload.Set("body", content)
	payload.Set("withVet", "true")

	url := "https://go.dev/_/compile?backend&" + payload.Encode()

	reqBody, _ := json.Marshal(map[string]string{"code": content})
	req, err := http.NewRequest("POST", url, strings.NewReader(string(reqBody)))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json, text/javascript, */*; q=0.01")
	req.Header.Set("User-Agent", "Thunder Client (https://www.thunderclient.com)")
	req.Header.Set("Origin", "https://go.dev")

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	return result, nil
}

func main() {
	r := gin.Default()

	r.POST("/compile", func(c *gin.Context) {
		var jsonData map[string]string
		if err := c.ShouldBindJSON(&jsonData); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"errors": err.Error(), "output": ""})
			return
		}

		result, err := sendPostRequest(jsonData["code"])
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"errors": err.Error(), "output": ""})
			return
		}

		output, errors := extractOutputAndErrors(result)
		c.JSON(http.StatusOK, gin.H{"output": output, "errors": errors})
	})

	r.Run()
}
