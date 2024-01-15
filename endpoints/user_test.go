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

func TestUserGetByID(t *testing.T) {
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

	userModel := db_models.User{
		Name:         "TestUser",
		Surname:      "TEst",
		Nickname:     "User",
		PasswordHash: 5665,
	}

	result := conn.Create(&userModel)
	if result.Error != nil {
		t.Errorf("Create user error: %v", result.Error)
	}

	req := internal_api_protos.UserRequest{
		ID: userModel.ID,
	}
	resp, err := us.Get(ctx, &req)

	if err != nil {
		t.Errorf("Get by id request error: %v", err)
	}

	if resp.ID != userModel.ID {
		t.Errorf("ID differ: expected: %d, got: %d", resp.ID, userModel.ID)
	}
	if resp.Nickname != userModel.Nickname {
		t.Errorf("Nickname differ: expected: %s, got: %s", resp.Nickname, userModel.Nickname)
	}
	if resp.PasswordHash != userModel.PasswordHash {
		t.Errorf("PasswordHash differ: expected: %d, got: %d", resp.PasswordHash, userModel.PasswordHash)
	}
	if resp.Name != userModel.Name {
		t.Errorf("Name differ: expected: %s, got: %s", resp.Name, userModel.Nickname)
	}
	if resp.Surname != userModel.Surname {
		t.Errorf("Surname differ: expected: %s, got: %s", resp.Surname, userModel.Nickname)
	}
}

func TestUserGetByIDFill(t *testing.T) {
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

	userModel := db_models.User{
		Name:         "TestUser",
		Surname:      "TEst",
		Nickname:     "User",
		PasswordHash: 5665,
	}

	err = fillUsers(0)
	if err != nil {
		t.Errorf("Fill error: %v", err)
	}

	result := conn.Create(&userModel)
	if result.Error != nil {
		t.Errorf("Create user error: %v", result.Error)
	}

	err = fillUsers(3)
	if err != nil {
		t.Errorf("Fill error: %v", err)
	}

	req := internal_api_protos.UserRequest{
		ID: userModel.ID,
	}
	resp, err := us.Get(ctx, &req)

	if err != nil {
		t.Errorf("Get by id request error: %v", err)
	}

	if resp.ID != userModel.ID {
		t.Errorf("ID differ: expected: %d, got: %d", resp.ID, userModel.ID)
	}
	if resp.Nickname != userModel.Nickname {
		t.Errorf("Nickname differ: expected: %s, got: %s", resp.Nickname, userModel.Nickname)
	}
	if resp.PasswordHash != userModel.PasswordHash {
		t.Errorf("PasswordHash differ: expected: %d, got: %d", resp.PasswordHash, userModel.PasswordHash)
	}
	if resp.Name != userModel.Name {
		t.Errorf("Name differ: expected: %s, got: %s", resp.Name, userModel.Nickname)
	}
	if resp.Surname != userModel.Surname {
		t.Errorf("Surname differ: expected: %s, got: %s", resp.Surname, userModel.Nickname)
	}
}

func TestUserGetByNickname(t *testing.T) {
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

	userModel := db_models.User{
		Name:         "TestUser",
		Surname:      "TEst",
		Nickname:     "User",
		PasswordHash: 5665,
	}

	result := conn.Create(&userModel)
	if result.Error != nil {
		t.Errorf("Create user error: %v", result.Error)
	}

	req := internal_api_protos.UserRequest{
		Nickname: userModel.Nickname,
	}
	resp, err := us.Get(ctx, &req)

	if err != nil {
		t.Errorf("Get by id request error: %v", err)
	}

	if resp.ID != userModel.ID {
		t.Errorf("ID differ: expected: %d, got: %d", resp.ID, userModel.ID)
	}
	if resp.Nickname != userModel.Nickname {
		t.Errorf("Nickname differ: expected: %s, got: %s", resp.Nickname, userModel.Nickname)
	}
	if resp.PasswordHash != userModel.PasswordHash {
		t.Errorf("PasswordHash differ: expected: %d, got: %d", resp.PasswordHash, userModel.PasswordHash)
	}
	if resp.Name != userModel.Name {
		t.Errorf("Name differ: expected: %s, got: %s", resp.Name, userModel.Nickname)
	}
	if resp.Surname != userModel.Surname {
		t.Errorf("Surname differ: expected: %s, got: %s", resp.Surname, userModel.Nickname)
	}
}

