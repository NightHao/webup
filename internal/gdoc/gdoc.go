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

	for _, se := range doc.Body.Content {
		se.setType()
		se.print()
	}
	return ""
}
