package bot

import (
	"context"

	"github.com/2mf8/Go-QQ-Client/dto"
	"github.com/2mf8/Go-QQ-Client/onebot"
	"github.com/2mf8/Go-QQ-Client/openapi"
)

func HandleSendGroupMsg(c context.Context, api openapi.OpenAPI, req *onebot.Api_Params) {
	api.PostGroupMessage(c, req.GroupId, &dto.GroupMessageToCreate{})
}
