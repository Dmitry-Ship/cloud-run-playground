package users

import (
	"os"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Repository *UserStorage
var Mock sqlmock.Sqlmock

func TestMain(m *testing.M) {
	db, mock, _ := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	Mock = mock

	gdb, _ := gorm.Open(postgres.New(postgres.Config{
		Conn: db,
	}), &gorm.Config{})

	Repository = &UserStorage{db: gdb}
	os.Exit(m.Run())
}

func TestGetAllUsers(t *testing.T) {
	const sqlSelectAll = `SELECT * FROM "users" LIMIT 50`

	Mock.ExpectQuery(sqlSelectAll).
		WillReturnRows(sqlmock.NewRows(nil))

	result, err := Repository.GetAllUsers(0)

	assert.Equal(t, nil, err, "Error occurred")
	assert.Equal(t, 0, len(result), "The two words should be the same.")
}
