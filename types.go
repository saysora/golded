package golded

import (
	"encoding/json"
	"fmt"
)

type SocketEvent struct {
	Op int             `json:"op"`
	T  string          `json:"t"`
	S  string          `json:"s"`
	D  json.RawMessage `json:"d"`
}

func (se *SocketEvent) String() string {
	return fmt.Sprintf("%d, %s, %s, %s", se.Op, string(se.T), string(se.S), se.D)
}

type Announcement struct {
	Id        string   `json:"id"`
	ServerId  string   `json:"serverId"`
	GroupId   string   `json:"groupId"`
	ChannelId string   `json:"channelId"`
	CreatedAt string   `json:"createdAt"`
	CreatedBy string   `json:"createdBy"`
	Content   string   `json:"content"`
	Mentions  Mentions `json:"mentions"`
	Title     string   `json:"title"`
}

type PostAnnouncement struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

// TODO: Finish data sets
type Category struct {
	Id        int     `json:"id"`
	ServerId  string  `json:"serverId"`
	GroupId   string  `json:"groupId"`
	CreatedAt string  `json:"createdAt"`
	UpdatedAt *string `json:"updatedAt"`
	Name      string  `json:"name"`
	Priority  *int    `json:"priority"`
}

type PostCategory struct {
	Name    string  `json:"name"`
	GroupId *string `json:"groupId"`
}

type PatchCategory struct {
	Name     *string `json:"name"`
	Priority *int    `json:"priority"`
}

// TODO: Finish data sets
type Channel struct {
	Id         string  `json:"id"`
	Type       string  `json:"type"`
	Name       string  `json:"name"`
	Topic      *string `json:"topic"`
	CreatedAt  string  `json:"createdAt"`
	CreatedBy  string  `json:"createdBy"`
	UpdatedAt  *string `json:"updatedAt"`
	ServerId   string  `json:"serverId"`
	RootId     *string `json:"rootId"`
	ParentId   *string `json:"parentId"`
	MessageId  *string `json:"messageId"`
	CategoryId *int    `json:"categoryId"`
	GroupId    string  `json:"groupId"`
	Visibility *string `json:"visibility"`
	ArchivedBy *string `json:"archivedBy"`
	ArchivedAt *string `json:"archivedAt"`
}

type PostChannel struct {
	Name       string  `json:"name"`
	Topic      *string `json:"topic"`
	Visibility *string `json:"visibility"`
	Type       string  `json:"type"`
	ServerId   *string `json:"serverId"`
	GroupId    *string `json:"groupId"`
	CategoryId *string `json:"categoryId"`
	ParentId   *string `json:"parentId"`
	MessageId  *string `json:"messageId"`
}

type PatchChannel struct {
	Name       *string `json:"name"`
	Topic      *string `json:"topic"`
	Visibility *string `json:"visibility"`
}

type ChatEmbedFooter struct {
	IconUrl *string `json:"icon_url"`
	Text    string  `json:"text"`
}

type ChatEmbedMedia struct {
	Url string `json:"url"`
}

type ChatEmbedAuthor struct {
	Name    *string `json:"name"`
	Url     *string `json:"url"`
	IconUrl *string `json:"icon_url"`
}

type ChatEmbedField struct {
	Name   string `json:"name"`
	Value  string `json:"value"`
	Inline *bool  `json:"inline"`
}

type ChatEmbed struct {
	Title       string            `json:"title,omitempty"`
	Description string            `json:"description,omitempty"`
	Url         string            `json:"url,omitempty"`
	Color       int               `json:"color,omitempty"`
	Footer      *ChatEmbedFooter  `json:"footer,omitempty"`
	Timestamp   string            `json:"timestamp,omitempty"`
	Thumbnail   *ChatEmbedMedia   `json:"thumbnail,omitempty"`
	Image       *ChatEmbedMedia   `json:"image,omitempty"`
	Author      *ChatEmbedAuthor  `json:"author,omitempty"`
	Fields      *[]ChatEmbedField `json:"fields,omitempty"`
}

