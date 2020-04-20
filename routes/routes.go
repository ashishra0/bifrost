package routes

import (
	"github.com/ashishra0/raspberry-pi-service/httpd/handler"
	"github.com/gin-gonic/gin"
	"golang.org/x/sync/errgroup"
	"log"
	"net/http"
	"os"
	"time"
)

var (
	g errgroup.Group
)

// StartService runs the gin application.
// It is the single point of entry in the application.
func StartTeslaService() http.Handler {
	// PORT for prod environment
	port := "5000"
	if port == "" {
		log.Fatal("$PORT must be set")
	}
	router := gin.Default()
	// To trigger the webhook
	router.GET("/api/power", handler.TriggerWebhook)
	// To save the response to db
	router.POST("/api/power/status", handler.HandlePost)
	// To return the saved response to user
	router.GET("/api/power/status", handler.GetStatus)
	// To handle the case of incorrect route
	router.NoRoute(func(context *gin.Context) {
		context.AbortWithStatus(http.StatusNotFound)
	})
	return router
}


func StartAlfredService() http.Handler {
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("$PORT must be set")
	}
	router := gin.Default()
	// To get the data from hasura
	router.GET("/api/expense", handler.GetExpense)
	// To forward the event response from hasura
	router.POST("/api/expense", handler.PostExpense)
	router.NoRoute(func(context *gin.Context) {
		context.AbortWithStatus(http.StatusNotFound)
	})
	return router
}

func StartAllService() {
	port01 := os.Getenv("PORT01")
	port02 := os.Getenv("PORT02")

	server01 := &http.Server{
		Addr: ":" + port01,
		Handler: StartTeslaService(),
		ReadTimeout: 5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	server02 := &http.Server{
		Addr: ":" + port02,
		Handler: StartAlfredService(),
		ReadTimeout: 5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	g.Go(func() error {
		err := server01.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			log.Fatal(err)
		}
		return err
	})

	g.Go(func() error {
		err := server02.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			log.Fatal(err)
		}
		return err
	})

	if err := g.Wait(); err != nil {
		log.Fatal(err)
	}
}