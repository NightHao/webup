package gsheet

import (
	"io"
	"net/http"
	"webup/internal/gdoc"
)

func Request(id, rangeStr string) ([]byte, error) {
	client := gdoc.ClientMustFromFile("cred.json")
	req, _ := http.NewRequest("GET", "https://sheets.googleapis.com/v4/spreadsheets/"+id+"/values/"+rangeStr, nil)
	resp, err := client.Do(req)
	if err != nil {
		return []byte{}, err
	}

	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	return body, nil
}
