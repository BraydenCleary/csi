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
	passedUrl := "https://golang.org"
	parsedUrl, err := url.Parse(passedUrl)
	if err != nil {
		log.Fatal(err)
	}

	baseUrl := fmt.Sprintf("%s://%s", parsedUrl.Scheme, parsedUrl.Host)

	workChan := make(chan string)
	wg := sync.WaitGroup{}

	wg.Add(1)
	go saveUrl(passedUrl, baseUrl, workChan, &wg)

	go func() {
		wg.Wait()
		close(workChan)
	}()

	savedUrls := make(map[string]bool)
	for url := range workChan {
		if !savedUrls[url] && strings.Index(url, "https://golang.org/dl/") != 0 {
			wg.Add(1)
			savedUrls[url] = true
			go saveUrl(url, baseUrl, workChan, &wg)
		}
		fmt.Println(len(savedUrls))
	}
}

func saveUrl(url string, baseUrl string, workChan chan<- string, wg *sync.WaitGroup) {
	fmt.Println("savingUrl")
	defer wg.Done()
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	defer func() { resp.Body.Close() }()

	if nonHtmlResponse(resp) {
		fmt.Printf("Html in response for url %s not found. Exiting.\n", url)
	}

	generateFilePath(url)

	rootNode, err := html.Parse(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	wg.Add(1)
	crawlPageForLinks(rootNode, baseUrl, workChan, wg)

	writeFile, err := os.Create(generateFilePath(url))
	defer func() { writeFile.Close() }()
	if err != nil {
		log.Fatal(err)
	}
	err = html.Render(writeFile, rootNode)
	if err != nil {
		log.Fatal(err)
	}
}

func generateFilePath(url string) string {
	urlHash := sha1.Sum([]byte(url))
	return filepath.Join("/var/tmp/test", fmt.Sprintf("%x%s", urlHash, ".html"))
}

func crawlPageForLinks(n *html.Node, baseUrl string, workChan chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()
	if n.Type == html.ElementNode && n.Data == "a" {
		for idx, a := range n.Attr {
			if a.Key == "href" {
				if strings.Index(a.Val, "/") == 0 && strings.Index(a.Val, "//") != 0 {
					parseID := strings.Split(a.Val, "#")
					urlToSave := ""
					if len(parseID) > 1 {
						urlToSave = fmt.Sprintf("%s%s", baseUrl, parseID[0])
						n.Attr[idx].Val = fmt.Sprintf("%s#%s", generateFilePath(urlToSave), parseID[1])
					} else {
						urlToSave = fmt.Sprintf("%s%s", baseUrl, a.Val)
						n.Attr[idx].Val = generateFilePath(urlToSave)
					}

					fmt.Println(urlToSave)
					workChan <- urlToSave
				}
			}
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		wg.Add(1)
		go crawlPageForLinks(c, baseUrl, workChan, wg)
	}
}

func nonHtmlResponse(response *http.Response) bool {
	htmlResponse := false
	for _, contentType := range response.Header["Content-Type"] {
		if strings.Index(contentType, "text/html") >= 0 {
			htmlResponse = true
			break
		}
	}

	return !htmlResponse
}
