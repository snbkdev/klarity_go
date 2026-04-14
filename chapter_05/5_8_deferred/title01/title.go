// 5.8 Отложенные вызовы функций
package title01

import (
	"fmt"
	"net/http"
	"program_lang/chapter_05/5_2_recursion/findlinks1/html"
	"strings"
)

func title(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}

	ct := resp.Header.Get("Content-Type")
	if ct != "text/html" && !strings.HasPrefix(ct, "text/html;") {
		resp.Body.Close()
		return fmt.Errorf("%s имеет тип %s, не text/html", url, ct)
	}

	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		return fmt.Errorf("анализ %s как HTML: %v", url, err)
	}

	visitNode := func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "title" && n.FirstChild != nil {
			fmt.Println(n.FirstChild.Data)
		}
	}
	ForEachNode(doc, visitNode, nil)
	return nil
}

func ForEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		ForEachNode(c, pre, post)
	}

	if post != nil {
		post(n)
	}
}