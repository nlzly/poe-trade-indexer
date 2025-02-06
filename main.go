package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

func main() {
	client := &http.Client{}
	var data = strings.NewReader(`{"query":{"status":{"option":"online"},"have":["exalted"],"want":["annul"]},"sort":{"have":"asc"}}`)
	req, err := http.NewRequest("POST", "https://www.pathofexile.com/api/trade2/exchange/poe2/Standard", data)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("accept", "*/*")
	req.Header.Set("accept-language", "en-US,en;q=0.9,fr-FR;q=0.8,fr;q=0.7")
	req.Header.Set("content-type", "application/json")
	req.Header.Set("dnt", "1")
	req.Header.Set("origin", "https://www.pathofexile.com")
	req.Header.Set("priority", "u=1, i")
	req.Header.Set("referer", "https://www.pathofexile.com/trade2/exchange/poe2/Standard")
	req.Header.Set("x-requested-with", "XMLHttpRequest")
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	bodyText, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", bodyText)
}
