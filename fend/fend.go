package fend

import (
	"context"
)

type Fend func(ctx context.Context, mode Mode) error
