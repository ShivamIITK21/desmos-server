package main

import(
    "github.com/ShivamIITK21/desmos-server/http-listener"
)

func main(){
    httpListener := httplistener.NewHttpListener()
    httpListener.Run(":8080")
}
