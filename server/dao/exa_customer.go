package dao

import (
    "g-admin/dao/model"
    "time"
    "fmt"
    "github.com/jmoiron/sqlx"
)

type ExaCustomerOp struct {
}
var IExaCustomer = ExaCustomerOp{}

func (ExaCustomerOp)CreateExaCustomer(customer *model.ExaCustomer)(err error){
    sqlStr := "insert into exa_customers(created_at, updated_at,,customer_name, customer_phone_data, sys_user_id, sys_user_authority_id) Values (:created_at, :updated_at, :customer_name, :customer_phone_data, :sys_user_id, :sys_user_authority_id)"
    _, err  = _gDB.NamedExec(sqlStr, customer)
    return
}

func (ExaCustomerOp)DeleteExaCustomer(id int64)(err error){
    sqlStr := "update exa_customers set deleted_at = ? where  id = ?"
    _, err  = _gDB.Exec(sqlStr, time.Now(), id)
    
    return
}

func (ExaCustomerOp)UpdateExaCustomer(e *model.ExaCustomer) (err error) {
    sqlStr := "update exa_customers set updated_at=?, customer_name=?, customer_phone_data=? where id=?"
    _, err  = _gDB.Exec(sqlStr, time.Now(), e.CustomerName, e.CustomerPhoneData, e.ID)
    return nil
}

func (ExaCustomerOp)GetExaCustomer(id int64) (err error, customer model.ExaCustomer) {
    sqlStr := "select c.id, c.created_at, c.updated_at, c.deleted_at, c.customer_name, c.customer_phone_data, c.sys_user_id, c.sys_user_authority_id, " +
        "u.id as uid, u.username, u.nick_name, u.header_img, u.authority_id from exa_customers c join sys_users u on c.sys_user_id=u.id where c.id=?"
    fmt.Printf("sqlStr=%+v,id=%d\n", sqlStr, id)
    err = _gDB.Get(&customer, sqlStr, id)
    return
}

func (ExaCustomerOp)GetCustomerInfoList(sysUserAuthorityID string, info model.PageInfo) (err error, list interface{}, total int64) {
    limit := info.PageSize
    offset := info.PageSize * (info.Page - 1)
    
    // 获取有权限的资源列表
    err, mList := ISysAuthority.loadSysDataAuthorityId([]string{sysUserAuthorityID})
    if err != nil {
        return
    }
    //获取总数
    query, args, err := sqlx.In("SELECT count(*) FROM exa_customers where sys_user_authority_id in (?)", mList[sysUserAuthorityID])
    if err != nil {
        fmt.Printf("111111111111111111 err=%+v\n", err)
        return
    }
    query = _gDB.Rebind(query)
    err = _gDB.Get(&total, query, args...)
    if err != nil {
        fmt.Printf("222222222 err=%+v\n", err)
        return
    }
    fmt.Printf("333333333333 err=%+v,  mList[sysUserAuthorityID]=%+v\n", err, mList[sysUserAuthorityID])
    inStr := fmt.Sprintf("select c.id, c.created_at, c.updated_at, c.deleted_at, c.customer_name, c.customer_phone_data, c.sys_user_id, c.sys_user_authority_id, " +
        "u.id as uid, u.username, u.nick_name, u.header_img, u.authority_id from exa_customers c join sys_users u on c.sys_user_id=u.id where c.sys_user_authority_id in (?) limit %d,%d", offset, limit)
    query, args, err = sqlx.In(inStr,  mList[sysUserAuthorityID])

    if err != nil {
        fmt.Printf("4444444444444 err=%+v\n", err)
        return
    }
    query = _gDB.Rebind(query)
    var rlist []model.ExaCustomer
    fmt.Printf("query=%v, args=%+v\n", query, args)
    err = _gDB.Select(&rlist, query, args...)
    if err != nil {
        fmt.Printf("555555555555 err=%+v\n", err)
        return
    }
    list = rlist
 
    return
}