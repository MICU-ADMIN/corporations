```go

package main

import (
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

func sendToCloudflare(mdFilename string) error {
	content, err := ioutil.ReadFile(mdFilename)
	if err != nil {
		return err
	}

	// Include file content within the user's content message
	payload := fmt.Sprintf(`{
	"messages": [
		{"role": "system", "content": "You are a friendly assistant that helps write mermaid overviews and outlooks on go projects"},
		{"role": "user", "content": "Write a mermaid equivalent about this go file%s}"}
	]
}`, content)


	req, err := http.NewRequest("POST", cloudflareAPIEndpoint, strings.NewReader(payload))
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


```
