package database

import (
	"fmt"
	"time"

	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"github.com/school/utils"
)

// RegisterDriverAndDatabase _
func RegisterDriverAndDatabase() {
	// register db driver (mysql)
	if err := orm.RegisterDriver("mysql", orm.DRMySQL); err != nil {
		panic(err)
	}
	// enable or disable SQL logger
	orm.Debug = utils.GetAppConfig("OrmDebug", "false") == "true"
	// register database
	orm.RegisterDataBase("default", "mysql", getDbString())
	// set database timezone
	orm.SetDataBaseTZ("default", time.UTC)
}

// build db connection string
func getDbString() string {
	fmt.Println(fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8&tls=%s",
		utils.GetAppConfig("dbuser", "root"),
		utils.GetAppConfig("dbpass", "password"),
		utils.GetAppConfig("dbhost", "localhost"),
		utils.GetAppConfig("dbport", "3306"),
		utils.GetAppConfig("dbname", "school_service_mysql_dev"),
		utils.GetAppConfig("dbusetls", "false"),
	))

	return fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8&tls=%s",
		utils.GetAppConfig("dbuser", "root"),
		utils.GetAppConfig("dbpass", "password"),
		utils.GetAppConfig("dbhost", "localhost"),
		utils.GetAppConfig("dbport", "3306"),
		utils.GetAppConfig("dbname", "school_service_mysql_dev"),
		utils.GetAppConfig("dbusetls", "false"),
	)
}
