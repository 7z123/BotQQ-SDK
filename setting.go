package main

import (
	"fmt"
	"os"
	"time"

	"github.com/2mf8/Go-QQ-Client/log"
	"github.com/BurntSushi/toml"
)

type Setting struct {
	QQ        uint64
	AppId     uint64
	Token     string
	AppSecret string
	IsPrivate bool
	IsSandBox bool

	// IntentGuilds 包含
	// - GUILD_CREATE
	// - GUILD_UPDATE
	// - GUILD_DELETE
	// - GUILD_ROLE_CREATE
	// - GUILD_ROLE_UPDATE
	// - GUILD_ROLE_DELETE
	// - CHANNEL_CREATE
	// - CHANNEL_UPDATE
	// - CHANNEL_DELETE
	// - CHANNEL_PINS_UPDATE
	IntentGuilds bool

	// IntentGuildMembers 包含
	// - GUILD_MEMBER_ADD
	// - GUILD_MEMBER_UPDATE
	// - GUILD_MEMBER_REMOVE
	IntentGuildMembers bool

/* 	不支持
	IntentGuildBans         bool
	IntentGuildEmojis       bool
	IntentGuildIntegrations bool
	IntentGuildWebhooks     bool
	IntentGuildInvites      bool
	IntentGuildVoiceStates  bool
	IntentGuildPresences    bool */
	IntentGuildMessages     bool

	// IntentGuildMessageReactions 包含
	// - MESSAGE_REACTION_ADD
	// - MESSAGE_REACTION_REMOVE
	IntentGuildMessageReactions bool

	/* 不支持 IntentGuildMessageTyping     bool */
	IntentDirectMessages         bool
	/* 不支持  IntentDirectMessageReactions bool
	IntentDirectMessageTyping    bool */

	// IntentQQ 包含
	// - C2C_MESSAGE_CREATE
	// - GROUP_AT_MESSAGE_CREATE
	IntentQQ bool

	IntentInteraction bool // 互动事件
	IntentAudit       bool // 审核事件
	// IntentForum 论坛事件
	//  - THREAD_CREATE     // 当用户创建主题时
	//  - THREAD_UPDATE     // 当用户更新主题时
	//  - THREAD_DELETE     // 当用户删除主题时
	//  - POST_CREATE       // 当用户创建帖子时
	//  - POST_DELETE       // 当用户删除帖子时
	//  - REPLY_CREATE      // 当用户回复评论时
	//  - REPLY_DELETE      // 当用户回复评论时
	//  - FORUM_PUBLISH_AUDIT_RESULT      // 当用户发表审核通过时
	IntentForum bool // 论坛事件

	// IntentAudio
	//  - AUDIO_START           // 音频开始播放时
	//  - AUDIO_FINISH          // 音频播放结束时
	IntentAudio          bool // 音频机器人事件
	IntentGuildAtMessage bool // 只接收@消息事件
}

var SettingPath = "setting"
var AllSetting *Setting = &Setting{}
var NullSetting = &Setting{}

func AllSettings() *Setting {
	_, err := toml.DecodeFile("setting/setting.toml", AllSetting)
	if err != nil {
		return AllSetting
	}
	return AllSetting
}

func ReadSetting() Setting {
	tomlData := `QQ = 0 # 机器人QQ
AppId = 0 # 机器人AppId
Token = "" # 机器人Token
AppSecret = "" # 机器人Secret
IsPrivate = false # 是否是私域
IsSandBox = false # 是否是沙箱

# IntentGuilds 包含
# - GUILD_CREATE
# - GUILD_UPDATE
# - GUILD_DELETE
# - GUILD_ROLE_CREATE
# - GUILD_ROLE_UPDATE
# - GUILD_ROLE_DELETE
# - CHANNEL_CREATE
# - CHANNEL_UPDATE
# - CHANNEL_DELETE
# - CHANNEL_PINS_UPDATE
IntentGuilds = false

# IntentGuildMembers 包含
# - GUILD_MEMBER_ADD
# - GUILD_MEMBER_UPDATE
# - GUILD_MEMBER_REMOVE
IntentGuildMembers = false

# IntentGuildMessages // 仅私域能够设置
# - MESSAGE_CREATE
# - MESSAGE_DELETE
IntentGuildMessages = false

# IntentGuildMessageReactions 包含
# - MESSAGE_REACTION_ADD
# - MESSAGE_REACTION_REMOVE
IntentGuildMessageReactions = false

IntentDirectMessages = false

# IntentQQ 包含
# - C2C_MESSAGE_CREATE
# - GROUP_AT_MESSAGE_CREATE
# - GROUP_ADD_ROBBOT
# - GROUP_DEL_ROBBOT
# - GROUP_MSG_REJECT
# - GROUP_MSG_RECEIVE
# - FRIEND_ADD
# - FRIEND_DEL
# - C2C_MSG_REJECT
# - C2C_MSG_RECEIVE
IntentQQ = false

IntentInteraction = false # 互动事件
IntentAudit = false # 审核事件
#  IntentForum 论坛事件
#   - THREAD_CREATE     // 当用户创建主题时
#   - THREAD_UPDATE     // 当用户更新主题时
#   - THREAD_DELETE     // 当用户删除主题时
#   - POST_CREATE       // 当用户创建帖子时
#   - POST_DELETE       // 当用户删除帖子时
#   - REPLY_CREATE      // 当用户回复评论时
#   - REPLY_DELETE      // 当用户回复评论时
#   - FORUM_PUBLISH_AUDIT_RESULT      // 当用户发表审核通过时
IntentForum = false #  论坛事件 仅私域

#  IntentAudio
#   - AUDIO_START           // 音频开始播放时
#   - AUDIO_FINISH          // 音频播放结束时
IntentAudio = false #  音频机器人事件
IntentGuildAtMessage = false #  只接收@消息事件
	`
	if !PathExists(SettingPath) {
		if err := os.MkdirAll(SettingPath, 0777); err != nil {
			log.Warnf("failed to mkdir")
			return *AllSetting
		}
	}
	_, err := os.Stat(fmt.Sprintf("%s/setting.toml", SettingPath))
	if err != nil {
		err = os.WriteFile(fmt.Sprintf("%s/setting.toml", SettingPath), []byte(tomlData), 0644)
		if err != nil {
			log.Warn(err)
			log.Warn("生成配置文件 conf.toml 失败,请稍后重试。")
			return *AllSetting
		}
		log.Warn("已生成配置文件 conf.toml ,请修改后重新启动程序。")
		log.Info("该程序将于5秒后退出！")
		time.Sleep(time.Second * 5)
		os.Exit(1)
	}
	AllSetting = AllSettings()
	return *AllSetting
}

func PathExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil || os.IsExist(err)
}
