package mysql_test

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/netologist/frodao"
	"github.com/netologist/frodao/drivers/mysql"
	"github.com/netologist/frodao/tableid"

	"context"
	"testing"
	"time"
)

const testTableName = "test_table"

type TestTable struct {
	frodao.Table[tableid.Int]
	AField string `db:"a_field"`
}

func init() {
	mysql.TimeNow = func() time.Time { return time.Date(1970, time.January, 1, 2, 3, 4, 5, time.UTC) }
}

func TestGetTableName(t *testing.T) {
	dao := mysql.NewDAO[TestTable, tableid.Int](testTableName)
	if dao.GetTableName() != testTableName {
		t.Errorf("table %s is not matched with expected", testTableName)
	}
}

func TestCreate(t *testing.T) {
	mock := connectDB(t)
	defer closeDB()
	dao := mysql.NewDAO[TestTable, tableid.Int]("test")

	mock.ExpectExec("INSERT INTO `test` (`a_field`) VALUES ('some_field_value')").
		WillReturnResult(sqlmock.NewResult(967, 1))

	mock.ExpectQuery("SELECT * FROM `test` WHERE ((`id` = 967) AND (`deleted` IS FALSE)) LIMIT 1").
		WillReturnRows(sqlmock.NewRows([]string{"id", "a_field"}).AddRow(967, "some_field_value"))

	r, err := dao.Create(context.Background(), &TestTable{AField: "some_field_value"})

	if err != nil {
		t.Errorf("Failed with error: %s", err)
	}
	if r == nil || r.ID != frodao.TableIDFromInt(967) || r.AField != "some_field_value" {
		t.Errorf("Record is not matched %v", r)
	}
}

func TestUpdate(t *testing.T) {
	mock := connectDB(t)
	defer closeDB()
	dao := mysql.NewDAO[TestTable, tableid.Int]("test")

	mock.ExpectExec("UPDATE `test` SET `a_field`='new_value',`updated_at`='1970-01-01 02:03:04' WHERE (`id` = 123)").WillReturnResult(sqlmock.NewResult(1, 1))

	record := &TestTable{AField: "some_field_value"}
	record.SetID(frodao.TableIDFromInt(123))
	record.AField = "new_value"

	err := dao.Update(context.Background(), record)

	if err != nil {
		t.Errorf("Failed with error: %s", err)
	}
}

func TestDelete(t *testing.T) {
	mock := connectDB(t)
	defer closeDB()
	dao := mysql.NewDAO[TestTable, tableid.Int]("test")

	mock.ExpectExec("UPDATE `test` SET `deleted`=1,`updated_at`='1970-01-01 02:03:04' WHERE (`id` = 612)").WillReturnResult(sqlmock.NewResult(1, 1))

	err := dao.Delete(context.Background(), frodao.TableIDFromInt(612))

	if err != nil {
		t.Errorf("Failed with error: %s", err)
	}
}

func TestFindByID(t *testing.T) {
	mock := connectDB(t)
	defer closeDB()
	dao := mysql.NewDAO[TestTable, tableid.Int]("test")

	mock.ExpectQuery("SELECT * FROM `test` WHERE ((`id` = 612) AND (`deleted` IS FALSE)) LIMIT 1").
		WillReturnRows(sqlmock.NewRows([]string{"id", "a_field"}).AddRow(612, "another_field_value"))

	r, err := dao.FindByID(context.Background(), frodao.TableIDFromInt(612))

	if err != nil {
		t.Errorf("Failed with error: %s", err)
	}
	if r == nil || r.ID != frodao.TableIDFromInt(612) || r.AField != "another_field_value" {
		t.Errorf("Record is not matched %v", r)
	}
}

func connectDB(t *testing.T) sqlmock.Sqlmock {

	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))

	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	mysql.SetSession(db)
	return mock
}

func closeDB() {
	mysql.Close()
}
