package commonService

import (
	"encoding/json"
	_logUtils "github.com/easysoft/zagent/internal/pkg/lib/log"
	serverConst "github.com/easysoft/zagent/internal/server/utils/const"
	"github.com/kataras/iris/v12/websocket"
	"github.com/kataras/neffos"
)

var (
	wsConn *neffos.Conn
)

type WebSocketService struct {
}

func NewWebSocketService() *WebSocketService {
	return &WebSocketService{}
}

func (s *WebSocketService) UpdateTask(taskId uint, msg string) {
	data := map[string]interface{}{"action": serverConst.TaskUpdate, "taskId": taskId, "msg": msg}
	s.Broadcast(serverConst.WsNamespace, serverConst.WsDefaultRoom, serverConst.WsEvent, data)

}

func (s *WebSocketService) Broadcast(namespace, room, event string, data interface{}) {
	bytes, _ := json.Marshal(data)

	if wsConn == nil {
		_logUtils.Warnf("WebSocket connection not init")
		return
	}

	wsConn.Server().Broadcast(nil, websocket.Message{
		Namespace: namespace,
		Room:      room,
		Event:     event,
		Body:      bytes,
	})
}

func (s *WebSocketService) SetConn(conn *neffos.Conn) {
	wsConn = conn
}
