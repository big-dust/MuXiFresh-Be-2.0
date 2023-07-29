package model

import (
	"context"
	"github.com/zeromicro/go-zero/core/stores/mon"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var _ CommentModel = (*customCommentModel)(nil)

type (
	// CommentModel is an interface to be customized, add more methods here,
	// and implement the added methods in customCommentModel.
	CommentModel interface {
		commentModel
		FindBySubmissionID(ctx context.Context, submissionID string) ([]*Comment, error)
	}

	customCommentModel struct {
		*defaultCommentModel
	}
)

// NewCommentModel returns a model for the mongo.
func NewCommentModel(url, db, collection string) CommentModel {
	conn := mon.MustNewModel(url, db, collection)
	return &customCommentModel{
		defaultCommentModel: newDefaultCommentModel(conn),
	}
}

func (m *customCommentModel) FindBySubmissionID(ctx context.Context, submissionID string) ([]*Comment, error) {
	var comments []*Comment
	sid, err := primitive.ObjectIDFromHex(submissionID)
	if err != nil {
		return nil, err
	}
	err = m.conn.Find(ctx, &comments, bson.M{"submission_id": sid})

	switch err {
	case nil:
		return comments, nil
	case mon.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}
