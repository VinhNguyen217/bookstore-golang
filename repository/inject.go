package repository

import "github.com/samber/do"

func Inject(di *do.Injector) {
	do.Provide(di, newBillDetailRepository)
	do.Provide(di, newBillRepository)
	do.Provide(di, newBookRepository)
	do.Provide(di, newCartRepository)
	do.Provide(di, newUserRepository)
}
