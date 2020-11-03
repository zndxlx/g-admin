package dao

import (
    "g-admin/dao/model"
    "strconv"
    "errors"
    "time"
    "fmt"
    "github.com/jmoiron/sqlx"
)

type SysMenuOp struct {
}
var ISysMenu = SysMenuOp{}

func (SysMenuOp)getBaseMenuTreeMap() (err error, treeMap map[string][]model.SysBaseMenu) {
    var allMenus []model.SysBaseMenu
    treeMap = make(map[string][]model.SysBaseMenu)
    sqlStr := "select * from sys_base_menus "
    err = _gDB.Select(&allMenus, sqlStr)
    if err != nil {
        return
    }
    err, mList := ISysMenu.loadAllSysBaseMenuParameters()
    if err != nil {
        return
    }
    //fmt.Printf("mList=%+v\n")
    for index, v := range allMenus {
        
        allMenus[index].Parameters = mList[v.ID]
        
        treeMap[v.ParentId] = append(treeMap[v.ParentId],  allMenus[index])
    }
    //fmt.Printf("%+v\n", treeMap)
    return err, treeMap
}

func getBaseChildrenList(menu *model.SysBaseMenu, treeMap map[string][]model.SysBaseMenu) (err error) {
    menu.Children = treeMap[strconv.Itoa(int(menu.ID))]
    for i := 0; i < len(menu.Children); i++ {
        err = getBaseChildrenList(&menu.Children[i], treeMap)
    }
    return err
}

func (SysMenuOp)GetInfoList() (err error, list interface{}, total int64) {
    var menuList []model.SysBaseMenu
    err, treeMap := ISysMenu.getBaseMenuTreeMap()
    menuList = treeMap["0"]
    for i := 0; i < len(menuList); i++ {
        err = getBaseChildrenList(&menuList[i], treeMap)
    }
    return err, menuList, total
}

func (SysMenuOp)loadAllSysBaseMenuParameters() (err error, mList map[int64][]model.SysBaseMenuParameter) {
    var list []model.SysBaseMenuParameter
    sqlStr := "select id as p_id, created_at as p_created_at,  updated_at as p_updated_at, sys_base_menu_id, `type`, `key`, `value` from sys_base_menu_parameters"
    err = _gDB.Select(&list, sqlStr)
    if err != nil {
        return
    }
   
    mList = make(map[int64][]model.SysBaseMenuParameter)
    for _, l := range list {
        mList[l.SysBaseMenuID] = append(mList[l.SysBaseMenuID], l)
    }
   // fmt.Printf("111111111111 %+v\n",mList )
    return
}

func (SysMenuOp)GetBaseMenuById(id int64) (err error, menu model.SysBaseMenu) {
    sqlStr := "select * from sys_base_menus where id = ?"
    err = _gDB.Get(&menu, sqlStr, id)
    if err != nil {
        return
    }
    
    err, mList := ISysMenu.loadAllSysBaseMenuParameters()
    if err != nil {
        return
    }
   // fmt.Printf("menu.id=%d, mlist=%+v\n", menu.ID, mList)
    menu.Parameters = mList[menu.ID]
    return
}

func (SysMenuOp)GetMenuAuthority(authorityId string) (err error, allMenus []model.SysBaseMenu) {
    sqlStr := "select  m.* from sys_authority_menus a join sys_base_menus m on a.sys_base_menu_id = m.id WHERE sys_authority_authority_id=?"
    err = _gDB.Select(&allMenus, sqlStr, authorityId)
    if err != nil {
        return
    }
    
    err, mList := ISysMenu.loadAllSysBaseMenuParameters()
    if err != nil {
        return
    }
    for index, v := range allMenus {
        allMenus[index].Parameters = mList[v.ID]
    }
    return err, allMenus
}

func (SysMenuOp)AddBaseMenu(menu model.SysBaseMenu) (err error) {
    sqlStr := "insert into sys_base_menus(created_at, updated_at, menu_level, parent_id, path, name, hidden, component, sort, keep_alive, default_menu, title, icon) " +
        "Values (:created_at, :updated_at, :menu_level, :parent_id, :path, :name, :hidden, :component, :sort, :keep_alive, :default_menu, :title, :icon)"
    _, err  = _gDB.NamedExec(sqlStr, menu)
    return err

}


func (SysMenuOp)UpdateBaseMenu(menu model.SysBaseMenu) (err error) {
    sqlStr := "update sys_base_menus set updated_at=?, menu_level=?, parent_id=?, path=?, name=?, hidden=?, component=?, sort=?, keep_alive=?, default_menu=?, title=?, icon=?  where id=?"
    fmt.Printf("UpdateBaseMenu sqlStr=%v， menu=%+v\n", sqlStr, menu)
    r, err  := _gDB.Exec(sqlStr, time.Now(), menu.MenuLevel, menu.ParentId, menu.Path, menu.Name, menu.Hidden, menu.Component, menu.Sort, menu.KeepAlive,
        menu.DefaultMenu, menu.Title, menu.Icon, menu.ID)
    if err != nil {
        return err
    }
    rows, err := r.RowsAffected()
    if err != nil {
        return err
    }
    if rows == 0 {
        return errors.New("0行被修改")
    }

    return err
}

func (SysMenuOp)DeleteBaseMenu(id int64) (err error) {
    //查询是否有子菜单
    var total int64
    err = _gDB.Get(&total, "SELECT count(*) FROM sys_base_menus where parent_id = ?;", id)
    if err != nil {
        return
    }
    if total != 0 {
        return errors.New("此菜单存在子菜单不可删除")
    }
    //删除sys_authority_menus记录
    _, err  = _gDB.Exec("delete from sys_authority_menus where sys_base_menu_id=?", id)
    if err != nil {
        return err
    }
    //删除SysBaseMenuParameter记录
    _, err  = _gDB.Exec("delete from sys_base_menu_parameters where sys_base_menu_id=?", id)
    if err != nil {
        return err
    }
    //删除sys_base_menus记录
    _, err  = _gDB.Exec("delete from sys_base_menus where id=?", id)
    if err != nil {
        return err
    }
    
    return err
}




func (SysMenuOp)AddMenuAuthority(menus []model.SysAuthorityMenus, authorityId string) (err error) {
    _, err  = _gDB.Exec("delete from sys_authority_menus where sys_authority_authority_id=?", authorityId)
    if err != nil {
        return err
    }
    var mlist = make([]interface{}, len(menus))
    for index, v := range menus {
        mlist[index] = v
    }
    query, args, err := sqlx.In(
        "INSERT INTO sys_authority_menus (sys_authority_authority_id, sys_base_menu_id) VALUES (?), (?)",
        mlist...,
    )
    fmt.Printf("query=%+v, args=%+v, err=%+v\n", query, args, err)
    _, err = _gDB.Exec(query, args...) //...参数展开
    fmt.Printf("err%+v\n", err)
    return err
}