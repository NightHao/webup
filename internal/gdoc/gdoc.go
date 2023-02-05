package gdoc

import (
	"encoding/json"
	"log"
	"strings"
)

func Parse(data []byte) string {
	doc := Document{}
	err := json.Unmarshal(data, &doc)
	if err != nil {
		log.Fatalln(err)
	}

	result := ""
	for _, se := range doc.Body.Content {
		if se.Paragraph.exists() {
			run := ""
			for _, elem := range se.Paragraph.Elements {
				if elem.TextRun.exists() {
					text := elem.TextRun.Content
					text = strings.Replace(text, "\n", "<br />", -1)
					text = strings.Replace(text, "\u000b", "<br />", -1)
					if elem.TextRun.TextStyle.Bold {
						text = "<b>" + text + "</b>"
					}
					if elem.TextRun.TextStyle.Italic {
						text = "<i>" + text + "</i>"
					}
					if elem.TextRun.TextStyle.Underline {
						text = "<u>" + text + "</u>"
					}
					run += "<span>" + text + "</span>"
				}
			}
			switch se.Paragraph.ParagraphStyle.NamedStyleType {
			case "HEADING_1":
				result += "<h1>" + run + "</h1>"
			case "HEADING_2":
				result += "<h2>" + run + "</h2>"
			case "HEADING_3":
				result += "<h3>" + run + "</h3>"
			default:
				result += "<p>" + run + "</p>"
			}
		}
	}

	return "<html><body>" + result + "</body></html>"
}
