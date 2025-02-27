package websocket

import (
	"flag"
	"fmt"
	"log"
	"net/url"

	"github.com/gorilla/websocket"
)

const addr string = "api.exchange.cryptomkt.com"


// wsManager deals with the server communication, it sends and recieves data
// the way to use it is to snd via its send channel and to recieve in a loop
// via its rcv channel. creation and connection are separated. closable
type wsManager struct {
	streamPath string
	conn       *websocket.Conn
	snd        chan []byte
	rcv        chan []byte
	isOpen     bool
}

func newWSManager(path string) *wsManager {
	return &wsManager{
		streamPath: path,
		snd:        make(chan []byte, 1),
		rcv:        make(chan []byte, 1),
		isOpen:     false,
	}
}

func (ws *wsManager) connect() error {
	flag.Parse()
	log.SetFlags(0)

	u := url.URL{Scheme: "wss", Host: addr, Path: ws.streamPath}
	log.Printf("connecting to %s", u.String())

	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		return fmt.Errorf("dial: %v", err)
	}
	ws.conn = c

	go ws.rcvLoop()
	go ws.sndLoop()

	ws.isOpen = true
	return nil
}

func (ws *wsManager) close() {
	ws.isOpen = false
	close(ws.snd)
}

func (ws *wsManager) sndLoop() {
	for msg := range ws.snd {
		err := ws.conn.WriteMessage(websocket.TextMessage, msg)
		if err != nil {
			fmt.Println("write:", err)
			return
		}
	}
	// send close msg to server
	err := ws.conn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
	if err != nil {
		fmt.Println("write close:", err)
		return
	}
}

// TODO: A close after a timeout since the send of the close message. => close even if the close response is not recieved
func (ws *wsManager) rcvLoop() {
	defer close(ws.rcv)
	for {
		_, message, err := ws.conn.ReadMessage()
		if err != nil {
			fmt.Println("read:", err)
			ws.conn.Close()
			return
		}
		ws.rcv <- message
	}
}
