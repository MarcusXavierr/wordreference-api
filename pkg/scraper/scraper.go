package scraper

import (
	"github.com/antchfx/htmlquery"
	"golang.org/x/net/html"
)

type Page struct {
	Title       string
	Definitions []Definition
}

type Definition struct {
	Word     string
	Examples []string
	Meanings []Meaning
}

type Meaning struct {
	Original, Translation string
}

func ParsePage(buffer *html.Node) Page {
	titleElement, err := htmlquery.Query(buffer, "//h3[@class='headerWord']")

	if err != nil {
		return Page{}
	}

	words, err := htmlquery.QueryAll(buffer, "//tr[@class!='langHeader']/td[@class='FrWrd']/strong")

	if err != nil {
		return Page{}
	}

	definitions, err := populateDefinitions(buffer, words)

	if err != nil {
		return Page{}
	}

	return Page{
		Title:       titleElement.FirstChild.Data,
		Definitions: definitions,
	}
}

func populateDefinitions(buffer *html.Node, nodes []*html.Node) ([]Definition, error) {
	definitions := []Definition{}
	for _, word := range nodes {
		def := Definition{Word: htmlquery.InnerText(word)}
		definitions = append(definitions, def)
	}

	return definitions, nil
}
