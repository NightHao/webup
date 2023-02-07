package gdrive

import (
	"encoding/json"
	"io"
	"net/http"
	"webup/internal/gdoc"
)

type ListResponse struct {
	// PageToken string
	Files []File `json:"files"`
}

type File struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

func List(folder string) ([]File, error) {
	client := gdoc.ClientMustFromFile("cred.json")
	url := "https://www.googleapis.com/drive/v3/files?q='" + folder + "'+in+parents"
	req, _ := http.NewRequest("GET", url, nil)
	var err error
	resp, err := client.Do(req)
	if err != nil {
		return []File{}, err
	}

	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	l := ListResponse{}
	err = json.Unmarshal(body, &l)
	if err != nil {
		return []File{}, err
	}
	return l.Files, nil
}
