package main

import (
	"fmt"

	"github.com/2mf8/Go-QQ-Client/dto"
	"github.com/2mf8/Go-QQ-Client/event"
)

// ThreadEventHandler 论坛主贴事件
func ThreadEventHandler() event.ThreadEventHandler {
	return func(event *dto.WSPayload, data *dto.WSThreadData) error {
		fmt.Println(event, data)
		return nil
	}
}
