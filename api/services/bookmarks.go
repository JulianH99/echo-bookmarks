package services

import (
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

type Meta struct {
	PageTitle   string
	Image       string
	Description string
}

func GetMetadataFromUrl(url string) (*Meta, error) {
	response, err := http.Get(url)

	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	doc, err := goquery.NewDocumentFromReader(response.Body)

	if err != nil {
		return nil, err
	}

	meta := &Meta{}
	doc.Find("meta").Each(func(_ int, s *goquery.Selection) {

		var metaName string
		name, existsName := s.Attr("name")
		property, _ := s.Attr("property")

		if existsName {
			metaName = name
		} else {
			metaName = property
		}

		content, _ := s.Attr("content")

		switch metaName {
		case "og:image":
			meta.Image = content

		case "og:title":
			meta.PageTitle = content

		case "og:description":
			meta.Description = content

		case "description":
			meta.Description = content

		}

	})

	// if no og:title was found, then get the <title> tag from the doc
	if meta.PageTitle == "" {
		titleNode := doc.Find("title").First()

		meta.PageTitle = titleNode.Text()
	}

	return meta, nil

}
