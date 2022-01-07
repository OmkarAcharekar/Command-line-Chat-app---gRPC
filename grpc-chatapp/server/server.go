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

