package tests

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHashPassword(t *testing.T) {
	user := User{
		Password: "secret",
	}

	err := user.HashPassword(user.Password)
	assert.NoError(t, err)

	os.Setenv("passwordHash", user.Password)
}

func TestCreateUserRecord(t *testing.T) {
	var userResult User

	err := InitDatabase()
	if err != nil {
		t.Error(err)
	}

	err = GlobalDB.AutoMigrate(&User{})
	assert.NoError(t, err)

	user := User{
		Name:     "Test User",
		Email:    "test@email.com",
		Password: os.Getenv("passwordHash"),
	}

	err = user.CreateUserRecord()
	assert.NoError(t, err)

	GlobalDB.Where("email = ?", user.Email).Find(&userResult)

	GlobalDB.Unscoped().Delete(&user)

	assert.Equal(t, "Test User", userResult.Name)
	assert.Equal(t, "test@email.com", userResult.Email)

}

func TestCheckPassword(t *testing.T) {
	hash := "$2y$10$viTMrllQZStSgNcznSS.6uxUNUZZGKbQlRf0C4OTeGajz8ONUnqIW"

	user := User{
		Password: hash,
	}

	err := user.CheckPassword("samiroquai")
	assert.NoError(t, err)
}
