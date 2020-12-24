package handlers

import (
	pb "article"
	"article/internal/logics"
	"article/internal/utils"
	"context"
)

// NewService returns a naïve, stateless implementation of Service.
func NewService() pb.ArticleServer {
	return articleService{}
}

type articleService struct{}

//获取文章详情
func (s articleService) Detail(ctx context.Context, in *pb.DetailRequest) (*pb.DetailResponse, error) {
	article_id_hash := in.ArticleId
	if len(article_id_hash) == 0 {
		return nil, utils.NewParamErr("article_id")
	}
	article_id, err := utils.DecodeID(article_id_hash)
	if err != nil {
		return nil, utils.NewErrMsg(err.Error())
	}
	if article_id == 0 {
		return nil, utils.NewParamErr("article_id")
	}

	var resp pb.DetailResponse
	detail, err := logics.GetArticleDetail(article_id)
	if err != nil {
		return nil, utils.NewErrMsg(err.Error())
	}

	resp.Code = 1
	resp.Msg = "success"
	resp.Data = detail
	return &resp, nil
}

func (s articleService) Records(ctx context.Context, in *pb.RecordsRequest) (*pb.RecordsResponse, error) {
	var resp pb.RecordsResponse
	return &resp, nil
}

func (s articleService) Remove(ctx context.Context, in *pb.RemoveRequest) (*pb.RemoveResponse, error) {
	var resp pb.RemoveResponse
	return &resp, nil
}

func (s articleService) Top(ctx context.Context, in *pb.TopRequest) (*pb.TopResponse, error) {
	var resp pb.TopResponse
	return &resp, nil
}

func (s articleService) Recommends(ctx context.Context, in *pb.RecommendsRequest) (*pb.RecommendsResponse, error) {
	var resp pb.RecommendsResponse
	return &resp, nil
}

func (s articleService) Review(ctx context.Context, in *pb.ReviewRequest) (*pb.ReviewResponse, error) {
	var resp pb.ReviewResponse
	return &resp, nil
}

func (s articleService) Publish(ctx context.Context, in *pb.PublishRequest) (*pb.PublishResponse, error) {
	var resp pb.PublishResponse
	return &resp, nil
}

func (s articleService) Edit(ctx context.Context, in *pb.EditRequest) (*pb.EditResponse, error) {
	var resp pb.EditResponse
	return &resp, nil
}
