package asterisk

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"net/url"

	"github.com/gorilla/websocket"
)

// Нужные типы определяются здесь (замена удалённого ari-ws.go)
type ARIConfig struct {
    URL      string
    Username string
    Password string
    AppName  string
}

type SimpleARIClient struct {
    baseURL    string
    username   string
    password   string
    appName    string
    httpClient *http.Client
    wsConn     *websocket.Conn
}

// ConnectAndListen создает клиент и подключается к WebSocket, возвращает клиент.
func ConnectAndListen(cfg ARIConfig) (*SimpleARIClient, error) {
    client, err := NewSimpleARIClient(cfg)
    if err != nil {
        return nil, err
    }

    if err := client.ConnectWebSocket(); err != nil {
        return nil, err
    }

    return client, nil
}

// NewSimpleARIClient создает базовый клиент (без healthcheck)
func NewSimpleARIClient(cfg ARIConfig) (*SimpleARIClient, error) {
    return &SimpleARIClient{
        baseURL: cfg.URL,
        username: cfg.Username,
        password: cfg.Password,
        appName: cfg.AppName,
        httpClient: &http.Client{},
    }, nil
    
}

// ConnectWebSocket подключается к WebSocket событий ARI
func (c *SimpleARIClient) ConnectWebSocket() error {
    wsURL := c.baseURL + "/events?app=" + url.QueryEscape(c.appName)
    if len(wsURL) >= 4 && wsURL[:4] == "http" {
        wsURL = "ws" + wsURL[4:]
    }

    header := http.Header{}
    header.Add("Authorization", "Basic "+basicAuth(c.username, c.password))

    conn, resp, err := websocket.DefaultDialer.Dial(wsURL, header)
    if err != nil {
        if resp != nil {
            return fmt.Errorf("ws dial HTTP %d", resp.StatusCode)
        }
        return err
    }
    c.wsConn = conn
    return nil
}

// ReadEvents читает события из WS и возвращает каналы событий и ошибок
func (c *SimpleARIClient) ReadEvents() (<-chan map[string]interface{}, <-chan error) {
    events := make(chan map[string]interface{}, 100)
    errors := make(chan error, 1)

    go func() {
        defer close(events)
        defer close(errors)
        for {
            var evt map[string]interface{}
            if err := c.wsConn.ReadJSON(&evt); err != nil {
                errors <- err
                return
            }
            events <- evt
        }
    }()

    return events, errors
}

// Close закрывает WS соединение
func (c *SimpleARIClient) Close() {
    if c.wsConn != nil {
        c.wsConn.Close()
    }
}

// base64 helper — простой вызов из стандартной библиотеки
func basicAuth(username, password string) string {
    auth := username + ":" + password
    return base64.StdEncoding.EncodeToString([]byte(auth))
}

