package gdoc

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)

type ErrorDoc struct {
	Err *struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
		Status  string `json:"status"`
	} `json:"error"`
}

func CheckError(data []byte) error {
	edoc := ErrorDoc{}
	err := json.Unmarshal(data, &edoc)
	if err != nil {
		return err
	}
	if e := edoc.Err; e != nil {
		return errors.New(fmt.Sprintf("%d %s: %s",
			e.Code, e.Status, e.Message))
	}
	return nil
}

func Parse(data []byte) (string, error) {
	doc := Document{}
	err := json.Unmarshal(data, &doc)
	if err != nil {
		return "", err
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
					if link := textRun.TextStyle.Link; link != nil {
						text = fmt.Sprintf("<a href=\"%s\">%s</a>",
							link.Url, text)
					}
					style := ""
					if c := textRun.TextStyle.ForegroundColor; c != nil {
						style = " style=\"color: " + c.toCssRgb() + "\" "
					}
					run += "<span" + style + ">" + text + "</span>"
				} else if ioe := elem.InlineObjectElement; ioe != nil {
					ioeKey := ioe.InlineObjectId
					inlineObject := doc.InlineObjects[ioeKey]
					embeddedObject := inlineObject.InlineObjectProperties.EmbeddedObject
					if image := embeddedObject.ImageProperties; image != nil {
						run += fmt.Sprintf("<img src=\"%s\" width=\"%f\" height=\"%f\"",
							image.ContentUri, embeddedObject.Size.Width.Magnitude, embeddedObject.Size.Height.Magnitude)
						run += "<img src=\"" + image.ContentUri + "\" "
						run += "width="
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

	return "<html><body>" + result + "</body></html>", nil
}
