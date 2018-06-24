package main

import (
	"errors"
	"flag"
	"fmt"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"sync"

	"github.com/PuerkitoBio/goquery"
	"github.com/aymerick/douceur/parser"
)

var (
	bihua = regexp.MustCompile("bihua_([0-9]+).html")
)

func getChar(charCode int64) (err error) {
	var req *http.Request
	req, err = http.NewRequest("GET", fmt.Sprintf("https://zidian.911cha.com/zi%x.html", charCode), nil)
	if err != nil {
		return
	}

	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/67.0.3396.87 Safari/537.36")
	client := &http.Client{}

	var res *http.Response
	res, err = client.Do(req)
	if err != nil {
		return
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		err = errors.New("status code != 200")
		return
	}

	var doc *goquery.Document
	doc, err = goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return
	}

	style := doc.Find("style").Text()

	// for anti-crawler purpose, the web page randomly generates fake content and use stylesheet to make fake content hidden
	// we parse the stylesheet and collect the class names having the "display: none" rule
	var fakes []string
	stylesheet, _ := parser.Parse(style)
	for _, rule := range stylesheet.Rules {
		decls := (*rule).Declarations
		for i := len(decls) - 1; i > -1; i-- {
			if strings.Contains(decls[i].String(), "none") {
				for _, sel := range (*rule).Selectors {
					fakes = append(fakes, strings.TrimLeft(sel, "."))
				}
				break
			}
		}
	}

	// remove variant
	doc.Find(".mcon > p").Each(func(_ int, s *goquery.Selection) {
		if strings.Contains(s.Text(), "繁体部首") {
			s.Remove()
		}
	})

	var codes []string
	doc.Find(".z_d").Each(func(_ int, s *goquery.Selection) {
		// if the element contains "fake" classes, remove it
		for _, fake := range fakes {
			if s.HasClass(fake) {
				return
			}
		}
		href, exists := s.Find("a").Attr("href")
		if exists {
			if v := bihua.FindStringSubmatch(href); len(v) > 1 {
				codes = append(codes, v[1])
			}
		}
	})
	fmt.Println(charCode, strings.Join(codes, " "))
	return
}

func main() {
	flag.Parse()

	start, err := strconv.ParseInt(flag.Arg(0), 16, 64)
	if err != nil {
		panic(err)
	}

	end, err := strconv.ParseInt(flag.Arg(1), 16, 64)
	if err != nil {
		panic(err)
	}

	concurrency := 5
	codes := make(chan int64)

	go func() {
		defer close(codes)

		for i := start; i <= end; i++ {
			codes <- i
		}
	}()

	var wg sync.WaitGroup
	wg.Add(concurrency)
	for i := 0; i < concurrency; i++ {
		go func() {
			defer wg.Done()
			for code := range codes {
				if err := getChar(code); err != nil {
					fmt.Println("error while getting char code =", code)
					panic(err)
				}
			}
		}()
	}
	wg.Wait()
}
