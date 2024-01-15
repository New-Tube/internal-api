package endpoints

import (
	"context"
	"internal-api/db"
	db_models "internal-api/db/models"

	pb "github.com/New-Tube/internal-api-protos"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
)

type userServer struct {
	pb.UnimplementedUserServer
}

func (s *userServer) Get(ctx context.Context, request *pb.UserRequest) (*pb.UserResponse, error) {
	conn, err := db.GetDBConnection()
	if err != nil {
		return nil, errors.Errorf("DB error occured: %v", err)
	}
	var model db_models.User
	if request.GetID() != 0 {
		model = db_models.User{
			ID: request.GetID(),
		}
		result := conn.Limit(1).Find(&model)
		if result.RowsAffected < 1 {
			return nil, errors.Errorf("DB error occured: User not found")
		}
	} else {
		model = db_models.User{}
		_, rows_affected := applyFilters[db_models.User](map[string]any{
			"nickname": request.GetNickname(),
		}, conn, &model)
		if rows_affected < 1 {
			return nil, errors.Errorf("DB error occured: User not found")
		}
	}

	return &pb.UserResponse{
		ID:           model.ID,
		Name:         model.Name,
		Surname:      model.Surname,
		Nickname:     model.Nickname,
		PasswordHash: model.PasswordHash,
		CreatedAt:    uint64(model.CreatedAt.Unix()),
		EditedAt:     uint64(model.EditedAt.Unix()),
	}, nil
}

func (s *userServer) Update(ctx context.Context, request *pb.UserUpdateRequest) (*pb.StatusResponse, error) {
	conn, err := db.GetDBConnection()
	if err != nil {
		return nil, errors.Errorf("DB error occured: %v", err)
	}

	model := db_models.User{
		ID: request.GetID(),
	}
	result := conn.Limit(1).Find(&model)
	if result.Error != nil {
		return nil, errors.Errorf("DB error occured: %v", result.Error)
	}

	model.Name = request.GetName()
	model.Surname = request.GetSurname()
	model.Nickname = request.GetNickname()
	model.PasswordHash = request.GetPasswordHash()

	updateUser := conn.Save(&model)

	if updateUser.Error != nil {
		return nil, errors.Errorf("DB error occured: %v", updateUser.Error)
	}

	return &pb.StatusResponse{
		Success: true,
		Message: "OK",
	}, nil
}

func (s *userServer) Create(ctx context.Context, request *pb.UserCreateRequest) (*pb.UserResponse, error) {
	conn, err := db.GetDBConnection()
	if err != nil {
		return nil, errors.Errorf("DB error occured: %v", err)
	}

	model := db_models.User{
		Name:         request.GetName(),
		Surname:      request.GetSurname(),
		Nickname:     request.GetNickname(),
		PasswordHash: request.GetPasswordHash(),
	}

	result := conn.Create(&model)
	if result.Error != nil {
		return nil, errors.Errorf("DB error occured: %v", result.Error)
	}
	return &pb.UserResponse{
		ID:           model.ID,
		Name:         model.Name,
		Surname:      model.Surname,
		Nickname:     model.Nickname,
		PasswordHash: model.PasswordHash,
		CreatedAt:    uint64(model.CreatedAt.Unix()),
		EditedAt:     uint64(model.EditedAt.Unix()),
	}, nil
}

func RegisterUserService(s *grpc.Server) {
	pb.RegisterUserServer(s, &userServer{})
}
