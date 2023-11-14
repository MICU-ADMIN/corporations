package main

import (
	"fmt"
	"io/ioutil"
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

func copyFile(src, dest string) error {
	content, err := ioutil.ReadFile(src)
	if err != nil {
		return err
	}

	return ioutil.WriteFile(dest, content, 0644)
}

func sendToCloudflare(mdFilename string) error {
	content, err := ioutil.ReadFile(mdFilename)
	if err != nil {
		return err
	}

	// Include file content within the user's content message
	payload := fmt.Sprintf(`{
		"messages": [
			{"role": "system", "content": "You are a friendly assistant that helps write mermaid overviews and outlooks on go projects"},
			{"role": "user", "content": "Write a mermaid equivalent about this go file + insert go file content:\n\n%s"}
		]
	}`, string(content))

	req, err := http.NewRequest("POST", cloudflareAPIEndpoint, strings.NewReader(payload))
	if err != nil {
		return err
	}

	req.Header.Set("Authorization", "Bearer "+accessToken)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Handle the response as needed

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

	if err := copyFile(path, mdFilename); err != nil {
		return err
	}

	if err := sendToCloudflare(mdFilename); err != nil {
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
