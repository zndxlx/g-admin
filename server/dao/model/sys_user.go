package model

import (
    "github.com/satori/go.uuid"
)

type SysUser struct {
    Model
    UUID        uuid.UUID `db:"uuid" json:"uuid"`
    Username    string    `db:"username" json:"userName"`
    Password    string    `db:"password" json:"-" `
    NickName    string    `db:"nick_name" json:"nickName"`
    HeaderImg   string    `db:"header_img" json:"headerImg"`
    AuthorityId string    `db:"authority_id" json:"authorityId`
}