package line

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

type Client struct {
	client *http.Client

	baseURL string
}

func New() *Client {
	return &Client{
		client:  http.DefaultClient,
		baseURL: "https://api.line.me",
	}
}

func (c *Client) request(method, route string, body io.Reader, headers map[string]string) ([]byte, error) {
	req, err := http.NewRequest(method, c.baseURL+route, body)
	if err != nil {
		return nil, err
	}

	for key, value := range headers {
		req.Header.Set(key, value)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	byteResp, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	switch resp.StatusCode {
	case http.StatusOK:
	case http.StatusCreated:
	default:
		return nil, fmt.Errorf("line clinet http error %d: %s", resp.StatusCode, byteResp)
	}

	return byteResp, nil
}

func (c *Client) GetToken(code, redirectURI, clientID, clientSecret string) (st *GetTokenResponse, err error) {
	data := url.Values{}
	data.Set("grant_type", "authorization_code")
	data.Set("code", code)
	data.Set("redirect_uri", redirectURI)
	data.Set("client_id", clientID)
	data.Set("client_secret", clientSecret)

	resp, err := c.request(
		http.MethodPost,
		"/oauth2/v2.1/token",
		strings.NewReader(data.Encode()),
		map[string]string{
			"Content-Type": "application/x-www-form-urlencoded",
		},
	)
	if err != nil {
		return nil, err
	}

	json.Unmarshal(resp, &st)

	return st, nil
}

func (c *Client) GetUser(accessToken string) (st *GetUserResponse, err error) {
	resp, err := c.request(
		http.MethodGet,
		"/v2/profile",
		nil,
		map[string]string{
			"Authorization": fmt.Sprintf("Bearer %s", accessToken),
		},
	)
	if err != nil {
		return nil, err
	}

	json.Unmarshal(resp, &st)

	return st, nil
}
