package service

import (
    "strings"
    "g-admin/config"
    "github.com/casbin/casbin/v2"
    "github.com/casbin/casbin/v2/util"
    "fmt"
    gormadapter "github.com/casbin/gorm-adapter/v3"
)

var _enforcer *casbin.Enforcer

func initCasbin() {
    mysqlconf := &config.Conf.Mysql
    dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?%s",
        mysqlconf.Username, mysqlconf.Password, mysqlconf.Path, mysqlconf.Dbname, mysqlconf.Config)
    
    a, _ := gormadapter.NewAdapter("mysql", dsn, true)
    e, _ := casbin.NewEnforcer(config.Conf.Casbin.ModelPath, a)
    e.AddFunction("ParamsMatch", ParamsMatchFunc)
    _ = e.LoadPolicy()
    
    _enforcer = e
}

func Casbin() *casbin.Enforcer {
    return _enforcer
}

func GetPolicyPathByAuthorityId(authorityId string) {
    e := Casbin()
    list := e.GetFilteredPolicy(0, authorityId)
    for _, v := range list {
        fmt.Printf("11111111 v=%+v\n", v)
    }
    return
}

func ParamsMatch(fullNameKey1 string, key2 string) bool {
    key1 := strings.Split(fullNameKey1, "?")[0]
    // 剥离路径后再使用casbin的keyMatch2
    return util.KeyMatch2(key1, key2)
}

func ParamsMatchFunc(args ...interface{}) (interface{}, error) {
    name1 := args[0].(string)
    name2 := args[1].(string)
    
    return ParamsMatch(name1, name2), nil
}
