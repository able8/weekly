package main

import (
	"bufio"
	"crypto/tls"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

type Article struct {
	Title       string `json:"title,omitempty"`
	TargetUrl   string `json:"target_url,omitempty"`
	SourceUrl   string `json:"source_url,omitempty"`
	SourceName  string `json:"source_name,omitempty"`
	Description string `json:"description,omitempty"`
}

var client *http.Client

func init() {
	tr := &http.Transport{
		MaxIdleConnsPerHost: 100,
		// Necessary to skip certificates validation
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	}
	client = &http.Client{
		Transport: tr,
		Timeout:   time.Second * 30,
	}
}

func check(e error) {
	if e != nil {
		log.Fatalln(e)
	}
}

func main() {
	// for i := 205; i < 376; i++ {
	// 	err := ScrapGolangWeeklyArticles(strconv.Itoa(i))
	// 	time.Sleep(time.Microsecond * 500)
	// 	if err != nil {
	// 		log.Fatalln(err)
	// 	}
	// }

	for i := 3; i < 284; i++ {
		ScrapSreWeekly(strconv.Itoa(i))
	}
}

func ScrapSreWeekly(Number string) {
	f, err := os.Create("sreweekly/sreweekly-" + Number + ".md")
	check(err)
	defer f.Close()
	w := bufio.NewWriter(f)

	url := "https://sreweekly.com/sre-weekly-issue-" + Number
	log.Println("getting: ", url)
	resp, err := client.Get(url)
	check(err)
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	check(err)

	var headerTitle, headerUrl, updatedTime string
	header := doc.Find("#content").Find("header").First()
	headerTitle = header.Find("a").Text()
	headerUrl, _ = header.Find("a").Attr("href")
	updatedTime = doc.Find("#content").Find(".updated").First().Text()

	fmt.Fprintf(w, "## [%s](%s) - %s\n", headerTitle, headerUrl, updatedTime)
	fmt.Fprintf(w, "### Articles\n\n")

	doc.Find("#content").Find(".sreweekly-entry").Each(func(i int, s *goquery.Selection) {
		var title, titleURL, description string
		title = s.Find(".sreweekly-title>a").Text()
		titleURL, _ = s.Find(".sreweekly-title>a").Attr("href")
		description = s.Find(".sreweekly-description>p").Text()

		fmt.Fprintf(w, "1. [%s](%s)\n\n    %s\n", title, titleURL, description)
	})

	fmt.Fprintf(w, "### Outages\n\n")
	doc.Find("#content").Find(".entry-content.cf>.sreweekly-outages>li").Each(func(i int, s *goquery.Selection) {
		var title, titleURL, description string
		title = s.Find("a").Text()
		titleURL, _ = s.Find("a").Attr("href")
		description = s.Find("li").Text()
		fmt.Fprintf(w, "1. [%s](%s)\n", title, titleURL)
		if description != "" {
			fmt.Fprintf(w, "    %s\n", strings.TrimSpace(description))
		}
	})

	IntNumber, _ := strconv.Atoi(Number)
	fmt.Fprintf(w, "\n### [ << Prev ](sreweekly-%d.md) ------------- [ Next >> ](sreweekly-%d.md)", IntNumber-1, IntNumber+1)
	w.Flush()
}

func ScrapGolangWeeklyArticles(Number string) {
	f, err := os.Create("golangweekly/golangweekly-" + Number + ".md")
	check(err)
	defer f.Close()
	w := bufio.NewWriter(f)

	url := "https://golangweekly.com/issues/" + Number
	resp, err := client.Get(url)
	log.Println("getting: ", url)
	check(err)
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	check(err)

	// articles := make([]Article, 0)
	var header string
	header = doc.Find("#content").Find("table").First().Find("p").First().Text()
	fmt.Fprintf(w, "## [%s](%s)\n\n", header, url)

	doc.Find("#content").Find(".el-item,.miniitem,.el-heading").Each(func(i int, s *goquery.Selection) {
		var isSponsored bool
		isSponsored = s.Find("span").Last().HasClass("tag-sponsor")

		if s.HasClass("el-heading") {
			header = s.Find("p").First().Text()
			fmt.Fprintf(w, "### %s\n\n", header)
		}

		if !isSponsored {
			var title, targetUrl, desc string

			desc = s.Find(".desc").First().Contents().FilterFunction(func(i int, s *goquery.Selection) bool {
				return !s.Is("span")
			}).Text()
			title = s.Find(".desc").First().Find("a").Text()
			targetUrl, _ = s.Find(".desc").First().Find("a").Attr("href")

			if targetUrl != "" {
				fmt.Fprintf(w, "1. [%s](%s)\n\n", title, targetUrl)
			}

			if strings.TrimSpace(desc) != "" {
				fmt.Fprintf(w, "    %s\n", strings.TrimRight(desc, " "))
			}
		}
	})

	IntNumber, _ := strconv.Atoi(Number)
	fmt.Fprintf(w, "\n### [ << Prev ](golangweekly-%d.md) ------------- [ Next >> ](golangweekly-%d.md)", IntNumber-1, IntNumber+1)
	w.Flush()
}
