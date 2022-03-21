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
