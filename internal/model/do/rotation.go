// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Rotation is the golang structure of table gf_rotation for DAO operations like Where/Data.
type Rotation struct {
	g.Meta   `orm:"table:gf_rotation, do:true"`
	Id       interface{} //
	PicUrl   interface{} //
	Link     interface{} //
	Sort     interface{} //
	CreateAt *gtime.Time //
	UpdateAt *gtime.Time //
	DeleteAt *gtime.Time //
}
