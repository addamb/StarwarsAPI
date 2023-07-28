package service

import (
	"io"
	"net/http"
)

type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

type SwapiRequest interface {
	SendSwapiRequest(reqURL string) (io.ReadCloser, error)
}

type SwapiClient struct {
	Client HTTPClient
}

func NewSwapiClient(client HTTPClient) SwapiClient {
	return SwapiClient{
		Client: client,
	}
}

func (s *SwapiClient) SendSwapiRequest(reqURL string) (io.ReadCloser, error) {
	request, _ := http.NewRequest(http.MethodGet, reqURL, nil)

	resp, err := s.Client.Do(request)
	if err != nil {
		return nil, err
	}

	return resp.Body, nil
}
