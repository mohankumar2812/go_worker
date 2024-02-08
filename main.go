package main

import (
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/mohankumar2812/go_worker/model"
	"github.com/mohankumar2812/go_worker/worker"
)


func main() {

	r := gin.Default()

	inChannel := make(chan model.Input)
	outChannel := make(chan model.Output)
	var wg sync.WaitGroup
	
	r.GET("/getData", func(ctx *gin.Context) {
		var inp model.Input

		if err := ctx.ShouldBindJSON(&inp); err != nil {
			ctx.Status(http.StatusBadGateway)
			return
		}
	
		wg.Add(1)

		go worker.StartWorker(inChannel, outChannel, &wg)

		inChannel <- inp
		response := <- outChannel

		ctx.JSON(http.StatusOK, response)
	})

	r.Run(":8080")

	wg.Wait()

}