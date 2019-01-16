## What It Is

Basically a wrapper for the `golang.org/x/net/html` package.

## Why?

Because writing recursive functions to iterate through `html.Node`s sucks. Also, I believe I haven't found a good alternative to Python's `BeautifulSoup` in Go.

## Usage Example

```go
package main

import (
	"github.com/dhaninugraha/gadogado"
	"encoding/json"
	"net/http"
	"time"
	"fmt"
	"log"
)

func main() {
	var url = "http://motherfuckingwebsite.com"

	httpClient := &http.Client{
		Timeout: 10 * time.Second,
	}

	resp, err := httpClient.Get(url)
	if err != nil {
		log.Fatalf("Error fetching %q: %s\n", url, err.Error())
	}
	defer resp.Body.Close()

	gado2, err := gadogado.Make(resp.Body, nil)
	/* alternately, if you wanna exclude certain tags; eg. <meta> and <style> */
	// g, err := gadogado.Make(resp.Body, gadogado.ExcludeTags("meta", "style"))

	if err != nil {
		log.Fatalf("Error making gado-gado: %s\n", err.Error())
	}

	asJson, err := json.MarshalIndent(gado2, "", " ")
	if err != nil {
		log.Fatalf("Error marshaling to JSON: %s\n", err.Error())
	}

	fmt.Println(string(asJson))


	// cherry-pick a certain element
	styleTag := gado2.CherryPick("style")
	asJson, err = json.MarshalIndent(styleTag, "", " ")
	if err != nil {
		log.Fatalf("Error marshaling to JSON: %s\n", err.Error())
	}

	fmt.Println(string(asJson))	
}
```