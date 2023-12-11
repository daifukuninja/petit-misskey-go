package websocket

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/url"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/daifukuninja/petit-misskey-go/model/misskey"
	"github.com/fatih/color"
	"github.com/google/uuid"
	"github.com/sacOO7/gowebsocket"
)

type (
	Client struct {
		baseUrl     string
		accessToken string
		writer      io.Writer
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

	// TODO: このパッケージはwebsocketによる通信処理の責務を負う。tmplなどの見た目の部分は別パッケージに移管する
	//go:embed template/note.tmpl
	NoteTmpl string

	//go:embed template/renote.tmpl
	RenoteTmpl string
)

func NewClient(baseUrl string, accessToken string, writeTo io.Writer) *Client {
	return &Client{
		baseUrl:     baseUrl,
		accessToken: accessToken,
		writer:      writeTo,
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
		// TODO: このあたりの描画処理はまるごとwriterへ委譲する
		// Write -> SetContentへ流すViewを作ってClientへDIする
		note := &misskey.Note{}
		if err := json.Unmarshal([]byte(message), &note); err != nil {
			log.Printf("note marshalize error %v", err)
		}
		var data map[string]interface{}
		if note.Body.Body.RenoteID != "" {
			t, err := template.New("note").Parse(RenoteTmpl)
			if err != nil {
				log.Printf("template error: %v", err)
			}
			data = map[string]interface{}{
				"renotedName":     color.HiBlackString(note.Body.Body.User.Name),
				"renotedUsername": color.HiBlackString(note.Body.Body.User.Username),
				"name":            color.HiGreenString(note.Body.Body.Renote.User.Name),
				"username":        color.HiBlueString(note.Body.Body.Renote.User.Username),
				"text":            note.Body.Body.Renote.Text,
				"createdAt":       note.Body.Body.Renote.CreatedAt.Format(time.RFC3339),
			}
			if err := t.Execute(c.writer, data); err != nil {
				log.Printf("template execute error: %v", err)
			}
		} else {
			t, err := template.New("note").Parse(NoteTmpl)
			if err != nil {
				log.Printf("template error: %v", err)
			}
			data = map[string]interface{}{
				"name":      color.HiGreenString(note.Body.Body.User.Name),
				"username":  color.HiBlueString(note.Body.Body.User.Username),
				"text":      note.Body.Body.Text,
				"createdAt": note.Body.Body.CreatedAt.String(),
			}
			if err := t.Execute(c.writer, data); err != nil {
				log.Printf("template execute error: %v", err)
			}
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
