package logic

import (
	"Gim/internal/server"
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
	"gopkg.in/fatih/set.v0"
	"gorm.io/gorm"
)

type MessageInfo struct {
	gorm.Model
	Sender   string
	Receiver string
	Type     string // private / group / broadcast
	Media    string // text / image / file
	Content  string
}

func (table *MessageInfo) TableName() string {
	return "message_info"
}

func InitMessageTable() {
	server.DB.AutoMigrate(&MessageInfo{})
}

type Node struct {
	Conn      *websocket.Conn
	Messages  chan []byte
	GroupSets set.Interface
}

var clientMap map[string]*Node = make(map[string]*Node)
var clientMapLock sync.RWMutex

func Chat(writer http.ResponseWriter, request *http.Request) {
	query := request.URL.Query()
	sender := query.Get("sender")
	sender = "test"
	// receiver := query.Get("receiver")
	// msgType := query.Get("Type")
	// media := query.Get("media")
	// content := query.Get("content")

	conn, err := (&websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}).Upgrade(writer, request, nil)
	if err != nil {
		fmt.Println("Upgrade failed: ", err)
		return
	}

	client := &Node{
		Conn:      conn,
		Messages:  make(chan []byte, 4096),
		GroupSets: set.New(set.ThreadSafe),
	}

	clientMapLock.Lock()
	clientMap[sender] = client
	clientMapLock.Unlock()

	go client.read()
	go client.write()

	sendMessage(sender, []byte("Welcome to the chat room!"))
}

func (client *Node) read() {
	for {
		_, message, err := client.Conn.ReadMessage()
		if err != nil {
			fmt.Println("Read message failed: ", err)
			return
		}
		broadcastMessage(message)
	}
}
func (client *Node) write() {
	for {
		select {
		case message := <-client.Messages:
			if err := client.Conn.WriteMessage(websocket.TextMessage, message); err != nil {
				fmt.Println("Write message failed: ", err)
				return
			}
		}
	}
}

var udpsendChan chan []byte = make(chan []byte, 4096)

func broadcastMessage(msg []byte) {
	udpsendChan <- msg
}

func init() {
	go udpSend()
	go udpRecv()
}

func udpSend() {
	conn, err := net.DialUDP("udp", nil, &net.UDPAddr{
		IP:   net.IPv4(127, 0, 0, 1),
		Port: 8080,
	})
	if err != nil {
		fmt.Println("UDP dial failed: ", err)
		return
	}
	defer conn.Close()

	for {
		msg := <-udpsendChan
		_, err := conn.Write(msg)
		if err != nil {
			fmt.Println("UDP send failed: ", err)
			return
		}
	}
}

func udpRecv() {
	conn, err := net.ListenUDP("udp", &net.UDPAddr{
		Port: 8080,
	})
	if err != nil {
		fmt.Println("UDP listen failed: ", err)
		return
	}
	defer conn.Close()

	for {
		buf := make([]byte, 4096)
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("UDP receive failed: ", err)
			return
		}
		dispatchMessage(buf[:n])
	}
}

func dispatchMessage(data []byte) {
	var msg MessageInfo
	if err := json.Unmarshal(data, &msg); err != nil {
		fmt.Println("Unmarshal failed: ", err)
		return
	}

	switch msg.Type {
	case "private":
		sendMessage(msg.Receiver, data)
		// case "group":
		// 	sendGroupMessage(msg.Receiver, data)
		// case "broadcast":
		// 	sendBroadcastMessage(data)
	}
}

func sendMessage(receiver string, message []byte) {
	clientMapLock.RLock()
	client, ok := clientMap[receiver]
	clientMapLock.RUnlock()

	if ok {
		client.Messages <- message
	}
}
