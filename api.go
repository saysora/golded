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

var APIURL = "https://www.guilded.gg/api/v1/"

const (
	SERVERURL               = "servers/%v"
	GROUPSURL               = "servers/%v/groups"
	GROUPURL                = "servers/%v/groups/%v"
	CHANNELSURL             = "channels/"
	CHANNELURL              = "channels/%v"
	CHANNELARCHIVE          = "channels/%v/archive"
	CATEGORIESURL           = "servers/%v/categories"
	CATEGORYURL             = "servers/%v/categories/%v"
	MESSAGESURL             = "channels/%v/messages"
	MESSAGEURL              = "channels/%v/messages/%v"
	MESSAGEPINURL           = "channels/%v/messages/%v/pin"
	MEMBERSNICKNAMEURL      = "servers/%v/members/%/nickname"
	MEMBERSURL              = "servers/%v/members"
	MEMBERURL               = "servers/%v/members/%v"
	GROUPMEMBERURL          = "groups/%v/members/%v"
	ROLEMEMBERURL           = "servers/%v/members/%v/roles/%v"
	MEMBERBANSURL           = "servers/%v/bans"
	MEMBERBANURL            = "servers/%v/bans/%v"
	ANNOUNCEMENTSURL        = "channels/%v/announcements"
	ANNOUNCEMENTURL         = "channels/%v/announcements/%v"
	ANNOUNCEMENTCOMMENTSURL = "channels/%v/announcements/%/comments"
	ANNOUNCEMENTCOMMENTURL  = "channels/%v/announcements/%/comments/%v"
	CALENDAREVENTSURL       = "channels/%v/events"
	CALENDAREVENTURL        = "channels/%v/events/%v"
)

type API struct {
	Token string
}

