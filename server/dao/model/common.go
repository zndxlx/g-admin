package model

import (
    "time"
    "database/sql"
)

type Model struct {
    ID        int64 `db:"id"`
    CreatedAt time.Time  `db:"created_at"`
    UpdatedAt time.Time   `db:"updated_at"`
    DeletedAt sql.NullTime `db:"deleted_at"`
}

// Paging common input parameter structure
type PageInfo struct {
    Page     int `json:"page" form:"page"`
    PageSize int `json:"pageSize" form:"pageSize"`
}

// Find by id structure
type GetById struct {
    Id int64 `json:"id" form:"id"`
}

type IdsReq struct {
    Ids []int `json:"ids" form:"ids"`
}

type PageResult struct {
    List     interface{} `json:"list"`
    Total    int64       `json:"total"`
    Page     int         `json:"page"`
    PageSize int         `json:"pageSize"`
}