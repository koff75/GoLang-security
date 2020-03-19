package main

import (
	"archive/zip"
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"

	"os"

	"../metadata"
	"github.com/PuerkitoBio/goquery"
)

// i: link number, s: html link
func handler(i int, s *goquery.Selection) {
	// Find the href URL
	url, ok := s.Find("a").Attr("href")
	if !ok {
		return
	}
	// Print the found URL
	fmt.Printf("%d: %s\n", i, url)

	// GET URL
	res, err := http.Get(url)
	if err != nil {
		return
	}
	// Read the file
	buf, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return
	}
	defer res.Body.Close()

	// Unzip the file to find metadatas
	r, err := zip.NewReader(bytes.NewReader(buf), int64(len(buf)))
	if err != nil {
		return
	}

	// openxml.go : Fill the structs with element found in the file
	// cp, app structs are filled
	cp, ap, err := metadata.NewProperties(r)
	if err != nil {
		return
	}

	// Print the result from the file
	log.Printf(
		"%25s %25s - %s %s\n",
		cp.Creator,
		cp.LastModifiedBy,
		ap.Application,
		ap.GetMajorVersion())
}

func main() {
	if len(os.Args) != 3 {
		log.Fatalln("Missing arg. Usage: main.go domain filetype")
	}
	domain := os.Args[1]
	filetype := os.Args[2]

	q := fmt.Sprintf(
		"site:%s && filetype:%s && instreamset:(url title):%s",
		domain,
		filetype,
		filetype)

	search := fmt.Sprintf("http://www.bing.com/search?q=%s", url.QueryEscape(q))
	doc, err := goquery.NewDocument(search)
	if err != nil {
		log.Panicln(err)
	}

	// Used devtools on the link
	s := "html body div#b_content main ol#b_results li.b_algo h2"
	// For each link we send "<a href..."
	doc.Find(s).Each(handler)
}
