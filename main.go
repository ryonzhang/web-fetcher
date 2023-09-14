package main

import (
	"errors"
	"fetch/downloader"
	"fmt"
	"net/url"
	"os"
)

func main() {
	args := os.Args
	urls, err := validateArgs(args)
	if err != nil {
		fmt.Print(err.Error())
	}
	for _, uri := range urls {
		err := downloader.DownloadWebPage(uri)
		if err != nil {
			fmt.Print(err.Error())
		}
	}
}

func validateArgs(args []string) (validUrls []*url.URL, err error) {
	if len(args) <= 1 {
		return nil, errors.New("there is no url provided")
	}
	errorString := ""
	for _, arg := range args[1:] {
		uri, err := url.ParseRequestURI(arg)
		if err != nil {
			errorString += err.Error() + ";"
			continue
		}
		validUrls = append(validUrls, uri)
	}
	if errorString != "" {
		err = errors.New(errorString)
	}
	return
}
