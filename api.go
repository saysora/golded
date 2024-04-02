package golded

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

var APIURL = "https://guilded.gg/api/v1/"

const (
	SERVERURL      = "servers/%v"
	GROUPSURL      = "servers/%v/groups"
	GROUPURL       = "servers/%v/groups/%v"
	POSTMESSAGEURL = "channels/%v/messages"
)

type API struct {
	Token string
}

func api(url string) string {
	return APIURL + url
}

func (a *API) oldget(url string) (*http.Request, error) {
	req, err := http.NewRequestWithContext(context.Background(), "GET", url, nil)

	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %v", a.Token))

	return req, nil
}

func (a *API) oldreq(method, url string, payload io.Reader) (*http.Request, error) {
	req, err := http.NewRequestWithContext(context.Background(), method, url, payload)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %v", a.Token))

	return req, nil
}

func (a *API) get(url string) (*[]byte, error) {
	req, err := http.NewRequestWithContext(context.Background(), "GET", url, nil)

	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %v", a.Token))

  res, err := http.DefaultClient.Do(req)
  if err != nil {
    return nil, err
  }

  defer res.Body.Close()

  body, err := io.ReadAll(res.Body)

  if err != nil {
    return nil, err
  }

	return &body, nil
}

func (a *API) del(url string) (*http.Response, error) {
	req, err := http.NewRequestWithContext(context.Background(), "DELETE", url, nil)

	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %v", a.Token))

  res, err := http.DefaultClient.Do(req)
  if err != nil {
    return nil, err
  }

  defer res.Body.Close()

  return res, nil
}

func (a *API) req(method, url string, payload []byte) (*[]byte, error) {
	req, err := http.NewRequestWithContext(context.Background(), method, url, bytes.NewReader(payload))
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %v", a.Token))

  res, err := http.DefaultClient.Do(req)

  if err != nil {
    return nil, err
  }

  defer res.Body.Close()

  body, err := io.ReadAll(res.Body)
  if err != nil {
    return nil, err
  }

  return &body, nil
}

// Server requests

func (a *API) GetServer(serverId string) (*ServerGetRes, error) {
	body, err := a.get(fmt.Sprintf(api(SERVERURL), serverId))
	if err != nil {
		return nil, err
	}

	var server ServerGetRes

	err = json.Unmarshal(*body, &server)

	if err != nil {
		return nil, err
	}

	return &server, nil
}

func (a *API) CreateGroup(serverId string, newGroup PostGroup) (*GetGroupRes, error) {
	payload, err := json.Marshal(newGroup)

	if err != nil {
		return nil, err
	}

	body, err := a.req("POST", fmt.Sprintf(api(GROUPSURL), serverId), payload)
  if err != nil {
    return nil, err
  }

  var group GetGroupRes

  err = json.Unmarshal(*body, &group)

  if err != nil {
    return nil, err
  }

  return &group, nil
}

func (a *API) GetGroups(serverId string) (*GetGroupsRes, error) {
  body, err := a.get(fmt.Sprintf(api(GROUPSURL), serverId))
  if err != nil {
    return nil, err
  }

  var groups GetGroupsRes

  err = json.Unmarshal(*body, &groups)
  if err != nil {
    return nil, err
  }

  return &groups, nil
}

func (a *API) GetGroup(serverId, groupId string) (*GetGroupRes, error) {
  body, err := a.get(fmt.Sprintf(api(GROUPURL), serverId, groupId))

  if err != nil {
    return nil, err
  }

  var group GetGroupRes

  err = json.Unmarshal(*body, &group)
  if err != nil {
    return nil, err
  }
  return &group, nil
}

func (a *API) UpdateGroup(serverId, groupId string, updatedGroup PatchGroup) (*GetGroupRes, error) {
  payload, err := json.Marshal(updatedGroup)

  if err != nil {
    return nil, err
  }

  body, err := a.req("PATCH", fmt.Sprintf(GROUPURL, serverId, groupId), payload)

  if err != nil {
    return nil, err
  }

  var group GetGroupRes

  err = json.Unmarshal(*body, &group)
  if err != nil {
    return nil, err
  }

  return &group, nil
}

func (a *API) DeleteGroup(serverId, groupId string) (error) {
  res, err := a.del(fmt.Sprintf(api(GROUPURL), serverId, groupId))
  if err != nil {
    return err
  }

  if res.StatusCode != http.StatusOK {
    return errors.New("could not delete the group")
  }

  return nil
}

// So much more :sweat_emote:
func (a *API) SendMessage(chanID string, msg PostMessage) (*GetMessageRes, error) {

	payload, err := json.Marshal(msg)

	if err != nil {
		return nil, err
	}

	body, err := a.req("POST", fmt.Sprintf(api(POSTMESSAGEURL), chanID), payload)

	if err != nil {
		return nil, err
	}

	var newMsg GetMessageRes

	err = json.Unmarshal(*body, &newMsg)
	if err != nil {
		return nil, err
	}

	return &newMsg, nil
}
