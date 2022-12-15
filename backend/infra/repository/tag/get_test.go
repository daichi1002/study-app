package tag

import (
	"backend/domain/repository"
	"backend/infra"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDbMock() (repository.TagRepository, sqlmock.Sqlmock, error) {
	sqlDB, mock, _ := sqlmock.New()
	mockDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn: sqlDB,
	}), &gorm.Config{})

	// repository 初期化
	gormHandler := new(infra.GormHandler)
	gormHandler.DB = mockDB
	repository := NewTagRepository(gormHandler)
	return repository, mock, err
}

func TestListUsers(t *testing.T) {
	repository, mock, err := NewDbMock()
	if err != nil {
		t.Errorf("Failed to initialize mock DB: %v", err)
	}

	// mock設定
	rows := sqlmock.NewRows([]string{"id", "name", "created_at", "updated_at"}).
		AddRow(1, "test1", time.Now(), time.Now()).
		AddRow(2, "test2", time.Now(), time.Now())

		// 期待されるSQL
	mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "tags"`)).
		WillReturnRows(rows)

	// テスト対象の関数実行
	_, err = repository.ListTags()

	if err != nil {
		t.Fatal(err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Test Find Users: %v", err)
	}
}
