package main

import (
	"context"
	"strings"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/tencent-connect/botgo/dto"
	"github.com/tencent-connect/botgo/event"
	"github.com/tencent-connect/botgo/token"
	"github.com/tencent-connect/botgo/websocket"
)

func main() {
	s := ReadSetting()
	if s == *NullSetting {
		log.Warn("请配置 setting.toml 重新启动")
	} else {
		token := token.BotToken(s.AppId, s.Token, string(token.TypeBot))
		api := NewSandboxOpenAPI(token).WithTimeout(3 * time.Second)
		ctx := context.Background()
		ws, err := api.WS(ctx, nil, "")
		if err != nil {
			log.Warn("登录失败，请检查 appid 和 AccessToken 是否正确。")
			log.Info("该程序将于5秒后退出！")
			time.Sleep(time.Second * 5)
		}
		var g event.GroupAtMessageEventHandler = func(event *dto.WSPayload, data *dto.WSGroupATMessageData) error {
			if strings.TrimSpace(data.Content) == "测试" {
				resp, _ := api.PostGroupRichMediaMessage(ctx, data.GroupId, &dto.GroupRichMediaMessageToCreate{FileType: 1, Url: "https://www.2mf8.cn/static/image/cube3/b1.png", SrvSendMsg: false})
				if resp != nil {
					newMsg := &dto.GroupMessageToCreate{
						Content: "msg",
						Media: &dto.FileInfo{
							FileInfo: resp.FileInfo,
						},
						MsgID:   data.MsgId,
						MsgType: 7,
						MsgReq:  1,
					}
					api.PostGroupMessage(ctx, data.GroupId, newMsg)
				}
			}
			return nil
		}

		var c2cMessage event.C2CMessageEventHandler = func(event *dto.WSPayload, data *dto.WSC2CMessageData) error {
			if data.Content == "测试" {
				resp, err := api.PostC2CRichMediaMessage(ctx, data.Author.UserOpenId, &dto.C2CRichMediaMessageToCreate{FileType: 1, Url: "https://www.2mf8.cn/static/image/cube3/b1.png", SrvSendMsg: false})
				log.Info(err, resp.FileInfo, resp.FileUuid)
				if resp != nil {
					newMsg := &dto.C2CMessageToCreate{
						Content: "msg",
						Media: &dto.FileInfo{
							FileInfo: resp.FileInfo,
						},
						MsgID:   data.Id,
						MsgType: 7,
						MsgReq:  1,
					}
					_, err := api.PostC2CMessage(ctx, data.Author.UserOpenId, newMsg)
					log.Info(err)
				}
				return nil
			}
			return nil
		}
		intent := websocket.RegisterHandlers(g, c2cMessage)
		NewSessionManager().Start(ws, token, &intent)
	}
}
