package storage

import (
	"context"
	"mime/multipart"
)

func (s *Storage) UploadProofPayment(ctx context.Context, file *multipart.FileHeader) (string, error) {
	const maxIconSize = 3000 * 1024
	return s.UploadFile(ctx, file, "bukti", []string{".jpeg", ".png", ".webp", "jpg"}, maxIconSize)
}
