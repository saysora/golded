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

// Announcements
type Announcement struct {
	Id        string    `json:"id"`
	ServerId  string    `json:"serverId"`
	GroupId   string    `json:"groupId"`
	ChannelId string    `json:"channelId"`
	CreatedAt string    `json:"createdAt"`
	CreatedBy string    `json:"createdBy"`
	Content   string    `json:"content"`
	Mentions  *Mentions `json:"mentions,omitempty"`
	Title     string    `json:"title"`
}

type PostAnnouncement struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

// TODO: Think over structure for get requests

type GetAnnouncementsRes struct {
	Announcements []Announcement `json:"announcements"`
}

// NOTE: Reused for responses for Post / Patch
type GetAnnouncementRes struct {
	Announcement Announcement `json:"announcement"`
}

type PatchAnnouncement struct {
	Title   string `json:"title,omitempty"`
	Content string `json:"content,omitempty"`
}

type AnnouncementComment struct {
	Id             uint      `json:"id"`
	Content        string    `json:"content"`
	CreatedAt      string    `json:"createdAt"`
	UpdatedAt      *string   `json:"updatedAt,omitempty"`
	CreatedBy      string    `json:"createdBy"`
	ChannelId      string    `json:"channelId"`
	AnnouncementId string    `json:"announcementId"`
	Mentions       *Mentions `json:"mentions,omitempty"`
}

// Duplicated by Patch
type PostAnnouncementComment struct {
	Comment string `json:"content"`
}

type GetAnnouncementComments struct {
	AnnouncementComments []AnnouncementComment `json:"announcementComments"`
}

type GetAnnouncementComment struct {
	AnnouncementComment AnnouncementComment `json:"announcementComment"`
}

// Calendar Events

type CalendarEvent struct {
	Id               string                     `json:"id"`
	ServerId         string                     `json:"serverId"`
	GroupId          string                     `json:"groupId"`
	ChannelId        string                     `json:"channelId"`
	Name             string                     `json:"name"`
	Description      *string                    `json:"description,omitempty"`
	Location         *string                    `json:"location,omitempty"`
	Url              *string                    `json:"url,omitempty"`
	Color            *string                    `json:"color,omitempty"`
	Repeats          *bool                      `json:"repeats,omitempty"`
	SeriesId         *string                    `json:"seriesId,omitempty"`
	RoleIds          *string                    `json:"roleIds,omitempty"`
	RsvpDisabled     *bool                      `json:"rsvpDisabled,omitempty"`
	IsAllDay         *string                    `json:"isAllDay,omitempty"`
	RsvpLimit        *uint                      `json:"rsvpLimit,omitempty"`
	AutofillWaitlist *bool                      `json:"autofillWaitlist,omitempty"`
	StartsAt         string                     `json:"startsAt"`
	Duration         *uint                      `json:"duration,omitempty"`
	IsPrivate        *bool                      `json:"isPrivate,omitempty"`
	Mentions         *Mentions                  `json:"mentions,omitempty"`
	CreatedAt        string                     `json:"createdAt"`
	CreatedBy        string                     `json:"createdBy"`
	Cancellation     *CalendarEventCancellation `json:"cancellation,omitempty"`
}

type CalendarEventCancellation struct {
	Description *string `json:"description,omitempty"`
	CreatedBy   string  `json:"createdBy"`
}

type CalendarEventRsvp struct {
	CalendarEventId uint    `json:"calendarEventId"`
	ChannelId       string  `json:"channelId"`
	ServerId        string  `json:"serverId"`
	UserId          string  `json:"userId"`
	Status          string  `json:"status"`
	CreatedBy       string  `json:"createdBy"`
	UpdatedBy       *string `json:"updatedBy,omitempty"`
	UpdatedAt       *string `json:"updatedAt,omitempty"`
}

// Categories
type Category struct {
	Id        int     `json:"id"`
	ServerId  string  `json:"serverId"`
	GroupId   string  `json:"groupId"`
	CreatedAt string  `json:"createdAt"`
	UpdatedAt *string `json:"updatedAt,omitempty"`
	Name      string  `json:"name"`
	Priority  *int    `json:"priority,omitempty"`
}

type PostCategory struct {
	Name    string  `json:"name"`
	GroupId *string `json:"groupId,omitempty"`
}

type PatchCategory struct {
	Name     *string `json:"name,omitempty"`
	Priority *int    `json:"priority,omitempty"`
}