func api(url string) string {
	return APIURL + url
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

func (a *API) emptyUpdate(method, url string) (*http.Response, error) {
	req, err := http.NewRequestWithContext(context.Background(), method, url, nil)
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

	body, err := a.req("POST", api(CHANNELSURL), payload)

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
	res, err := a.emptyUpdate("PUT", fmt.Sprintf(api(CHANNELARCHIVE), channelId))
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

func (a *API) GetCategory(serverId, categoryId string) (*GetCategoryRes, error) {
	body, err := a.get(fmt.Sprintf(api(CATEGORYURL), serverId, categoryId))

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

func (a *API) UpdateCategory(serverId, categoryId string, updateCategory PatchCategory) (*GetCategoryRes, error) {
	payload, err := json.Marshal(updateCategory)
	if err != nil {
		return nil, err
	}

	body, err := a.req("PATCH", fmt.Sprintf(api(CATEGORYURL), serverId, categoryId), payload)
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

func (a *API) DeleteCategory(serverId, categoryId string) error {
	res, err := a.del(fmt.Sprintf(api(CATEGORYURL), serverId, categoryId))

	if err != nil {
		return err
	}

	if res.StatusCode != http.StatusOK {
		return errors.New("error deleting category")
	}

	return nil
}

// Messages

func (a *API) SendMessage(chanId string, msg PostMessage) (*GetMessageRes, error) {

	payload, err := json.Marshal(msg)

	if err != nil {
		return nil, err
	}

	body, err := a.req("POST", fmt.Sprintf(api(MESSAGESURL), chanId), payload)

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

func (a *API) GetMessages(chanId string) (*GetMessagesRes, error) {
	body, err := a.get(fmt.Sprintf(api(MESSAGESURL), chanId))

	if err != nil {
		return nil, err
	}

	var messages GetMessagesRes

	err = json.Unmarshal(*body, &messages)
	if err != nil {
		return nil, err
	}

	return &messages, err
}

func (a *API) GetMessage(chanId, msgId string) (*GetMessageRes, error) {
	body, err := a.get(fmt.Sprintf(api(MESSAGEURL), chanId, msgId))

	if err != nil {
		return nil, err
	}

	var message GetMessageRes

	err = json.Unmarshal(*body, &message)
	if err != nil {
		return nil, err
	}

	return &message, nil
}

func (a *API) UpdateMessage(chanId, msgId string, updateMsg PatchMessage) (*GetMessageRes, error) {
	payload, err := json.Marshal(updateMsg)
	if err != nil {
		return nil, err
	}

	body, err := a.req("PUT", fmt.Sprintf(api(MESSAGEURL), chanId, msgId), payload)
	if err != nil {
		return nil, err
	}

	var message GetMessageRes

	err = json.Unmarshal(*body, &message)
	if err != nil {
		return nil, err
	}

	return &message, nil
}

func (a *API) DeleteMessage(chanId, msgId string) error {
	res, err := a.del(fmt.Sprintf(MESSAGEURL, chanId, msgId))
	if err != nil {
		return err
	}

	if res.StatusCode != http.StatusOK {
		return errors.New("error deleting message")
	}

	return nil
}

func (a *API) PinMessage(chanId, msgId string) error {
	res, err := a.emptyUpdate("POST", fmt.Sprintf(api(MESSAGEPINURL), chanId, msgId))
	if err != nil {
		return err
	}

	if res.StatusCode != http.StatusOK {
		return errors.New("could not pin message")
	}

	return nil
}

func (a *API) UnpinMessage(chanId, msgId string) error {
	res, err := a.del(fmt.Sprintf(api(MESSAGEPINURL), chanId, msgId))
	if err != nil {
		return err
	}

	if res.StatusCode != http.StatusOK {
		return errors.New("could not unpin message")
	}

	return nil
}

// Members

func (a *API) UpdateMemberNickname(serverId, userId string, nickname PutMemberNickname) (*PutMemberNickname, error) {
	payload, err := json.Marshal(nickname)
	if err != nil {
		return nil, err
	}

	body, err := a.req("PUT", fmt.Sprintf(api(MEMBERSNICKNAMEURL), serverId, userId), payload)
	if err != nil {
		return nil, err
	}

	var nicknameRes PutMemberNickname

	err = json.Unmarshal(*body, &nicknameRes)
	if err != nil {
		return nil, err
	}

	return &nicknameRes, nil
}

func (a *API) DeleteMemberNickname(serverId, userId string) error {
	res, err := a.del(fmt.Sprintf(api(MEMBERSNICKNAMEURL), serverId, userId))
	if err != nil {
		return err
	}

	if res.StatusCode != http.StatusOK {
		return errors.New("could not reset nickname")
	}

	return nil
}

func (a *API) GetMembers(serverId string) (*GetMembersRes, error) {
	body, err := a.get(fmt.Sprintf(api(MEMBERSURL), serverId))
	if err != nil {
		return nil, err
	}

	var members GetMembersRes

	err = json.Unmarshal(*body, &members)
	if err != nil {
		return nil, err
	}

	return &members, nil
}

func (a *API) GetMember(serverId, userId string) (*GetMemberRes, error) {
	body, err := a.get(fmt.Sprintf(api(MEMBERURL), serverId, userId))
	if err != nil {
		return nil, err
	}

	var member GetMemberRes
	err = json.Unmarshal(*body, &member)
	if err != nil {
		return nil, err
	}

	return &member, nil
}

func (a *API) KickMember(serverId, userId string) error {
	res, err := a.del(fmt.Sprintf(api(MEMBERURL), serverId, userId))
	if err != nil {
		return err
	}

	if res.StatusCode != http.StatusOK {
		return errors.New("could not kick member")
	}

	return nil
}

// Groups
func (a *API) AddMemberToGroup(groupId, userId string) error {
	res, err := a.emptyUpdate("PUT", fmt.Sprintf(api(GROUPMEMBERURL), groupId, userId))

	if err != nil {
		return err
	}

	if res.StatusCode != http.StatusOK {
		return errors.New("could not add member to group")
	}

	return nil
}

func (a *API) RemoveMemberFromGroup(groupId, userId string) error {
	res, err := a.del(fmt.Sprintf(api(GROUPMEMBERURL), groupId, userId))
	if err != nil {
		return err
	}

	if res.StatusCode != http.StatusOK {
		return errors.New("could not remove member from group")
	}

	return nil
}

// Roles
func (a *API) AddRoleToMember(serverId, userId, roleId string) error {
	res, err := a.emptyUpdate("PUT", fmt.Sprintf(api(ROLEMEMBERURL), serverId, userId, roleId))
	if err != nil {
		return err
	}

	if res.StatusCode != http.StatusOK {
		return errors.New("could not add role to member")
	}

	return nil
}

func (a *API) RemoveRoleFromMember(serverId, userId, roleId string) error {
	res, err := a.del(fmt.Sprintf(api(ROLEMEMBERURL), serverId, userId, roleId))
	if err != nil {
		return err
	}

	if res.StatusCode != http.StatusOK {
		return errors.New("could not remove role from member")
	}

	return nil
}

// Ban Members

func (a *API) BanMember(serverId, userId, reason *string) (*MemberBanRes, error) {
	payload, err := json.Marshal(PostMemberBan{Reason: reason})
	if err != nil {
		return nil, err
	}

	body, err := a.req("POST", fmt.Sprintf(api(MEMBERBANURL), serverId, userId), payload)
	if err != nil {
		return nil, err
	}

	var memberBan MemberBanRes

	err = json.Unmarshal(*body, &memberBan)
	if err != nil {
		return nil, err
	}

	return &memberBan, nil
}

func (a *API) GetMemberBan(serverId, userId string) (*MemberBanRes, error) {
	body, err := a.get(fmt.Sprintf(api(MEMBERBANURL), serverId, userId))
	if err != nil {
		return nil, err
	}

	var bannedMember MemberBanRes

	err = json.Unmarshal(*body, &bannedMember)
	if err != nil {
		return nil, err
	}

	return &bannedMember, nil
}

func (a *API) GetMemberBans(serverId string) (*MemberBansRes, error) {
	body, err := a.get(fmt.Sprintf(api(MEMBERBANSURL), serverId))
	if err != nil {
		return nil, err
	}

	var memberBans MemberBansRes

	err = json.Unmarshal(*body, &memberBans)

	if err != nil {
		return nil, err
	}

	return &memberBans, nil
}

func (a *API) RemoveMemberBan(serverId, userId string) error {
	res, err := a.del(fmt.Sprintf(api(MEMBERBANURL), serverId, userId))
	if err != nil {
		return err
	}

	if res.StatusCode != http.StatusOK {
		return errors.New("could not remove member ban")
	}

	return nil
}

// Announcements

func (a *API) CreateAnnouncement(channelId string, newAnnouncement PostAnnouncement) (*GetAnnouncementRes, error) {
	payload, err := json.Marshal(newAnnouncement)

	if err != nil {
		return nil, err
	}

	body, err := a.req("POST", fmt.Sprintf(api(ANNOUNCEMENTSURL), channelId), payload)
	if err != nil {
		return nil, err
	}

	var announcement GetAnnouncementRes

	err = json.Unmarshal(*body, &announcement)
	if err != nil {
		return nil, err
	}

	return &announcement, nil

}

func (a *API) GetAnnouncements(channelId string) (*GetAnnouncementsRes, error) {
	body, err := a.get(fmt.Sprintf(api(ANNOUNCEMENTSURL), channelId))
	if err != nil {
		return nil, err
	}

	var announcements GetAnnouncementsRes

	err = json.Unmarshal(*body, &announcements)

	if err != nil {
		return nil, err
	}

	return &announcements, nil
}

func (a *API) GetAnnouncement(channelId, announcementId string) (*GetAnnouncementRes, error) {
	body, err := a.get(fmt.Sprintf(api(ANNOUNCEMENTURL), channelId, announcementId))
	if err != nil {
		return nil, err
	}

	var announcement GetAnnouncementRes

	err = json.Unmarshal(*body, &announcement)

	if err != nil {
		return nil, err
	}

	return &announcement, nil
}

func (a *API) UpdateAnnouncement(channelId, announcementId string, updateAnnouncement PatchAnnouncement) (*GetAnnouncementRes, error) {
	payload, err := json.Marshal(updateAnnouncement)
	if err != nil {
		return nil, err
	}

	body, err := a.req("PATCH", fmt.Sprintf(api(ANNOUNCEMENTURL), channelId, announcementId), payload)
	if err != nil {
		return nil, err
	}

	var announcement GetAnnouncementRes

	err = json.Unmarshal(*body, &announcement)
	if err != nil {
		return nil, err
	}

	return &announcement, nil
}

func (a *API) DeleteAnnouncement(channelId string) error {
	res, err := a.del(fmt.Sprintf(api(ANNOUNCEMENTURL), channelId))
	if err != nil {
		return err
	}

	if res.StatusCode != http.StatusOK {
		return errors.New("could not delete announcement")
	}

	return nil
}

func (a *API) CreateAnnouncementComment(channelId, announcementId string, newComment PostAnnouncementComment) (*GetAnnouncementComment, error) {
	payload, err := json.Marshal(newComment)
	if err != nil {
		return nil, err
	}

	body, err := a.req("POST", fmt.Sprintf(api(ANNOUNCEMENTCOMMENTSURL), channelId, announcementId), payload)

	if err != nil {
		return nil, err
	}

	var announcementComment GetAnnouncementComment

	err = json.Unmarshal(*body, &announcementComment)
	if err != nil {
		return nil, err
	}

	return &announcementComment, nil
}

func (a *API) GetAnnouncementComments(channelId, announcementId string) (*GetAnnouncementComments, error) {
	body, err := a.get(fmt.Sprintf(api(ANNOUNCEMENTCOMMENTSURL), channelId, announcementId))
	if err != nil {
		return nil, err
	}

	var comment GetAnnouncementComments

	err = json.Unmarshal(*body, &comment)
	if err != nil {
		return nil, err
	}

	return &comment, nil
}

func (a *API) GetAnnouncementComment(channelId, announcementId, commentId string) (*GetAnnouncementComments, error) {
	body, err := a.get(fmt.Sprintf(api(ANNOUNCEMENTCOMMENTURL), channelId, announcementId, commentId))
	if err != nil {
		return nil, err
	}

	var comments GetAnnouncementComments

	err = json.Unmarshal(*body, &comments)
	if err != nil {
		return nil, err
	}

	return &comments, nil
}

func (a *API) UpdateAnnouncementComment(channelId, announcementId, commentId string, newComment PostAnnouncementComment) (*GetAnnouncementComment, error) {
	payload, err := json.Marshal(newComment)
	if err != nil {
		return nil, err
	}

	body, err := a.req("PATCH", fmt.Sprintf(api(ANNOUNCEMENTCOMMENTURL), channelId, announcementId, commentId), payload)
	if err != nil {
		return nil, err
	}

	var comment GetAnnouncementComment

	err = json.Unmarshal(*body, &comment)

	if err != nil {
		return nil, err
	}

	return &comment, err
}

func (a *API) DeleteAnnouncementComment(channelId, announcementId, commentId string) error {
	res, err := a.del(fmt.Sprintf(api(ANNOUNCEMENTCOMMENTURL), channelId, announcementId, commentId))
	if err != nil {
		return err
	}

	if res.StatusCode != http.StatusOK {
		return errors.New("could not delete comment")
	}

	return nil
}

// Events
func (a *API) CreateEvent(channelId string, newEvent PostCalendarEvent) (*GetCalendarEventRes, error) {
	payload, err := json.Marshal(newEvent)
	if err != nil {
		return nil, err
	}

	body, err := a.req("POST", fmt.Sprintf(api(CALENDAREVENTSURL), channelId), payload)

  if err != nil {
    return nil, err
  }

  var event GetCalendarEventRes

  err = json.Unmarshal(*body, &event)
  if err != nil {
    return nil, err
  }

  return &event, nil
}

func (a *API) GetEvents(channelId string) (*GetCalendarEventsRes, error) {
  body, err := a.get(fmt.Sprintf(api(CALENDAREVENTSURL), channelId))
  if err != nil {
    return nil, err
  }

  var events GetCalendarEventsRes

  err = json.Unmarshal(*body, &events)
  if err != nil {
    return nil, err
  }

  return &events, nil
}

func (a *API) GetEvent(channelId, eventId string) (*GetCalendarEventRes, error) {
  body, err := a.get(fmt.Sprintf(api(CALENDAREVENTURL), channelId, eventId))
  if err != nil {
    return nil, err
  }

  var event GetCalendarEventRes

  err = json.Unmarshal(*body, &event)
  if err != nil {
    return nil, err
  }

  return &event, nil
}

// So much more :sweat_emote:
