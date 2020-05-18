package studentHttp

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/erDong01/micro-kit/cmd/artisan/http/errGroup"
	student "github.com/erDong01/micro-kit/cmd/rxstudent/bootstrap"
	"github.com/erDong01/micro-kit/cmd/rxstudent/route"
	"github.com/erDong01/micro-kit/internal/core"
	"time"
)

func Start() *http.Server {
	httpServer := &http.Server{
		Addr:         ":5003",
		Handler:      studentRoute(),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	errGroup.ErrGroup.Go(func() error {
		student.App().SetPort(5003)
		fmt.Println(core.New().GetPort())
		defer core.Close()
		return httpServer.ListenAndServe()
	})
	return httpServer
}
func studentRoute() http.Handler {
	e := gin.Default()
	route.Init(e) //注册路由
	return e
}
