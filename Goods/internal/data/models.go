package data

import (
	"database/sql"
	"errors"
)

var (
	ErrRecordNotFound = errors.New("record not found")
	ErrEditConflict   = errors.New("edit conflict")
)

type Models struct {
	LeatherGoods LeatherGoodsModel
	Permissions  PermissionModel
	Tokens       TokenModel
	Users        UserModel
}

func NewModels(db *sql.DB) Models {
	return Models{
		LeatherGoods: LeatherGoodsModel{DB: db},
		Permissions:  PermissionModel{DB: db},
		Tokens:       TokenModel{DB: db},
		Users:        UserModel{DB: db},
	}
}
