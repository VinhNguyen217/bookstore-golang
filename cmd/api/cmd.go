package api

import (
	"book-store/conf"
	"book-store/connection"
	"book-store/repository"
	"book-store/router"
	"book-store/service"
	"book-store/utils"
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
	fmt.Println("start api server at : 8088")
	_ = r.Run(":8088")
}
