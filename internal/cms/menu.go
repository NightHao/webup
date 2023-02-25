package cms

import (
	"encoding/json"
	"log"
	"webup/internal/gsheet"
)

type MenuItem struct {
	//	LabelEN string `json:"labelEN"`
	LabelTW string `json:"labelTW"`
	Link    string `json:"link"`
}

func GetMenu(id string) ([]MenuItem, error) {
	raw, err := gsheet.Request(id, "A1:B100")
	if err != nil {
		log.Fatalln(err)
	}
	val := gsheet.ValueRange{}
	err = json.Unmarshal(raw, &val)
	if err != nil {
		return []MenuItem{}, err
	}

	items := make([]MenuItem, len(val.Values))
	for i, item := range val.Values {
		// assert item length
		items[i] = MenuItem{
			LabelTW: item[0],
			Link:    item[1],
		}
	}
	return items, nil
}
