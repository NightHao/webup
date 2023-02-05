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
		if paragraph := se.Paragraph; paragraph != nil {
			run := ""
			for _, elem := range paragraph.Elements {
				if textRun := elem.TextRun; textRun != nil {
					text := textRun.Content
					text = strings.Replace(text, "\n", "<br />", -1)
					text = strings.Replace(text, "\u000b", "<br />", -1)
					if textRun.TextStyle.Bold {
						text = "<b>" + text + "</b>"
					}
					if textRun.TextStyle.Italic {
						text = "<i>" + text + "</i>"
					}
					if textRun.TextStyle.Underline {
						text = "<u>" + text + "</u>"
					}
					run += "<span>" + text + "</span>"
				} else if ioe := elem.InlineObjectElement; ioe != nil {
					ioeKey := ioe.InlineObjectId
					inlineObject := doc.InlineObjects[ioeKey]
					if image := inlineObject.InlineObjectProperties.EmbeddedObject.ImageProperties; image != nil {
						run += "<img src=\"" + image.ContentUri + "\" />"
					}
				}
			}
			switch paragraph.ParagraphStyle.NamedStyleType {
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
