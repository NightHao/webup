package gdoc

import (
	"context"
	"io"
	"log"
	"net/http"
	"os"

	"golang.org/x/oauth2/google"
)

func ClientMustFromFile(fn string) *http.Client {
	cred, err := os.ReadFile(fn)
	if err != nil {
		log.Fatalln(err)
	}
	conf, err := google.JWTConfigFromJSON(cred,
		"https://www.googleapis.com/auth/documents.readonly",
		"https://www.googleapis.com/auth/drive.metadata.readonly")
	if err != nil {
		log.Fatalln(err)
	}
	return conf.Client(context.Background())
}

func Request(id string) ([]byte, error) {
	client := ClientMustFromFile("cred.json")
	req, _ := http.NewRequest("GET", "https://docs.googleapis.com/v1/documents/"+id, nil)
	resp, err := client.Do(req)
	if err != nil {
		return []byte{}, err
	}

	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	return body, nil
}
