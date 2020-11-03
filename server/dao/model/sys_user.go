package model

import (
    "github.com/satori/go.uuid"
)

type SysAuthorityInUser struct {
    AuthorityId   string `json:"authorityId" db:"aid"`  // 重命名为aid为了和外面不冲突，
    AuthorityName string `json:"authorityName" db:"authority_name"`
    ParentId      string `json:"parentId" db:"parent_id"`
}

type SysUser struct {
    Model
    UUID               uuid.UUID `db:"uuid" json:"uuid"`
    Username           string    `db:"username" json:"userName"`
    Password           string    `db:"password" json:"-" `
    NickName           string    `db:"nick_name" json:"nickName"`
    HeaderImg          string    `db:"header_img" json:"headerImg"`
    AuthorityId        string    `db:"authority_id" json:"authorityId`
    SysAuthorityInUser `json:"authority"`
}
