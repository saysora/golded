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

type GetAnnouncementsRes struct {
	Announcements []Announcement `json:"announcements"`
}

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

type PatchCalendarEventCancellation struct {
	Description *string `json:"description"`
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

type CalendarEventRepeatInfoEvery struct {
	Count    int    `json:"count"`
	Interval string `json:"interval"`
}

type CalendarEventRepeatInfo struct {
	Type               string                        `json:"type"`
	Every              *CalendarEventRepeatInfoEvery `json:"every,omitempty"`
	EndAfterOccurences *int                          `json:"endAfterOccurences,omitempty"`
	EndDate            *string                       `json:"endDate,omitempty"`
	On                 *[]string                     `json:"on,omitempty"`
}

type PatchCalendarEvent struct {
	Name             string                          `json:"name"`
	Description      *string                         `json:"description,omitempty"`
	Location         *string                         `json:"location,omitempty"`
	StartsAt         *string                         `json:"startsAt,omitempty"`
	Url              *string                         `json:"url,omitempty"`
	Color            *uint                           `json:"color,omitempty"`
	IsAllday         *bool                           `json:"isAllDay,omitempty"`
	RsvpDisabled     *bool                           `json:"rsvpDisabled,omitempty"`
	RsvpLimit        *int                            `json:"rsvpLimit,omitempty"`
	AutofillWaitlist *bool                           `json:"autofillWaitlist,omitempty"`
	Duration         *uint                           `json:"duration,omitempty"`
	IsPrivate        *bool                           `json:"isPrivate,omitempty"`
	RoleIds          *[]int                          `json:"roleIds,omitempty"`
	Cancellation     *PatchCalendarEventCancellation `json:"cancellation,omitempty"`
}

type GetCalendarEventRes struct {
	CalendarEvent CalendarEvent `json:"calendarEvent"`
}

type GetCalendarEventsRes struct {
	CalendarEvents []CalendarEvent `json:"calendarEvents"`
}

type GetCalendarEventRsvpRes struct {
	CalendarEventRsvp CalendarEventRsvp `json:"calendarEventRsvp"`
}

type GetCalendareventRsvpsRes struct {
	CalendarEventRsvps []CalendarEventRsvp `json:"calendarEventRsvps"`
}

// Calendar Event Comments
type CalendarEventComment struct {
	Id              uint      `json:"id"`
	Content         string    `json:"content"`
	CreatedAt       string    `json:"createdAt"`
	UpdatedAt       string    `json:"updatedAt"`
	CalendarEventId uint      `json:"calendarEventId"`
	ChannelId       string    `json:"channelId"`
	Mentions        *Mentions `json:"mentions,omitempty"`
}

type GetCalendarEventCommentsRes struct {
	CalendarEventComments []CalendarEventComment `json:"calendarEventComments"`
}

type GetCalendarEventCommentRes struct {
	CalendarEventComment CalendarEventComment `json:"calendarEventComment"`
}

// Calendar Event Series
type CalendarEventSeries struct {
	Id        string `json:"id"`
	ServerId  string `json:"serverId"`
	ChannelId string `json:"channelId"`
}

// WARN: Post missing from documentation

type PatchCalendarEventSeries struct {
	Name             *string                  `json:"name,omitempty"`
	Decsription      *string                  `json:"description,omitempty"`
	Location         *string                  `json:"location,omitempty"`
	StartsAt         *string                  `json:"startsAt,omitempty"`
	Url              *string                  `json:"url,omitempty"`
	Color            *int                     `json:"color,omitempty"`
	IsAllDay         *bool                    `json:"isAllDay,omitempty"`
	RsvpDisabled     *bool                    `json:"rsvpDisabled,omitempty"`
	RsvpLimit        *uint                    `json:"rsvpLimit,omitempty"`
	AutofillWaitlist *bool                    `json:"autofillWaitlist,omitempty"`
	Duration         *uint                    `json:"duration,omitempty"`
	IsPrivate        *bool                    `json:"isPrivate,omitempty"`
	RoleIds          *[]int                   `json:"roleIds,omitempty"`
	RepeatInfo       *CalendarEventRepeatInfo `json:"repeatInfo,omitempty"`
	CalendarEventId  *uint                    `json:"calendarEventId,omitempty"`
}

type DelCalendarEventSeries struct {
	CalendarEventId *uint `json:"calendarEventId,omitempty"`
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

type GetCategoryRes struct {
	Category Category `json:"category"`
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

type GetChannelRes struct {
	Channel Channel `json:"channel"`
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

type GetDocsRes struct {
	Docs []Doc `json:"docs"`
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

type PatchGroup struct {
	Name        *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	EmoteId     *int    `json:"emoteId,omitempty"`
	IsPublic    *bool   `json:"isPublic,omitempty"`
}

type GetGroupRes struct {
	Group Group `json:"group"`
}

type GetGroupsRes struct {
	Groups []Group `json:"groups"`
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

type MemberSummary struct {
	User    UserSum `json:"user"`
	RoleIds []int   `json:"roleIds"`
}

type GetMemberRes struct {
	Member Member `json:"member"`
}

type MemberRolesRes struct {
	RoleIds int `json:"roleIds"`
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

type GetMessageRes struct {
	Message Message `json:"message"`
}

type GetMessagesRes struct {
	Messages []Message `json:"messages"`
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

type ChatMessageUpdated = ChatMessageCreated

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

type MemberRoleId struct {
	UserId  string `json:"userId"`
	RoleIds []int  `json:"roleIds"`
}

type ServerRolesUpdated struct {
	ServerId      string         `json:"serverId"`
	MemberRoleIds []MemberRoleId `json:"memberRoleIds"`
}

type ServerChannelCreated struct {
	ServerId string  `json:"serverId"`
	Channel  Channel `json:"channel"`
}

type ServerChannelUpdated = ServerChannelCreated
type ServerChannelDeleted = ServerChannelUpdated

type ServerMemberSocialLinkCreated struct {
	ServerId   string     `json:"serverId"`
	SocialLink SocialLink `json:"socialLink"`
}

type ServerMemberSocialLinkUpdated = ServerMemberSocialLinkCreated
type ServerMemberSocialLinkDeleted = ServerMemberSocialLinkCreated

type ServerWebhookCreated struct {
	ServerId string  `json:"serverId"`
	Webhook  Webhook `json:"webhook"`
}

type ServerWebhookUpdated = ServerWebhookCreated

type DocCreated struct {
	ServerId string `json:"serverId"`
	Doc      Doc    `json:"doc"`
}

type DocUpdated = DocCreated
type DocDeleted = DocUpdated

type DocCommentCreated struct {
	ServerId   string     `json:"serverId"`
	DocComment DocComment `json:"docComment"`
}

type DocCommentUpdated = DocCommentCreated
type DocCommentDeleted = DocCommentUpdated

type CalendarEventCreated struct {
	ServerId      string        `json:"serverId"`
	CalendarEvent CalendarEvent `json:"calendarEvent"`
}

type CalendarEventUpdated = CalendarEventCreated
type CalendarEventDeleted = CalendarEventUpdated

type ForumTopicCreated struct {
	ServerId   string     `json:"serverId"`
	ForumTopic ForumTopic `json:"forumTopic"`
}

type ForumTopicUpdated = ForumTopicCreated
type ForumTopicDeleted = ForumTopicUpdated
type ForumTopicPinned = ForumTopicCreated
type ForumTopicUnpinned = ForumTopicPinned
type ForumTopicLocked = ForumTopicCreated
type ForumTopicUnlocked = ForumTopicCreated

type ForumTopicReactionCreated struct {
	ServerId string             `json:"serverId"`
	Reaction ForumTopicReaction `json:"reaction"`
}

type ForumTopicReactionDeleted = ForumTopicReactionCreated

type ForumTopicCommentCreated struct {
	ServerId          string            `json:"serverId"`
	ForumTopicComment ForumTopicComment `json:"forumTopicComment"`
}

type ForumTopicCommentUpdated = ForumTopicCommentCreated
type ForumTopicCommentDeleted = ForumTopicCommentUpdated

type CalendarEventRsvpUpdated struct {
	ServerId          string            `json:"serverId"`
	CalendarEventRsvp CalendarEventRsvp `json:"calendarEventRsvp"`
}

type CalendarEventRsvpManyUpdated struct {
	ServerId           string              `json:"serverId"`
	CalendarEventRsvps []CalendarEventRsvp `json:"calendarEventRsvps"`
}

type CalendarEventRsvpDeleted = CalendarEventRsvpUpdated

type ListItemCreated struct {
	ServerId string   `json:"serverId"`
	ListItem ListItem `json:"listItem"`
}

type ListItemUpdated = ListItemCreated
type ListItemDeleted = ListItemUpdated
type ListItemCompleted = ListItemCreated
type ListItemUncompleted = ListItemCompleted

type ChannelMessageReactionCreated struct {
	ServerId string       `json:"serverId"`
	Reaction ChatReaction `json:"chatMessageReaction"`
}

type ChannelMessageReactionDeleted struct {
	ServerId  string       `json:"serverId"`
	DeletedBy string       `json:"deletedBy"`
	Reaction  ChatReaction `json:"reaction"`
}

type ChannelMessageReactionManyDeleted struct {
	ServerId  string `json:"serverId"`
	ChannelId string `json:"channelId"`
	MessageId string `json:"messageId"`
	DeletedBy string `json:"deletedBy"`
	Count     int    `json:"count"`
	Emote     *Emote `json:"emote,omitempty"`
}

// type ForumTopicCommentReactionCreated struct {
//   ServerId string `json:"serverId"`
//   Reaction ForumTopicCommentReaction `json:"reaction"`
// }

// type ForumTopicCommentReactionDeleted = ForumTopicCommentReactionCreated

// type CalendarEventCommentCreated struct {
//   ServerId string `json:"serverId"`
//   CalendarEventComment CalendarEventComment `json:"calendarEventComment"`
// }

//type CalendarEventCommentUpdated = CalendareEventCommentCreated
//type CalendarEventCommentDeleted = CalendareEventCommentCreated

type CalendarEventReactionCreated struct {
	ServerId string                `json:"serverId"`
	Reaction CalendarEventReaction `json:"reaction"`
}

type CalendarEventReactionDeleted = CalendarEventReactionCreated

type CalendarEventCommentReactionCreated struct {
	ServerId string                       `json:"serverId"`
	Reaction CalendarEventCommentReaction `json:"reaction"`
}

type CalendarEventCommentReactionDeleted = CalendarEventCommentReactionCreated

type DocReactionCreated struct {
	ServerId string      `json:"serverId"`
	Reaction DocReaction `json:"docReaction"`
}

type DocReactionDeleted = DocReactionCreated

type DocCommentReactionCreated struct {
	ServerId string             `json:"serverId"`
	Reaction DocCommentReaction `json:"reaction"`
}

type DocCommentReactionDeleted = DocCommentReactionCreated

type CalendarEventSeriesUpdated struct {
	ServerId            string              `json:"serverId"`
	CalendarEventSeries CalendarEventSeries `json:"calendarEventSeries"`
	CalendarEventId     *uint               `json:"calendareventId"`
}

type CalendarEventSeriesDeleted struct {
	ServerId            string              `json:"serverId"`
	CalendarEventSeries CalendarEventSeries `json:"calendarEventSeries"`
	CalendarEventId     *uint               `json:"calendarEventId"`
}

type GroupCreated struct {
	ServerId string `json:"serverId"`
	Group    Group  `json:"group"`
}

type GroupUpdated struct {
	ServerId string `json:"serverId"`
	Group    Group  `json:"group"`
}

type GroupDeleted struct {
	ServerId string `json:"serverId"`
	Group    Group  `json:"group"`
}

// Announcement

type AnnouncementCreated struct {
	ServerId     string       `json:"serverId"`
	Announcement Announcement `json:"announcement"`
}

type AnnouncementUpdated struct {
	ServerId     string       `json:"serverId"`
	Announcement Announcement `json:"announcement"`
}

type AnnouncementDeleted struct {
	ServerId     string       `json:"serverId"`
	Announcement Announcement `json:"announcement"`
}

type AnnouncementReactionCreated struct {
	ServerId string               `json:"serverId"`
	Reaction AnnouncementReaction `json:"reaction"`
}

type AnnouncementReactionDeleted struct {
	ServerId string               `json:"serverId"`
	Reaction AnnouncementReaction `json:"reaction"`
}

type AnnouncementCommentCreated struct {
	ServerId            string              `json:"serverId"`
	AnnouncementComment AnnouncementComment `json:"annuncementComment"`
}

type AnnouncementCommentUpdated struct {
	ServerId            string              `json:"serverId"`
	AnnouncementComment AnnouncementComment `json:"annuncementComment"`
}

type AnnouncementCommentDeleted struct {
	ServerId            string              `json:"serverId"`
	AnnouncementComment AnnouncementComment `json:"annuncementComment"`
}

type AnnouncementCommentReactionCreated struct {
	ServerId string                      `json:"serverId"`
	Reaction AnnouncementCommentReaction `json:"reaction"`
}

type AnnouncementCommentReactionDeleted struct {
	ServerId string                      `json:"serverId"`
	Reaction AnnouncementCommentReaction `json:"reaction"`
}

// User status

type UserStatusCreated struct {
	ExpiresAt  *string    `json:"expiresAt,omitempty"`
	UserId     string     `json:"userId"`
	UserStatus UserStatus `json:"userStatus"`
}

type UserStatusDeleted struct {
	UserId     string     `json:"userId"`
	UserStatus UserStatus `json:"userStatus"`
}

// Roles
type RoleCreated struct {
	ServerId string `json:"serverId"`
	Role     Role   `json:"role"`
}

type RoleUpdated struct {
	ServerId string `json:"serverId"`
	Role     Role   `json:"role"`
}

type RoleDeleted struct {
	ServerId string `json:"serverId"`
	Role     Role   `json:"role"`
}

type Permission map[string]bool

type ChannelRolePermission struct {
	Permissions []Permission `json:"permissions"`
	CreatedAt   string       `json:"createdAt"`
	UpdatedAt   *string      `json:"updatedAt,omitempty"`
	RoleId      int          `json:"roleId"`
	ChannelId   string       `json:"channelId"`
}

type ChannelUserPermission struct {
	Permissions []Permission `json:"permissions"`
	CreatedAt   string       `json:"createdAt"`
	UpdatedAt   *string      `json:"updatedAt,omitempty"`
	UserId      string       `json:"userId"`
	ChannelId   string       `json:"channelId"`
}

type ChannelCategoryRolePermission struct {
	Permissions []Permission `json:"permissions"`
	CreatedAt   string       `json:"createdAt"`
	UpdatedAt   *string      `json:"updatedAt,omitempty"`
	RoleId      int          `json:"roleId"`
	CategoryId  string       `json:"categoryId"`
}

type ChannelCategoryUserPermission struct {
	Permissions []Permission `json:"permissions"`
	CreatedAt   string       `json:"createdAt"`
	UpdatedAt   *string      `json:"updatedAt,omitempty"`
	UserId      string       `json:"userId"`
	CategoryId  string       `json:"categoryId"`
}

// Channels

type ChannelArchived struct {
	ServerId string  `json:"string"`
	Channel  Channel `json:"channel"`
}

type ChannelRestored struct {
	ServerId string  `json:"string"`
	Channel  Channel `json:"channel"`
}

// Category
type CategoryCreated struct {
	ServerId string   `json:"serverId"`
	Category Category `json:"category"`
}

type CategoryUpdated struct {
	ServerId string   `json:"serverId"`
	Category Category `json:"category"`
}

type CategoryDeleted struct {
	ServerId string   `json:"serverId"`
	Category Category `json:"category"`
}

// Channel Messages

type ChannelMessagePinned struct {
	ServerId string  `json:"serverId"`
	Message  Message `json:"message"`
}

type ChannelMessageUnpinned struct {
	ServerId string  `json:"serverId"`
	Message  Message `json:"message"`
}

type ChannelRolePermissionUpdated struct {
	ServerId              string                `json:"serverId"`
	ChannelRolePermission ChannelRolePermission `json:"channelRolePermission"`
}

type ChannelRolePermissionDeleted struct {
	ServerId              string                `json:"serverId"`
	ChannelRolePermission ChannelRolePermission `json:"channelRolePermission"`
}

type ChannelRolePermissionCreated struct {
	ServerId              string                `json:"serverId"`
	ChannelRolePermission ChannelRolePermission `json:"channelRolePermission"`
}

type ChannelUserPermissionCreated struct {
	ServerId              string                `json:"serverId"`
	ChannelUserPermission ChannelUserPermission `json:"channelUserPermission"`
}

type ChannelUserPermissionUpdated struct {
	ServerId              string                `json:"serverId"`
	ChannelUserPermission ChannelUserPermission `json:"channelUserPermission"`
}

type ChannelUserPermissionDeleted struct {
	ServerId              string                `json:"serverId"`
	ChannelUserPermission ChannelUserPermission `json:"channelUserPermission"`
}

type ChannelCategoryUserPermissionCreated struct {
	ServerId                      string                        `json:"serverId"`
	ChannelCategoryUserPermission ChannelCategoryUserPermission `json:"channelCategoryUserPermission"`
}

type ChannelCategoryUserPermissionUpdated struct {
	ServerId                      string                        `json:"serverId"`
	ChannelCategoryUserPermission ChannelCategoryUserPermission `json:"channelCategoryUserPermission"`
}

type ChannelCategoryUserPermissionDeleted struct {
	ServerId                      string                        `json:"serverId"`
	ChannelCategoryUserPermission ChannelCategoryUserPermission `json:"channelCategoryUserPermission"`
}

type ChannelCategoryRolePermissionCreated struct {
	ServerId                      string                        `json:"serverId"`
	ChannelCategoryRolePermission ChannelCategoryRolePermission `json:"channelCategoryRolePermission"`
}

type ChannelCategoryRolePermissionUpdated struct {
	ServerId                      string                        `json:"serverId"`
	ChannelCategoryRolePermission ChannelCategoryRolePermission `json:"channelCategoryRolePermission"`
}

type ChannelCategoryRolePermissionDeleted struct {
	ServerId                      string                        `json:"serverId"`
	ChannelCategoryRolePermission ChannelCategoryRolePermission `json:"channelCategoryRolePermission"`
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

// TODO: Add in all the permissions
