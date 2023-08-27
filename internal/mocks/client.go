package mocks

import "net/http"

type HTTPClient interface {
	Get(url string) (*http.Response, error)
	// Add more methods as needed
}

type MockHTTPClient struct {
	Response *http.Response
	Err      error
}

func (m *MockHTTPClient) Get(url string) (*http.Response, error) {
	return m.Response, m.Err
}
