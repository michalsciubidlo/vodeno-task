package app

import (
	"context"
	"vodeno-task/app/customermailing"
)

type service interface {
	Add(ctx context.Context, msg customermailing.Message) error
}
