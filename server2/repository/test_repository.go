package repository

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/LucienVen/go-web-demo/server2/domain"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

type TestRepository struct {
	database *sqlx.DB
	table    string
}

func NewTestRepository(db *sqlx.DB, table string) domain.TestRepository {
	return &TestRepository{
		database: db,
		table:    table,
	}
}

func (tr *TestRepository) Create(c context.Context, test *domain.TestStruct) (int64, error) {
	sqlStr := fmt.Sprintf("insert into %s(name, age, status, create_time, update_time) values (?, ?, ?, ?, ?)", tr.table)
	ret, err := tr.database.Exec(sqlStr, test.Name, test.Age, test.Status, test.CreateTime, test.UpdateTime)
	if err != nil {
		return 0, err
	}
	theId, err := ret.LastInsertId()
	if err != nil {
		return 0, err
	}
	return theId, nil
}

func (tr *TestRepository) List(c context.Context) ([]domain.TestStruct, error) {
	return nil, nil
}

func (tr *TestRepository) Update(c context.Context, test *domain.TestStruct) error {
	return nil
}

func (tr *TestRepository) GetById(c context.Context, id int64) (domain.TestStruct, error) {
	sqlStr := fmt.Sprintf("select * from %s where id=?", tr.table)
	var ret domain.TestStruct
	err := tr.database.Get(&ret, sqlStr, id)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return ret, err
	}
	return ret, nil
}
