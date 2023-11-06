package main

import (
	// "container/list"
	"fmt"
	"os"

	"strings"

	"golang.org/x/net/html"
)

type Link struct {
	Href string
	Text string
}

func main() {
	file, err := os.Open("index.html")
	if err != nil {
		fmt.Println("Error opening file: ", err)
		return 
	}

	doc, err := html.Parse(file)
	if err != nil {
		fmt.Println("Error parsing html: ", err)
	}
	
	nodes := findLinkNodes(doc)
	var links []Link
	for _, node := range nodes {

		links = append(links, createLink(node))
	}

	for _, link := range links {
		fmt.Printf("{\nHref: %s\nText: %s\n}\n",link.Href, link.Text)
	}
}

func findLinkNodes(n *html.Node) []*html.Node {
	var linkNodes []*html.Node
	
	if n.Type == html.ElementNode && n.Data == "a" {
		return []*html.Node{n}
	}
	

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		linkNodes = append(linkNodes, findLinkNodes(c)...)
	}

	return linkNodes
}

func createLink(n *html.Node) Link {
	var link Link 
	for _, node := range n.Attr {
		if node.Key == "href" {
			link.Href = node.Val
		}
	}
	link.Text = getText(n)
	// fmt.Println(link.Text)
	return link
}


func getText(n *html.Node) string {
	if n.Type == html.TextNode {
		return n.Data
	}
	if n.Type != html.ElementNode {
		return ""
	}

	var text string

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		text += getText(c)
	}

	return strings.Join(strings.Fields(text), " ")
	// return "hello"
}