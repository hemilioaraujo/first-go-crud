package repository

import (
	"fmt"
	"os"
	"testing"

	"github.com/hemilioaraujo/first-go-crud/src/model/repository/entity"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/integration/mtest"
)

func TestUserRepository_FindUserByEmail(t *testing.T) {
	database_name := "user_database_test"
	collection_name := "user_collection_test"
	err := os.Setenv("MONGODB_USER_DATABASE", collection_name)
	if err != nil {
		t.FailNow()
	}
	defer os.Clearenv()

	mtestDb := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))

	mtestDb.Run("when_sending_a_valid_email_returns_success", func(mt *mtest.T) {
		userEntity := entity.UserEntity{
			ID:       primitive.NewObjectID(),
			Email:    "john.doe@example.com",
			Name:     "John Doe",
			Password: "password",
			Age:      15,
		}
		mt.AddMockResponses(mtest.CreateCursorResponse(
			1,
			fmt.Sprintf("%s.%s", database_name, collection_name),
			mtest.FirstBatch,
			convertEntityToBson(userEntity),
		))
		databaseMock := mt.Client.Database(database_name)
		repo := NewUserRepository(databaseMock)
		userDomain, err := repo.FindUserByEmail(userEntity.Email)

		assert.Nil(t, err)
		assert.EqualValues(t, userDomain.GetEmail(), userEntity.Email)
		assert.EqualValues(t, userDomain.GetName(), userEntity.Name)
		assert.EqualValues(t, userDomain.GetAge(), userEntity.Age)
		assert.EqualValues(t, userDomain.GetPassword(), userEntity.Password)
		assert.EqualValues(t, userDomain.GetId(), userEntity.ID.Hex())
	})

	mtestDb.Run("when_mongodb_returns_error", func(mt *mtest.T) {
		mt.AddMockResponses(bson.D{
			{Key: "ok", Value: 0},
		})
		databaseMock := mt.Client.Database(database_name)
		repo := NewUserRepository(databaseMock)
		userDomain, err := repo.FindUserByEmail("john.doe@example.com")

		assert.NotNil(t, err)
		assert.Equal(t, err.Message, "error trying to find user with email john.doe@example.com")
		assert.Nil(t, userDomain)
	})

	mtestDb.Run("when_mongodb_returns_no_results", func(mt *mtest.T) {
		userEntity := entity.UserEntity{
			ID:       primitive.NewObjectID(),
			Email:    "john.doe@example.com",
			Name:     "John Doe",
			Password: "password",
			Age:      15,
		}
		mt.AddMockResponses(mtest.CreateCursorResponse(
			0,
			fmt.Sprintf("%s.%s", database_name, collection_name),
			mtest.FirstBatch,
		))
		databaseMock := mt.Client.Database(database_name)
		repo := NewUserRepository(databaseMock)
		userDomain, err := repo.FindUserByEmail(userEntity.Email)

		assert.NotNil(t, err)
		assert.Equal(t, err.Message, "user not found with email john.doe@example.com")
		assert.Nil(t, userDomain)
	})

}

func TestUserRepository_FindUserByID(t *testing.T) {
	database_name := "user_database_test"
	collection_name := "user_collection_test"
	err := os.Setenv("MONGODB_USER_DATABASE", collection_name)
	if err != nil {
		t.FailNow()
	}
	defer os.Clearenv()

	mtestDb := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))

	mtestDb.Run("when_sending_a_valid_id_returns_success", func(mt *mtest.T) {
		userEntity := entity.UserEntity{
			ID:       primitive.NewObjectID(),
			Email:    "john.doe@example.com",
			Name:     "John Doe",
			Password: "password",
			Age:      15,
		}
		mt.AddMockResponses(mtest.CreateCursorResponse(
			1,
			fmt.Sprintf("%s.%s", database_name, collection_name),
			mtest.FirstBatch,
			convertEntityToBson(userEntity),
		))
		databaseMock := mt.Client.Database(database_name)
		repo := NewUserRepository(databaseMock)
		userDomain, err := repo.FindUserById(userEntity.ID.Hex())

		assert.Nil(t, err)
		assert.EqualValues(t, userDomain.GetEmail(), userEntity.Email)
		assert.EqualValues(t, userDomain.GetName(), userEntity.Name)
		assert.EqualValues(t, userDomain.GetAge(), userEntity.Age)
		assert.EqualValues(t, userDomain.GetPassword(), userEntity.Password)
		assert.EqualValues(t, userDomain.GetId(), userEntity.ID.Hex())
	})

	mtestDb.Run("when_mongodb_returns_error", func(mt *mtest.T) {
		mt.AddMockResponses(bson.D{
			{Key: "ok", Value: 0},
		})
		databaseMock := mt.Client.Database(database_name)
		repo := NewUserRepository(databaseMock)
		userDomain, err := repo.FindUserById("678901234567890123456789")

		assert.NotNil(t, err)
		assert.Equal(t, err.Message, "error trying to find user with id 678901234567890123456789")
		assert.Nil(t, userDomain)
	})

	mtestDb.Run("when_mongodb_returns_no_results", func(mt *mtest.T) {
		mt.AddMockResponses(mtest.CreateCursorResponse(
			0,
			fmt.Sprintf("%s.%s", database_name, collection_name),
			mtest.FirstBatch,
		))
		databaseMock := mt.Client.Database(database_name)
		repo := NewUserRepository(databaseMock)
		userDomain, err := repo.FindUserById("678901234567890123456789")

		assert.NotNil(t, err)
		assert.Equal(t, err.Message, "user not found with id 678901234567890123456789")
		assert.Nil(t, userDomain)
	})
}

