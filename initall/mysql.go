package initall

import (
	"fmt"
	"hello/global"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func Mysql() {
	if db, err := gorm.Open("mysql", "root:root@(127.0.0.1:3306)/ggbeng?charset=utf8&parseTime=True&loc=Local"); err != nil {
		fmt.Println("数据库连接失败")
	} else {
		db.SingularTable(true)
		global.DB = db
	}
	// defer db.Close()
}
