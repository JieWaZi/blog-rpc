package api

import (
	"golang.org/x/net/context"
	pb "github.com/blog/proto"
	"github.com/blog/api/models"
	"github.com/blog/utils"
)

func (*BlogService) ListArticles(ctx context.Context, req *pb.ListArticlesRequest) (*pb.ListArticlesResponse, error) {

	articles, err := models.TArticle.ListArticle(req.PageIndex, req.PageSize)
	if err != nil {
		return nil, err
	}

	var protoArticles []*pb.Article
	for _, article := range articles {
		protoArticle := &pb.Article{
			Id:             article.Id,
			Title:          article.Title,
			CreatedAt:      article.CreatedAt.Format(utils.Time_Format),
			CoverImg:       article.CoverImg,
			ReadCount:      int32(article.ReadCount),
			FavouriteCount: int32(article.FavouriteCount),
			Abstract:       article.Abstract,
			CommentCount:   article.CommentCount,
			FilePath:       article.FilePath,
			Tag:            article.Tag,
		}
		protoArticles = append(protoArticles, protoArticle)
	}

	return &pb.ListArticlesResponse{Articles: protoArticles}, nil
}
