package service

import (
    "g-admin/dao"
    "g-admin/dao/model"
    "g-admin/utils"
    "github.com/pkg/errors"
    
    "fmt"
    uuid "github.com/satori/go.uuid"
    "time"
)

func Login(u string, p string) (err error, userInter *model.SysUser) {
    err, userInter = dao.ISysUser.GetSysUserInfoByName(u)
    if err == nil {
        md5P := utils.MD5V([]byte(p))
        fmt.Println(userInter.Password)
        if userInter.Password != md5P {
            err = errors.New("password err")
        }
    }
    return
}

func Register(u *model.SysUser) (err error, userInter *model.SysUser) {
    //var user model.SysUser
    u.Password = utils.MD5V([]byte(u.Password))
    u.UUID = uuid.NewV4()
    now := time.Now()
    u.CreatedAt = now
    u.UpdatedAt = now
    err = dao.ISysUser.AddUser(u)
    return err, u
}

func ChangePassword(u *model.SysUser, newPassword string) (err error, userInter *model.SysUser) {
    //var user model.SysUser
    u.Password = utils.MD5V([]byte(u.Password))
    newPassword = utils.MD5V([]byte(newPassword))
    err = dao.ISysUser.ChangePassword(u, newPassword)
   
    return err, u
}

// func FindUserByUuid(uuid string)(err error, userInter *dao.SysUser) {
//     err, userInter = dao.ISysUser.FindUserByUuid(uuid)
// }