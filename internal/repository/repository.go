package repository

import "context"


type ExampleRepo interface {
Count(ctx context.Context) (int, error)
}