package storage

import (
	"context"
	"mime/multipart"
)

func (s *Storage) UploadProofPayment(ctx context.Context, file *multipart.FileHeader) (string, error) {
	const maxIconSize = 5000 * 1024
	return s.UploadFile(ctx, file, "bukti", []string{".jpeg", ".png", ".webp", ".jpg", ".heic"}, maxIconSize)
}