type Emote struct {
	Id       int     `json:"id"`
	Name     string  `json:"name"`
	Url      string  `json:"url"`
	ServerId *string `json:"serverId"`
}

// TODO: Finish data sets
type Group struct {
	Id          string  `json:"id"`
	ServerId    string  `json:"serverId"`
	Name        string  `json:"name"`
	Description *string `json:"description"`
	Avatar      *string `json:"avatar"`
	IsHome      *bool   `json:"isHome"`
	EmoteId     *int    `json:"emoteId"`
	IsPublic    *bool   `json:"isPublic"`
	CreatedAt   string  `json:"createdAt"`
	CreatedBy   string  `json:"createdBy"`
	UpdatedAt   *string `json:"updatedAt"`
	UpdatedBy   *string `json:"updatedBy"`
	ArchivedAt  *string `json:"archivedAt"`
	ArchivedBy  *string `json:"archivedBy"`
}

type PostGroup struct {
	Name        string  `json:"name"`
	Description *string `json:"description"`
	EmoteId     *int    `json:"emoteId"`
	IsPublic    *bool   `json:"isPublic"`
}

// TODO: Finish data sets
type Member struct {
	User     User    `json:"user"`
	RoleIds  *[]int  `json:"roleIds"`
	Nickname *string `json:"nickname"`
	JoinedAt string  `json:"joinedAt"`
	IsOwner  *bool   `json:"isOwner"`
}

type MemberSum struct {
	User    UserSum `json:"user"`
	RoleIds []int   `json:"roleIds"`
}

type MentionUser struct {
	Id string `json:"id"`
}

type MentionChannel = MentionUser

type MentionRole struct {
	Id int `json:"id"`
}

type Mentions struct {
	Users    *[]MentionUser    `json:"users"`
	Channels *[]MentionChannel `json:"channels"`
	Roles    *[]MentionRole    `json:"roles"`
	Everyone *bool             `json:"everyone"`
	Here     *bool             `json:"here"`
}

type Message struct {
	Id                    string      `json:"id"`
	Type                  string      `json:"type"`
	ServerId              string      `json:"serverId,omitempty"`
	GroupId               string      `json:"groupId,omitempty"`
	ChannelId             string      `json:"channelId"`
	Content               string      `json:"content,omitempty"`
	HiddenLinkPreviewUrls []string    `json:"hiddenLinkPreviewUrls,omitempty"`
	Embeds                []ChatEmbed `json:"embeds,omitempty"`
	ReplyMessageIds       []string    `json:"replyMessageIds,omitempty"`
	IsPrivate             bool        `json:"isPrivate,omitempty"`
	IsSilent              bool        `json:"isSilent,omitempty"`
	IsPinned              bool        `json:"isPinned,omitempty"`
	Mentions              Mentions    `json:"mentions,omitempty"`
	CreatedAt             string      `json:"createdAt"`
	CreatedBy             string      `json:"createdBy"`
	CreatedByWebhookId    string      `json:"createdByWebhookId,omitempty"`
	UpdatedAt             string      `json:"updatedAt,omitempty"`
}

type PostMessage struct {
	IsPrivate             bool        `json:"isPrivate,omitempty"`
	IsSilent              bool        `json:"isSilent,omitempty"`
	ReplyMessageIds       []string    `json:"replyMessageIds,omitempty"`
	Content               string      `json:"content,omitempty"`
	HiddenLinkPreviewUrls []string    `json:"hiddenLinkPreviewUrls,omitempty"`
	Embeds                []ChatEmbed `json:"embeds,omitempty"`
}

type MessageRes struct {
	Message Message `json:"message"`
}

type Reaction struct {
	ChannelId string `json:"channelId"`
	CreatedBy string `json:"createdBy"`
	Emote     Emote  `json:"emote"`
}

type ChatReaction struct {
	Reaction
	MessageId string `json:"messageId"`
}

