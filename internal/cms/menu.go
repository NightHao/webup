package cms

import (
	"encoding/json"
	"errors"
	"log"
	"strconv"

	"webup/internal/gsheet"
)

type MenuItem struct {
	Title    string     `json:"title"`
	Link     string     `json:"link"`
	Data     string     `json:"data"`
	Type     string     `json:"type"`
	Lang     string     `json:"lang"`
	Children []MenuItem `json:"children"`
	driveId  string
}

func parseItem(vals []gsheet.Value,
	startIdx, length int) (int, []MenuItem) {
	// TITLE | LINK | DATA | TYPE | LANG | LEVEL
	//   0   |  1   |  2   |  3   |  4   |  5
	arr := make([]MenuItem, 0)
	for i := startIdx; i < length; {
		item := vals[i]
		mi := MenuItem{
			Title:    item[0],
			Link:     item[1],
			Data:     strconv.Itoa(i), // index of the menu entry
			Type:     item[3],
			Lang:     item[4],
			Children: nil,
			driveId:  "",
		}
		i++
		switch item[5] {
		case "<":
			// last item of this level
			arr = append(arr, mi)
			return i, arr
		case ">":
			// nest
			i, mi.Children = parseItem(vals, i, length)
			arr = append(arr, mi)
			break
		default:
			arr = append(arr, mi)
		}
	}
	return length, arr
}

func GetMenu(id string) ([]MenuItem, error) {
	raw, err := gsheet.Request(id, "A2:F100")
	if err != nil {
		log.Fatalln(err)
	}
	val := gsheet.ValueRange{}
	err = json.Unmarshal(raw, &val)
	if err != nil {
		return []MenuItem{}, err
	}

	_, items := parseItem(val.Values, 0, len(val.Values))
	return items, nil
}

func findFromNested(vals []gsheet.Value,
	startIdx, targetIdx, length int) (int, string) {
	for i := startIdx; i < length; {
		item := vals[i]
		if i == targetIdx {
			return i, item[2]
		}
		i++
		switch item[5] {
		case "<":
			return i, ""
		case ">":
			if ni, r := findFromNested(vals, i, targetIdx, length); r != "" {
				return ni, r
			} else {
				i = ni
			}
			break
		}
	}
	return length, ""
}

func ResolveDriveId(cmsBase, safeId string) (string, error) {
	var err error
	var intId int

	intId, err = strconv.Atoi(safeId)
	if err != nil {
		return "", err
	}

	raw, err := gsheet.Request(cmsBase, "A2:F100")
	if err != nil {
		log.Fatalln(err)
	}
	val := gsheet.ValueRange{}
	err = json.Unmarshal(raw, &val)
	if err != nil {
		return "", err
	}

	_, driveId := findFromNested(val.Values, 0, intId, len(val.Values))
	if driveId == "" {
		return "", errors.New("out of range")
	}

	return driveId, nil
}
