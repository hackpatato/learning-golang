package main

import (
	"io"
	"net/http"
	

)

func ContentScarper() {
	if resp, err := http.Get("https://test.test.com/test/Web_scraping"); err == nil {
	defer resp.Body.Close()
	if body, err := io.ReadAll(resp.Body); err == nil {
		println(string(body))
		}
	}
}