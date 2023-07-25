package bancocentral

import (
	"compress/gzip"
	"fmt"
	"io/ioutil"
	"net/http"
)

type AcessarBancoCentral struct {
	URL string
}

func NewAcessarBancoCentral(url string) *AcessarBancoCentral {
	return &AcessarBancoCentral{
		URL: url,
	}
}

func (ac *AcessarBancoCentral) GetURL() ([]byte, error) {
	headers := map[string]string{
		"Host":            "conteudo.bcb.gov.br",
		"Connection":      "keep-alive",
		"Cache-Control":   "max-age=0",
		"User-Agent":      "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.100 Safari/537.36",
		"DNT":             "1",
		"Content-Type":    "application/atom+xml",
		"Accept":          "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8",
		"Accept-Encoding": "gzip, deflate, br",
		"Accept-Language": "pt-BR,pt;q=0.9,en-US;q=0.8,en;q=0.7,mt;q=0.6",
	}

	client := &http.Client{}
	req, err := http.NewRequest("GET", ac.URL, nil)
	if err != nil {
		return nil, err
	}

	for key, value := range headers {
		req.Header.Set(key, value)
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("request failed with status code: %d", resp.StatusCode)
	}

	if resp.Header.Get("Content-Encoding") == "gzip" {
		reader, err := gzip.NewReader(resp.Body)
		if err != nil {
			return nil, err
		}
		defer reader.Close()
		respBody, err := ioutil.ReadAll(reader)
		if err != nil {
			return nil, err
		}
		return respBody, nil
	}

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return respBody, nil
}
