package repository

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/LucienVen/go-web-demo/server2/domain"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

type ProductRepository struct {
	database *sqlx.DB
	table    string
}

func NewProductRepository(db *sqlx.DB, table string) domain.ProductRepository {
	return &ProductRepository{
		database: db,
		table:    table,
	}
}

func (pr *ProductRepository) Create(c context.Context, data *domain.Product) (int64, error) {
	sql := fmt.Sprintf("insert into %s(name, price, status, create_time, update_time, is_delete) valus(?, ?, ?, ?, ?, ?)", pr.table)
	ret, err := pr.database.Exec(sql, data.Name, data.Price, data.Status, data.CreateTime, data.UpdateTime, data.IsDelete)
	if err != nil {
		return 0, err
	}

	thisId, err := ret.LastInsertId()
	if err != nil {
		return 0, err
	}

	return thisId, nil
}

func (pr *ProductRepository) Update(c context.Context, data *domain.Product) error {
	sqlStr := fmt.Sprintf("update %s set name=?, price=?, update_time=?, status=? where id=?", pr.table)
	_, err := pr.database.Exec(sqlStr, data.Name, data.Price, data.UpdateTime, data.Status, data.Id)
	if err != nil {
		return err
	}
	return nil
}

func (pr *ProductRepository) List(c context.Context, param *domain.SearchProductParam) ([]domain.Product, error) {
	sqlStr := fmt.Sprintf("select * from %s where 1=1 limit ? offset ?", pr.table)
	ret := make([]domain.Product, 0)
	err := pr.database.Select(&ret, sqlStr, param.PageSize, param.Page-1*param.PageSize)
	if err != nil {
		return nil, err
	}
	return ret, nil
}

func (pr ProductRepository) GetById(c context.Context, id string) (domain.Product, error) {
	sqlStr := fmt.Sprintf("select * from %s where id=?", pr.table)
	var ret domain.Product
	err := pr.database.Get(&ret, sqlStr, id)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return ret, err
	}

	return ret, nil
}
