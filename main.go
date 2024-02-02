package main

import(
    "net/http"
    "github.com/ShivamIITK21/desmos-server/http-listener"
    "github.com/ShivamIITK21/desmos-server/websocket-server"
)

func main(){
    httpListener := httplistener.NewHttpListener()
    websocketserver.NewWebSockConnection()
    go http.ListenAndServe(":8081", nil)
    go httpListener.Run(":8080")
    for {}
}
