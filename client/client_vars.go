package client

import "time"

var (
	t time.Time
)

const (
	noRetries     = 5
	retryInterval = 10 * time.Second
)
