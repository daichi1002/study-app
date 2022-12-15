package article

import (
	"backend/domain/model"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestCreateArticle(t *testing.T) {
	repository, mock, err := NewDbMock()
	if err != nil {
		t.Errorf("Failed to initialize mock DB: %v", err)
	}
	// mock設定
	testData := &model.Article{
		Id:        "01GKTYV0YZ3MK4ZK34HVWFHAZH",
		Title:     "テストタイトル1",
		Content:   "テストコンテンツ1",
		UserId:    "01GKTYV0YZ3MK4ZK34HVWFHAZD",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta(
		`INSERT INTO "articles" ("id","title","content","user_id","created_at","updated_at") VALUES ($1,$2,$3,$4,$5,$6)`)).
		WithArgs("01GKTYV0YZ3MK4ZK34HVWFHAZH", "テストタイトル1", "テストコンテンツ1", "01GKTYV0YZ3MK4ZK34HVWFHAZD", sqlmock.AnyArg(), sqlmock.AnyArg()).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	// テスト対象の関数実行
	err = repository.CreateArticle(testData)

	if err != nil {
		t.Fatal(err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Test Create User: %v", err)
	}
}

// TODO
func TestUpdateArticle(t *testing.T) {
}

// TODO
func TestDeleteArticle(t *testing.T) {
}
