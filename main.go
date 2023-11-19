package main

import (
	"example/test-cd/api"
	"fmt"
	"sync"

	"github.com/gin-gonic/gin"
)

func main() {
    wg := sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer wg.Done()

		router := gin.New()
		router.Use(gin.Recovery())

		router.ForwardedByClientIP = true
		router.SetTrustedProxies([]string{"127.0.0.1"})
    
		router = RunServer(router)

		fmt.Println("Server is running on port 8080")

		err := router.Run(":8080")
		if err != nil {
			panic(err)
		}
	}()

	wg.Wait()
}

func RunServer(engine *gin.Engine) *gin.Engine {
    HelloAPIHandler := api.CreateHelloAPIImpl()

    version := engine.Group("/api")
    {
        version.GET("/hello", HelloAPIHandler.Hello)
        version.GET("/secret", HelloAPIHandler.Secret)
    }

    return engine
}