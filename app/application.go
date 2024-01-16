package app

import (
	"context"

	"github.com/michalsciubidlo/vodeno-task/app/customermailing"
)

type service interface {
	Add(ctx context.Context, msg customermailing.Message) error
}
