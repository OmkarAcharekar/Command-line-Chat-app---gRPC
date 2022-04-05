package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net"
	"sync"

	chat "github.com/omkaracharekar/grpc-chat-app/grpc-chatapp/schema"
	"google.golang.org/grpc"
)
// Move this to a separate file
type Queue struct {
	container     []*chat.Message
	containerLock sync.RWMutex
}


func (s *server) generateToken() (string, error) {

	level.Debug(s.logger).Log("message", "started generating token")
	txt := make([]byte, tokenSize)
	_, err := rand.Read(txt)
	if err != nil {
		level.Error(s.logger).Log("error", "error while generating the token")
		return "", err
	}
	level.Debug(s.logger).Log("message", "finished generating token")
	return fmt.Sprintf("%x", txt), nil
}



func (s *server) addClientName(username string, tkn string) {

	s.nameMutex.RLock()
	defer s.nameMutex.RUnlock()
	level.Debug(s.logger).Log("message", "adding the client name", "client", username, "token", tkn)
	s.ClientName[tkn] = username

}

func (s *server) getClientName(tkn string) (string, bool) {

	s.nameMutex.RLock()
	defer s.nameMutex.RUnlock()
	level.Debug(s.logger).Log("message", "getting the client name", "token", tkn)
	name, ok := s.ClientName[tkn]
	return name, ok
}



func (s *server) Login(ctx context.Context, req *chat.LoginRequest) (*chat.LoginResponse, error) {

	// TODO: handle same name people in the chat
	// Generate a token
	level.Info(s.logger).Log("message", "new client login request", "req", req)
	tkn, err := s.generateToken()
	if err != nil {
		level.Error(s.logger).Log("error", "login failed for the request", req)
	}
	// Add the token in the client name
	s.addClientName(req.Username, tkn)
	// Send in a notif that broadcast is successful
	level.Info(s.logger).Log("message", "login is successful", "req", req)
	s.CommonChannel <- chat.StreamResponse{
		Timestamp: ptypes.TimestampNow(),
		Event: &chat.StreamResponse_ClientLogin{
			&chat.StreamResponse_Login{
				Name: req.Username,
			},
		},
	}