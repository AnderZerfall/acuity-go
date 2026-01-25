// Package analyzer contains gRPC/Connect service for analyzing posts
package analyzer

import (
	"Acuity/gen/api/analyzer/v1"
	"Acuity/gen/api/analyzer/v1/analyzerconnect"
	"context"

	"connectrpc.com/connect"
)

type AnalyzeServiceServer struct {
	analyzerconnect.AnalyzerServiceHandler
}

func (service *AnalyzeServiceServer) AnalyzePost(_ context.Context,
	req *connect.Request[analyzer.AnalyzePostRequest]) (*connect.Response[analyzer.AnalyzePostResponse], error) {
	post := req.Msg.GetPost()
	_ = post

	responseBody := &analyzer.AnalyzePostResponse{
		TrustCoefficient: 0,
		Details:          []string{"That's a test value"},
	}

	return connect.NewResponse(responseBody), nil
}