func TestUserGetByNicknameFill(t *testing.T) {
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

	userModel := db_models.User{
		Name:         "TestUser",
		Surname:      "TEst",
		Nickname:     "User",
		PasswordHash: 5665,
	}

	err = fillUsers(0)
	if err != nil {
		t.Errorf("Fill error: %v", err)
	}

	result := conn.Create(&userModel)
	if result.Error != nil {
		t.Errorf("Create user error: %v", result.Error)
	}

	err = fillUsers(3)
	if err != nil {
		t.Errorf("Fill error: %v", err)
	}

	req := internal_api_protos.UserRequest{
		Nickname: userModel.Nickname,
	}
	resp, err := us.Get(ctx, &req)

	if err != nil {
		t.Errorf("Get by id request error: %v", err)
	}

	if resp.ID != userModel.ID {
		t.Errorf("ID differ: expected: %d, got: %d", resp.ID, userModel.ID)
	}
	if resp.Nickname != userModel.Nickname {
		t.Errorf("Nickname differ: expected: %s, got: %s", resp.Nickname, userModel.Nickname)
	}
	if resp.PasswordHash != userModel.PasswordHash {
		t.Errorf("PasswordHash differ: expected: %d, got: %d", resp.PasswordHash, userModel.PasswordHash)
	}
	if resp.Name != userModel.Name {
		t.Errorf("Name differ: expected: %s, got: %s", resp.Name, userModel.Nickname)
	}
	if resp.Surname != userModel.Surname {
		t.Errorf("Surname differ: expected: %s, got: %s", resp.Surname, userModel.Nickname)
	}
}

func TestUserUpdate(t *testing.T) {
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

	userModel := db_models.User{
		Name:         "TestUser",
		Surname:      "TEst",
		Nickname:     "User",
		PasswordHash: 5665,
	}

	result := conn.Create(&userModel)
	if result.Error != nil {
		t.Errorf("Create user error: %v", result.Error)
	}

	req := internal_api_protos.UserUpdateRequest{
		ID:           userModel.ID,
		Name:         "UserTest",
		Surname:      "NO",
		Nickname:     "None",
		PasswordHash: 9999,
	}
	resp, err := us.Update(ctx, &req)

	if err != nil {
		t.Errorf("Update request error: %v", err)
	}

	if resp.Success == false {
		t.Errorf("Update request success false")
	}

	newModel := db_models.User{
		ID: userModel.ID,
	}

	result = conn.Limit(1).Find(&newModel)

	if result.Error != nil {
		t.Errorf("Get user after update request error: %v", err)
	}

	if newModel.ID != req.ID {
		t.Errorf("ID differ: expected: %d, got: %d", newModel.ID, req.ID)
	}
	if newModel.Nickname != req.Nickname {
		t.Errorf("Nickname differ: expected: %s, got: %s", newModel.Nickname, req.Nickname)
	}
	if newModel.PasswordHash != req.PasswordHash {
		t.Errorf("PasswordHash differ: expected: %d, got: %d", newModel.PasswordHash, req.PasswordHash)
	}
	if newModel.Name != req.Name {
		t.Errorf("Name differ: expected: %s, got: %s", newModel.Name, req.Nickname)
	}
	if newModel.Surname != req.Surname {
		t.Errorf("Surname differ: expected: %s, got: %s", newModel.Surname, req.Nickname)
	}
}

func TestUserUpdateFill(t *testing.T) {
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

	userModel := db_models.User{
		Name:         "TestUser",
		Surname:      "TEst",
		Nickname:     "User",
		PasswordHash: 5665,
	}

	err = fillUsers(0)
	if err != nil {
		t.Errorf("Fill error: %v", err)
	}

	result := conn.Create(&userModel)
	if result.Error != nil {
		t.Errorf("Create user error: %v", result.Error)
	}

	err = fillUsers(3)
	if err != nil {
		t.Errorf("Fill error: %v", err)
	}

	req := internal_api_protos.UserUpdateRequest{
		ID:           userModel.ID,
		Name:         "UserTest",
		Surname:      "NO",
		Nickname:     "None",
		PasswordHash: 9999,
	}
	resp, err := us.Update(ctx, &req)

	if err != nil {
		t.Errorf("Update request error: %v", err)
	}

	if resp.Success == false {
		t.Errorf("Update request success false")
	}

	newModel := db_models.User{
		ID: userModel.ID,
	}

	result = conn.Limit(1).Find(&newModel)

	if result.Error != nil {
		t.Errorf("Get user after update request error: %v", err)
	}

	if newModel.ID != req.ID {
		t.Errorf("ID differ: expected: %d, got: %d", newModel.ID, req.ID)
	}
	if newModel.Nickname != req.Nickname {
		t.Errorf("Nickname differ: expected: %s, got: %s", newModel.Nickname, req.Nickname)
	}
	if newModel.PasswordHash != req.PasswordHash {
		t.Errorf("PasswordHash differ: expected: %d, got: %d", newModel.PasswordHash, req.PasswordHash)
	}
	if newModel.Name != req.Name {
		t.Errorf("Name differ: expected: %s, got: %s", newModel.Name, req.Nickname)
	}
	if newModel.Surname != req.Surname {
		t.Errorf("Surname differ: expected: %s, got: %s", newModel.Surname, req.Nickname)
	}
}
