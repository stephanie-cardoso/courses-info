package client

import (
	"io"
	"net/http"
)

func DoRequest(url string) ([]byte, error) {
	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("accept", "application/json")
	req.Header.Add("apiKey", "7JbSA3ep6XOMV3t8t7QXuXq9HS79Dwnr")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}
