package postgres

import (
	"context"
	"errors"
	"reflect"
	"time"

	"github.com/doug-martin/goqu/v9"
	_ "github.com/doug-martin/goqu/v9/dialect/postgres"
	"github.com/georgysavva/scany/sqlscan"

	"github.com/hasanozgan/frodao"
)

var TimeNow = time.Now

type DAO[T frodao.Record] struct {
	frodao.DAO[T]
	TableName string
}

func NewDAO[T frodao.Record](tableName string) DAO[T] {
	return DAO[T]{
		TableName: tableName,
	}
}

func (d *DAO[T]) Create(ctx context.Context, t *T) (*T, error) {
	var lastInsertID frodao.ID
	query, _, _ := goqu.Dialect("postgres").Insert(d.TableName).Returning(goqu.I("id")).Rows(t).ToSQL()
	err := SESSION.QueryRow(query).Scan(&lastInsertID)
	if err != nil {
		return nil, err
	}

	return d.FindByID(ctx, frodao.TableID(lastInsertID))
}

func (d *DAO[T]) Update(ctx context.Context, t *T) error {
	if !instanceofTable(t) {
		return errors.New("it should be type of table")
	}
	setUpdatedAt(t)
	query, _, _ := goqu.Dialect("postgres").Update(d.TableName).Set(t).Where(goqu.Ex{"id": id(t)}).ToSQL()
	_, err := SESSION.Exec(query)
	if err != nil {
		return err
	}
	return nil
}

func (d *DAO[T]) Delete(ctx context.Context, id frodao.ID) error {
	query, _, _ := goqu.Dialect("postgres").Update(d.TableName).
		Set(goqu.Record{"deleted": true, "updated_at": TimeNow()}).
		Where(goqu.Ex{"id": id}).
		ToSQL()

	_, err := SESSION.Exec(query)
	if err != nil {
		return err
	}
	return nil
}

func (d *DAO[T]) FindByID(ctx context.Context, id frodao.ID) (*T, error) {
	var dest []*T

	query, _, _ := goqu.Dialect("postgres").From(d.TableName).Where(goqu.Ex{"deleted": false, "id": id}).Limit(1).ToSQL()

	err := sqlscan.Select(ctx, SESSION, &dest, query)
	if err != nil {
		return nil, err
	}

	if len(dest) == 1 {
		return dest[0], nil
	}
	return nil, nil
}

func (d *DAO[T]) SelectQuery() *goqu.SelectDataset {
	return goqu.Dialect("postgres").From(d.TableName)
}

func (d *DAO[T]) FindByQuery(ctx context.Context, q *goqu.SelectDataset) (*T, error) {
	var dest []*T

	query, _, _ := q.Where(goqu.Ex{"deleted": false}).ToSQL()

	err := sqlscan.Select(ctx, SESSION, &dest, query)
	if err != nil {
		return nil, err
	}

	if len(dest) == 1 {
		return dest[0], nil
	}
	return nil, nil
}

func (d DAO[T]) GetTableName() string {
	return d.TableName
}

func findTable[T frodao.Record](t *T) (frodao.Table, bool) {
	val := reflect.ValueOf(t).Elem()
	for i := 0; i < val.NumField(); i++ {
		valueField := val.Field(i)
		typeField := val.Type().Field(i)
		if typeField.Name == "Table" {
			return valueField.Interface().(frodao.Table), true
		}
	}
	return frodao.Table{}, false
}

func instanceofTable[T frodao.Record](t *T) bool {
	_, ok := findTable(t)
	return ok
}

func setUpdatedAt[T frodao.Record](t *T) {
	reflect.ValueOf(t).MethodByName("SetUpdatedAt").Call([]reflect.Value{
		reflect.ValueOf(TimeNow()),
	})
}

func id[T frodao.Record](t *T) frodao.ID {
	if table, ok := findTable(t); ok {
		return table.ID
	}
	return 0
}
