package model

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/cache"

	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ UsersModel = (*customUsersModel)(nil)

type (
	// UsersModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUsersModel.
	UsersModel interface {
		usersModel
		AllUser(ctx context.Context) ([]*Users, error)
		FindOneByName(ctx context.Context, name string) (*Users, error)
	}

	customUsersModel struct {
		*defaultUsersModel
	}
)

func (c customUsersModel) FindOneByName(ctx context.Context, name string) (*Users, error) {
	var resp Users
	query := fmt.Sprintf("select %s from %s where `name` = ?", usersRows, c.table)
	err := c.QueryRowNoCacheCtx(ctx, &resp, query, name)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (c customUsersModel) AllUser(ctx context.Context) ([]*Users, error) {
	query := fmt.Sprintf("select %s from %s", usersRows, c.table)

	var resp []*Users
	err := c.QueryRowsNoCacheCtx(ctx, &resp, query)
	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}
}

// NewUsersModel returns a model for the database table.
func NewUsersModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) UsersModel {
	return &customUsersModel{
		defaultUsersModel: newUsersModel(conn, c, opts...),
	}
}
