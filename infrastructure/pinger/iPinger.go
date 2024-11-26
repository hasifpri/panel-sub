package infrastructurepinger

import "context"

type IPinger interface {
	Ping(ctx context.Context) error
}
