package model


//为了防止sqlx中查询的冲突重新定义
type SysUserInCustomer struct {
	UID        int64       `db:"uid" json:"id"`
	Username    string    `db:"username" json:"userName"`
	NickName    string    `db:"nick_name" json:"nickName"`
	HeaderImg   string    `db:"header_img" json:"headerImg"`
	AuthorityId string    `db:"authority_id" json:"authorityId`
}

type ExaCustomer struct {
	Model
	CustomerName       string  `db:"customer_name" json:"customerName"`
	CustomerPhoneData  string  `db:"customer_phone_data" json:"customerPhoneData"`
	SysUserID          int64    `db:"sys_user_id" json:"sysUserId"`
	SysUserAuthorityID string  `db:"sys_user_authority_id" json:"sysUserAuthorityID"`
	SysUserInCustomer         `json:"sysUser"`
}
