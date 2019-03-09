package database

import (
	"fmt"
	"time"

	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"github.com/school/utils"
)

// Database interface for DB operations
type Database interface {
	Write(data interface{}) error
	Read(model interface{}) error
	ReadAll(ptrStructOrTableName, model interface{}) error
	Delete(model interface{}) error
}

type database struct {
	orm.Ormer
}

// Write inserts a new record into db
func (d *database) Write(data interface{}) error {
	if _, err := d.Insert(data); err != nil {
		return err
	}
	return nil
}

// Read return a record from db
// make sure you set primary key fields in the struct, ie: &events{id:1}
// also make sure to provide a pointer
func (d *database) Read(model interface{}) error {
	return d.Ormer.Read(model)
}

// ReadAll return all records related to a table
// table name can be string or struct.
// e.g. QueryTable("user"), QueryTable(&user{}) or QueryTable((*User)(nil)),
// TODO: This must implement pagination
func (d *database) ReadAll(ptrStructOrTableName, model interface{}) error {
	_, err := d.QueryTable(ptrStructOrTableName).All(model)
	return err
}

// Delete removes a record from db
func (d *database) Delete(model interface{}) error {
	_, err := d.Ormer.Delete(model)
	return err
}

// NewDatabase creates a new database instance that uses beego.Ormer as orm
func NewDatabase() Database {
	initDB()
	return &database{orm.NewOrm()}
}

// InitDB register orm driver and database
func initDB() {
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
