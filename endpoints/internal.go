package endpoints

import (
	"internal-api/db"
	db_models "internal-api/db/models"

	pb "github.com/New-Tube/internal-api-protos"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type reactionSearchParams struct {
	VideoID   uint64
	CommentID uint64
	UserID    uint64
}

func getReaction(source_db_model interface{}, params reactionSearchParams) (*pb.ReactionResponse, error) {
	conn, err := db.GetDBConnection()
	if err != nil {
		return nil, errors.Errorf("DB error occured: %v", err)
	}

	result := conn.Limit(1).Find(&source_db_model)

	if result.RowsAffected != 1 {
		return nil, errors.Errorf("Comment with provided id not found")
	}

	reaction_db_model := db_models.Reaction{
		VideoID:   params.VideoID,
		CommentID: params.CommentID,
		UserID:    params.UserID,
	}

	result = conn.Limit(1).Find(&reaction_db_model)

	if result.RowsAffected != 1 {
		return &pb.ReactionResponse{
			IsLike:    false,
			IsDislike: false,
		}, nil
	}

	return &pb.ReactionResponse{
		IsLike:    reaction_db_model.IsLike,
		IsDislike: reaction_db_model.IsDislike,
	}, nil
}

func updateReaction(source_db_model interface{}, params reactionSearchParams, request *pb.UpdateReactionRequest) (*pb.StatusResponse, error) {
	conn, err := db.GetDBConnection()
	if err != nil {
		return nil, errors.Errorf("DB error occured: %v", err)
	}

	result := conn.Limit(1).Find(&source_db_model)

	if result.RowsAffected != 1 {
		return nil, errors.Errorf("Comment with provided id not found")
	}

	reaction_db_model := db_models.Reaction{
		VideoID:   params.VideoID,
		CommentID: params.CommentID,
		UserID:    params.UserID,
	}

	result = conn.Limit(1).Find(&reaction_db_model)

	if request.GetIsLike() && request.GetIsDislike() {
		return nil, errors.Errorf("You cannot set isLike=true and isDislike=true")
	}

	type common struct {
		ID uint64
	}

	err = conn.Transaction(func(tx *gorm.DB) error {
		// Using raw SQL query in case multiple users change commentModel at the same time,
		//  so we want to apply these changes for live data in DB
		if reaction_db_model.IsLike && !request.GetIsLike() {
			result := tx.Exec("UPDATE comments SET likes = likes - 1 WHERE id = ?", source_db_model.(common).ID)
			if result.Error != nil {
				return result.Error
			}
		}
		if !reaction_db_model.IsLike && request.GetIsLike() {
			result := tx.Exec("UPDATE comments SET likes = likes + 1 WHERE id = ?", source_db_model.(common).ID)
			if result.Error != nil {
				return result.Error
			}
		}
		if reaction_db_model.IsDislike && !request.GetIsDislike() {
			result := tx.Exec("UPDATE comments SET dislikes = dislikes - 1 WHERE id = ?", source_db_model.(common).ID)
			if result.Error != nil {
				return result.Error
			}
		}
		if !reaction_db_model.IsDislike && request.GetIsDislike() {
			result := tx.Exec("UPDATE comments SET dislikes = dislikes + 1 WHERE id = ?", source_db_model.(common).ID)
			if result.Error != nil {
				return result.Error
			}
		}

		reaction_db_model.IsLike = request.GetIsLike()
		reaction_db_model.IsDislike = request.GetIsDislike()

		if result.RowsAffected != 1 {
			// No reacord in the database, so we create one
			createResult := tx.Create(&reaction_db_model)

			if createResult.Error != nil {
				return errors.Errorf("DB error occured: %v", createResult.Error)
			}
		} else {
			// Updating existing record
			updateResult := tx.Save(&reaction_db_model)

			if updateResult.Error != nil {
				return errors.Errorf("DB error occured: %v", updateResult.Error)
			}
		}

		return nil
	})

	if err != nil {
		return nil, errors.Errorf("DB error occured: %v", err)
	}

	return &pb.StatusResponse{
		Success: true,
		Message: "",
	}, nil
}
