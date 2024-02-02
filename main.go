package main

import(
    "github.com/ShivamIITK21/desmos-server/http-listener"
    "github.com/ShivamIITK21/desmos-server/websocket-server"
)

func main(){
    httpListener := httplistener.NewHttpListener()
    wsconn := websocketserver.NewWebSockConnection()
    go wsconn.Run(":8081") 
    go httpListener.Run(":8080")
    for {}
}
