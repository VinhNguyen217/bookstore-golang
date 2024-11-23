package connection

import (
	"book-store/conf"
	"fmt"
	"github.com/samber/do"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"time"
)

func NewMysqlConnection(di *do.Injector) (*gorm.DB, error) {
	cf := do.MustInvoke[*conf.Config](di)

	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True",
		cf.MYSQL.User, cf.MYSQL.Password,
		cf.MYSQL.Host, cf.MYSQL.Port,
		cf.MYSQL.DBName,
	)
	dbOrm, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return nil, err
	}

	//err = dbOrm.AutoMigrate(
	//	&model.User{},
	//	&model.Book{},
	//	&model.Bill{},
	//	&model.BillDetail{},
	//	&model.Cart{},
	//)
	//if err != nil {
	//	return nil, err
	//}

	db, err := dbOrm.DB()
	if err != nil {
		return nil, err
	}

	db.SetMaxIdleConns(1)
	db.SetMaxOpenConns(10)
	db.SetConnMaxIdleTime(10 * time.Minute)
	db.SetConnMaxLifetime(30 * time.Minute)

	return dbOrm, nil
}
