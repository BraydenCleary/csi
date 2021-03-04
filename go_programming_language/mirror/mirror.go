package main

import (
	"crypto/sha1"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"golang.org/x/net/html"
)

func main() {
	// TODO: make this passedURL dynamic
	passedURL := "https://golang.org"
	parsedURL, err := url.Parse(passedURL)
	if err != nil {
		log.Fatal(fmt.Errorf("Unable to parse url, please try again %s", err))
	}

	baseURL := fmt.Sprintf("%s://%s", parsedURL.Scheme, parsedURL.Host)

	workChan := make(chan string)
	wg := sync.WaitGroup{}

	wg.Add(1)
	go saveURL(passedURL, baseURL, workChan, &wg)

	go func() {
		wg.Wait()
		close(workChan)
	}()

	savedUrls := make(map[string]bool)
	for url := range workChan {
		// Figure out how to handle pages like "downloads" page for golang
		// Possibly set a timeout on our call to http.Get
		if !savedUrls[url] && strings.Index(url, "https://golang.org/dl/") != 0 {
			wg.Add(1)
			savedUrls[url] = true
			go saveURL(url, baseURL, workChan, &wg)
		}
	}
}

func saveURL(url string, baseURL string, workChan chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()

	resp, err := http.Get(url)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	if nonHTMLResponse(resp) {
		fmt.Printf("Html not found in response for url %s. Exiting.\n", url)
		return
	}

	rootNode, err := html.Parse(resp.Body)
	if err != nil {
		fmt.Println(fmt.Errorf("Error in parsing html response for url %s", err))
		return
	}

	// TODO: Call this as a go routine
	crawlPageForLinks(rootNode, baseURL, workChan)

	writeFile, err := os.Create(generateFilePath(url))
	if err != nil {
		fmt.Println(fmt.Errorf("Error in creating write file %s", generateFilePath(url)))
		return
	}
	defer writeFile.Close()

	err = html.Render(writeFile, rootNode)
	if err != nil {
		fmt.Println(fmt.Errorf("Error writing to write file %s", generateFilePath(url)))
		return
	}
}

func generateFilePath(url string) string {
	urlHash := sha1.Sum([]byte(url))
	// TODO: Make the write dir dynamic based on the passedURL
	return filepath.Join("/var/tmp/test", fmt.Sprintf("%x%s", urlHash, ".html"))
}

func crawlPageForLinks(n *html.Node, baseURL string, workChan chan<- string) {
	if n.Type == html.ElementNode && n.Data == "a" {
		for idx, a := range n.Attr {
			if a.Key == "href" {
				if strings.Index(a.Val, "/") == 0 && strings.Index(a.Val, "//") != 0 {
					urlToSave := ""
					if strings.Index(a.Val, "#") >= 0 {
						parseID := strings.Split(a.Val, "#")
						urlToSave = fmt.Sprintf("%s%s", baseURL, parseID[0])
						n.Attr[idx].Val = fmt.Sprintf("%s#%s", generateFilePath(urlToSave), parseID[1])
					} else {
						urlToSave = fmt.Sprintf("%s%s", baseURL, a.Val)
						n.Attr[idx].Val = generateFilePath(urlToSave)
					}
					workChan <- urlToSave
				}
			}
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		crawlPageForLinks(c, baseURL, workChan)
	}
}

func nonHTMLResponse(response *http.Response) bool {
	htmlResponse := false
	for _, contentType := range response.Header["Content-Type"] {
		if strings.Index(contentType, "text/html") >= 0 {
			htmlResponse = true
			break
		}
	}

	return !htmlResponse
}
