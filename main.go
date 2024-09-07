package main

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/2mf8/Go-QQ-Client/dto"
	"github.com/2mf8/Go-QQ-Client/log"
	"github.com/2mf8/Go-QQ-Client/token"
	"github.com/2mf8/Go-QQ-Client/websocket"
)

func main() {
	Ws := &dto.WebsocketAP{}
	s := ReadSetting()
	token := token.BotToken(s.AppId, s.Token, string(token.TypeBot))
	b, e := json.Marshal(s)
	fmt.Println(string(b), e)
	if s.IsSandBox {
		api := NewSandboxOpenAPI(token).WithTimeout(3 * time.Second)
		ctx := context.Background()
		ws, err := api.WS(ctx, nil, "")
		if err != nil {
			log.Warn("登录失败，请检查 appid 和 AccessToken 是否正确。")
			time.Sleep(time.Second * 5)
		} else {
			Ws = ws
		}
	} else {
		api := NewOpenAPI(token).WithTimeout(3 * time.Second)
		ctx := context.Background()
		ws, err := api.WS(ctx, nil, "")
		if err != nil {
			log.Warn("登录失败，请检查 appid 和 AccessToken 是否正确。")
			time.Sleep(time.Second * 5)
		} else {
			Ws = ws
		}
	}
	intent := websocket.RegisterHandlers()
	if s.IsPrivate {
		if s.IntentGuildMessages {
			intent = intent | dto.IntentGuildMessages
			fmt.Println(intent)
		}
		if s.IntentForum {
			intent = intent | dto.IntentForum
		}
	} else {
		if s.IntentGuildAtMessage {
			intent = intent | dto.IntentGuildAtMessage
		}
	}
	if s.IntentGuilds {
		intent = intent | dto.IntentGuilds
	}
	if s.IntentGuildMembers {
		intent = intent | dto.IntentGuildMembers
	}
	if s.IntentGuildMessageReactions {
		intent = intent | dto.IntentGuildMessageReactions
	}
	if s.IntentDirectMessages {
		intent = intent | dto.IntentDirectMessages
	}
	if s.IntentQQ {
		intent = intent | dto.IntentQQ
	}
	if s.IntentInteraction {
		intent = intent | dto.IntentInteraction
	}
	if s.IntentAudit {
		intent = intent | dto.IntentAudit
	}
	if s.IntentAudio {
		intent = intent | dto.IntentAudio
	}
	NewSessionManager().Start(Ws, token, &intent)
}
