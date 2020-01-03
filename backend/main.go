package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/lawtech0902/go_gin_blog/backend/middleware/logger"
	"github.com/lawtech0902/go_gin_blog/backend/pkg/setting"
	"github.com/lawtech0902/go_gin_blog/backend/routers"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	gin.SetMode(setting.ServerInfo.RunMode)
	r := routers.InitRouter()
	
	defer logger.CloseLogFile()
	
	s := &http.Server{
		Addr:           setting.ServerInfo.ServerAddr,
		Handler:        r,
		ReadTimeout:    setting.ServerInfo.ReadTimeout,
		WriteTimeout:   setting.ServerInfo.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}
	
	go func() {
		if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %v\n", err)
		}
	}()
	
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutdown Server ...")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	if err := s.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	
	log.Println("Server exiting")
}
