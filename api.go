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
	CHANNELSURL    = "channels/"
	CHANNELURL     = "channels/%v"
	CHANNELARCHIVE = "channels/%v/archive"
	CATEGORIESURL  = "servers/%v/categories"
	CATEGORYURL    = "servers/%v/categories/%v"
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

func (a *API) emptyput(url string) (*http.Response, error) {
	req, err := http.NewRequestWithContext(context.Background(), "PUT", url, nil)
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

func (a *API) DeleteGroup(serverId, groupId string) error {
	res, err := a.del(fmt.Sprintf(api(GROUPURL), serverId, groupId))
	if err != nil {
		return err
	}

	if res.StatusCode != http.StatusOK {
		return errors.New("could not delete the group")
	}

	return nil
}

// Channels

func (a *API) CreateChannel(newChannel PostChannel) (*GetChannelRes, error) {
	payload, err := json.Marshal(newChannel)

	if err != nil {
		return nil, err
	}

	body, err := a.req("POST", fmt.Sprintf(api(CHANNELSURL)), payload)

	if err != nil {
		return nil, err
	}

	var serverChannel GetChannelRes

	err = json.Unmarshal(*body, &body)
	if err != nil {
		return nil, err
	}

	return &serverChannel, nil
}

func (a *API) GetChannel(channelId string) (*GetChannelRes, error) {
	body, err := a.get(fmt.Sprintf(api(CHANNELURL), channelId))

	if err != nil {
		return nil, err
	}

	var channel GetChannelRes

	err = json.Unmarshal(*body, &channel)
	if err != nil {
		return nil, err
	}
	return &channel, nil
}

func (a *API) UpdateChannel(channelId string, updatedChannel PatchChannel) (*GetMessageRes, error) {
	payload, err := json.Marshal(updatedChannel)

	if err != nil {
		return nil, err
	}

	body, err := a.req("PATCH", fmt.Sprintf(api(CHANNELURL), channelId), payload)

	if err != nil {
		return nil, err
	}

	var channel GetMessageRes

	err = json.Unmarshal(*body, &channel)
	if err != nil {
		return nil, err
	}

	return &channel, nil
}

func (a *API) DeleteChannel(channelId string) error {
	res, err := a.del(fmt.Sprintf(api(CHANNELURL), channelId))
	if err != nil {
		return err
	}

	if res.StatusCode != http.StatusOK {
		return errors.New("could not delete channel")
	}

	return nil
}

func (a *API) ArchiveChannel(channelId string) error {
	res, err := a.emptyput(fmt.Sprintf(api(CHANNELARCHIVE), channelId))
	if err != nil {
		return err
	}

	if res.StatusCode != http.StatusOK {
		return errors.New("could not archive channel")
	}

	return nil
}

func (a *API) UnarchiveChannel(channelId string) error {
	res, err := a.del(fmt.Sprintf(api(CHANNELARCHIVE), channelId))
	if err != nil {
		return err
	}

	if res.StatusCode != http.StatusOK {
		return errors.New("could not unarchive channel")
	}

	return nil
}

// Categories

func (a *API) CreateCategory(serverId string, newCategory PostCategory) (*GetCategoryRes, error) {
	payload, err := json.Marshal(newCategory)

  if err != nil {
    return nil, err
  }

  body, err := a.req("POST", fmt.Sprintf(api(CATEGORIESURL), serverId), payload)
  if err != nil {
    return nil, err
  }

  var category GetCategoryRes

  err = json.Unmarshal(*body, &category)

  if err != nil {
    return nil, err
  }

  return &category, nil
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
