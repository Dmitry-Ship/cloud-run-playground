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

func TestCreateUser(t *testing.T) {

	user := User{
		ID:        1,
		FirstName: "John",
		LastName:  "Doe",
		Email:     "john@doe.com",
		Gender:    "male",
		IPAddress: "123",
	}

	const sqlInsert = `INSERT INTO "users" ("first_name","last_name","email","gender","ip_address","id")
                                        VALUES ($1,$2,$3,$4,$5,$6) RETURNING "id"`
	const newId = 1
	Mock.ExpectBegin() // begin transaction
	Mock.ExpectQuery(sqlInsert).
		WithArgs(user.FirstName, user.LastName, user.Email, user.Gender, user.IPAddress, user.ID).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(newId))
	Mock.ExpectCommit() // commit transaction

	result, err := Repository.CreateUser(user)

	assert.Equal(t, nil, err, "Error occurred")
	assert.Equal(t, user, result, "The two words should be the same.")
}

func TestGetUserById(t *testing.T) {
	user := User{
		ID:        1,
		FirstName: "John",
		LastName:  "Doe",
		Email:     "john@doe.com",
		Gender:    "male",
		IPAddress: "123",
	}

	rows := sqlmock.
		NewRows([]string{"id", "first_name", "last_name", "email", "gender", "ip_address"}).
		AddRow(user.ID, user.FirstName, user.LastName, user.Email, user.Gender, user.IPAddress)

	const sqlSelectOne = `SELECT * FROM "users" WHERE "users"."id" = $1`
	Mock.ExpectQuery(sqlSelectOne).WithArgs(user.ID).WillReturnRows(rows)

	result, err := Repository.GetUserById(user.ID)

	assert.Equal(t, nil, err, "Error occurred")
	assert.Equal(t, user, result, "The two users should be the same.")
}
