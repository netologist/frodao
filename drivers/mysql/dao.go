package mysql

import (
	"context"
	"errors"
	"reflect"
	"time"

	"github.com/doug-martin/goqu/v9"
	_ "github.com/doug-martin/goqu/v9/dialect/mysql"
	"github.com/georgysavva/scany/sqlscan"

	"github.com/hasanozgan/frodao"
	"github.com/hasanozgan/frodao/tableid"
)

var TimeNow = time.Now

type DAO[T frodao.Record, I tableid.Constraint] struct {
	frodao.DAO[T, I]
	TableName string
}

func NewDAO[T frodao.Record, I tableid.Constraint](tableName string) DAO[T, I] {
	return DAO[T, I]{
		TableName: tableName,
	}
}

func (d *DAO[T, I]) ConvertID(v int64) I {
	rv := reflect.ValueOf(v)
	return rv.Interface().(I)
}

func (d *DAO[T, I]) Create(ctx context.Context, t *T) (*T, error) {
	query, _, _ := goqu.Dialect("mysql").Insert(d.TableName).Rows(t).ToSQL()
	result, err := SESSION.Exec(query)
	if err != nil {
		return nil, err
	}
	lastInsertID, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}
	rv := reflect.ValueOf(lastInsertID)
	return d.FindByID(ctx, frodao.TableID(rv.Interface().(I)))
}

func (d *DAO[T, I]) Update(ctx context.Context, t *T) error {
	if !instanceofTable[T, I](t) {
		return errors.New("it should be type of table")
	}
	setUpdatedAt[T, I](t)
	query, _, _ := goqu.Dialect("mysql").Update(d.TableName).Set(t).Where(goqu.Ex{"id": id[T, I](t)}).ToSQL()
	_, err := SESSION.Exec(query)
	if err != nil {
		return err
	}
	return nil
}

func (d *DAO[T, I]) Delete(ctx context.Context, id frodao.ID[I]) error {
	query, _, _ := goqu.Dialect("mysql").Update(d.TableName).
		Set(goqu.Record{"deleted": true, "updated_at": TimeNow()}).
		Where(goqu.Ex{"id": id}).
		ToSQL()

	_, err := SESSION.Exec(query)
	if err != nil {
		return err
	}
	return nil
}

func (d *DAO[T, I]) FindByID(ctx context.Context, id frodao.ID[I]) (*T, error) {
	if dest, err := d.FindByQuery(ctx, d.SelectQuery().Where(goqu.Ex{"id": id}).Limit(1)); err != nil {
		return nil, err
	} else if len(dest) == 1 {
		return dest[0], nil
	}

	return nil, nil
}

func (d *DAO[T, I]) FirstRow(rows []*T, err error) (*T, error) {
	if err != nil {
		return nil, err
	}
	if len(rows) == 1 {
		return rows[0], nil
	}
	return nil, nil
}

func (d *DAO[T, I]) SelectQuery() *goqu.SelectDataset {
	return goqu.Dialect("mysql").From(d.TableName)
}

func (d *DAO[T, I]) FindByQuery(ctx context.Context, q *goqu.SelectDataset) ([]*T, error) {
	var dest []*T

	query, _, _ := q.Where(goqu.Ex{"deleted": false}).ToSQL()

	err := sqlscan.Select(ctx, SESSION, &dest, query)
	if err != nil {
		return nil, err
	}

	return dest, nil
}

func (d DAO[T, I]) GetTableName() string {
	return d.TableName
}

func findTable[T frodao.Record, I tableid.Constraint](t *T) (frodao.Table[I], bool) {
	val := reflect.ValueOf(t).Elem()
	for i := 0; i < val.NumField(); i++ {
		valueField := val.Field(i)
		typeField := val.Type().Field(i)
		if typeField.Name == "Table" {
			return valueField.Interface().(frodao.Table[I]), true
		}
	}
	return frodao.Table[I]{}, false
}

func instanceofTable[T frodao.Record, I tableid.Constraint](t *T) bool {
	_, ok := findTable[T, I](t)
	return ok
}

func setUpdatedAt[T frodao.Record, I tableid.Constraint](t *T) {
	reflect.ValueOf(t).MethodByName("SetUpdatedAt").Call([]reflect.Value{
		reflect.ValueOf(TimeNow()),
	})
}

func id[T frodao.Record, I tableid.Constraint](t *T) frodao.ID[I] {
	if table, ok := findTable[T, I](t); ok {
		return table.ID
	}
	return frodao.ID[I]{}
}
