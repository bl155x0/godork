package dork

import (
	"fmt"
	"strings"
)

const baseUrl = "https://www.google.com/search?q="

// DorkForUrl takes the given dorkfile and replaces {{URL}} with the given url
func DorkForUrl(url, dorkfile string) ([]string, error) {
	dorks, err := readLinesFromFile(dorkfile)
	if err != nil {
		return nil, err
	}

	var ret []string
	for _, dork := range dorks {
		newUrl := fmt.Sprintf("%s%s", baseUrl, strings.ReplaceAll(dork, "{{URL}}", url))
		ret = append(ret, newUrl)
	}

	return ret, nil
}

// AppendDorkForUrl takes the given dorkfile, replaces {{URL}} with the given url and appends the result to the given slice
func AppendDorkForUrl(url, dorkfile string, dorks *[]string) error {
	resultDorks, err := DorkForUrl(url, dorkfile)
	if err != nil {
		return err
	}
	*dorks = append(*dorks, resultDorks...)
	return nil
}
