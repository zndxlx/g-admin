package model

import (
    "time"
    "database/sql"
    "database/sql/driver"
)

type SysBaseMenu struct {
    Model
    MenuLevel  uint   `json:"-" db:"menu_level"`
    ParentId   string `json:"parentId" db:"parent_id"`
    Path       string `json:"path" db:"path"`
    Name       string `json:"name" db:"name"`
    Hidden     bool   `json:"hidden" db:"hidden"`
    Component  string `json:"component" db:"component"`
    Sort       int    `json:"sort" db:"sort"`
    Meta       `json:"meta"`
    Children   []SysBaseMenu          `json:"children" gorm:"-"`
    Parameters []SysBaseMenuParameter `json:"parameters"`
}

type Meta struct {
    KeepAlive   bool   `json:"keepAlive" db:"keep_alive"`
    DefaultMenu bool   `json:"defaultMenu" db:"default_menu"`
    Title       string `json:"title" db:"title"`
    Icon        string `json:"icon" db:"icon"`
}

type SysBaseMenuParameter struct {
    ID            int64        `db:"p_id"`
    CreatedAt     time.Time    `db:"p_created_at"`
    UpdatedAt     time.Time    `db:"p_updated_at"`
    DeletedAt     sql.NullTime `db:"p_deleted_at"`
    SysBaseMenuID int64         `json:"-" db:"sys_base_menu_id"`
    Type          string       `json:"type" db:"type"`
    Key           string       `json:"key" db:"key"`
    Value         string       `json:"value" db:"value"`
}

type SysAuthorityMenus struct {
    SysAuthorityAuthorityId    string `db:"sys_authority_authority_id"`
    SysBaseMenuId int64 `db:"sys_base_menu_id"`
}

func (s SysAuthorityMenus) Value() (driver.Value, error) {
    return []interface{}{s.SysAuthorityAuthorityId, s.SysBaseMenuId}, nil
}