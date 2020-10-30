package dao

import (
    "github.com/jmoiron/sqlx"
    _ "github.com/go-sql-driver/mysql"
    "fmt"
    "g-admin/config"
)

var _gDB *sqlx.DB

func initMysql(){
    mysqlconf := &config.Conf.Mysql
    dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?%s",
        mysqlconf.Username, mysqlconf.Password, mysqlconf.Path, mysqlconf.Dbname, mysqlconf.Config)
    // 也可以使用MustConnect连接不成功就panic
    db, err := sqlx.Connect("mysql", dsn)
    if err != nil {
        panic(err)
    }
    db.SetMaxOpenConns(mysqlconf.MaxOpenConns)
    db.SetMaxIdleConns(mysqlconf.MaxIdleConns)
    _gDB = db
    return
}

func Init(){
    initMysql()
    return
}
