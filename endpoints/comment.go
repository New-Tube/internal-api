package endpoints

import (
	"context"
	"internal-api/db"
	db_models "internal-api/db/models"

	pb "github.com/New-Tube/internal-api-protos"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
)

const COMMENTS_GET_MAX_LIMIT uint32 = 100

type commentService struct {
	pb.UnimplementedCommentServer
}

func (s *commentService) Get(ctx context.Context, request *pb.CommentRequest) (*pb.CommentResponse, error) {
	conn, err := db.GetDBConnection()
	if err != nil {
		return nil, errors.Errorf("DB error occured: %v", err)
	}

	model := db_models.Comment{
		ID: request.GetID(),
	}
	result := conn.Limit(1).Find(&model)

	if result.RowsAffected != 1 {
		return nil, errors.Errorf("Comment with provided id not found")
	}

	return &pb.CommentResponse{
		ID:        model.ID,
		Text:      model.Text,
		VideoID:   model.VideoID,
		CommentID: model.CommentID,
		UserID:    model.UserID,
		Likes:     model.Likes,
		Dislikes:  model.Dislikes,
		CreatedAt: uint64(model.CreatedAt.Unix()),
		EditedAt:  uint64(model.EditedAt.Unix()),
	}, nil
}

func (s *commentService) GetMany(ctx context.Context, request *pb.GetManyRequest) (*pb.GetManyResponse, error) {
	// Check for cases when both ids specified or none of them
	if request.GetVideoID() == 0 && request.GetCommentID() == 0 ||
		request.GetVideoID() != 0 && request.GetCommentID() != 0 {
		return nil, errors.Errorf("You should specify one of possible sources. VideoID or CommentID")
	}

	conn, err := db.GetDBConnection()
	if err != nil {
		return nil, errors.Errorf("DB error occured: %v", err)
	}

	var commentModels []db_models.Comment

	result := conn.Limit(int(min(request.GetLimit(), COMMENTS_GET_MAX_LIMIT))).
		Offset(int(request.GetOffset())).
		Where("video_id = ? AND comment_id = ?", request.GetVideoID(), request.GetCommentID()).
		Find(&commentModels)

	if result.Error != nil {
		return nil, errors.Errorf("DB error occured: %v", result.Error)
	}

	var comments []*pb.CommentResponse

	for _, model := range commentModels {
		comments = append(comments, &pb.CommentResponse{
			ID:        model.ID,
			Text:      model.Text,
			VideoID:   model.VideoID,
			CommentID: model.CommentID,
			UserID:    model.UserID,
			Likes:     model.Likes,
			Dislikes:  model.Dislikes,
			CreatedAt: uint64(model.CreatedAt.Unix()),
			EditedAt:  uint64(model.EditedAt.Unix()),
		})
	}

	return &pb.GetManyResponse{
		Count:    uint32(result.RowsAffected),
		Comments: comments,
	}, nil
}

func (s *commentService) Create(ctx context.Context, request *pb.CommentCreateRequest) (*pb.StatusResponse, error) {
	// Check for cases when both ids specified or none of them
	if request.GetVideoID() == 0 && request.GetCommentID() == 0 ||
		request.GetVideoID() != 0 && request.GetCommentID() != 0 {
		return nil, errors.Errorf("You should specify one of possible sources. VideoID or CommentID")
	}

	conn, err := db.GetDBConnection()
	if err != nil {
		return nil, errors.Errorf("DB error occured: %v", err)
	}

	model := db_models.Comment{
		Text:      request.GetText(),
		VideoID:   request.GetVideoID(),
		CommentID: request.GetCommentID(),
		UserID:    request.GetUserID(),
	}
	result := conn.Create(model)

	if result.Error != nil {
		return nil, errors.Errorf("DB error occured: %v", result.Error)
	}

	return &pb.StatusResponse{
		Success: true,
		Message: "",
	}, nil
}

func (s *commentService) GetReaction(ctx context.Context, request *pb.CommentRequest) (*pb.CommentReactionResponse, error) {
	conn, err := db.GetDBConnection()
	if err != nil {
		return nil, errors.Errorf("DB error occured: %v", err)
	}

	commentModel := db_models.Comment{
		ID: request.GetID(),
	}
	result := conn.Limit(1).Find(&commentModel)

	if result.RowsAffected != 1 {
		return nil, errors.Errorf("Comment with provided id not found")
	}

	reactionModel := db_models.Reaction{
		VideoID:   commentModel.VideoID,
		CommentID: commentModel.CommentID,
		UserID:    commentModel.UserID,
	}

	result = conn.Limit(1).Find(&reactionModel)

	if result.RowsAffected != 1 {
		return nil, errors.Errorf("Given user hadn't reacted to this video")
	}

	return &pb.CommentReactionResponse{
		IsLike:    reactionModel.IsLike,
		IsDislike: reactionModel.IsDislike,
	}, nil
}

func (s *commentService) UpdateReaction(ctx context.Context, request *pb.CommentUpdateReactionRequest) (*pb.StatusResponse, error) {
	conn, err := db.GetDBConnection()
	if err != nil {
		return nil, errors.Errorf("DB error occured: %v", err)
	}

	commentModel := db_models.Comment{
		ID: request.GetID(),
	}
	result := conn.Limit(1).Find(&commentModel)

	if result.RowsAffected != 1 {
		return nil, errors.Errorf("Comment with provided id not found")
	}

	reactionModel := db_models.Reaction{
		VideoID:   commentModel.VideoID,
		CommentID: commentModel.CommentID,
		UserID:    commentModel.UserID,
	}

	result = conn.Limit(1).Find(&reactionModel)

	if result.RowsAffected != 1 {
		return nil, errors.Errorf("Given user hadn't reacted to this video")
	}

	if request.GetIsLike() && request.GetIsDislike() {
		return nil, errors.Errorf("You cannot set isLike=true and isDislike=true")
	}

	reactionModel.IsLike = request.GetIsLike()
	reactionModel.IsDislike = request.GetIsDislike()

	updateResult := conn.Save(&reactionModel)

	if updateResult.Error != nil {
		return nil, errors.Errorf("DB error occured: %v", updateResult.Error)
	}

	return &pb.StatusResponse{
		Success: true,
		Message: "",
	}, nil
}

func (s *commentService) Delete(ctx context.Context, request *pb.CommentRequest) (*pb.StatusResponse, error) {
	conn, err := db.GetDBConnection()
	if err != nil {
		return nil, errors.Errorf("DB error occured: %v", err)
	}

	model := db_models.Comment{
		ID: request.GetID(),
	}
	result := conn.Limit(1).Find(&model)

	if result.RowsAffected != 1 {
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

func RegisterCommentService(s *grpc.Server) {
	pb.RegisterCommentServer(s, &commentService{})
}
