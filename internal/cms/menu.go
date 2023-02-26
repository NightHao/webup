package cms

import (
	"encoding/json"
	"log"
	"webup/internal/gsheet"
)

type MenuItem struct {
	Label string `json:"label"`
	Link  string `json:"link"`
}

func GetMenu(id string) (map[string][]MenuItem, error) {
	raw, err := gsheet.Request(id, "A1:C100")
	if err != nil {
		log.Fatalln(err)
	}
	val := gsheet.ValueRange{}
	err = json.Unmarshal(raw, &val)
	if err != nil {
		return map[string][]MenuItem{}, err
	}

	items := make(map[string][]MenuItem, 2)
	items["en"] = make([]MenuItem, len(val.Values))
	items["zh"] = make([]MenuItem, len(val.Values))

	for i, item := range val.Values {
		// assert item length
		items["en"][i] = MenuItem{
			Label: item[2],
			Link:  item[0],
		}
		items["zh"][i] = MenuItem{
			Label: item[1],
			Link:  item[0],
		}
	}
	return items, nil
}
