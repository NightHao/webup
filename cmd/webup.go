package main

import (
	"context"
	"golang.org/x/oauth2/google"
	"log"
	"webup/internal/gdoc"

	"fmt"
	"io"
	"net/http"
	"os"
)

func clientMustFromFile(fn string) *http.Client {
	cred, err := os.ReadFile(fn)
	if err != nil {
		log.Fatalln(err)
	}
	conf, err := google.JWTConfigFromJSON(cred, "https://www.googleapis.com/auth/documents.readonly")
	if err != nil {
		log.Fatalln(err)
	}
	return conf.Client(context.Background())
}

func request(id string) {
	//id := "1zDMdNztd20tlLttEWEBuYOTH6-qUSmgMiZmykmMfjhQ"

	client := clientMustFromFile("cred.json")
	req, _ := http.NewRequest("GET", "https://docs.googleapis.com/v1/documents/"+id, nil)
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	fmt.Println(string(body))
}

func main() {
	raw, _ := os.ReadFile("dump.json")
	gdoc.Parse(raw)
}
