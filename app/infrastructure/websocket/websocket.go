package websocket

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/url"
	"os"
	"os/signal"
	"syscall"

	"github.com/daifukuninja/petit-misskey-go/model/misskey"
	"github.com/google/uuid"
	"github.com/sacOO7/gowebsocket"
)

type (
	Client struct {
		baseUrl     string
		accessToken string
	}
	ConnectChannelPayload struct {
		Type string      `json:"type"`
		Body PayloadBody `json:"body"`
	}
	PayloadBody struct {
		Channel string `json:"channel,omitempty"`
		Id      string `json:"id"`
	}
)

var (
	ChannelTypeMain  = "main"
	ChannelTypeHome  = "homeTimeline"
	ChannelTypeLocal = "localTimeline"

	//go:embed template/note.tmpl
	NoteTmpl string
)

func NewClient(baseUrl string, accessToken string) *Client {
	return &Client{
		baseUrl:     baseUrl,
		accessToken: accessToken,
	}
}

func (c *Client) Start() error {
	urlInfo, err := url.Parse(c.baseUrl)
	if err != nil {
		panic(err) // TODO: エラー処理
	}
	wsUrl := fmt.Sprintf("wss://%s/streaming?i=%s", urlInfo.Host, c.accessToken)

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	socket := gowebsocket.New(wsUrl)

	socket.OnConnected = func(socket gowebsocket.Socket) {
		log.Println("Connected to server")
	}

	socket.OnConnectError = func(err error, socket gowebsocket.Socket) {
		log.Println("Received connect error ", err)
	}

	socket.OnTextMessage = func(message string, socket gowebsocket.Socket) {
		// log.Println("Received message " + message)
		note := &misskey.Note{}
		if err := json.Unmarshal([]byte(message), &note); err != nil {
			log.Printf("note marshalize error %v", err)
		}
		t, err := template.New("note").Parse(NoteTmpl)
		if err != nil {
			log.Printf("template error: %v", err)
		}
		if err := t.Execute(os.Stdout, note); err != nil {
			log.Printf("template execute error: %v", err)
		}
	}

	socket.OnDisconnected = func(err error, socket gowebsocket.Socket) {
		log.Println("Disconnected from server ")
		socket.Close()
	}

	socket.Connect()

	uu, err := uuid.NewRandom()
	if err != nil {
		return err
	}
	tlChId := uu.String()

	connectLocalBody := &PayloadBody{
		Channel: ChannelTypeLocal,
		Id:      tlChId,
	}
	homeText, _ := json.Marshal(&ConnectChannelPayload{Type: "connect", Body: *connectLocalBody})
	socket.SendText(string(homeText))

	for {
		<-interrupt
		log.Println("interrupt")

		disconnectBody := &PayloadBody{
			Id: tlChId,
		}
		disconnectText, _ := json.Marshal(&ConnectChannelPayload{Type: "disconnect", Body: *disconnectBody})
		socket.SendText(string(disconnectText))

		// socket.Close()
		return nil
	}
}
