package client

import (
	"net/http"
	"time"

	"github.com/photoprism/photoprism-places/internal/event"
)

var Retries = 2
var Timeout = time.Second * 15
var Agent = "photoprism.org"

var log = event.Log

func Request(url string) (res *http.Response, err error) {
	c := http.Client{
		Timeout: Timeout,
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)

	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("User-Agent", Agent)

	for i := 0; i < Retries; i++ {
		res, err = c.Do(req)

		if err == nil {
			break
		}
	}

	return res, err
}
