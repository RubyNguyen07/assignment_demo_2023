package main

import (
	"context"
	"errors"
	"testing"

	"github.com/TikTokTechImmersion/assignment_demo_2023/rpc-server/kitex_gen/rpc"
	"github.com/stretchr/testify/assert"
)

func TestIMServiceImpl_Send(t *testing.T) {
	type args struct {
		ctx context.Context
		req *rpc.SendRequest
	}
	tests := []struct {
		name    string
		args    args
		wantErr error
	}{
		{
			name: "success",
			args: args{
				ctx: context.Background(),
				req: &rpc.SendRequest{
					Message: &rpc.Message{
						Chat:   chatName,
						Text:   "test",
						Sender: "test1",
					},
				},
			},
			wantErr: nil,
		},
		{
			name: "success",
			args: args{
				ctx: context.Background(),
				req: &rpc.SendRequest{
					Message: &rpc.Message{
						Chat:   "test1:test2",
						Text:   "very random message?@",
						Sender: "test1",
					},
				},
			},
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &IMServiceImpl{}
			got, err := s.Send(tt.args.ctx, tt.args.req)
			assert.True(t, errors.Is(err, tt.wantErr))
			assert.NotNil(t, got)
		})
	}

}
