package postgres

import (
	"context"
	"github.com/alexrondon89/prosigliere-rest-api/internal/model"
	"github.com/alexrondon89/prosigliere-rest-api/internal/util"
)

func (pg *PostgresRepo) CreateComment(ctx context.Context, input *model.Comment) (*model.Comment, error) {
	ctx = util.SafeCtx(ctx)
	query := `
		INSERT INTO blog.comment (blog_post_id, username, content, created_at)
		VALUES($1, $2, $3, NOW())
		RETURNING *;
	`

	if err := pg.pool.QueryRow(ctx, query, input.PostId, input.Username, input.Content).Scan(&input.Id, &input.PostId, &input.Username, &input.Content, &input.CreatedAt); err != nil {
		return nil, err
	}
	return input, nil
}
