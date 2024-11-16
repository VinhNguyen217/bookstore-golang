package service

import "github.com/samber/do"

func Inject(di *do.Injector) {
	do.Provide(di, newAuthService)
	do.Provide(di, newBillService)
	do.Provide(di, newBookService)
	do.Provide(di, newCartService)
	do.Provide(di, newUserService)
}
