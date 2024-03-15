package gold

// TODO:
// + Implement Heartbeat
// + Better optimize ws reading

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var WSURL = "wss://www.guilded.gg/websocket/v1"

type Client struct {
	Token    string
	Conn     *websocket.Conn
	Callback map[string]func(data interface{})
}

func (c *Client) Connect() {
	auth := http.Header{
		"Authorization": {fmt.Sprintf("Bearer %s", c.Token)},
	}

	ws, _, err := websocket.DefaultDialer.Dial(WSURL, auth)

	if err != nil {
		panic(err)
	}

	c.Conn = ws
	defer c.Conn.Close()

	done := make(chan struct{})

	go func() {
		defer close(done)
		event := SocketEvent{}

		for {
			err := c.Conn.ReadJSON(&event)
			if err != nil {
				log.Println("read:", err)
				return
			}

			c.OnEvent(event.T, event.D)
		}

	}()

	<-done

}

func (c *Client) OnEvent(eType string, data json.RawMessage) {
	_, ok := c.Callback[eType]

	if !ok {
		fmt.Printf("%s does not exist in map\n", eType)
		return
	}

	switch eType {
	case "ChatMessageCreated":
		var msg ChatMessageCreated
		err := json.Unmarshal(data, &msg)
		if err != nil {
			fmt.Println(err)
			return
		}
		c.Callback["ChatMessageCreated"](msg)
	default:
		return
	}

}

func (c *Client) On(eType string, action func(data any)) {
	c.Callback[eType] = action
}

// Creates client and api
func NewClient(token string) (Client, API) {
	return Client{Token: token, Callback: make(map[string]func(data any))}, API{Token: token}
}
