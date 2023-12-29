package endpoints

import (
	"context"
	"internal-api/db"
	db_models "internal-api/db/models"
	"log"
	"testing"

	internal_api_protos "github.com/New-Tube/internal-api-protos"
	"github.com/joho/godotenv"
)

func TestUserCreate(t *testing.T) {
	err := godotenv.Load("../.env.tests")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	resetDB()
	db.Migrate()

	ctx := context.TODO()
	us := userServer{}

	conn, err := db.GetDBConnection()
	if err != nil {
		t.Errorf("DB error: %v", err)
	}

	model := db_models.User{
		ID: 1,
	}
	result := conn.Limit(1).Find(&model)

	if result.RowsAffected != 0 {
		t.Error("User already found")
	}

	req := internal_api_protos.UserCreateRequest{
		Nickname:     "user_nickname",
		PasswordHash: 5678,
		Name:         "FirstName",
		Surname:      "LastName",
	}
	resp, err := us.Create(ctx, &req)

	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	result = conn.Limit(1).Find(&model)
	if result.RowsAffected != 1 {
		t.Error("User wasn't created")
	}

	if model.Nickname != req.Nickname {
		t.Errorf("Nickname differ: expected: %s, got: %s", req.Nickname, model.Nickname)
	}
	if model.PasswordHash != req.PasswordHash {
		t.Errorf("PasswordHash differ: expected: %d, got: %d", req.PasswordHash, model.PasswordHash)
	}
	if model.Name != req.Name {
		t.Errorf("Name differ: expected: %s, got: %s", req.Name, model.Nickname)
	}
	if model.Surname != req.Surname {
		t.Errorf("Surname differ: expected: %s, got: %s", req.Surname, model.Nickname)
	}

	_ = resp
}
