# gold

A simple and lightweight go library for working with the Guilded API

:warning: In Development :warning:

This library is currently in development. Using it is not recommended until this library meets


### Usage

```go
package main

import gold "github.com/saysora/gold"

func main() {
    // client is used for handling the websocket client, api is used to communicate to the guilded rest api
    client, api := gold.NewClient("<TOKEN>")

    // On method is used to define what websocket event to respond to
    client.On("ChatMessageCreated", func(event any) {
        payload, ok := event.(gold.ChatMessageCreated)
        
        if !ok {
            return
        }

        if payload.Message.Content == "ping" {
            api.SendMessage(payload.Message.ChannelId, gold.PostMessage{Content: "pong"});
        }
    })

    // Connects to the websocket
    client.Connect()
}
```

More details to come.
