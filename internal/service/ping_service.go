package service

import "context"


type Pinger interface { Ping(ctx context.Context) string }


type pingService struct{}


func NewPingService() Pinger { return &pingService{} }


func (s *pingService) Ping(ctx context.Context) string { return "pong" }