package database

import (
	"crypto/tls"
	"fmt"
	"time"

	"github.com/astaxie/beego/orm"
	"github.com/go-sql-driver/mysql"
	"github.com/school/utils"
)

func RegisterDriverAndDatabase(driver string) {
	switch driver {
	case "postgres":
		orm.RegisterDriver(driver, orm.DRPostgres)
	case "mysql":
		orm.RegisterDriver(driver, orm.DRMySQL)
	}

	// enable ssl
	// https://github.com/go-sql-driver/mysql/blob/749ddf1598b47e3cd909414bda735fe790ef3d30/utils.go#L58
	mysql.RegisterTLSConfig("ssl", &tls.Config{
		ServerName: utils.GetAppConfig("dbhost", "localhost"),
	})

	maxIdle := utils.GetAppIntConfig("db.maxidleconnections", 10)
	maxConn := utils.GetAppIntConfig("db.maxconnections", 10)
	maxLifetime := utils.GetAppIntConfig("db.maxlifetimeconnection", 120)

	orm.RegisterDataBase("default", driver, getDbString(driver), maxIdle, maxConn)

	// Apply MaxLifetime
	defaultDB, err := orm.GetDB("default")
	if err != nil {
		panic("Get default DB error: " + err.Error())
	}
	defaultDB.SetConnMaxLifetime(time.Duration(maxLifetime) * time.Second)

	// enable or disable SQL logger
	orm.Debug = utils.GetAppBoolConfig("OrmDebug", false)
	orm.DefaultTimeLoc = time.UTC
}

func getDbString(driver string) string {
	switch driver {
	case "postgres":
		return fmt.Sprintf(
			"user=%s password=%s host=%s port=%s sslmode=disable",
			utils.GetAppConfig("dbuser", "postgres"),
			utils.GetAppConfig("dbpass", "postgres"),
			utils.GetAppConfig("dbhost", "localhost"),
			utils.GetAppConfig("dbport", "5432"),
		)
	case "mysql":
		return fmt.Sprintf(
			"%s:%s@tcp(%s:%s)/%s?charset=utf8&tls=%s",
			utils.GetAppConfig("dbuser", "root"),
			utils.GetAppConfig("dbpass", "password"),
			utils.GetAppConfig("dbhost", "localhost"),
			utils.GetAppConfig("dbport", "3306"),
			utils.GetAppConfig("dbname", "school_service_mysql_dev"),
			utils.GetAppConfig("usetls", "false"),
		)
	default:
		return ""
	}
}
