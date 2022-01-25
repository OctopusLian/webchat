package controllers

import (
	"webchat/models"

	"github.com/astaxie/beego"
	"github.com/gorilla/websocket"
)

var clients map[int]*websocket.Conn = make(map[int]*websocket.Conn)

var boardcast chan models.Message = make(chan models.Message, 100)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

type WebSocketController struct {
	beego.Controller
}

func (c *WebSocketController) Msg() {
	// 检查是否登陆, 若未登陆则跳转到登陆页面
	suser := c.GetSession("user")
	client, err := upgrader.Upgrade(c.Ctx.ResponseWriter, c.Ctx.Request, nil)
	if err == nil {
		defer client.Close()
		if suser == nil {
			//未登录
			client.WriteJSON(models.Message{Code: 400})
		} else {
			user, _ := suser.(models.User)
			defer func() {
				// 离线
				beego.Debug("offline:", user)
				delete(clients, user.Id)
				boardcast <- models.Message{Type: "offline", User: user, Date: models.Now()}

			}()

			// 在线
			beego.Debug("online:", user)
			clients[user.Id] = client
			boardcast <- models.Message{Type: "online", User: user, Date: models.Now()}
			for {
				var msg models.Message
				err := client.ReadJSON(&msg)
				if err != nil {
					break
				}
				msg.User = user
				msg.Date = models.Now()
				beego.Debug("msg:", msg)
				boardcast <- msg
			}
		}
	}
}

func init() {
	go func() {
		for {
			msg := <-boardcast
			beego.Debug("boardcast:", msg)
			for _, client := range clients {
				msg.Code = 200
				client.WriteJSON(msg)
			}
		}
	}()
}
