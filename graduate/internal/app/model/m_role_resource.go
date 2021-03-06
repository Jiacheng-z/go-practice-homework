package model

import (
	"context"

	"github.com/LyricTian/gin-admin/v6/internal/app/schema"
)

// IRoleResource 角色资源存储接口
type IRoleResource interface {
	// 查询数据
	Query(ctx context.Context, params schema.RoleResourceQueryParam, opts ...schema.RoleResourceQueryOptions) (*schema.RoleResourceQueryResult, error)
	//// 查询指定数据
	//Get(ctx context.Context, id string, opts ...schema.RoleMenuQueryOptions) (*schema.RoleMenu, error)
	//// 创建数据
	//Create(ctx context.Context, item schema.RoleMenu) error
	//// 更新数据
	//Update(ctx context.Context, id string, item schema.RoleMenu) error
	//// 删除数据
	//Delete(ctx context.Context, id string) error
	//// 根据角色ID删除数据
	//DeleteByRoleID(ctx context.Context, roleID string) error
}