type ForumTopicReaction struct {
	Reaction
	ForumTopicId int `json:"forumTopicId"`
}

type ForumCommentReaction struct {
	ForumTopicReaction
	ForumTopicCommentId int `json:"forumTopicCommentId"`
}

type DocReaction struct {
	Reaction
	DocId int `json:"docId"`
}

type DocCommentReaction struct {
	DocReaction
	DocCommentId int `json:"docCommentId"`
}

type CalendarEventReaction struct {
	Reaction
	CalendarEventId int `json:"calendarEventId"`
}

type CalendarEventCommentReaction struct {
	CalendarEventReaction
	CalendarEventCommentId int `json:"calendarEventCommentId"`
}

type AnnouncementReaction struct {
	Reaction
	AnnouncementId string `json:"announcementId"`
}

type AnnouncementCommentReaction struct {
	AnnouncementReaction
	AnnouncementCommentId int `json:"announcementCommentId"`
}

// TODO: Finish data sets
type Role struct {
	Id                    int      `json:"id"`
	ServerId              string   `json:"serverId"`
	CreatedAt             string   `json:"createdAt"`
	UpdatedAt             *string  `json:"updatedAt"`
	Name                  string   `json:"name"`
	IsDisplayedSeparately *bool    `json:"isDisplayedSeparately"`
	IsSelfAssignable      *bool    `json:"isSelfAssignable"`
	IsMentionable         *bool    `json:"isMentionable"`
	Permissions           []string `json:"permissions"`
	Colors                *[]int   `json:"colors"`
	Icon                  *string  `json:"icon"`
	Priority              *int     `json:"priority"`
	Position              int      `json:"position"`
	IsBase                *bool    `json:"isBase"`
	BotUserId             *string  `json:"botUserId"`
}

type PostRole struct {
	Name                  string   `json:"name"`
	IsDisplayedSeparately *bool    `json:"isDisplayedSeparately"`
	IsSelfAssignable      *bool    `json:"isSelfAssignable"`
	IsMentionable         *bool    `json:"isMentionable"`
	Permissions           []string `json:"permissions"`
	Colors                *[]int   `json:"colors"`
}

// TODO: Finish data sets
type Server struct {
	Id               string  `json:"id"`
	OwnerId          string  `json:"ownerId"`
	Type             *string `json:"type"`
	Name             string  `json:"name"`
	Url              *string `json:"url"`
	About            *string `json:"about"`
	Avatar           *string `json:"avatar"`
	Banner           *string `json:"banner"`
	Timezone         *string `json:"timezone"`
	IsVerified       *bool   `json:"isVerified"`
	DefaultChannelId *string `json:"defaultChannelId"`
	CreatedAt        string  `json:"createdAt"`
}

type ServerGetRes struct {
	Server            Server `json:"server"`
	ServerMemberCount int    `json:"serverMemberCount"`
}

type ServerSubscriptionTier struct {
	Type        string  `json:"type"`
	ServerId    string  `json:"serverId"`
	Description *string `json:"description"`
	RoleId      *int    `json:"roleId"`
	Cost        int     `json:"cost"`
	CreatedAt   string  `json:"createdAt"`
}

// TODO: Finish data sets
type User struct {
	Id        string      `json:"id"`
	Type      *string     `json:"type"`
	Name      string      `json:"name"`
	Avatar    *string     `json:"avatar"`
	Banner    *string     `json:"banner"`
	CreatedAt string      `json:"createdAt"`
	Status    *UserStatus `json:"status"`
}

type UserSum struct {
	Id     string  `json:"id"`
	Type   *string `json:"type"`
	Name   string  `json:"name"`
	Avatar *string `json:"avatar"`
}

type UserServers struct {
	Servers []Server `json:"servers"`
}

type UserStatus struct {
	Content *string `json:"content"`
	EmoteId int     `json:"emoteId"`
}

// SocketEvents

type ChatMessageCreated struct {
	ServerID string  `json:"serverId"`
	Message  Message `json:"message"`
}
