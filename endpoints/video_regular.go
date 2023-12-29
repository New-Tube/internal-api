package endpoints

import (
	"context"
	"internal-api/db"
	db_models "internal-api/db/models"

	pb "github.com/New-Tube/internal-api-protos"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
)

type videoRegularUserService struct {
	pb.UnimplementedVideoRegularUserServer
}

func dbPrivacyToEnum(privacy uint16) db_models.Privacy {
	switch privacy {
	case 1:
		return db_models.PrivacyPrivate
	case 2:
		return db_models.PrivacyLink
	case 3:
		return db_models.PrivacyPublic
	default:
		return db_models.PrivacyNone
	}
}

func (s *videoRegularUserService) GetInfo(ctx context.Context, request *pb.VideoRequest) (*pb.VideoResponse, error) {
	conn, err := db.GetDBConnection()
	if err != nil {
		return nil, errors.Errorf("DB error occured: %v", err)
	}

	model := db_models.Video{
		ID: request.GetID(),
	}

	result := conn.Limit(1).Find(&model)

	if result.RowsAffected != 1 {
		return nil, errors.Errorf("DB error occured: %v", result.Error)
	}

	return &pb.VideoResponse{
		ID:          model.ID,
		UserID:      model.UserID,
		Title:       model.Title,
		Description: model.Description,
		Privacy:     pb.Privacy(dbPrivacyToEnum(model.Privacy)),
		Link:        model.Link,
		Likes:       model.Likes,
		Dislikes:    model.Dislikes,
		CreatedAt:   uint64(model.CreatedAt.Unix()),
		EditedAt:    uint64(model.EditedAt.Unix()),
	}, nil
}

func (s *videoRegularUserService) GetVideoStream(ctx context.Context, request *pb.VideoStreamRequest) (*pb.StatusResponse, error) {
	return &pb.StatusResponse{
		Success: false,
		Message: "NOT IMPLEMENTED",
	}, nil
}

func (s *videoRegularUserService) GetReaction(ctx context.Context, request *pb.ReactionRequest) (*pb.ReactionResponse, error) {
	return getReaction(
		ReactionSearchParams{
			VideoID:   request.GetID(),
			CommentID: 0,
			UserID:    request.GetUserID(),
		},
	)
}

func (s *videoRegularUserService) UpdateReaction(ctx context.Context, request *pb.UpdateReactionRequest) (*pb.StatusResponse, error) {
	return updateReaction(
		ReactionSearchParams{
			VideoID:   request.GetID(),
			CommentID: 0,
			UserID:    request.GetUserID(),
		},
		request,
	)
}

func RegisterVideoRegularUserService(s *grpc.Server) {
	pb.RegisterVideoRegularUserServer(s, &videoRegularUserService{})
}
