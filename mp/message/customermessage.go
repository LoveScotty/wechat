package message

// 客服消息
type CustomerMessage struct {
	ToUser          string                `json:"touser"`                    //接收者OpenID
	Msgtype         MsgType               `json:"msgtype"`                   //客服消息类型
	Text            *MediaText            `json:"text,omitempty"`            //可选
	Image           *MediaResource        `json:"image,omitempty"`           //可选
	Voice           *MediaResource        `json:"voice,omitempty"`           //可选
	Video           *MediaVideo           `json:"video,omitempty"`           //可选
	Music           *MediaMusic           `json:"music,omitempty"`           //可选
	News            *MediaNews            `json:"news,omitempty"`            //可选
	Mpnews          *MediaResource        `json:"mpnews,omitempty"`          //可选
	Wxcard          *MediaWxcard          `json:"wxcard,omitempty"`          //可选
	Msgmenu         *MediaMsgmenu         `json:"msgmenu,omitempty"`         //可选
	Miniprogrampage *MediaMiniprogrampage `json:"miniprogrampage,omitempty"` //可选
}

// 文本消息结构体构造方法
func NewCustomerTextMessage(toUser, text string) *CustomerMessage {
	return &CustomerMessage{
		ToUser:  toUser,
		Msgtype: MsgTypeText,
		Text: &MediaText{
			text,
		},
	}
}

// 图片消息的构造方法
func NewCustomerImgMessage(toUser, mediaID string) *CustomerMessage {
	return &CustomerMessage{
		ToUser:  toUser,
		Msgtype: MsgTypeImage,
		Image: &MediaResource{
			mediaID,
		},
	}
}

// 语音消息的构造方法
func NewCustomerVoiceMessage(toUser, mediaID string) *CustomerMessage {
	return &CustomerMessage{
		ToUser:  toUser,
		Msgtype: MsgTypeVoice,
		Voice: &MediaResource{
			mediaID,
		},
	}
}

// 图文消息的构造方法
func NewCustomerNewsMessage(toUser string, articleList []MediaArticles) *CustomerMessage {
	return &CustomerMessage{
		ToUser:  toUser,
		Msgtype: MsgTypeNews,
		News:    &MediaNews{Articles: articleList},
	}
}

// 图文消息的内容的文章列表中的单独一条构造
func NewNewsMessage(title, description, picUrl, url string) MediaArticles {
	return MediaArticles{
		Title:       title,
		Description: description,
		URL:         url,
		Picurl:      picUrl,
	}
}

// 文本消息的文字
type MediaText struct {
	Content string `json:"content"`
}

// 消息使用的永久素材id
type MediaResource struct {
	MediaID string `json:"media_id"`
}

// 视频消息包含的内容
type MediaVideo struct {
	MediaID      string `json:"media_id"`
	ThumbMediaID string `json:"thumb_media_id"`
	Title        string `json:"title"`
	Description  string `json:"description"`
}

// 音乐消息包括的内容
type MediaMusic struct {
	Title        string `json:"title"`
	Description  string `json:"description"`
	Musicurl     string `json:"musicurl"`
	Hqmusicurl   string `json:"hqmusicurl"`
	ThumbMediaID string `json:"thumb_media_id"`
}

// 图文消息的内容
type MediaNews struct {
	Articles []MediaArticles `json:"articles"`
}

// 图文消息的内容的文章列表中的单独一条
type MediaArticles struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	URL         string `json:"url"`
	Picurl      string `json:"picurl"`
}

// 菜单消息的内容
type MediaMsgmenu struct {
	HeadContent string        `json:"head_content"`
	List        []MsgmenuItem `json:"list"`
	TailContent string        `json:"tail_content"`
}

// 菜单消息的菜单按钮
type MsgmenuItem struct {
	ID      string `json:"id"`
	Content string `json:"content"`
}

// 卡券的id
type MediaWxcard struct {
	CardID string `json:"card_id"`
}

// 小程序消息
type MediaMiniprogrampage struct {
	Title        string `json:"title"`
	Appid        string `json:"appid"`
	Pagepath     string `json:"pagepath"`
	ThumbMediaID string `json:"thumb_media_id"`
}
