package golded

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

var APIURL = "https://guilded.gg/api/v1"

const (
	POSTMESSAGEURL = "https://www.guilded.gg/api/v1/channels/%v/messages"
)

type API struct {
	Token string
}

func (a *API) req(method, url string, payload io.Reader) (*http.Request, error) {
	req, err := http.NewRequestWithContext(context.Background(), method, url, payload)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %v", a.Token))

	return req, nil
}

func (a *API) SendMessage(chanID string, msg PostMessage) (*MessageRes, error) {

	payload, err := json.Marshal(msg)

	if err != nil {
		return nil, err
	}

	req, err := a.req("POST", fmt.Sprintf(POSTMESSAGEURL, chanID), bytes.NewReader(payload))

	if err != nil {
		return nil, err
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var newMsg MessageRes

	err = json.Unmarshal(body, &newMsg)
	if err != nil {
		return nil, err
	}

	return &newMsg, nil
}
