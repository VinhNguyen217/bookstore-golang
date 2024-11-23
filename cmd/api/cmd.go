package api

import (
	"book-store/conf"
	"book-store/connection"
	"book-store/log"
	"book-store/repository"
	"book-store/router"
	"book-store/service"
	"book-store/utils"
	"context"
	"fmt"
	"github.com/samber/do"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "api",
	Short: "api",
	Long:  `api`,
	Run: func(cmd *cobra.Command, args []string) {
		startApi()
	},
}

func startApi() {
	injector := do.New()
	defer func() {
		_ = injector.Shutdown()
	}()
	conf.Inject(injector)
	utils.Inject(injector)
	connection.Inject(injector)
	service.Inject(injector)
	repository.Inject(injector)

	r, err := router.InitRouter(injector)
	if err != nil {
		panic(err)
	}

	cf := do.MustInvoke[*conf.Config](injector)
	addr := fmt.Sprintf(":%v", cf.ApiService.Port)
	log.Infow(context.Background(), fmt.Sprintf("start api server at %v", addr))
	_ = r.Run(addr)
}
