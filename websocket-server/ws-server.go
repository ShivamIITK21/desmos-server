package websocketserver

import (
	"fmt"
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
    ReadBufferSize: 1024,
    WriteBufferSize: 1024,
}

type WSConnection struct{
    Conn    *websocket.Conn
    Mut     sync.RWMutex 
}

var GConn *websocket.Conn

func NewWebSockConnection() *WSConnection{
    wsconn := WSConnection{}
    http.HandleFunc("/echo", func(w http.ResponseWriter, r *http.Request){
        upgrader.CheckOrigin = func(r *http.Request) bool { return true }
        conn, err := upgrader.Upgrade(w, r, nil) 
        if err != nil {
            log.Fatalln(err)
        }
        wsconn.Conn = conn
        GConn = conn
        
        for{
            msgType, msg, err := conn.ReadMessage()
            if err != nil {
                return
            }

            // Print the message to the console
            fmt.Printf("%s sent: %s\n", conn.RemoteAddr(), string(msg))

            // Write message back to browser
            if err = conn.WriteMessage(msgType, msg); err != nil {
                return
            }
        }

    })
    return &wsconn
}

func (*WSConnection) Run(port string){
    fmt.Printf("Starting the websocket connection at %s", port)
    http.ListenAndServe(port, nil)
}
