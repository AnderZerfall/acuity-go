// Package analyzer contains gRPC/Connect service for analyzing posts
package transport

import (
	"Acuity/gen/api/analyzer/v1"
	"Acuity/gen/api/analyzer/v1/analyzerconnect"
	"Acuity/internal/domain"
	"context"

	"connectrpc.com/connect"
)

type AnalyzeServiceServer struct {
	analyzerconnect.AnalyzerServiceHandler
	Analyzer *domain.PostAnalyzer
}

func (service *AnalyzeServiceServer) AnalyzePost(ctx context.Context,
	req *connect.Request[analyzer.AnalyzePostRequest]) (*connect.Response[analyzer.AnalyzePostResponse], error) {
	post := req.Msg.GetPost()
	_ = post

	result, err := service.Analyzer.Analyze(ctx, post.Content)

	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	responseBody := &analyzer.AnalyzePostResponse{
		TrustCoefficient: float32(result.TrustCoefficient),
		Details:          result.Details,
	}

	return connect.NewResponse(responseBody), nil
}
