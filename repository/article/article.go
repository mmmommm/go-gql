package article

import (
	"context"

	"github.com/mmmommm/go-gql/database/mysql"
	domain "github.com/mmmommm/go-gql/domain/article"
)

type ArticleRepository interface {
	FindAll(ctx context.Context) ([]*domain.Article, error)
	Create(ctx context.Context, in *domain.Article) error
}

type articleRepository struct {
	c *mysql.MysqlClient
	// sql.dbとか、userの取得に必要な情報
}

func (r *articleRepository) FindAll(ctx context.Context) ([]*domain.Article, error) {
	rows, err := r.c.QueryContext(ctx,
		`SELECT
			id,
			title,
			content,
			FROM article
		`)
	if err != nil {
		return nil, err
	}
	var articles []*domain.Article
	for rows.Next() {
		article := &domain.Article{}
		if err := rows.Scan(
			&article.ID,
			&article.Title,
			&article.Content,
		); err != nil {
			return nil, err
		}
		articles = append(articles, article)
	}
	return articles, nil
}

func (r *articleRepository) Create(ctx context.Context, in *domain.Article) error {
	_, err := r.c.ExecContext(ctx,
		`INSERT INTO article(
			id,
			title,
			content,
		) VALUES (?,?,?);`,
		&in.ID,
		&in.Title,
		&in.Content,
	)
	return err
}

func ProvideArticleRepository(c *mysql.MysqlClient) ArticleRepository {
	return &articleRepository{c: c}
}
