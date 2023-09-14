package downloader

import (
	"errors"
	"fetch/parser"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func DownloadWebPage(uri *url.URL) error {
	fileName := fmt.Sprintf("%s.html", uri.Host)
	file, err := os.Create(fileName)
	defer file.Close()

	if err != nil {
		return err
	}
	client := http.Client{
		CheckRedirect: func(r *http.Request, via []*http.Request) error {
			r.URL.Opaque = r.URL.Path
			return nil
		},
	}
	fetchedTime := time.Now()
	resp, err := client.Get(uri.String())
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	bodyContent, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	_, err = file.Write(bodyContent)
	if err != nil {
		return err
	}

	links, err := parser.GetLinksFromBody(bodyContent)
	if err != nil {
		return err
	}
	imgs, err := parser.GetImgsFromBody(bodyContent)
	if err != nil {
		return err
	}
	allAssetsToDownload, err := parser.GetAllAssets(bodyContent)
	if err != nil {
		fmt.Printf("Error during getting assets:%s", err.Error())
	}
	for _, asset := range allAssetsToDownload {
		err := DownloadAsset(uri, asset)
		if err != nil {
			fmt.Printf("Error during downloading asset %s:%s", asset, err.Error())
		}
	}

	fmt.Printf("url:%s\n", uri.String())
	fmt.Printf("num_links:%d\n", links)
	fmt.Printf("images:%d\n", imgs)
	fmt.Printf("last_fetch:%s\n", fetchedTime.Format(time.RFC822))
	return nil
}

func DownloadAsset(uri *url.URL, asset string) error {
	if !strings.HasPrefix(asset, "/") {
		asset = "/" + asset
	}
	if strings.HasSuffix(asset, "/") {
		asset += "index.html"
	}
	if strings.HasPrefix(asset, "/") {
		asset = "." + asset
	}
	if _, err := os.Stat(asset); errors.Is(err, os.ErrNotExist) {
		err := os.MkdirAll(filepath.Dir(asset), os.ModePerm)
		if err != nil {
			log.Println(err)
		}
	}
	file, err := os.Create(asset)
	defer file.Close()

	if err != nil {
		return err
	}
	client := http.Client{
		CheckRedirect: func(r *http.Request, via []*http.Request) error {
			r.URL.Opaque = r.URL.Path
			return nil
		},
	}
	resp, err := client.Get(uri.String())
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	_, err = io.Copy(file, resp.Body)
	if err != nil {
		return err
	}
	return nil
}
