package trello

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"reflect"
	"strings"
)

const TrelloAPI = "https://api.trello.com/1"

type Client struct {
	client *http.Client
	key    string
	token  string
}

// Create a Trello API client.
func NewClient(key string, token string) *Client {
	return &Client{&http.Client{}, key, token}
}

// Build API url.
func (client *Client) ApiURL(resource string) string {
	params := url.Values{}
	params.Add("key", client.key)
	params.Add("token", client.token)
	return TrelloAPI + resource + "?" + params.Encode()
}

// Query the API.
func (client *Client) Query(method string, resource string, params url.Values, output interface{}) error {
	var requestURL = client.ApiURL(resource)
	var requestHeaders = map[string]string{}
	var req *http.Request
	var err error

	// Prepare the request.
	if method == "GET" || method == "DELETE" {
		if params != nil {
			requestURL += "&" + params.Encode()
		}
		req, err = http.NewRequest(method, requestURL, nil)
	}
	if method == "POST" || method == "PUT" {
		if params == nil {
			params = url.Values{}
		}
		requestHeaders["Content-Type"] = "application/x-www-form-urlencoded"
		req, err = http.NewRequest(method, requestURL, strings.NewReader(params.Encode()))
	}

	// Query the API.
	if err != nil {
		return err
	}
	for key, val := range requestHeaders {
		req.Header.Set(key, val)
	}
	req.Close = true
	resp, err := client.client.Do(req)
	if err != nil {
		return err
	}

	// Read the body.
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("Unexpected response when querying '%s' with method %s: %s", resource, method, string(body))
	}

	// Parse the json response.
	if output != nil && json.Unmarshal(body, output) != nil {
		return fmt.Errorf("Couldn't unmarshal body %s into %s", string(body), reflect.TypeOf(output))
	}
	return nil
}

// GET, PUT, POST, DELETE queries to the Trello API.
func (client *Client) Get(resource string, params url.Values, output interface{}) error {
	return client.Query("GET", resource, params, output)
}
func (client *Client) Put(resource string, params url.Values, output interface{}) error {
	return client.Query("PUT", resource, params, output)
}
func (client *Client) Post(resource string, params url.Values, output interface{}) error {
	return client.Query("POST", resource, params, output)
}
func (client *Client) Delete(resource string) error {
	return client.Query("DELETE", resource, nil, nil)
}

// Board.
func (client *Client) GetBoard(id string, params url.Values) (board Board, err error) {
	return board, client.Get("/boards/"+id, params, &board)
}
func (client *Client) PutBoard(id string, params url.Values) (board Board, err error) {
	return board, client.Put("/boards/"+id, params, &board)
}
func (client *Client) PostBoard(params url.Values) (board Board, err error) {
	return board, client.Post("/boards", params, &board)
}
func (client *Client) DeleteBoard(id string) error {
	return client.Delete("/boards/" + id)
}

// List.
func (client *Client) GetList(id string, params url.Values) (list List, err error) {
	return list, client.Get("/lists/"+id, params, &list)
}
func (client *Client) PutList(id string, params url.Values) (list List, err error) {
	return list, client.Put("/lists/"+id, params, &list)
}
func (client *Client) PostList(params url.Values) (list List, err error) {
	return list, client.Post("/lists", params, &list)
}
func (client *Client) DeleteList(id string) error {
	return client.Delete("/lists/" + id)
}

// Card.
func (client *Client) GetCard(id string, params url.Values) (card Card, err error) {
	return card, client.Get("/cards/"+id, params, &card)
}
func (client *Client) PutCard(id string, params url.Values) (card Card, err error) {
	return card, client.Put("/cards/"+id, params, &card)
}
func (client *Client) PostCard(params url.Values) (card Card, err error) {
	return card, client.Post("/cards", params, &card)
}
func (client *Client) DeleteCard(id string) error {
	return client.Delete("/cards/" + id)
}

// Webhook.
func (client *Client) GetWebhooks() (webhooks []Webhook, err error) {
	return webhooks, client.Get("/token/"+client.token+"/webhooks", params, &webhooks)
}
func (client *Client) GetWebhook(id string, params url.Values) (webhook Webhook, err error) {
	return webhook, client.Get("/webhooks/"+id, params, &webhook)
}
func (client *Client) PutWebhook(id string, params url.Values) (webhook Webhook, err error) {
	return webhook, client.Put("/webhooks/"+id, params, &webhook)
}
func (client *Client) PostWebhook(params url.Values) (webhook Webhook, err error) {
	return webhook, client.Post("/webhooks", params, &webhook)
}
func (client *Client) DeleteWebhook(id string) error {
	return client.Delete("/webhooks/" + id)
}
