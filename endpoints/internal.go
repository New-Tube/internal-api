package endpoints

import (
	"fmt"
	"internal-api/db"
	db_models "internal-api/db/models"

	pb "github.com/New-Tube/internal-api-protos"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

func applyFilters[T any](filters map[string]any, conn *gorm.DB, model *T) (error, int64) {
	sql := ""
	var vals []any

	if len(filters) > 0 {
		for k, v := range filters {
			sql += k + "=? AND "
			vals = append(vals, v)
		}

		sql = sql[:len(sql)-5]
	}

	result := conn.Where(sql, vals...).Limit(1).Find(model)

	return result.Error, result.RowsAffected
}

type ReactionSearchParams struct {
	VideoID   uint64
	CommentID uint64
	UserID    uint64
}

func getReaction(params ReactionSearchParams) (*pb.ReactionResponse, error) {
	conn, err := db.GetDBConnection()
	if err != nil {
		return nil, errors.Errorf("DB error occured: %v", err)
	}

	if params.VideoID == 0 {
		result := conn.Limit(1).Find(&db_models.Comment{
			ID: params.CommentID,
		})
		if result.RowsAffected != 1 {
			return nil, errors.Errorf("Comment with provided id not found")
		}
	} else {
		result := conn.Limit(1).Find(&db_models.Video{
			ID: params.VideoID,
		})
		if result.RowsAffected != 1 {
			return nil, errors.Errorf("Video with provided id not found")
		}
	}

	reaction_db_model := db_models.Reaction{}

	_, rows_affected := applyFilters[db_models.Reaction](map[string]any{
		"video_id":   params.VideoID,
		"comment_id": params.CommentID,
		"user_id":    params.UserID,
	}, conn, &reaction_db_model)

	if rows_affected != 1 {
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

func updateReaction(params ReactionSearchParams, request *pb.UpdateReactionRequest) (*pb.StatusResponse, error) {
	conn, err := db.GetDBConnection()
	if err != nil {
		return nil, errors.Errorf("DB error occured: %v", err)
	}

	var table string
	var SourceID uint64
	if params.VideoID == 0 {
		SourceID = params.CommentID
		table = "comments"
		result := conn.Limit(1).Find(&db_models.Comment{
			ID: SourceID,
		})
		if result.RowsAffected != 1 {
			return nil, errors.Errorf("Comment with provided id not found")
		}
	} else {
		SourceID = params.VideoID
		table = "videos"
		result := conn.Limit(1).Find(&db_models.Video{
			ID: SourceID,
		})
		if result.RowsAffected != 1 {
			return nil, errors.Errorf("Video with provided id not found")
		}
	}

	reaction_db_model := db_models.Reaction{}

	err, rows_affected := applyFilters[db_models.Reaction](map[string]any{
		"video_id":   params.VideoID,
		"comment_id": params.CommentID,
		"user_id":    params.UserID,
	}, conn, &reaction_db_model)

	if request.GetIsLike() && request.GetIsDislike() {
		return nil, errors.Errorf("You cannot set isLike=true and isDislike=true")
	}

	err = conn.Transaction(func(tx *gorm.DB) error {
		// Using raw SQL query in case multiple users change commentModel at the same time,
		//  so we want to apply these changes for live data in DB
		if reaction_db_model.IsLike && !request.GetIsLike() {
			result := tx.Exec("UPDATE "+table+" SET likes = likes - 1 WHERE id = ?", SourceID)
			if result.Error != nil {
				return result.Error
			}
		}
		if !reaction_db_model.IsLike && request.GetIsLike() {
			result := tx.Exec("UPDATE "+table+" SET likes = likes + 1 WHERE id = ?", SourceID)
			if result.Error != nil {
				return result.Error
			}
		}
		if reaction_db_model.IsDislike && !request.GetIsDislike() {
			result := tx.Exec("UPDATE "+table+" SET dislikes = dislikes - 1 WHERE id = ?", SourceID)
			if result.Error != nil {
				return result.Error
			}
		}
		if !reaction_db_model.IsDislike && request.GetIsDislike() {
			result := tx.Exec("UPDATE "+table+" SET dislikes = dislikes + 1 WHERE id = ?", SourceID)
			if result.Error != nil {
				return result.Error
			}
		}

		reaction_db_model.IsLike = request.GetIsLike()
		reaction_db_model.IsDislike = request.GetIsDislike()

		if rows_affected != 1 {
			// No reacord in the database, so we create one
			createResult := tx.Create(&reaction_db_model)

			if createResult.Error != nil {
				fmt.Println("ERRRIRR!!!!")
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
