package api

import (
    "fmt"
    
    "github.com/gin-gonic/gin"
    "g-admin/utils/response"
    "g-admin/dao/model"
    "g-admin/dao"
    "time"
    "g-admin/utils/jwt"
)

type CreateExaCustomerReq struct {
    CustomerName      string `json:"customerName" binding:"required"`
    CustomerPhoneData string `json:"customerPhoneData" binding:"required"`
}

func CreateExaCustomer(c *gin.Context) {
    var R CreateExaCustomerReq
    if err := c.ShouldBindJSON(&R); err != nil {
        errReqFormatRespone(err, c)
        return
    }
    
    claims, _ := c.Get("claims")
    waitUse := claims.(*jwt.CustomClaims)
    err := dao.IExaCustomer.CreateExaCustomer(&model.ExaCustomer{
        Model: model.Model{
            CreatedAt: time.Now(),
            UpdatedAt: time.Now(),
        },
        SysUserAuthorityID: waitUse.AuthorityId,
        SysUserID:          waitUse.ID,
        CustomerName:       R.CustomerName,
        CustomerPhoneData:  R.CustomerPhoneData,
    })
    if err != nil {
        response.FailWithMessage(fmt.Sprintf("创建失败：%v", err), c)
    } else {
        response.OkWithMessage("创建成功", c)
    }
}

func DeleteExaCustomer(c *gin.Context) {
    var R model.GetById
    _ = c.ShouldBindJSON(&R)
    
    err := dao.IExaCustomer.DeleteExaCustomer(R.Id)
    if err != nil {
        response.FailWithMessage(fmt.Sprintf("删除失败：%v", err), c)
    } else {
        response.OkWithMessage("删除成功", c)
    }
}

type UpdateExaCustomerReq struct {
    CustomerName      string `json:"customerName" binding:"required"`
    CustomerPhoneData string `json:"customerPhoneData" binding:"required"`
}

func UpdateExaCustomer(c *gin.Context) {
    var R CreateExaCustomerReq
    if err := c.ShouldBindJSON(&R); err != nil {
        errReqFormatRespone(err, c)
        return
    }
    
    err := dao.IExaCustomer.UpdateExaCustomer(&model.ExaCustomer{
        Model: model.Model{
            UpdatedAt: time.Now(),
        },
        CustomerName:      R.CustomerName,
        CustomerPhoneData: R.CustomerPhoneData,
    })
    if err != nil {
        response.FailWithMessage(fmt.Sprintf("更新失败：%v", err), c)
    } else {
        response.OkWithMessage("更新成功", c)
    }
}

type ExaCustomerRsp struct {
    Customer model.ExaCustomer `json:"customer"`
}

func GetExaCustomer(c *gin.Context) {
    var R model.GetByID
    _ = c.ShouldBindQuery(&R)

    err, customer := dao.IExaCustomer.GetExaCustomer(R.ID)
    if err != nil {
        response.FailWithMessage(fmt.Sprintf("获取失败：%v", err), c)
    } else {
        response.OkWithData(ExaCustomerRsp{Customer: customer}, c)
    }
}

func GetExaCustomerList(c *gin.Context) {
    var pageInfo model.PageInfo
    _ = c.ShouldBindQuery(&pageInfo)
    
    claims, _ := c.Get("claims")
    waitUse := claims.(*jwt.CustomClaims)
    
    err, customerList, total := dao.IExaCustomer.GetCustomerInfoList(waitUse.AuthorityId, pageInfo)
    if err != nil {
        response.FailWithMessage(fmt.Sprintf("获取失败：%v", err), c)
    } else {
        response.OkWithData(model.PageResult{
            List:     customerList,
            Total:    total,
            Page:     pageInfo.Page,
            PageSize: pageInfo.PageSize,
        }, c)
    }
}
