package bot

import (
	"context"
	"encoding/json"
	"net/http"
	"regexp"
	"sync"
	"time"

	"github.com/2mf8/Go-QQ-Client/bot/safe_ws"
	"github.com/2mf8/Go-QQ-Client/openapi"
	"github.com/fanliao/go-promise"

	"github.com/gorilla/websocket"
	log "github.com/sirupsen/logrus"
)

var (
	// RemoteServers key是botId，value是map（key是serverName，value是server）
	RemoteServers  RemoteMap
	forwardServers = make(map[string]*ForwardServer, 0)
	lock           = new(sync.RWMutex)
	SafeForwardMap = NewForwards()
)

type SafeForwards struct {
	Map map[string]*ForwardServer
	RW  *sync.RWMutex
}

type LifeTime struct {
	Time          int64  `json:"time,omitempty"`
	SelfId        int64  `json:"self_id,omitempty"`
	PostType      string `json:"post_type,omitempty"`
	MetaEventType string `json:"meta_event_type"`
	SubType       string `json:"sub_type"`
}

type actionResp struct {
	BotId   int64  `json:"bot_id,omitempty"`
	Ok      bool   `json:"ok,omitempty"`
	Status  any    `json:"status"`
	RetCode int32  `json:"retcode"`
	Data    any    `json:"data"`
	Echo    string `json:"echo"`
}

type WsServer struct {
	*safe_ws.SafeWebSocket        // 线程安全的ws
	wsUrl                  string // 随机抽中的url
	regexp                 *regexp.Regexp
}

type ForwardServer struct {
	Url           string
	Session       *safe_ws.ForwardSafeWebSocket
	Mu            sync.Mutex
	WaitingFrames map[string]*promise.Promise
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func NewForwards() *SafeForwards {
	return &SafeForwards{
		Map: forwardServers,
		RW:  lock,
	}
}

func (f *SafeForwards) SetForward(key string, value *ForwardServer) {
	f.RW.Lock()
	defer f.RW.Unlock()
	f.Map[key] = value
}

func (f *SafeForwards) GetForward(key string) *ForwardServer {
	f.RW.RLock()
	defer f.RW.RUnlock()
	return f.Map[key]
}

func (f *SafeForwards) DeleteForward(key string) {
	f.RW.Lock()
	defer f.RW.Unlock()
	delete(f.Map, key)
}

func UpgradeWebsocket(w http.ResponseWriter, r *http.Request) error {
	_, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		return err
	}
	return nil
}

func ForwradConnect(url string, conn *websocket.Conn) *ForwardServer {
	lt := &LifeTime{
		Time:          time.Now().Unix(),
		PostType:      "meta_event",
		MetaEventType: "lifecycle",
		SubType:       "connect",
	}
	messageHandler := func(messageType int, data []byte) {
	}
	closeHandler := func(code int, message string) {
		SafeForwardMap.GetForward(url).Session.Conn.Close()
		SafeForwardMap.DeleteForward(url)
		log.Infof("正向 WebSocket 已断连，断连 账号_地址 为 %s", url)
	}
	safews := safe_ws.NewForwardSafeWebSocket(conn, messageHandler, closeHandler)
	fs := &ForwardServer{
		Url:           url,
		Session:       safews,
		WaitingFrames: make(map[string]*promise.Promise),
	}
	SafeForwardMap.SetForward(url, fs)
	b, err := json.Marshal(lt)
	if err == nil {
		fs.Mu.Lock()
		defer fs.Mu.Unlock()
		fs.Session.ForwardSend(websocket.TextMessage, b)
	}
	return fs
}

func ConnectUniversal(c context.Context, api openapi.OpenAPI) {
}
