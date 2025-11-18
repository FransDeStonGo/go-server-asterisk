package http

import (
	"net/http"
	"time"
)

// Client простой HTTP-клиент для ARI REST API
type Client struct {
    BaseURL    string
    HTTPClient *http.Client
}

// NewClient возвращает новый клиент
func NewClient(baseURL string) *Client {
    return &Client{
        BaseURL: baseURL,
        HTTPClient: &http.Client{
            Timeout: 10 * time.Second,
        },
    }
}

// DoGet выполняет GET-запрос к указанному пути (относительно BaseURL)
func (c *Client) DoGet(path string) (*http.Response, error) {
    return c.HTTPClient.Get(c.BaseURL + path)
}
