package model

import "context"

type Todo struct{}

func (*Todo) Broken(ctx context.Context) string {
	return ""
}
