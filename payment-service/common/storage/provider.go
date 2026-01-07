package storage

import (
	"context"
)

type Provider interface {
	UploadFile(context.Context, string, []byte) (string, error)
}
