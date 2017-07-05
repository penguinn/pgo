package jsonrpc2

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"
)

type Client struct {
	Url string
	id  int
	c   *http.Client
}

type Resp struct {
	Id     int            `json:"id"`
	Result interface{}    `json:"result"`
	Error  *RespError `json:"error"`
}

type RespError struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func (this *RespError) Error() string {
	b, err := json.Marshal(this)
	if err != nil {
		return err.Error()
	}
	return string(b)
}

func NewClient(url string) *Client {
	client := new(Client)
	client.Url = url
	client.c = &http.Client{}
	return client
}

func (c *Client) Call(method string, params interface{}, id int) (*Resp, error) {
	return c.CallTimeout(method, params, id, 0)
}

func (c *Client) CallTimeout(method string, params interface{}, id int, timeout time.Duration) (*Resp, error) {
	var payload = map[string]interface{}{
		"jsonrpc": "2.0",
		"method":  method,
		"params":  params,
		"id":      id,
	}

	c.id += 1
	data, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}
	buf := bytes.NewBuffer(data)
	c.c.Timeout = timeout
	resp, err := c.c.Post(c.Url, "application/json", buf)

	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	decoder := json.NewDecoder(resp.Body)
	var respPayload *Resp
	err = decoder.Decode(&respPayload)

	if err != nil {
		return nil, errors.New(fmt.Sprintf("Invalid Resp from server. Status: %s. %s", resp.Status, err.Error()))
	}

	return respPayload, nil
}