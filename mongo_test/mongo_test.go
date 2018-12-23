package mongo_test

import (
	"log"
	"testing"

	"github.com/cassul/mongo"
	"github.com/cassul/pkg/mock"

	"github.com/cassul/root"
)

const (
	mongoUrl           = " mongo.zd-dev.com:27017"
	dbName             = "test_db"
	userCollectionName = "user"
)

func Test_UserService(t *testing.T) {
	t.Run("CreateUser", createUser_should_insert_user_into_mongo)
}

func createUser_should_insert_user_into_mongo(t *testing.T) {
	//Arrange
	session, err := mongo.NewSession(mongoUrl)
	if err != nil {
		log.Fatalf("Unable to connect to mongo: %s", err)
	}
	defer session.Close()
	userCollectionPassword := mock.Hash{}
	userService := mongo.NewUserService(session.Copy(), dbName, &userCollectionPassword, userCollectionName)

	testEmail := "integration_test_email"
	testPassword := "integration_test_password"
	user := root.User{
		Email:    testEmail,
		Password: testPassword}

	//Act
	err = userService.Create(&user)

	//Assert
	if err != nil {
		t.Error(err)
	}
	var results []root.User
	session.GetCollection(dbName, userCollectionName).Find(nil).All(&results)

	count := len(results)
	if count != 1 {
		t.Error(count)
	}
	if results[0].Email != user.Email {
		t.Error(testEmail, results[0].Email)
	}

	defer func() {
		err = session.DropDatabase(dbName)
		if err != nil {
			t.Error(err)
		}
		session.Close()
	}()
}
