package repository

import (
	"os"
	"testing"

	"github.com/hemilioaraujo/first-go-crud/src/model"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/integration/mtest"
)

func TestUserRepository_UpdateUser(t *testing.T) {
	database_name := "user_database_test"
	collection_name := "user_collection_test"
	err := os.Setenv("MONGODB_USER_DATABASE", collection_name)
	if err != nil {
		t.FailNow()
	}
	defer os.Clearenv()

	mtestDb := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))

	mtestDb.Run("when_sending_a_valid_domain_returns_success", func(mt *mtest.T) {
		mt.AddMockResponses(bson.D{
			{Key: "ok", Value: 1},
			{Key: "n", Value: 1},
			{Key: "acknowledged", Value: true},
		})
		databaseMock := mt.Client.Database(database_name)
		repo := NewUserRepository(databaseMock)
		domain := model.NewUserDomain(
			"john.doe@example.com",
			"John Doe",
			"password",
			15,
		)
		err := repo.UpdateUser("678901234567890123456789", domain)

		assert.Nil(t, err)
	})

	mtestDb.Run("when_mongodb_returns_error", func(mt *mtest.T) {
		mt.AddMockResponses(bson.D{
			{Key: "ok", Value: 0},
		})
		databaseMock := mt.Client.Database(database_name)
		repo := NewUserRepository(databaseMock)
		domain := model.NewUserDomain(
			"john.doe@example.com",
			"John Doe",
			"password",
			15,
		)
		err := repo.UpdateUser("678901234567890123456789", domain)

		assert.NotNil(t, err)
	})
}
