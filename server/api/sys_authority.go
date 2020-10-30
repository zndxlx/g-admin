package api

import (
    "g-admin/utils/response"
    "github.com/gin-gonic/gin"
    "g-admin/dao/model"
    "fmt"
    "g-admin/dao"
    "time"
)

func GetAuthorityList(c *gin.Context) {
    var pageInfo model.PageInfo
    _ = c.ShouldBindJSON(&pageInfo)
   
    err, list, total := dao.ISysAuthority.GetAuthorityInfoList(pageInfo)
    if err != nil {
        response.FailWithMessage(fmt.Sprintf("获取数据失败，%v", err), c)
    } else {
        response.OkWithData(model.PageResult{
            List:     list,
            Total:    total,
            Page:     pageInfo.Page,
            PageSize: pageInfo.PageSize,
        }, c)
    }
}

type CreateAuthorityReq struct {
    AuthorityId     string         `json:"authorityId" binding:"required"`
    AuthorityName   string         `json:"authorityName" binding:"required"`
    ParentId        string         `json:"parentId" binding:"required"`
}

type SysAuthorityResponse struct {
    Authority model.SysAuthority `json:"authority"`
}

func CreateAuthority(c *gin.Context) {
    var R CreateAuthorityReq
    if err := c.ShouldBindJSON(&R); err != nil {
        errReqFormatRespone(err, c)
        return
    }
    
    now := time.Now()
    sysAuthority := &model.SysAuthority{CreatedAt:now,UpdatedAt:now,
        AuthorityId:R.AuthorityId,
        AuthorityName:R.AuthorityName,
        ParentId:R.ParentId }
    err := dao.ISysAuthority.CreateAuthority(sysAuthority)
    if err != nil {
        response.FailWithMessage(fmt.Sprintf("创建失败，%v", err), c)
    } else {
        response.OkWithData(SysAuthorityResponse{Authority: *sysAuthority}, c)
    }
}

type DeleteAuthorityReq struct {
    AuthorityId     string         `json:"authorityId" binding:"required"`
}

func DeleteAuthority(c *gin.Context) {
    var R DeleteAuthorityReq
    if err := c.ShouldBindJSON(&R); err != nil {
        errReqFormatRespone(err, c)
        return
    }

    err := dao.ISysAuthority.DeleteAuthority(R.AuthorityId)
    if err != nil {
        response.FailWithMessage(fmt.Sprintf("删除失败，%v", err), c)
    } else {
        response.OkWithMessage("删除成功", c)
    }
}

type UpdateAuthorityReq struct {
    AuthorityId     string         `json:"authorityId" binding:"required"`
    AuthorityName   string         `json:"authorityName" binding:"required"`
    ParentId        string         `json:"parentId" binding:"required"`
}

func UpdateAuthority(c *gin.Context) {
    var R UpdateAuthorityReq
    if err := c.ShouldBindJSON(&R); err != nil {
        errReqFormatRespone(err, c)
        return
    }
    now := time.Now()
    sysAuthority := &model.SysAuthority{UpdatedAt:now,
        AuthorityId:R.AuthorityId,
        AuthorityName:R.AuthorityName,
        ParentId:R.ParentId }
    err := dao.ISysAuthority.UpdateAuthority(sysAuthority)
    if err != nil {
        response.FailWithMessage(fmt.Sprintf("更新失败，%v", err), c)
    } else {
        response.OkWithData(SysAuthorityResponse{Authority: *sysAuthority}, c)
    }
}