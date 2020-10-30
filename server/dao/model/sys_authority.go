package model

import (
    "time"
    "database/sql"
)

type SysAuthority struct {
    CreatedAt       time.Time      `db:"created_at"`
    UpdatedAt       time.Time      `db:"updated_at"`
    DeletedAt       sql.NullTime   `db:"deleted_at"`
    AuthorityId     string         `json:"authorityId" db:"authority_id"`
    AuthorityName   string         `json:"authorityName" db:"authority_name"`
    ParentId        string         `json:"parentId" db:"parent_id"`
    DataAuthorityId []string       `json:"dataAuthorityIdList"`
    Children        []SysAuthority `json:"children"`
    // SysBaseMenus    []SysBaseMenu  `json:"menus" `
}

type SysDataAuthorityId struct {
    SysAuthorityAuthorityId    string `db:"sys_authority_authority_id"`
    DataAuthorityIdAuthorityId string `db:"data_authority_id_authority_id"`
}
