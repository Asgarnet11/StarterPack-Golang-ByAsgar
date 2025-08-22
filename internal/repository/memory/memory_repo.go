package memory

import (
	"context"
)


type MemoryRepo struct{ items []int }


func New() *MemoryRepo { return &MemoryRepo{items: []int{1,2,3}} }


func (m *MemoryRepo) Count(ctx context.Context) (int, error) { return len(m.items), nil }