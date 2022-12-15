package user

import (
	"backend/infra"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDbMock() (*gorm.DB, sqlmock.Sqlmock, error) {
	sqlDB, mock, _ := sqlmock.New()
	mockDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn: sqlDB,
	}), &gorm.Config{})
	return mockDB, mock, err
}

func TestListUsers(t *testing.T) {
	mockDB, mock, err := NewDbMock()
	if err != nil {
		t.Errorf("Failed to initialize mock DB: %v", err)
	}

	// mock設定
	articleRows := sqlmock.NewRows([]string{"id", "title", "content", "user_id", "created_at", "updated_at"}).
		AddRow("01GKTYV0YZ3MK4ZK34HVWFHAZH", "テストタイトル1", "テストコンテンツ1", "01GKTYV0YZ3MK4ZK34HVWFHAZD", time.Now(), time.Now())

	userRows := sqlmock.NewRows([]string{"id", "user_name", "email", "password", "created_at", "updated_at"}).
		AddRow("01GKTYV0YZ3MK4ZK34HVWFHAZD", "テストユーザー1", "example@gmail.com", "testpassword", time.Now(), time.Now())

		// 期待されるSQL
	mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "users"`)).
		WillReturnRows(userRows)
	mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "articles" WHERE "articles"."user_id" = $1`)).
		WithArgs("01GKTYV0YZ3MK4ZK34HVWFHAZD").
		WillReturnRows(articleRows)

	// repository 初期化
	gormHandler := new(infra.GormHandler)
	gormHandler.DB = mockDB
	repository := NewUserRepository(gormHandler)

	// テスト対象の関数実行
	_, err = repository.ListUsers()

	if err != nil {
		t.Fatal(err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Test Find Users: %v", err)
	}
}

func TestGetLoginUser(t *testing.T) {
	mockDB, mock, err := NewDbMock()
	if err != nil {
		t.Errorf("Failed to initialize mock DB: %v", err)
	}

	// mock設定
	articleRows := sqlmock.NewRows([]string{"id", "title", "content", "user_id", "created_at", "updated_at"}).
		AddRow("01GKTYV0YZ3MK4ZK34HVWFHAZH", "テストタイトル2", "テストコンテンツ2", "01GKTYV0YZ3MK4ZK34HVWFHAZD", time.Now(), time.Now())

	userRows := sqlmock.NewRows([]string{"id", "user_name", "email", "password", "created_at", "updated_at"}).
		AddRow("01GKTYV0YZ3MK4ZK34HVWFHAZD", "テストユーザー2", "example@gmail.com", "testpassword", time.Now(), time.Now())

		// 期待されるSQL
	mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "users" WHERE email = $1 ORDER BY "users"."id" LIMIT 1`)).
		WithArgs("example@gmail.com").
		WillReturnRows(userRows)
	mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "articles" WHERE "articles"."user_id" = $1`)).
		WithArgs("01GKTYV0YZ3MK4ZK34HVWFHAZD").
		WillReturnRows(articleRows)

	// repository 初期化
	gormHandler := new(infra.GormHandler)
	gormHandler.DB = mockDB
	repository := NewUserRepository(gormHandler)

	// テスト対象の関数実行
	_, err = repository.GetLoginUser("example@gmail.com")

	if err != nil {
		t.Fatal(err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Test Find Users: %v", err)
	}
}