func TestUserRepository_FindUserByEmailAndPassword(t *testing.T) {
	database_name := "user_database_test"
	collection_name := "user_collection_test"
	err := os.Setenv("MONGODB_USER_DATABASE", collection_name)
	if err != nil {
		t.FailNow()
	}
	defer os.Clearenv()

	mtestDb := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))

	mtestDb.Run("when_sending_a_valid_email_and_password_returns_success", func(mt *mtest.T) {
		userEntity := entity.UserEntity{
			ID:       primitive.NewObjectID(),
			Email:    "john.doe@example.com",
			Name:     "John Doe",
			Password: "password",
			Age:      15,
		}
		mt.AddMockResponses(mtest.CreateCursorResponse(
			1,
			fmt.Sprintf("%s.%s", database_name, collection_name),
			mtest.FirstBatch,
			convertEntityToBson(userEntity),
		))
		databaseMock := mt.Client.Database(database_name)
		repo := NewUserRepository(databaseMock)
		userDomain, err := repo.FindUserByEmailAndPassword(userEntity.Email, userEntity.Password)

		assert.Nil(t, err)
		assert.EqualValues(t, userDomain.GetEmail(), userEntity.Email)
		assert.EqualValues(t, userDomain.GetName(), userEntity.Name)
		assert.EqualValues(t, userDomain.GetAge(), userEntity.Age)
		assert.EqualValues(t, userDomain.GetPassword(), userEntity.Password)
		assert.EqualValues(t, userDomain.GetId(), userEntity.ID.Hex())
	})

	mtestDb.Run("when_mongodb_returns_error", func(mt *mtest.T) {
		mt.AddMockResponses(bson.D{
			{Key: "ok", Value: 0},
		})
		databaseMock := mt.Client.Database(database_name)
		repo := NewUserRepository(databaseMock)
		userDomain, err := repo.FindUserByEmailAndPassword("john.doe@example.com", "password")

		assert.NotNil(t, err)
		assert.Equal(t, err.Message, "error trying to find user with email and password")
		assert.Nil(t, userDomain)
	})

	mtestDb.Run("when_mongodb_returns_no_results", func(mt *mtest.T) {
		userEntity := entity.UserEntity{
			ID:       primitive.NewObjectID(),
			Email:    "john.doe@example.com",
			Name:     "John Doe",
			Password: "password",
			Age:      15,
		}
		mt.AddMockResponses(mtest.CreateCursorResponse(
			0,
			fmt.Sprintf("%s.%s", database_name, collection_name),
			mtest.FirstBatch,
		))
		databaseMock := mt.Client.Database(database_name)
		repo := NewUserRepository(databaseMock)
		userDomain, err := repo.FindUserByEmailAndPassword(userEntity.Email, userEntity.Password)

		assert.NotNil(t, err)
		assert.Equal(t, err.Message, "email or password is incorrect")
		assert.Nil(t, userDomain)
	})

}

func convertEntityToBson(userEntity entity.UserEntity) bson.D {
	return bson.D{
		{Key: "email", Value: userEntity.Email},
		{Key: "name", Value: userEntity.Name},
		{Key: "password", Value: userEntity.Password},
		{Key: "age", Value: userEntity.Age},
		{Key: "_id", Value: userEntity.ID},
	}
}
