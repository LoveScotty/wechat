package message

import "encoding/xml"

// MsgType 基本消息类型
type MsgType string

// EventType 事件类型
type EventType string

// InfoType 第三方平台授权事件类型
type InfoType string

const (
	MsgTypeText       MsgType = "text"       // 文本消息
	MsgTypeImage      MsgType = "image"      // 图片消息
	MsgTypeVoice      MsgType = "voice"      // 语音消息
	MsgTypeVideo      MsgType = "video"      // 视频消息
	MsgTypeShortVideo MsgType = "shortvideo" // 小视频消息
	MsgTypeLocation   MsgType = "location"   // 地理位置消息
	MsgTypeLink       MsgType = "link"       // 链接消息
	MsgTypeMusic      MsgType = "music"      // 音乐消息
	MsgTypeNews       MsgType = "news"       // 图文消息
	MsgTypeEvent      MsgType = "event"      // 事件推送消息
)

func (m MsgType) String() string {
	return string(m)
}

const (
	EventTypeSubScribe   EventType = "subscribe"   // 关注事件
	EventTypeUnSubScribe EventType = "unsubscribe" // 取关事件
	EventTypeScan        EventType = "SCAN"        // 扫描带参数二维码事件
	EventTypeLocation    EventType = "LOCATION"    // 上报地理位置事件
	EventTypeClick       EventType = "CLICK"       // 自定义菜单点击事件
	EventTypeView        EventType = "VIEW"        // 点击菜单跳转链接时的事件
)

func (e EventType) String() string {
	return string(e)
}

const (
	InfoTypeVerifyTicket     InfoType = "component_verify_ticket" // 验证票据推送
	InfoTypeAuthorized       InfoType = "authorized"              // 授权
	InfoTypeUpdateAuthorized InfoType = "updateauthorized"        // 更新授权
	InfoTypeUnAuthorized     InfoType = "unauthorized"            // 取消授权
)

func (i InfoType) String() string {
	return string(i)
}

// CDATA  使用该类型,在序列化为 xml 文本时文本会被解析器忽略
type CDATA string

// MarshalXML 实现自己的序列化方法
func (c CDATA) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(struct {
		string `xml:",cdata"`
	}{string(c)}, start)
}

func (c CDATA) String() string {
	return string(c)
}

// 消息中通用的结构
type CommonMessage struct {
	XMLName      xml.Name `xml:"xml"`
	ToUserName   CDATA    `xml:"ToUserName"`
	FromUserName CDATA    `xml:"FromUserName"`
	CreateTime   int64    `xml:"CreateTime"`
	MsgType      MsgType  `xml:"MsgType"`
}

type Message struct {
	CommonMessage

	//基本消息
	MsgID         int64  `xml:"MsgId"` //其他消息推送过来是MsgId
	TemplateMsgID int64  `xml:"MsgID"` //模板消息推送成功的消息是MsgID
	Content       string `xml:"Content"`
	MediaID       string `xml:"MediaId"`
	ThumbMediaID  string `xml:"ThumbMediaId"`
	Title         string `xml:"Title"`
	Description   string `xml:"Description"`
	URL           string `xml:"Url"`

	//事件相关
	Event    EventType `xml:"Event"`
	EventKey string    `xml:"EventKey"`
	Ticket   string    `xml:"Ticket"`
	Status   string    `xml:"Status"`

	// 第三方平台相关
	InfoType                     InfoType `xml:"InfoType"`
	AppID                        string   `xml:"AppId"`
	ComponentVerifyTicket        string   `xml:"ComponentVerifyTicket"`
	AuthorizerAppid              string   `xml:"AuthorizerAppid"`
	AuthorizationCode            string   `xml:"AuthorizationCode"`
	AuthorizationCodeExpiredTime int64    `xml:"AuthorizationCodeExpiredTime"`
	PreAuthCode                  string   `xml:"PreAuthCode"`
}
