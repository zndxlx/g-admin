package api

import (
    "github.com/gin-gonic/gin"
    "g-admin/utils/response"
    "g-admin/dao/model"
    "g-admin/dao"
    "fmt"
    "time"
)



func GetMenuList(c *gin.Context) {
    var pageInfo model.PageInfo
    _ = c.ShouldBindJSON(&pageInfo)
    
    err, menuList, total := dao.ISysMenu.GetInfoList()
    if err != nil {
        response.FailWithMessage(fmt.Sprintf("获取数据失败，%v", err), c)
    } else {
        response.OkWithData(model.PageResult{
            List:     menuList,
            Total:    total,
            Page:     pageInfo.Page,
            PageSize: pageInfo.PageSize,
        }, c)
    }
}

type GetBaseMenuByIdRsp struct {
    Menu model.SysBaseMenu `json:"menu"`
}

func GetBaseMenuById(c *gin.Context) {
    var idInfo model.GetById
    _ = c.ShouldBindJSON(&idInfo)
   
    err, menu := dao.ISysMenu.GetBaseMenuById(idInfo.Id)
    if err != nil {
        response.FailWithMessage(fmt.Sprintf("查询失败：%v", err), c)
    } else {
        response.OkWithData(GetBaseMenuByIdRsp{Menu: menu}, c)
    }
}

type AddBaseMenuReq struct {
    ParentId   string `json:"parentId" binding:"required"`
    Path       string `json:"path" binding:"required"`
    Name       string `json:"name" binding:"required"`
    Hidden     bool   `json:"hidden"`
    Component  string `json:"component" binding:"required"`
    Sort       int    `json:"sort"`
    MenuLevel  uint   `json:"menu_level"`
    model.Meta       `json:"meta"`
    
}

func AddBaseMenu(c *gin.Context) {
    var R AddBaseMenuReq
    if err := c.ShouldBindJSON(&R); err != nil{
        errReqFormatRespone(err, c)
        return
    }
    now := time.Now()
    err := dao.ISysMenu.AddBaseMenu(model.SysBaseMenu{
        Model: model.Model{
            CreatedAt:now,
            UpdatedAt:now,
        },
        ParentId : R.ParentId,
        Path     : R.Path,
        Name    : R.Name,
        Hidden    : R.Hidden,
        Component : R.Component,
        Sort      : R.Sort,
        MenuLevel : R.MenuLevel,
        Meta      : R.Meta,
    })
    if err != nil {
        response.FailWithMessage(fmt.Sprintf("添加失败，%v", err), c)
    } else {
        response.OkWithMessage("添加成功", c)
    }
}


type UpdateBaseMenuReq struct {
    ID        int64    `binding:"required"`
    ParentId   string `json:"parentId" binding:"required"`
    Path       string `json:"path" binding:"required"`
    Name       string `json:"name" binding:"required"`
    Hidden     bool   `json:"hidden"`
    Component  string `json:"component" binding:"required"`
    Sort       int    `json:"sort"`
    MenuLevel  uint   `json:"menu_level"`
    model.Meta       `json:"meta" binding:"required"`
    
}

func UpdateBaseMenu(c *gin.Context) {
    var R UpdateBaseMenuReq
    if err := c.ShouldBindJSON(&R); err != nil{
        errReqFormatRespone(err, c)
        return
    }
    now := time.Now()
    err := dao.ISysMenu.UpdateBaseMenu(model.SysBaseMenu{
        Model: model.Model{
            UpdatedAt:now,
            ID: R.ID,
        },
        ParentId : R.ParentId,
        Path     : R.Path,
        Name    : R.Name,
        Hidden    : R.Hidden,
        Component : R.Component,
        Sort      : R.Sort,
        MenuLevel : R.MenuLevel,
        Meta      : R.Meta,
    })
    if err != nil {
        response.FailWithMessage(fmt.Sprintf("修改失败，%v", err), c)
    } else {
        response.OkWithMessage("修改成功", c)
    }
}

func DeleteBaseMenu(c *gin.Context) {
    var idInfo model.GetById
    _ = c.ShouldBindJSON(&idInfo)
    
    err := dao.ISysMenu.DeleteBaseMenu(idInfo.Id)
    if err != nil {
        response.FailWithMessage(fmt.Sprintf("删除失败：%v", err), c)
    } else {
        response.OkWithMessage("删除成功", c)
        
    }
}


type GetMenuAuthorityReq struct {
    AuthorityId string `json:"authorityId"  binding:"required"`
}

type GetMenuAuthorityRsp struct {
    Menus []model.SysBaseMenu `json:"menus"`
}

func GetMenuAuthority(c *gin.Context) {
    var R GetMenuAuthorityReq
    if err := c.ShouldBindJSON(&R); err != nil{
        errReqFormatRespone(err, c)
        return
    }
    
    err, menus := dao.ISysMenu.GetMenuAuthority(R.AuthorityId)
    if err != nil {
        response.FailWithDetailed(response.ERROR, GetMenuAuthorityRsp{Menus: menus}, fmt.Sprintf("添加失败，%v", err), c)
    } else {
        response.Result(response.SUCCESS, gin.H{"menus": menus}, "获取成功", c)
    }
}

type AddMenuAuthorityMenuReq struct {
    ID int64
}

type AddMenuAuthorityReq struct {
    Menus       []AddMenuAuthorityMenuReq `json:"menus"  binding:"required"`
    AuthorityId string  `json:"authorityId"  binding:"required"`
}



func AddMenuAuthority(c *gin.Context) {
    var R AddMenuAuthorityReq
    if err := c.ShouldBindJSON(&R); err != nil{
        errReqFormatRespone(err, c)
        return
    }
    var menus = make([]model.SysAuthorityMenus, len(R.Menus))
    for index, v := range R.Menus {
        menus[index].SysBaseMenuId = v.ID
        menus[index].SysAuthorityAuthorityId = R.AuthorityId
       
    }
    err := dao.ISysMenu.AddMenuAuthority(menus, R.AuthorityId)
    if err != nil {
        response.FailWithMessage(fmt.Sprintf("添加失败，%v", err), c)
    } else {
        response.OkWithMessage("添加成功", c)
    }
}