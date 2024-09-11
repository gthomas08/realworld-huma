package ping

import "context"

type Usecase interface {
	GetPingMessage(ctx context.Context) string
}
