package dork

import (
	"fmt"
	"net/url"
	"strings"
)

// -----------------------------------------------------

type DorkProcessor interface {
	Process(dorks []string) error
}

// -----------------------------------------------------
// DorkPrinter just prints the dorks out to STDOUT
type DorkPrinter struct {
	UrlEncode bool
}

func (p *DorkPrinter) Process(dorks []string) error {
	for _, dork := range dorks {
		dorkUrl, err := url.Parse(dork)
		if err != nil {
			return err
		}
		if p.UrlEncode {

			queryParams := dorkUrl.Query()
			dorkUrl.RawQuery = queryParams.Encode()
		}
		fmt.Println(dorkUrl.String())
	}
	return nil
}

func urlEncodeSpaces(s string) string {
	return strings.ReplaceAll(s, " ", "%20")
}

// -----------------------------------------------------
