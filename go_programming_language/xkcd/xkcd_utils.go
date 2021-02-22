package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"regexp"
	"strings"
	"sync"

	"golang.org/x/net/html"
)

const archiveUrl string = "https://xkcd.com/archive/"
const cachedXkcdsPath string = "./xkcds.json"

type Xkcd struct {
	Link       string
	Num        int
	Title      string
	Transcript string
}

func getAllXkcdLinks(n *html.Node, links []string) []string {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				matched, _ := regexp.MatchString(`^\/\d+\/$`, a.Val)
				if matched {
					links = append(links, fmt.Sprintf("https://xkcd.com%sinfo.0.json", a.Val))
				}
			}
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = getAllXkcdLinks(c, links)
	}

	return links
}

func writeXkcdsToCache(xkcds []Xkcd) error {
	cacheWriteFile, err := os.Create(cachedXkcdsPath)
	if err != nil {
		return err
	}
	defer cacheWriteFile.Close()
	b, err := json.MarshalIndent(xkcds, "", "\t")
	if err != nil {
		return err
	}

	r := bytes.NewReader(b)

	_, err = io.Copy(cacheWriteFile, r)
	if err != nil {
		return err
	}

	return nil
}

func readXkcdsFromCache() ([]Xkcd, error) {
	cacheReadFile, err := os.Open(cachedXkcdsPath)
	defer cacheReadFile.Close()
	if err != nil {
		return []Xkcd{}, err
	}
	xkcds := []Xkcd{}
	err = json.NewDecoder(cacheReadFile).Decode(&xkcds)
	if err != nil {
		return []Xkcd{}, err
	}

	return xkcds, nil
}

func closeXkcdChanAfterFetch(xkcdChan chan Xkcd, wg *sync.WaitGroup) {
	wg.Wait()
	close(xkcdChan)
}

func fetchXkcd(link string, writeChan chan<- Xkcd, requestChan chan int, wg *sync.WaitGroup) {
	wg.Add(1)

	requestChan <- 1
	resp, err := http.Get(link)

	defer func() {
		wg.Done()
		resp.Body.Close()
		<-requestChan
	}()

	if err != nil {
		fmt.Println("Error fetching %s: %s", link, err)
	}

	xkcd := Xkcd{}
	json.NewDecoder(resp.Body).Decode(&xkcd)
	xkcd.Link = link[:len(link)-12]
	writeChan <- xkcd
}

func fetchXkcdsFromArchive() []Xkcd {
	xkcds, err := readXkcdsFromCache()

	if err == nil && len(xkcds) > 0 {
		return xkcds
	} else {
		fmt.Printf("Xkcds not cached in %s. Fetching from %s\n", cachedXkcdsPath, archiveUrl)
	}

	resp, err := http.Get(archiveUrl)
	if err != nil {
		log.Fatalf("Error in fetching archives from %s: %s", archiveUrl, err)
	}

	archivePage, err := html.Parse(resp.Body)
	if err != nil {
		log.Fatalf("Error in parsing archive page response: %s", err)
	}

	// Todo: Can fetch data here too
	links := getAllXkcdLinks(archivePage, []string{})

	wg := sync.WaitGroup{}
	xkcdChan := make(chan Xkcd, len(links))
	openRequestChan := make(chan int, 20)

	for _, link := range links {
		go fetchXkcd(link, xkcdChan, openRequestChan, &wg)
	}

	go closeXkcdChanAfterFetch(xkcdChan, &wg)

	for xkcd := range xkcdChan {
		xkcds = append(xkcds, xkcd)
	}

	err = writeXkcdsToCache(xkcds)

	if err != nil {
		fmt.Printf("Error writing xkcds to cache file %s\n", cachedXkcdsPath)
	}

	return xkcds
}

func searchXkcds(xkcds *[]Xkcd, searchTerm string) []Xkcd {
	output := []Xkcd{}
	for _, xkcd := range *xkcds {
		titleMatch := strings.Index(strings.ToLower(xkcd.Title), strings.ToLower(searchTerm)) >= 0
		transcriptMatch := strings.Index(strings.ToLower(xkcd.Transcript), strings.ToLower(searchTerm)) >= 0

		if titleMatch || transcriptMatch {
			output = append(output, xkcd)
		}
	}

	return output
}

func main() {
	searchTermPtr := flag.String("search", "", "Use to search xkcd titles and transcripts. Eg \"star wars\"")
	flag.Parse()

	if len(*searchTermPtr) == 0 {
		fmt.Println("Please use -search flag...Eg -search=\"star wars\"")
		return
	}

	xkcds := fetchXkcdsFromArchive()
	filteredXkcds := searchXkcds(&xkcds, *searchTermPtr)

	fmt.Printf("Your search \"%s\" returned %d results:\n", *searchTermPtr, len(filteredXkcds))
	for _, xkcd := range filteredXkcds {
		fmt.Printf("%s: %s\n", xkcd.Title, xkcd.Link)
	}
}
