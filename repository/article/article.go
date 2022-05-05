package article

import (
	"context"

	"github.com/mmmommm/go-gql/database/mysql"
	domain "github.com/mmmommm/go-gql/domain/article"
)

type ArticleRepository interface {
	FindAll(context.Context) ([]*domain.Article, error)
}

type articleRepository struct {
	c *mysql.MysqlClient
	// sql.dbとか、userの取得に必要な情報
}

func (r *articleRepository) FindAll(context.Context) ([]*domain.Article, error) {
	// 実際に接続する処理
	return nil, nil
}

func ProvideArticleRepository(c *mysql.MysqlClient) ArticleRepository {
	return &articleRepository{c: c}
}
