package httplistener

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CORSMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
        c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
        c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, token, accept, origin, Cache-Control, X-Requested-With")
        c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

        if c.Request.Method == "OPTIONS" {
            c.AbortWithStatus(204)
            return
        }

        c.Next()
    }
}

func pingHandler() gin.HandlerFunc{
    return func(c *gin.Context){
        c.JSON(http.StatusOK, gin.H{"ok":"server is up"})
    }
}

func addExpressionHandler() gin.HandlerFunc {
    return func(c *gin.Context){
        id := c.Query("id")
        exp := c.Query("exp")
        log.Printf("got req to put %s as id %s", exp, id)
        c.JSON(http.StatusOK, gin.H{"response": "added exp"})
    }
}

func removeExpHandler() gin.HandlerFunc {
    return func(c *gin.Context){
        id := c.Query("id")
        log.Printf("got req to remove %s", id)
        c.JSON(http.StatusOK, gin.H{"response": "removed exp"})
    }
}

func NewHttpListener() *gin.Engine{
    router := gin.New()
    router.Use(gin.Logger())
    router.Use(CORSMiddleware())
    router.GET("/ping", pingHandler())
    router.GET("/add", addExpressionHandler())
    router.GET("/remove", removeExpHandler())
    return router
}



