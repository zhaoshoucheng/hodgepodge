package net

import (
	"context"
	"errors"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
)

const (
	AppUserAgent = "test/0.1"
	CTJson       = "application/json"
	CTUrlencoded = "application/x-www-form-urlencoded"
)

func PostQueryWithHeaders(ctx context.Context, url, auth, body, ct string, headers map[string]string) (string, error) {
	if headers == nil {
		headers = make(map[string]string)
	}
	headers["Content-Type"] = ct
	resp, err := DoQuery(ctx, "POST", url, auth, body, headers)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	buff, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	if resp.StatusCode != 200 {
		return string(buff), errors.New("status code is not 200")
	}
	return string(buff), nil
}

func DoQuery(ctx context.Context, method, url, auth, body string, headers map[string]string) (*http.Response, error) {
	var reader io.Reader
	if body != "" {
		reader = strings.NewReader(body)
	}
	req, err := http.NewRequest(method, url, reader)
	if err != nil {
		return nil, err
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}
	if auth != "" {
		// detect kind
		req.Header.Set("Authorization", auth)
	}
	req.Header.Set("User-Agent", AppUserAgent)
	for h, v := range headers {
		req.Header.Set(h, v)
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
