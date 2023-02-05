package gdoc

import (
	"encoding/json"
	"log"
)

func Parse(data []byte) string {
	doc := Document{}
	err := json.Unmarshal(data, &doc)
	if err != nil {
		log.Fatalln(err)
	}

	result := ""
	for _, se := range doc.Body.Content {
		se.setType()
		result += se.toHTML()
		// se.print()
	}
	return "<html><body>" + result + "</body></html>"
}
