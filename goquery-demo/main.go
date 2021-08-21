package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	html := `
<body>
	<div lang="zh">DIV1</div>
	<p>P1</p>
	<div lang="zh-cn">DIV2</div>
	<div lang="en">DIV3</div>
	<span>
		<div>DIV4</div>
	</span>
	<p>P2</p>
</body>`

	dom, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		log.Fatalln(err)
	}

	// https://www.flysnow.org/2018/01/20/golang-goquery-examples-selector.html#prevnext%E9%80%89%E6%8B%A9%E5%99%A8
	fmt.Println("prev+next相邻选择器")
	// 假设我们要筛选的元素没有规律，但是该元素的上一个元素有规律，我们就可以使用这种下一个相邻选择器来进行选择。
	dom.Find("div[lang=zh]+p").Each(func(i int, selection *goquery.Selection) {
		fmt.Println(selection.Text())
	})
	// 有相邻就有兄弟，兄弟选择器就不一定要求相邻了，只要他们共有一个父元素就可以。
	fmt.Println("prev~next选择器")
	dom.Find("div[lang=zh]~p").Each(func(i int, selection *goquery.Selection) {
		fmt.Println(selection.Text())
	})

	fmt.Println("内容过滤器")
	// Find(":contains(text)")表示筛选出的元素要包含指定的文本，我们例子中要求选择出的div元素要包含DIV2文本，那么只有一个DIV2元素满足要求。
	dom.Find("div:contains(DIV2)").Each(func(i int, selection *goquery.Selection) {
		fmt.Println(selection.Text())
	})

	dom.Find("span:has(div)").Each(func(i int, selection *goquery.Selection) {
		fmt.Println(selection.Text())
	})

	fmt.Println(":nth-child(n) 过滤器")
	// 这个表示筛选出的元素是其父元素的第n个元素，n以1开始。所以我们可以知道:first-child和:nth-child(1)是相等的。
	// 通过指定n，我们就很灵活的筛选出我们需要的元素。
	dom.Find("div:nth-child(3)").Each(func(i int, selection *goquery.Selection) {
		fmt.Println(selection.Text())
	})
	fmt.Println("选择器或(|)运算")
	// 如果我们想同时筛选出div,span等元素怎么办？这时候可以采用多个选择器进行组合使用，并且以逗号(,)分割，
	// Find("selector1, selector2, selectorN")表示，只要满足其中一个选择器就可以被筛选出来，也就是选择器的或(|)运算操作。
	dom.Find("div,span").Each(func(i int, selection *goquery.Selection) {
		fmt.Println(selection.Html())
	})

}
