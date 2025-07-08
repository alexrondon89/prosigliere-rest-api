package postgres

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/alexrondon89/prosigliere-rest-api/internal/model"
	"github.com/alexrondon89/prosigliere-rest-api/internal/util"
	"github.com/jackc/pgx/v5"
)

func (pg *PostgresRepo) CreatePostInDb(ctx context.Context, input *model.Post) (*model.Post, error) {
	ctx = util.SafeCtx(ctx)
	query := `
		INSERT INTO blog.blog_post (title, content, created_at)
		VALUES($1, $2, NOW())
		RETURNING *;
	`

	if err := pg.pool.QueryRow(ctx, query, input.Title, input.Content).Scan(&input.Id, &input.Title, &input.Content, &input.CreatedAt); err != nil {
		return nil, err
	}
	return input, nil
}

func (pg *PostgresRepo) GetPostFromDb(ctx context.Context, id string) (*model.Post, error) {
	ctx = util.SafeCtx(ctx)

	query := `
		WITH post AS (
			SELECT id, title, content, created_at
			FROM   blog.blog_post
			WHERE  id = $1
			LIMIT  1
		),
		comments_page AS (
			SELECT *
			FROM   blog.comment
			WHERE  blog_post_id = $1
			ORDER  BY created_at DESC
		)
		SELECT  p.id,
				p.title,
				p.content,
				p.created_at,
				COALESCE(
				  json_agg(
					  json_build_object(
						'id',        c.id,
						'username',  c.username,
						'content',   c.content,
						'createdAt', c.created_at
					  )
					  ORDER BY c.created_at DESC
				  ) FILTER (WHERE c.id IS NOT NULL),
				  '[]'
				) AS comments
		FROM    post p
		LEFT JOIN comments_page c
			   ON c.blog_post_id = p.id
		GROUP BY p.id, p.title, p.content, p.created_at
		ORDER BY p.created_at DESC;
	`
	var p model.Post
	var rawJSON []byte
	err := pg.pool.QueryRow(ctx, query, id).
		Scan(&p.Id, &p.Title, &p.Content, &p.CreatedAt, &rawJSON)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return &p, nil
		}
		return nil, err
	}

	if err := json.Unmarshal(rawJSON, &p.Comments); err != nil {
		return nil, fmt.Errorf("unmarshal comments: %w", err)
	}
	return &p, nil
}

func (pg *PostgresRepo) GetAllPostsFromDb(ctx context.Context) ([]model.Post, error) {
	ctx = util.SafeCtx(ctx)
	query := `
		WITH posts AS (
			SELECT id, title, content, created_at
			FROM   blog.blog_post
			ORDER  BY created_at DESC
		), comments_src AS (
			SELECT c.*,
				   ROW_NUMBER() OVER (
					   PARTITION BY c.blog_post_id
					   ORDER BY c.created_at DESC
				   ) AS rn
			FROM   blog.comment c
			WHERE  c.blog_post_id IN (SELECT id FROM posts)
		)
		SELECT  p.id,
				p.title,
				p.content,
				p.created_at,
				COALESCE(
				  json_agg(
					json_build_object(
					  'id',        c.id,
					  'username',  c.username,
					  'content',   c.content,
					  'createdAt', c.created_at
					)
					ORDER BY c.created_at
				  ) FILTER (WHERE c.id IS NOT NULL),
				  '[]'
				) AS comments
		FROM    posts p
		LEFT JOIN comments_src c
			   ON c.blog_post_id = p.id
			  AND c.rn <= 30
		GROUP BY p.id, p.title, p.content, p.created_at
		ORDER BY p.created_at DESC; `

	posts := []model.Post{}
	rows, err := pg.pool.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var p model.Post
		var rawJSON []byte
		if err := rows.Scan(&p.Id, &p.Title, &p.Content, &p.CreatedAt, &rawJSON); err != nil {
			return nil, err
		}
		err = json.Unmarshal(rawJSON, &p.Comments)
		if err != nil {
			return nil, fmt.Errorf("unmarshal comments: %w", err)
		}
		posts = append(posts, p)
	}
	return posts, nil
}
