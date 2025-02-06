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
	req.Header.Set("sec-ch-ua", `"Not A(Brand";v="8", "Chromium";v="132", "Google Chrome";v="132"`)
	req.Header.Set("sec-ch-ua-arch", `"x86"`)
	req.Header.Set("sec-ch-ua-bitness", `"64"`)
	req.Header.Set("sec-ch-ua-full-version", `"132.0.6834.162"`)
	req.Header.Set("sec-ch-ua-full-version-list", `"Not A(Brand";v="8.0.0.0", "Chromium";v="132.0.6834.162", "Google Chrome";v="132.0.6834.162"`)
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("sec-ch-ua-model", `""`)
	req.Header.Set("sec-ch-ua-platform", `"Windows"`)
	req.Header.Set("sec-ch-ua-platform-version", `"10.0.0"`)
	req.Header.Set("sec-fetch-dest", "empty")
	req.Header.Set("sec-fetch-mode", "cors")
	req.Header.Set("sec-fetch-site", "same-origin")
	req.Header.Set("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/132.0.0.0 Safari/537.36")
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
