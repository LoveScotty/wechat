package server

import (
	"net/http"

	"scotty/wechat/mp"
	"scotty/wechat/mp/message"
)

type Server struct {
	*mp.Config
	Writer  http.ResponseWriter
	Request *http.Request

	openID string

	messageHandler func(message.Message) *message.Reply

	RequestRawXMLMsg  []byte
	RequestMsg        message.Message
	ResponseRawXMLMsg []byte
	ResponseMsg       interface{}

	isSafeMode bool
	random     []byte
	nonce      string
	timestamp  int64
}

func NewServer(setters ...Setter) (*Server, error) {
	s := new(Server)
	for _, setter := range setters {
		if err := setter(s); err != nil {
			return nil, err
		}
	}
	return s, nil
}
