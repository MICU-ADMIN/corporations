package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

const cloudflareAPIEndpoint = "https://api.cloudflare.com/client/v4/accounts/0843a42dc7915eb7a5ca1e3bb05cfce2/ai/run/@cf/meta/llama-2-7b-chat-int8"
const accessToken = "wwJH-Ts41Wquh49IXlT50krftRa4E3w7SbAq7gL3"

func convertToMarkdown(filename string) string {
	// Your logic to convert file content to markdown goes here
	// For simplicity, let's just add a markdown extension to the original filename
	return strings.TrimSuffix(filename, filepath.Ext(filename)) + ".md"
}

func copyFileWithCodeBlock(src, dest, fileExtension string) error {
	if fileExtension == "md" {
		// Skip processing .md files
		return nil
	}

	content, err := ioutil.ReadFile(src)
	if err != nil {
		return err
	}

	// Wrap the content with the specified markdown code block without the dot in the file extension
	wrappedContent := fmt.Sprintf("```%s\n\n%s\n\n```\n", fileExtension, string(content))

	return ioutil.WriteFile(dest, []byte(wrappedContent), 0644)
}


type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

func sendToCloudflare(mdFilename string) error {
	content, err := ioutil.ReadFile(mdFilename)
	if err != nil {
		return err
	}

	messages := []Message{
		{"system", "You are a friendly assistant that helps write mermaid overviews and outlooks on go projects"},
		{"user", fmt.Sprintf("Write a mermaid equivalent about this go file%s}", content)},
	}

	inputs := map[string]interface{}{"messages": messages}
	inputJSON, err := json.Marshal(inputs)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", cloudflareAPIEndpoint+"@cf/meta/llama-2-7b-chat-int8", bytes.NewBuffer(inputJSON))
	if err != nil {
		return err
	}

	req.Header.Set("Authorization", "Bearer "+accessToken)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
		return err
	}
	defer resp.Body.Close()

	// Read the response body
	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
		return err
	}

	// Check for a successful response (status code 2xx)
	if resp.StatusCode/100 != 2 {
		log.Fatalf("Cloudflare API request failed with status code %d: %s", resp.StatusCode, string(responseBody))
		return fmt.Errorf("Cloudflare API request failed with status code %d", resp.StatusCode)
	}

	// Append the response to the Markdown file
	responseMarkdown := fmt.Sprintf("```mermaid\n\n%s\n\n```\n", string(responseBody))
    newContent := string(content) + "\n" + responseMarkdown

	// Update the Markdown file
	err = ioutil.WriteFile(mdFilename, []byte(newContent), 0644)
	if err != nil {
		return err
	}

	return nil
}


func processFile(path string, info os.FileInfo, err error) error {
	if err != nil {
		return err
	}

	if info.IsDir() {
		return nil
	}

	mdFilename := convertToMarkdown(path)
	fmt.Printf("Copying %s to %s\n", path, mdFilename)

	fileExtension := strings.TrimPrefix(filepath.Ext(path), ".")
	if err := copyFileWithCodeBlock(path, mdFilename, fileExtension); err != nil {
		return err
	}

	if err := sendToCloudflare(mdFilename); err != nil {
       log.Fatal(err)
		return err
	}

	return nil
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run main.go <repository_path>")
		return
	}

	repoPath := os.Args[1]
	err := filepath.Walk(repoPath, processFile)
	if err != nil {
		fmt.Println("Error:", err)
	}
}
