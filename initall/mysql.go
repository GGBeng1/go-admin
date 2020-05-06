package initall

import (
	"fmt"
	"hello/global"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func Mysql() {
	admin := global.Server.Mysql
	if db, err := gorm.Open("mysql", admin.Username+":"+admin.Password+"@("+admin.Path+")/"+admin.Dbname+"?"+admin.Config); err != nil {
		fmt.Println("数据库连接失败")
	} else {
		db.SingularTable(true)
		global.DB = db
		global.DB.DB().SetMaxIdleConns(admin.MaxIdleConns)
		global.DB.DB().SetMaxOpenConns(admin.MaxOpenConns)
	}
	// defer db.Close()
}
