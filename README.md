## What It Is

Basically a wrapper for the `golang.org/x/net/html` package.

## Why?

Because writing recursive functions to iterate through `html.Node`s sucks.

## Usage Example

```go
package main

import (
	"github.com/dhaninugraha/gadogado"
	// "encoding/json"
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

	g, err := gadogado.Make(resp.Body, nil)
	/* alternately, if you wanna exclude certain tags; eg. <meta> and <style> */
	// g, err := gadogado.Make(resp.Body, gadogado.ExcludeTags("meta", "style"))

	if err != nil {
		log.Fatalf("Error making gado-gado: %s\n", err.Error())
	}

	/* ugly-print the result */
	fmt.Printf("%#v", g)

	/* or, you know, you could always pretty-print it */
	// j, err := json.MarshalIndent(g, "", "  ")
	// if err != nil {
	// 	log.Fatalf("Error marshaling to JSON: %s\n", err.Error())
	// }

	// fmt.Println(string(j))
}
```