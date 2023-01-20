package rotation

import (
	"context"
	"demo02/internal/dao"
	"demo02/internal/model"
	"demo02/internal/service"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/encoding/ghtml"
	"github.com/gogf/gf/v2/frame/g"
)

type sRotation struct{}

func init() {
	service.RegisterRotation(New())
}

func New() *sRotation {
	return &sRotation{}
}
func (s *sRotation) Create(ctx context.Context, in model.RotationCreateInput) (out model.RotationCreateOutput, err error) {

	// 不允许HTML代码
	if err = ghtml.SpecialCharsMapOrStruct(in); err != nil {
		return out, err
	}
	lastInsertID, err := dao.Rotation.Ctx(ctx).Data(in).InsertAndGetId()
	if err != nil {
		return out, err
	}
	return model.RotationCreateOutput{RotationId: int(lastInsertID)}, err
}
func (s *sRotation) Delete(ctx context.Context, id uint) error {
	return dao.Rotation.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {

		// 删除内容
		_, err := dao.Rotation.Ctx(ctx).Where(g.Map{
			dao.Rotation.Columns().Id: id,
			dao.Rotation.Columns().Id: id,
		}).Delete()
		return err
	})
}

// Update 修改
func (s *sRotation) Update(ctx context.Context, in model.RotationUpdateInput) error {
	return dao.Rotation.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		// 不允许HTML代码
		if err := ghtml.SpecialCharsMapOrStruct(in); err != nil {
			return err
		}
		_, err := dao.Rotation.
			Ctx(ctx).
			Data(in).
			FieldsEx(dao.Rotation.Columns().Id).
			Where(dao.Rotation.Columns().Id, in.Id).
			Update()
		return err
	})
}
