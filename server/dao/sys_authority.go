package dao

import (
    "g-admin/dao/model"
    "fmt"
    
    "errors"
    "time"
    "github.com/jmoiron/sqlx"
)

type SysAuthorityOp struct {
}
var ISysAuthority = SysAuthorityOp{}


func (SysAuthorityOp)CreateAuthority(authority *model.SysAuthority) (err error) {
    sqlStr := "insert into  sys_authorities(created_at, updated_at, authority_id, authority_name, parent_id) values (:created_at, :updated_at, :authority_id, :authority_name, :parent_id)"
    _, err = _gDB.NamedExec(sqlStr, authority)

    return err
}


func (SysAuthorityOp) GetAuthorityInfoList(info model.PageInfo) (err error, list interface{}, total int64) {
    limit := info.PageSize
    offset := info.PageSize * (info.Page - 1)
    var authority []model.SysAuthority
    sqlStr := "select * from sys_authorities where parent_id=0 limit ?,?"
    err = _gDB.Select(&authority, sqlStr, offset, limit)
    if err != nil {
        return
    }
    if len(authority) > 0 {
        for k := range authority {
            err = ISysAuthority.findChildrenAuthority(&authority[k])
        }
    }
    //ISysAuthority.loadSysDataAuthorityId()
    return err, authority, total
}

func (SysAuthorityOp)findChildrenAuthority(authority *model.SysAuthority) (err error) {
    sqlStr := "select * from sys_authorities where parent_id=?"
    err = _gDB.Select(&authority.Children, sqlStr, authority.AuthorityId)
    if err != nil {
        return
    }
    if len(authority.Children) > 0 {
        for k := range authority.Children {
            err = ISysAuthority.findChildrenAuthority(&authority.Children[k])
        }
    }
    return err
}

// func findChildrenAuthority(authority *model.SysAuthority) (err error) {
//     //err = global.GVA_DB.Preload("DataAuthorityId").Where("parent_id = ?", authority.AuthorityId).Find(&authority.Children).Error
//     sqlStr := ""
//     if len(authority.Children) > 0 {
//         for k := range authority.Children {
//             err = findChildrenAuthority(&authority.Children[k])
//         }
//     }
//     return err
// }

func (SysAuthorityOp)loadSysDataAuthorityId(ids []string) (err error, mList map[string][]string) {
    var list []model.SysDataAuthorityId
    
    query, args, err := sqlx.In("select sys_authority_authority_id, data_authority_id_authority_id from sys_data_authority_id where sys_authority_authority_id in (?)", ids)
    if err != nil {
        return
    }
    query = _gDB.Rebind(query)
    err = _gDB.Select(&list, query, args...)
    if err != nil {
        return
    }
    mList = make(map[string][]string)
    for _, l := range list {
        mList[l.SysAuthorityAuthorityId] = append(mList[l.SysAuthorityAuthorityId], l.DataAuthorityIdAuthorityId)
    }
    fmt.Printf("111111111111 %+v\n",mList )
    return
}

func (SysAuthorityOp)DeleteAuthority(id string) (err error) {
    err, count := ISysUser.getCountByAuthorityId(id)
    if err != nil {
        return
    }
    if count != 0 {
        return errors.New("此角色有用户正在使用禁止删除")
    }
    var total int
    err = _gDB.Get(&total, "SELECT count(*) FROM sys_authorities where parent_id = ?;", id)
    if err != nil {
        return
    }

    if total != 0 {
        return errors.New("此角色存在子角色不允许删除")
    }
    
    //删除sys_authority_menus 中 角色id相关的menu
    sqlStr := "delete from sys_authority_menus where  sys_authority_authority_id = ?"
    _, err  = _gDB.Exec(sqlStr, id)
    if err != nil {
        return
    }
    //删除本角色
    sqlStr = "update sys_authorities  set deleted_at = ? where  authority_id = ?"
    _, err  = _gDB.Exec(sqlStr, time.Now(), id)
    if err != nil {
        return
    }
    //删除casbin
   
    return err
}

func (SysAuthorityOp)UpdateAuthority(authority *model.SysAuthority) (err error) {
    sqlStr := "update  sys_authorities set  updated_at=?, parent_id=?, authority_name=?  where authority_id=?"
    _, err  = _gDB.Exec(sqlStr, authority.UpdatedAt, authority.ParentId, authority.AuthorityName, authority.AuthorityId)
    if err != nil {
        return
    }
    return err
}

