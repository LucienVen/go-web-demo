package repository

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/LucienVen/go-web-demo/user/domain"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

type UserRepository struct {
	database *sqlx.DB
	table    string
}

func NewUserRepository(db *sqlx.DB, table string) domain.UserRepository {
	return &UserRepository{
		database: db,
		table:    table,
	}
}

func (tr *UserRepository) Create(c context.Context, user *domain.UserStruct) (int64, error) {
	sqlStr := fmt.Sprintf("insert into %s(name, age, status, create_time, update_time) values (?, ?, ?, ?, ?)", tr.table)
	ret, err := tr.database.Exec(sqlStr, user.Name, user.Age, user.Status, user.CreateTime, user.UpdateTime)
	if err != nil {
		return 0, err
	}
	theId, err := ret.LastInsertId()
	if err != nil {
		return 0, err
	}
	return theId, nil
}

func (tr *UserRepository) List(c context.Context) ([]domain.UserStruct, error) {
	return nil, nil
}

func (tr *UserRepository) Update(c context.Context, user *domain.UserStruct) error {
	return nil
}

func (tr *UserRepository) GetById(c context.Context, id int64) (domain.UserStruct, error) {
	sqlStr := fmt.Sprintf("select * from %s where id=?", tr.table)
	var ret domain.UserStruct
	err := tr.database.Get(&ret, sqlStr, id)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return ret, err
	}
	return ret, nil
}
