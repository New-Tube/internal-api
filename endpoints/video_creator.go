package endpoints

import (
	"context"
	"internal-api/db"
	db_models "internal-api/db/models"

	pb "github.com/New-Tube/internal-api-protos"
	"github.com/pkg/errors"
)

type videoCreatorUserService struct{ pb.UnimplementedUserServer }

func (s *videoCreatorUserService) Create(ctx context.Context, request *pb.VideoRequest) (*pb.VideoCreateResponse, error) {
	conn, err := db.GetDBConnection()
	if err != nil {
		return nil, errors.Errorf("DB error occured: %v", err)
	}
	model := db_models.Video{
		UserID: request.GetUserID(),
	}
	result := conn.Create(&model)
	if result.Error != nil {
		return nil, errors.Errorf("DB error occured: %v", result.Error)
	}
	return &pb.VideoCreateResponse{
		Success: true,
		ID:      model.ID,
	}, nil
}

func (s *videoCreatorUserService) UpdateInfo(ctx context.Context, request *pb.VideoUpdateRequest) (*pb.StatusResponse, error) {
	conn, err := db.GetDBConnection()
	if err != nil {
		return nil, errors.Errorf("DB error occured: %v", err)
	}

	model := db_models.Video{
		ID: request.GetID(),
	}
	result := conn.Limit(1).Find(&model)
	if result.Error != nil {
		return nil, errors.Errorf("DB error occured: %v", result.Error)
	}
	model.Title = request.GetTitle()
	model.Description = request.GetDescription()
	model.Privacy = uint16(request.GetPrivacy())
	model.Link = request.GetLink()

	updateVideo := result.Save(&model)

	if updateVideo.Error != nil {
		return nil, errors.Errorf("DB error occured: %v", updateVideo.Error)
	}

	return &pb.StatusResponse{
		Success: true,
		Message: "OK",
	}, nil
}

func (s *videoCreatorUserService) GetUploadLink(ctx context.Context, request *pb.VideoRequest) (*pb.StatusResponse, error) {
	return &pb.StatusResponse{
		Success: false,
		Message: "NOT IMPLEMENTED",
	}, nil
}

func (s *videoCreatorUserService) Delete(ctx context.Context, request *pb.VideoRequest) (*pb.StatusResponse, error) {
	conn, err := db.GetDBConnection()
	if err != nil {
		return nil, errors.Errorf("DB error occured: %v", err)
	}

	model := db_models.Video{
		ID: request.GetID(),
	}
	result := conn.Limit(1).Find(&model)

	if result.RowsAffected < 1 {
		return nil, errors.Errorf("Comment with provided id not found")
	}

	deleteResult := conn.Delete(&model)

	if deleteResult.Error != nil {
		return nil, errors.Errorf("DB error occured: %v", deleteResult.Error)
	}

	return &pb.StatusResponse{
		Success: true,
		Message: "",
	}, nil
}
