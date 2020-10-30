package dao

import (
    "g-admin/core/log"
    "go.uber.org/zap"
    "g-admin/dao/model"
    "fmt"
    "github.com/pkg/errors"
    uuid "github.com/satori/go.uuid"
    "time"
)

type SysUserOp struct {
}
var ISysUser = SysUserOp{}

func (SysUserOp) GetSysUserInfoById(id int64) (err error, user *model.SysUser) {
    sqlStr := "select id, created_at, updated_at, deleted_at, uuid, username, nick_name, password, nick_name,  header_img, authority_id, from sys_users where id=?"
    var u model.SysUser
    err = _gDB.Get(&u, sqlStr, id)
    if err != nil {
        log.Error("GetSysUserInfoById failed",
            zap.Int64("id", id),
            zap.String("sql", sqlStr),
            zap.Error(err))
        return
    }
    user = &u
    return
}

func (SysUserOp) GetSysUserInfoByName(name string) (err error, user *model.SysUser) {
    sqlStr := "select id, created_at, updated_at, deleted_at, uuid, username, nick_name, password, nick_name,  header_img, authority_id from sys_users where username=?"
    var u model.SysUser
    err = _gDB.Get(&u, sqlStr, name)
    if err != nil {
        log.Error("GetSysUserInfoByName failed",
            zap.String("name", name),
            zap.String("sql", sqlStr),
            zap.Error(err))
        return
    }
    user = &u
    return
}

func (SysUserOp) FindUserByUuid(uuid string)(err error, user *model.SysUser) {
    sqlStr := "select id, created_at, updated_at, deleted_at, uuid, username, nick_name, password, nick_name,  header_img, authority_id from sys_users where uuid=?"
    var u model.SysUser
    err = _gDB.Get(&u, sqlStr, uuid)
    if err != nil {
        log.Error("FindUserByUuid failed",
            zap.String("uuid", uuid),
            zap.String("sql", sqlStr),
            zap.Error(err))
        return
    }
    user = &u
    return
}


func (SysUserOp) AddUser(user *model.SysUser) (err error) {
    sqlStr := ""
    if user.HeaderImg == "" {
        sqlStr = "insert into sys_users(created_at, updated_at, uuid, username, password, nick_name, authority_id) Values (:created_at, :updated_at, :uuid, :username, :password, :nick_name, :authority_id)"
    }else{
        sqlStr = "insert into sys_users(created_at, updated_at,,uuid, username, password, nick_name, head_img, authority_id) Values (:created_at, :updated_at, :uuid, :username, :password, :nick_name, :head_img, :authority_id)"
    }
    
    _, err  = _gDB.NamedExec(sqlStr, user)
    return
}

func (SysUserOp)DeleteUser(id int64) (err error) {
    //var user model.SysUser
    sqlStr := "update sys_users set deleted_at = ? where  id = ?"
    _, err  = _gDB.Exec(sqlStr, time.Now(), id)
    
    return
}

func (SysUserOp)getCountByAuthorityId(authorityId string) (err error, total int64) {
    err = _gDB.Get(&total, "SELECT count(*) FROM sys_users where authority_id = ?;", authorityId)
    if err != nil {
        return
    }
    return
}

func (SysUserOp) ChangePassword(u *model.SysUser, newPassword string) (err error) {
    fmt.Printf("password=%s, newPassword=%s\n",  u.Password, newPassword)
    sqlStr := "update sys_users set password = ? where  username = ? and  password = ?"
    r, err  := _gDB.Exec(sqlStr, newPassword, u.Username, u.Password)
    if err != nil {
        return err
    }
    rows, err := r.RowsAffected()
    if err != nil {
        return err
    }
    if rows == 0 {
        return errors.New("修改失败")
    }
    return
}

func (SysUserOp)GetUserInfoList(info model.PageInfo) (err error, list []model.SysUser, total int64){
    err = _gDB.Get(&total, "SELECT count(*) FROM sys_users;")
    if err != nil {
        return
    }
    limit := info.PageSize
    offset := info.PageSize * (info.Page - 1)
    sqlStr := "select id, created_at, updated_at, deleted_at, uuid, username, nick_name, password, nick_name,  header_img, authority_id from sys_users limit ?, ?"
    err = _gDB.Select(&list, sqlStr, offset, limit)

    return

}

func (SysUserOp)SetUserAuthority(uuid uuid.UUID, authorityId string) (err error) {
    sqlStr := "update sys_users set authority_id=? where uuid=?"
    r, err  := _gDB.Exec(sqlStr, authorityId, uuid)
    if err != nil {
        return err
    }
    rows, err := r.RowsAffected()
    if err != nil {
        return err
    }
    if rows == 0 {
        return errors.New("修改失败")
    }
    return
}