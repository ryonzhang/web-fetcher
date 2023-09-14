package parser

import (
	"fmt"
	"golang.org/x/net/html"
	"net/url"
	"strings"
)

func GetAllAssets(bodyContent []byte) ([]string, error) {
	var assets []string
	content := string(bodyContent)
	doc, err := html.Parse(strings.NewReader(content))
	if err != nil {
		return assets, err
	}
	var f func(*html.Node)
	f = func(n *html.Node) {
		// Anchor pages
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, a := range n.Attr {
				if a.Key == "href" {
					assets = append(assets, a.Val)
					break
				}
			}
		}
		// Image sources
		if n.Type == html.ElementNode && n.Data == "img" {
			for _, a := range n.Attr {
				if a.Key == "src" {
					assets = append(assets, a.Val)
					break
				}
			}
		}
		// Stylesheet
		if n.Type == html.ElementNode && n.Data == "link" {
			for _, a := range n.Attr {
				if a.Key == "href" {
					assets = append(assets, a.Val)
					break
				}
			}
		}
		// Javascript
		if n.Type == html.ElementNode && n.Data == "script" {
			for _, a := range n.Attr {
				if a.Key == "src" {
					assets = append(assets, a.Val)
					break
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(doc)
	var result []string
	for _, asset := range assets {
		uri, err := url.ParseRequestURI(asset)
		if err != nil {
			fmt.Print(err.Error())
			continue
		}
		if !uri.IsAbs() {
			result = append(result, asset)
		}
	}
	return result, nil
}

func GetLinksFromBody(bodyContent []byte) (int, error) {
	links := 0
	content := string(bodyContent)
	doc, err := html.Parse(strings.NewReader(content))
	if err != nil {
		return links, err
	}
	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, a := range n.Attr {
				if a.Key == "href" {
					links += 1
					break
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(doc)
	return links, nil
}

func GetImgsFromBody(bodyContent []byte) (int, error) {
	imgs := 0
	content := string(bodyContent)
	doc, err := html.Parse(strings.NewReader(content))
	if err != nil {
		return imgs, err
	}
	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "img" {
			imgs += 1
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(doc)
	return imgs, nil
}
