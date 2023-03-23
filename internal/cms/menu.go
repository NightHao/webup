package cms

import (
	"encoding/json"
	"errors"
	"log"
	"strconv"
	"webup/internal/gsheet"
)

type MenuItem struct {
	Label   string `json:"label"`
	Link    string `json:"link"`
	Lang    string `json:"lang"`
	Id      string `json:"id"`
	driveId string
}

func GetMenu(id string) ([]MenuItem, error) {
	raw, err := gsheet.Request(id, "A1:D100")
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
		items[i] = MenuItem{
			Label:   item[0],
			Link:    item[1],
			Lang:    item[3],
			Id:      strconv.Itoa(i),
			driveId: item[2],
		}
	}
	return items, nil
}

func ResolveDriveId(cmsBase, safeId string) (string, error) {
	var err error
	var menu []MenuItem
	var intId int

	menu, err = GetMenu(cmsBase)
	if err != nil {
		return "", err
	}

	intId, err = strconv.Atoi(safeId)
	if err != nil {
		return "", err
	}

	if intId >= len(menu) {
		return "", errors.New("out of range")
	}
	return menu[intId].driveId, nil
}
