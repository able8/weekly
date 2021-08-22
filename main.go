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
	// 	ScrapGolangWeeklyArticles(strconv.Itoa(i))
	// 	time.Sleep(time.Microsecond * 500)
	// }

	// for i := 1; i < 231; i++ {
	// 	ScrapDevOpsWeekly(fmt.Sprintf("%03d", i))
	// 	time.Sleep(time.Microsecond * 500)
	// }

	// From 106
	// for i := 180; i < 231; i++ {
	// 	ScrapDevOpsWeeklyNotes(fmt.Sprintf("%03d", i))
	// 	time.Sleep(time.Microsecond * 500)
	// }

	for i := 50; i < 100; i++ {
		ScrapKubelist(fmt.Sprintf("%d", i))
		time.Sleep(time.Microsecond * 500)
	}
}

func ScrapKubelist(Number string) {
	path := "kubelist/"
	if _, err := os.Stat(path); os.IsNotExist(err) {
		err := os.Mkdir(path, 0755)
		check(err)
	}

	url := "https://kubelist.com/issue/" + Number
	resp, err := client.Get(url)
	log.Println("getting: ", url, resp.StatusCode)
	check(err)
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotFound {
		return
	}

	f, err := os.Create(path + "kubelist-" + Number + ".md")
	check(err)
	defer f.Close()
	w := bufio.NewWriter(f)

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	check(err)

	var header string
	header, _ = doc.Find("head meta[name=description]").Attr("content")
	firstDot := strings.Index(header, ".")
	header = header[:firstDot]

	fmt.Fprintf(w, "## [Kubelist %s](%s)\n\n", header, url)

	doc.Find(".gsZrho").Each(func(i int, s *goquery.Selection) {
		issueTitle := s.Find(".ckjJvh").Text()
		if issueTitle != "" {
			blockquote, _ := s.Find("blockquote p").Html()
			fmt.Fprintf(w, "#### %s\n\n> %s\n\n", issueTitle, blockquote)
			return
		}

		title := s.Find(".cCgpZH").Find("a").First().Text()
		titleURL, _ := s.Find(".cCgpZH").Find("a").First().Attr("href")
		description := s.Find(".jtHtiv").Text()
		fmt.Fprintf(w, "1. [%s](%s)\n\n", title, titleURL)
		fmt.Fprintf(w, "    %s\n", description)
	})

	IntNumber, _ := strconv.Atoi(Number)
	fmt.Fprintf(w, "\n### [ << Prev ](kubelist-%d.md) ------------- [ Next >> ](kubelist-%d.md)", IntNumber-1, IntNumber+1)
	w.Flush()
}

func ScrapDevOpsWeeklyNotes(Number string) {
	path := "devopsweeklynotes/"
	if _, err := os.Stat(path); os.IsNotExist(err) {
		err := os.Mkdir(path, 0755)
		check(err)
	}

	url := "https://devopsish.com/" + Number + "/notes/"
	resp, err := client.Get(url)
	log.Println("getting: ", url, resp.StatusCode)
	check(err)
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotFound {
		return
	}

	f, err := os.Create(path + "devopsweeklynotes-" + Number + ".md")
	check(err)
	defer f.Close()
	w := bufio.NewWriter(f)

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	check(err)

	var updatedTime string
	updatedTime = doc.Find("article .blog-post-meta time").Text()

	fmt.Fprintf(w, "## [DevOps'ish %s Notes](%s) - %s\n\n### Notes\n", Number, url, updatedTime)

	doc.Find("article>h2~p,article>h2~hr").EachWithBreak(func(i int, s *goquery.Selection) bool {
		if s.Is("hr") {
			return false
		}

		if s.Is("p") && s.Text() != "" {
			var title, titleURL string
			title = s.Find("a").Text()
			titleURL, _ = s.Find("a").Attr("href")

			fmt.Fprintf(w, "1. [%s](%s)\n\n", title, titleURL)
		}
		return true
	})

	IntNumber, _ := strconv.Atoi(Number)
	fmt.Fprintf(w, "\n### [ << Prev ](devopsweeklynotes-%03d.md) ------------- [ Next >> ](devopsweeklynotes-%03d.md)", IntNumber-1, IntNumber+1)
	w.Flush()
}

func ScrapDevOpsWeekly(Number string) {
	f, err := os.Create("devopsweekly/devopsweekly-" + Number + ".md")
	check(err)
	defer f.Close()
	w := bufio.NewWriter(f)

	url := "https://devopsish.com/" + Number
	log.Println("getting: ", url)
	resp, err := client.Get(url)
	check(err)
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	check(err)

	var updatedTime string
	updatedTime = doc.Find("article .blog-post-meta time").Text()

	fmt.Fprintf(w, "## [DevOps'ish %s](%s) - %s\n\n", Number, url, updatedTime)

	doc.Find("article").Contents().Not("header,dev").EachWithBreak(func(i int, s *goquery.Selection) bool {
		if s.Is("#people") {
			return false
		}

		if s.Is("p") && s.Text() != "" {
			description, _ := s.Html()
			fmt.Fprintf(w, "%s\n\n", description)
		}
		return true
	})

	doc.Find("article>h2, article>h2~p").EachWithBreak(func(i int, s *goquery.Selection) bool {
		if s.Is("#devops-ish-tweet-of-the-week") {
			return false
		}
		if s.Is("h2") {
			var title string
			title = s.Text()
			fmt.Fprintf(w, "### %s\n\n", title)
		}

		if s.Is("p") && s.Text() != "" {
			var title, titleURL, description string
			title = s.Find("a").Text()
			titleURL, _ = s.Find("a").Attr("href")

			fmt.Fprintf(w, "1. [%s](%s)\n\n", title, titleURL)
			description = s.Contents().Not("a").Text()
			description = strings.TrimSpace(description)
			description = strings.TrimPrefix(description, ":")
			description = strings.TrimPrefix(description, "—")
			fmt.Fprintf(w, "    %s\n", description)
		}
		return true
	})

	IntNumber, _ := strconv.Atoi(Number)
	fmt.Fprintf(w, "\n### [ << Prev ](devopsweekly-%03d.md) ------------- [ Next >> ](devopsweekly-%03d.md)", IntNumber-1, IntNumber+1)
	w.Flush()
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
