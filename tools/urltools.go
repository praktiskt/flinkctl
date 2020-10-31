package tools

import (
	"fmt"
	"net/url"
	"os"
)

func UrlOrFail(h url.URL, path string) url.URL {
	url, err := url.Parse(h.String() + path)
	if err != nil {
		fmt.Printf("Attempt to build url with host: %v and path: %v failed\n", h.Host, path)
		fmt.Println(err)
		os.Exit(1)
	}
	return *url
}