// Channels
type Channel struct {
	Id         string  `json:"id"`
	Type       string  `json:"type"`
	Name       string  `json:"name"`
	Topic      *string `json:"topic,omitempty"`
	CreatedAt  string  `json:"createdAt"`
	CreatedBy  string  `json:"createdBy"`
	UpdatedAt  *string `json:"updatedAt,omitempty"`
	ServerId   string  `json:"serverId"`
	RootId     *string `json:"rootId,omitempty"`
	ParentId   *string `json:"parentId,omitempty"`
	MessageId  *string `json:"messageId,omitempty"`
	CategoryId *int    `json:"categoryId,omitempty"`
	GroupId    string  `json:"groupId"`
	Visibility *string `json:"visibility,omitempty"`
	ArchivedBy *string `json:"archivedBy,omitempty"`
	ArchivedAt *string `json:"archivedAt,omitempty"`
}

type PostChannel struct {
	Name       string  `json:"name"`
	Topic      *string `json:"topic,omitempty"`
	Visibility *string `json:"visibility,omitempty"`
	Type       string  `json:"type"`
	ServerId   *string `json:"serverId,omitempty"`
	GroupId    *string `json:"groupId,omitempty"`
	CategoryId *string `json:"categoryId,omitempty"`
	ParentId   *string `json:"parentId,omitempty"`
	MessageId  *string `json:"messageId,omitempty"`
}

type PatchChannel struct {
	Name       *string `json:"name,omitempty"`
	Topic      *string `json:"topic,omitempty"`
	Visibility *string `json:"visibility,omitempty"`
}

// Chat Embed
type ChatEmbedFooter struct {
	IconUrl *string `json:"icon_url,omitempty"`
	Text    string  `json:"text"`
}

type ChatEmbedMedia struct {
	Url string `json:"url"`
}

type ChatEmbedAuthor struct {
	Name    *string `json:"name,omitempty"`
	Url     *string `json:"url,omitempty"`
	IconUrl *string `json:"icon_url,omitempty"`
}

