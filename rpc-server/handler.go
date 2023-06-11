package main

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/TikTokTechImmersion/assignment_demo_2023/rpc-server/kitex_gen/rpc"
)

// IMServiceImpl implements the last service interface defined in the IDL.
type IMServiceImpl struct{}

func (s *IMServiceImpl) Send(ctx context.Context, req *rpc.SendRequest) (*rpc.SendResponse, error) {
	if err := validateSendRequest(req); err != nil {
		return nil, err
	}

	timestamp := time.Now().Unix()
	message := &Message{
		Message:   req.Message.GetText(),
		Sender:    req.Message.GetSender(),
		Timestamp: timestamp,
	}

	roomID, err := getRoomID(req.Message.GetChat())
	if err != nil {
		return nil, err
	}

	err = rdb.SaveMessage(ctx, roomID, message)
	if err != nil {
		return nil, err
	}

	resp := rpc.NewSendResponse()
	resp.Code, resp.Msg = 0, "success"
	return resp, nil
}

func (s *IMServiceImpl) Pull(ctx context.Context, req *rpc.PullRequest) (*rpc.PullResponse, error) {
	roomID, err := getRoomID(req.GetChat())

	if err != nil {
		return nil, err
	}

	limit := int64(req.GetLimit())
	if limit == 0 {
		limit = 10 // Default limit 
	}

	start := req.GetCursor()
	end := start + limit
	order := req.GetReverse()

	messages, err := rdb.GetMessagesByRoomID(ctx, roomID, start, end, order)
	if err != nil {
		return nil, err
	}

	respMessages := make([]*rpc.Message, 0)
	var cnt int64 = 0
	var nextCursor int64 = 0
	hasMore := false

	for _, message := range messages {
		if cnt + 1 > limit {
			hasMore = true
			nextCursor = end
			break 
		}
		temp := &rpc.Message{
			Chat:     req.GetChat(),
			Text:     message.Message,
			Sender:   message.Sender,
			SendTime: message.Timestamp,
		}
		respMessages = append(respMessages, temp)
		cnt += 1
	}

	resp := rpc.NewPullResponse()
	resp.Messages = respMessages
	resp.Code = 0
	resp.Msg = "success"
	resp.HasMore = &hasMore
	resp.NextCursor = &nextCursor

	return resp, nil
}

func validateSendRequest(req *rpc.SendRequest) error {
	members := strings.Split(req.Message.Chat, ":")
	if len(members) != 2 {
		err := fmt.Errorf("invalid Chat ID '%s', should be in the format of user1:user2", req.Message.GetChat())
		return err
	}
	mem1, mem2 := members[0], members[1]

	if req.Message.GetSender() != mem1 && req.Message.GetSender() != mem2 {
		err := fmt.Errorf("sender '%s' not in the chat room", req.Message.GetSender())
		return err
	}

	return nil
}

func getRoomID(chat string) (string, error) {
	var roomID string

	formattedChat := strings.ToLower(chat)
	members := strings.Split(formattedChat, ":")

	if len(members) != 2 {
		err := fmt.Errorf("invalid Chat ID '%s', should be in the format of user1:user2", chat)
		return "", err
	}

	mem1, mem2 := members[0], members[1]
	// Compare the sender and receiver alphabetically, and sort it asc to form the room ID
	if comp := strings.Compare(mem1, mem2); comp == 1 {
		roomID = fmt.Sprintf("%s:%s", mem2, mem1)
	} else {
		roomID = fmt.Sprintf("%s:%s", mem1, mem2)
	}

	return roomID, nil
}
