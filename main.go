package main

import (
	_ "Outsourcing/docs"
	"Outsourcing/internal/router"
	"Outsourcing/pkg/color"
	"Outsourcing/pkg/config"
	"context"
	"fmt"
	"go.uber.org/zap"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// https://patorjk.com/software/taag/#p=display&f=ANSI%20Shadow&t=weichuang-gin-api
const logo = `

██╗   ██╗██╗   ██╗    ███████╗██╗  ██╗███████╗███╗   ██╗ ██████╗ 
╚██╗ ██╔╝██║   ██║    ██╔════╝██║  ██║██╔════╝████╗  ██║██╔════╝ 
 ╚████╔╝ ██║   ██║    ███████╗███████║█████╗  ██╔██╗ ██║██║  ███╗
  ╚██╔╝  ██║   ██║    ╚════██║██╔══██║██╔══╝  ██║╚██╗██║██║   ██║
   ██║   ╚██████╔╝    ███████║██║  ██║███████╗██║ ╚████║╚██████╔╝
   ╚═╝    ╚═════╝     ╚══════╝╚═╝  ╚═╝╚══════╝╚═╝  ╚═══╝ ╚═════╝ 

                                                                 
 
`

// swagger 页面地址:   http://localhost:8080/swagger/index.html
func main() {
	fmt.Println(color.Red(logo))

	s := router.NewHttpServer()

	server := &http.Server{
		Addr:    ":" + config.Get().ServerConfig.Port,
		Handler: s,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			zap.S().Errorf("listen:%s", err)
		}
	}()

	quit := make(chan os.Signal)

	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		zap.S().Errorf("shut down err:%s", err)
	}

	select {
	case <-ctx.Done():
		zap.L().Info("5秒超时退出")
	}

	zap.L().Info("服务器关闭")

}