type ChatEmbedField struct {
	Name   string `json:"name"`
	Value  string `json:"value"`
	Inline *bool  `json:"inline,omitempty"`
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

// Docs
type Doc struct {
	Id        uint      `json:"id"`
	ServerId  string    `json:"string"`
	GroupId   string    `json:"groupId"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	Mentions  *Mentions `json:"mentions,omitempty"`
	CreatedAt string    `json:"createdAt"`
	CreatedBy string    `json:"createdBy"`
	UpdatedAt string    `json:"updatedAt"`
	UpdatedBy *string   `json:"updatedBy,omitempty"`
}

type PostDoc struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

type GetDocRes struct {
	Doc Doc `json:"doc"`
}

type DocComment struct {
	Id        int       `json:"id"`
	Content   string    `json:"content"`
	CreatedAt string    `json:"createdAt"`
	CreatedBy string    `json:"createdBy"`
	UpdatedAt string    `json:"updatedAt"`
	ChannelId string    `json:"channelId"`
	DocId     int       `json:"docId"`
	Mentions  *Mentions `json:"mentions,omitempty"`
}

type PostDocComment struct {
	Content string `json:"content"`
}

type PatchDocComment struct {
	Content string `json:"content"`
}

type DocCommentRes struct {
	DocComment DocComment `json:"docComment"`
}

type PatchDoc struct {
	Title   *string `json:"title,omitempty"`
	Content *string `json:"content,omitempty"`
}

// Emotes
type Emote struct {
	Id       int     `json:"id"`
	Name     string  `json:"name"`
	Url      string  `json:"url"`
	ServerId *string `json:"serverId,omitempty"`
}

// Forum Topics
type ForumTopic struct {
	Id        int       `json:"id"`
	ServerId  string    `json:"serverId"`
	GroupId   string    `json:"groupId"`
	ChannelId string    `json:"channelId"`
	Title     string    `json:"title"`
	CreatedAt string    `json:"createdAt"`
	CreatedBy string    `json:"createdBy"`
	UpdatedAt *string   `json:"updatedAt,omitempty"`
	UpdatedBy *string   `json:"updatedBy,omitempty"`
	Content   string    `json:"content"`
	Mentions  *Mentions `json:"mentions,omitempty"`
}

type PostForumTopic struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

type PatchForumTopic struct {
	Title   *string `json:"title,omitempty"`
	Content *string `json:"content,omitempty"`
}

type ForumTopicRes struct {
	ForumTopic ForumTopic `json:"forumTopic"`
}

type ForumTopicSummary struct {
	Id        int     `json:"id"`
	ServerId  string  `json:"serverId"`
	GroupId   string  `json:"groupId"`
	ChannelId string  `json:"channelId"`
	Title     string  `json:"title"`
	CreatedAt string  `json:"createdAt"`
	CreatedBy string  `json:"createdBy"`
	UpdatedAt *string `json:"updatedAt,omitempty"`
	BumpedAt  *string `json:"bumpedAt,omitempty"`
	IsPinned  *bool   `json:"isPinned,omitempty"`
	IsLocked  *bool   `json:"isLocked,omitempty"`
}

type ForumTopicComment struct {
	Id           int       `json:"id"`
	Content      string    `json:"content"`
	CreatedAt    string    `json:"createdAt"`
	UpdatedAt    *string   `json:"updatedAt,omitempty"`
	ChannelId    string    `json:"channelId"`
	ForumTopicId int       `json:"forumTopicId"`
	CreatedBy    string    `json:"createdBy"`
	Mentions     *Mentions `json:"mentions,omitempty"`
}

type PostForumTopicComment struct {
	Content string `json:"content"`
}

type PatchForumTopicComment struct {
	Content string `json:"content"`
}

type ForumTopicCommentRes struct {
	ForumTopicComment ForumTopicComment `json:"forumTopicComment"`
}

// Groups
type Group struct {
	Id          string  `json:"id"`
	ServerId    string  `json:"serverId"`
	Name        string  `json:"name"`
	Description *string `json:"description,omitempty"`
	Avatar      *string `json:"avatar,omitempty"`
	IsHome      *bool   `json:"isHome,omitempty"`
	EmoteId     *int    `json:"emoteId,omitempty"`
	IsPublic    *bool   `json:"isPublic,omitempty"`
	CreatedAt   string  `json:"createdAt"`
	CreatedBy   string  `json:"createdBy"`
	UpdatedAt   *string `json:"updatedAt,omitempty"`
	UpdatedBy   *string `json:"updatedBy,omitempty"`
	ArchivedAt  *string `json:"archivedAt,omitempty"`
	ArchivedBy  *string `json:"archivedBy,omitempty"`
}

type PostGroup struct {
	Name        string  `json:"name"`
	Description *string `json:"description,omitempty"`
	EmoteId     *int    `json:"emoteId,omitempty"`
	IsPublic    *bool   `json:"isPublic,omitempty"`
}

// List Items
type ListItem struct {
	Id                 string        `json:"id"`
	ServerId           string        `json:"serverId"`
	GroupId            string        `json:"groupId"`
	ChannelId          string        `json:"channelId"`
	Message            string        `json:"message"`
	Mentions           *Mentions     `json:"mentions,omitempty"`
	CreatedAt          string        `json:"createdAt"`
	CreatedByWebhookId *string       `json:"createdByWebhookId,omitempty"`
	UpdatedAt          *string       `json:"updatedAt,omitempty"`
	UpdatedBy          *string       `json:"updatedBy,omitempty"`
	ParentListItemId   *string       `json:"parentListItemId,omitempty"`
	CompletedAt        *string       `json:"completedAt,omitempty"`
	CompletedBy        *string       `json:"completedBy,omitempty"`
	Note               *ListItemNote `json:"note,omitempty"`
}

type ListItemNote struct {
	CreatedAt string    `json:"createdAt"`
	CreatedBy string    `json:"createdBy"`
	UpdatedAt *string   `json:"updatedAt,omitempty"`
	UpdatedBy *string   `json:"updatedBy,omitempty"`
	Mentions  *Mentions `json:"mentions,omitempty"`
	Content   string    `json:"content"`
}

type ListItemSummary struct {
	Id                 string               `json:"id"`
	ServerId           string               `json:"serverId"`
	ChannelId          string               `json:"channelId"`
	Message            string               `json:"message"`
	Mentions           *Mentions            `json:"mentions,omitempty"`
	CreatedAt          string               `json:"createdAt"`
	CreatedBy          string               `json:"createdBy"`
	CreatedByWebhookId *string              `json:"createdByWebhookId,omitempty"`
	UpdatedAt          *string              `json:"updatedAt,omitempty"`
	UpdatedBy          *string              `json:"updatedBy,omitempty"`
	ParentListItemId   *string              `json:"parentListItemId,omitempty"`
	CompletedAt        *string              `json:"completedAt,omitempty"`
	CompletedBy        *string              `json:"completedBy,omitempty"`
	Note               *ListItemNoteSummary `json:"note,omitempty"`
}

type ListItemNoteSummary struct {
	CreatedAt string  `json:"createdAt"`
	CreatedBy string  `json:"createdBy"`
	UpdatedAt *string `json:"updatedAt,omitempty"`
	UpdatedBy *string `json:"updatedBy,omitempty"`
}

type PostListItem struct {
	Message string            `json:"message"`
	Note    *PostListItemNote `json:"note,omitempty"`
}

type PostListItemNote struct {
	Content string `json:"content"`
}

type PatchListItem struct {
	Message *string           `json:"message,omitempty"`
	Note    *PostListItemNote `json:"note,omitempty"`
}

type ListItemRes struct {
	ListItem ListItem `json:"listItem"`
}

// Members
type Member struct {
	User     User    `json:"user"`
	RoleIds  *[]int  `json:"roleIds,omitempty"`
	Nickname *string `json:"nickname,omitempty"`
	JoinedAt string  `json:"joinedAt"`
	IsOwner  *bool   `json:"isOwner,omitempty"`
}

type MemberSum struct {
	User    UserSum `json:"user"`
	RoleIds []int   `json:"roleIds"`
}

// Member Ban
type MemberBan struct {
	User      UserSum `json:"user"`
	Reason    *string `json:"reason,omitempty"`
	CreatedBy string  `json:"createdBy"`
	CreatedAt string  `json:"createdAt"`
}

type PostMemberBan struct {
	Reason *string `json:"reason,omitempty"`
}

// Mentions
type MentionUser struct {
	Id string `json:"id"`
}

type MentionChannel = MentionUser

type MentionRole struct {
	Id int `json:"id"`
}

type Mentions struct {
	Users    *[]MentionUser    `json:"users,omitempty"`
	Channels *[]MentionChannel `json:"channels,omitempty"`
	Roles    *[]MentionRole    `json:"roles,omitempty"`
	Everyone *bool             `json:"everyone,omitempty"`
	Here     *bool             `json:"here,omitempty"`
}

// Messages
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

type DeletedMessage struct {
	Id                    string       `json:"id"`
	Type                  string       `json:"type"`
	ServerId              *string      `json:"serverId,omitempty"`
	GroupId               *string      `json:"groupId,omitempty"`
	ChannelId             string       `json:"channelId"`
	Content               *string      `json:"content,omitempty"`
	HiddenLinkPreviewUrls *[]string    `json:"hiddenLinkPreviewUrls,omitempty"`
	Embeds                *[]ChatEmbed `json:"embeds,omitempty"`
	ReplyMessageIds       *[]string    `json:"replyMessageIds,omitempty"`
	IsPrivate             *bool        `json:"isPrivate,omitempty"`
	IsSilent              *bool        `json:"isSilent,omitempty"`
	IsPinned              *bool        `json:"isPinned,omitempty"`
	Mentions              *Mentions    `json:"mentions,omitempty"`
	CreatedAt             string       `json:"createdAt"`
	CreatedBy             string       `json:"createdBy"`
	CreatedByWebhookId    *string      `json:"createdByWebhookId,omitempty"`
	UpdatedAt             *string      `json:"updatedAt,omitempty"`
	DeletedAt             string       `json:"deletedAt"`
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

// Reactions
// NOTE: Way too many, will update later
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

// Roles
type Role struct {
	Id                    int      `json:"id"`
	ServerId              string   `json:"serverId"`
	CreatedAt             string   `json:"createdAt"`
	UpdatedAt             *string  `json:"updatedAt,omitempty"`
	Name                  string   `json:"name"`
	IsDisplayedSeparately *bool    `json:"isDisplayedSeparately,omitempty"`
	IsSelfAssignable      *bool    `json:"isSelfAssignable,omitempty"`
	IsMentionable         *bool    `json:"isMentionable,omitempty"`
	Permissions           []string `json:"permissions"`
	Colors                *[]int   `json:"colors,omitempty"`
	Icon                  *string  `json:"icon,omitempty"`
	Priority              *int     `json:"priority,omitempty"`
	Position              int      `json:"position"`
	IsBase                *bool    `json:"isBase,omitempty"`
	BotUserId             *string  `json:"botUserId,omitempty"`
}

type PostRole struct {
	Name                  string   `json:"name"`
	IsDisplayedSeparately *bool    `json:"isDisplayedSeparately,omitempty"`
	IsSelfAssignable      *bool    `json:"isSelfAssignable,omitempty"`
	IsMentionable         *bool    `json:"isMentionable,omitempty"`
	Permissions           []string `json:"permissions"`
	Colors                *[]int   `json:"colors,omitempty"`
}

// Servers
type Server struct {
	Id               string  `json:"id"`
	OwnerId          string  `json:"ownerId"`
	Type             *string `json:"type,omitempty"`
	Name             string  `json:"name"`
	Url              *string `json:"url,omitempty"`
	About            *string `json:"about,omitempty"`
	Avatar           *string `json:"avatar,omitempty"`
	Banner           *string `json:"banner,omitempty"`
	Timezone         *string `json:"timezone,omitempty"`
	IsVerified       *bool   `json:"isVerified,omitempty"`
	DefaultChannelId *string `json:"defaultChannelId,omitempty"`
	CreatedAt        string  `json:"createdAt"`
}

type ServerGetRes struct {
	Server            Server `json:"server"`
	ServerMemberCount int    `json:"serverMemberCount"`
}

// Server Subs
type ServerSubscriptionTier struct {
	Type        string  `json:"type"`
	ServerId    string  `json:"serverId"`
	Description *string `json:"description,omitempty"`
	RoleId      *int    `json:"roleId,omitempty"`
	Cost        int     `json:"cost"`
	CreatedAt   string  `json:"createdAt"`
}

// Social link
type SocialLink struct {
	Type      string  `json:"type"`
	UserId    string  `json:"userId"`
	Handle    *string `json:"handle,omitempty"`
	ServiceId *string `json:"serviceId,omitempty"`
	CreatedAt string  `json:"createdAt"`
}

type SocialLinkRes struct {
	SocialLink SocialLink `json:"socialLink"`
}

// Users
type User struct {
	Id        string      `json:"id"`
	Type      *string     `json:"type,omitempty"`
	Name      string      `json:"name"`
	Avatar    *string     `json:"avatar,omitempty"`
	Banner    *string     `json:"banner,omitempty"`
	CreatedAt string      `json:"createdAt"`
	Status    *UserStatus `json:"status,omitempty"`
}

type UserSum struct {
	Id     string  `json:"id"`
	Type   *string `json:"type,omitempty"`
	Name   string  `json:"name"`
	Avatar *string `json:"avatar,omitempty"`
}

type UserServers struct {
	Servers []Server `json:"servers"`
}

type UserStatus struct {
	Content *string `json:"content,omitempty"`
	EmoteId int     `json:"emoteId"`
}

// Webhooks

type Webhook struct {
	Id        string  `json:"id"`
	Name      string  `json:"name"`
	Avatar    *string `json:"avatar,omitempty"`
	ServerId  string  `json:"serverId"`
	ChannelId string  `json:"channelId"`
	CreatedAt string  `json:"createdAt"`
	CreatedBy string  `json:"createdBy"`
	DeletedAt *string `json:"deletedAt,omitempty"`
	Token     *string `json:"token,omitempty"`
}

type PostWebhook struct {
	Name      string `json:"name"`
	ChannelId string `json:"channelId"`
}

type PatchWebhook struct {
	Name      *string `json:"name,omitempty"`
	ChannelId *string `json:"channelId,omitempty"`
}

type WebhookRes struct {
	Webhook Webhook `json:"webhook"`
}

// XP
type PostXP struct {
	Amount int `json:"amount"`
}

// SocketEvents

type BotServerMembershipCreated struct {
	Server    Server `json:"server"`
	CreatedBy string `json:"createdBy"`
}

type BotServerMembershipDeleted struct {
	Server    Server `json:"server"`
	DeletedBy string `json:"deletedBy"`
}

type ChatMessageCreated struct {
	ServerId string  `json:"serverId"`
	Message  Message `json:"message"`
}

type ChatMessageUpdated struct {
	ServerId string  `json:"serverId"`
	Message  Message `json:"message"`
}

type ChatMessageDeleted struct {
	ServerId  string         `json:"serverId"`
	DeletedAt string         `json:"deletedAt"`
	Message   DeletedMessage `json:"deletedMessage"`
}

type ServerMemberJoined struct {
	ServerId          string `json:"serverId"`
	Member            Member `json:"member"`
	ServerMemberCount uint   `json:"serverMemberCount"`
}

type ServerMemberRemoved struct {
	ServerId string `json:"serverId"`
	UserId   string `json:"userId"`
	IsKick   *bool  `json:"isKick,omitempty"`
	IsBan    *bool  `json:"isBan,omitempty"`
}

type ServerMemberBanned struct {
	ServerId        string    `json:"serverId"`
	ServerMemberBan MemberBan `json:"serverMemberBan"`
}

type ServerMemberUnbanned struct {
	ServerId        string    `json:"serverId"`
	ServerMemberBan MemberBan `json:"serverMemberBan"`
}

type UserInfo struct {
	Id       string  `json:"id"`
	Nickname *string `json:"nickname,omitempty"`
}

type ServerMemberUpdated struct {
	ServerId string   `json:"serverId"`
	UserInfo UserInfo `json:"userInfo"`
}
