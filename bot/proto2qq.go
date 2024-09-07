package bot

import (
	"github.com/2mf8/Go-QQ-Client/log"
)

var MsgCount = make(map[string]int, 0)

/*
	 func ProtoMsgToQQMsg(msgList []*onebot.Message, notConvertText bool) []*dto.GroupMessageToCreate {
		MsgCount["text"] = 0
		gmc := &dto.GroupMessageToCreate{}
		for _, protoMsg := range msgList {
			switch protoMsg.Type {
			case "text":
				MsgCount["text"]++
				if notConvertText {
					gmc.Content = gmc.Content + "\n" + ProtoTextToQQText(protoMsg.Data)
				}
			}
		}
	}
*/
func ProtoTextToQQText(data map[string]string) string {
	text, ok := data["text"]
	if !ok {
		log.Warnf("text不存在")
		return ""
	}
	return text
}
