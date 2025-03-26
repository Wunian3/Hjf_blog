package api_chat

import (
	"blog_server/global"
	"blog_server/models"
	"blog_server/models/ctype"
	"blog_server/models/res"
	"blog_server/utils"
	"encoding/json"
	"fmt"
	"github.com/DanPlayer/randomname"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
	"net/http"
	"strings"
	"time"
)

type ChatUser struct {
	Conn     *websocket.Conn
	NickName string `json:"nick_name"`
	Avatar   string `json:"avatar"`
}

var ConnGroupMap = map[string]ChatUser{}

const (
	InRoomMsg  ctype.MsgType = 1
	TextMsg    ctype.MsgType = 2
	ImageMsg   ctype.MsgType = 3
	VioceMsg   ctype.MsgType = 4
	VideoMsg   ctype.MsgType = 5
	SystemMsg  ctype.MsgType = 6
	OutRoomMsg ctype.MsgType = 7
)

type GroupRequest struct {
	Content string        `json:"content"`  // 聊天的内容
	MsgType ctype.MsgType `json:"msg_type"` // 聊天类型
}
type GroupRes struct {
	NickName    string        `json:"nick_name"` // 前端自己生成
	Avatar      string        `json:"avatar"`    // 头像
	MsgType     ctype.MsgType `json:"msg_type"`  // 聊天类型
	Content     string        `json:"content"`   // 聊天的内容
	OnLineCount int           `json:"online_count"`
	Date        time.Time     `json:"date"` // 消息的时间
}

func (ApiChat) ChatGroup(c *gin.Context) {
	var upGrader = websocket.Upgrader{ // true放行，false拦截
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
	// websocket
	conn, err := upGrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}
	addr := conn.RemoteAddr().String()
	nickName := randomname.GenerateName()
	nickNameFirst := string([]rune(nickName)[0])
	avatar := fmt.Sprintf("uploads/chat_avatar/%s.png", nickNameFirst)

	chatUser := ChatUser{
		Conn:     conn,
		NickName: nickName,
		Avatar:   avatar,
	}
	ConnGroupMap[addr] = chatUser
	// 需要生成昵称，昵称首字关联头像地址
	logrus.Infof("%s %s 连接成功", addr, chatUser.NickName)
	for {
		_, p, err := conn.ReadMessage()
		if err != nil {
			SendGroupMsg(conn, GroupRes{
				NickName:    chatUser.NickName,
				Avatar:      chatUser.Avatar,
				MsgType:     OutRoomMsg,
				Content:     fmt.Sprintf("%s 离开聊天室", chatUser.NickName),
				OnLineCount: len(ConnGroupMap) - 1,
				Date:        time.Now(),
			})
			break
		}
		var request GroupRequest
		err = json.Unmarshal(p, &request)
		if err != nil {
			SendMsg(addr, GroupRes{
				NickName:    chatUser.NickName,
				Avatar:      chatUser.Avatar,
				MsgType:     SystemMsg,
				OnLineCount: len(ConnGroupMap),
				Content:     "参数绑定失败",
			})
			continue
		}
		switch request.MsgType {
		case TextMsg:
			if strings.TrimSpace(request.Content) == "" {
				SendMsg(addr, GroupRes{
					NickName:    chatUser.NickName,
					Avatar:      chatUser.Avatar,
					MsgType:     SystemMsg,
					Content:     "消息不能为空",
					OnLineCount: len(ConnGroupMap),
				})
				continue
			}
			SendGroupMsg(conn, GroupRes{
				NickName:    chatUser.NickName,
				Avatar:      chatUser.Avatar,
				Content:     request.Content,
				MsgType:     TextMsg,
				Date:        time.Now(),
				OnLineCount: len(ConnGroupMap),
			})
		case InRoomMsg:
			SendGroupMsg(conn, GroupRes{
				NickName:    chatUser.NickName,
				Avatar:      chatUser.Avatar,
				Content:     fmt.Sprintf("%s 进入聊天室", chatUser.NickName),
				Date:        time.Now(),
				OnLineCount: len(ConnGroupMap),
			})
		default:
			SendMsg(addr, GroupRes{
				NickName:    chatUser.NickName,
				Avatar:      chatUser.Avatar,
				MsgType:     SystemMsg,
				OnLineCount: len(ConnGroupMap),
				Content:     "消息类型错误",
			})
		}

	}
	defer conn.Close()
	delete(ConnGroupMap, addr)
}

// SendGroupMsg 群聊功能
func SendGroupMsg(conn *websocket.Conn, response GroupRes) {
	byteData, _ := json.Marshal(response)
	_addr := conn.RemoteAddr().String()
	ip, addr := getIPAndAddr(_addr)

	global.DB.Create(&models.ChatModel{
		NickName: response.NickName,
		Avatar:   response.Avatar,
		Content:  response.Content,
		IP:       ip,
		Addr:     addr,
		IsGroup:  true,
		MsgType:  response.MsgType,
	})
	for _, chatUser := range ConnGroupMap {
		chatUser.Conn.WriteMessage(websocket.TextMessage, byteData)
	}
}

// SendMsg 给某个用户发消息
func SendMsg(_addr string, response GroupRes) {
	byteData, _ := json.Marshal(response)
	chatUser := ConnGroupMap[_addr]
	ip, addr := getIPAndAddr(_addr)
	global.DB.Create(&models.ChatModel{
		NickName: response.NickName,
		Avatar:   response.Avatar,
		Content:  response.Content,
		IP:       ip,
		Addr:     addr,
		IsGroup:  false,
		MsgType:  response.MsgType,
	})
	chatUser.Conn.WriteMessage(websocket.TextMessage, byteData)
}

func getIPAndAddr(_addr string) (ip string, addr string) {
	addrList := strings.Split(_addr, ":")
	ip = addrList[0]
	addr = utils.GetAddr(ip)
	return ip, addr
}
