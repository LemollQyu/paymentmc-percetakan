package storage

import (
	"context"
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"paymentmc/utils"
	"strings"
	"time"
)

func IsValidExt(filename string, allowedExts []string) bool {
	ext := strings.ToLower(filepath.Ext(filename))
	for _, a := range allowedExts {
		if ext == a {
			return true
		}
	}
	return false
}

func (s *Storage) UploadFile(
	ctx context.Context,
	file *multipart.FileHeader,
	folder string,
	allowedExts []string,
	maxSize int64,
) (string, error) {

	if file == nil {
		return "", utils.FileRequired
	}

	if file.Size > maxSize {
		return "", utils.FileMaxSize
	}

	if !IsValidExt(file.Filename, allowedExts) {
		return "", utils.FileExtInvalid
	}

	dir := filepath.Join(s.BaseDir, folder)
	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		return "", err
	}

	filename := fmt.Sprintf("%d%s", time.Now().UnixNano(), filepath.Ext(file.Filename))
	fullPath := filepath.Join(dir, filename)

	src, err := file.Open()
	if err != nil {
		return "", err
	}
	defer src.Close()

	dst, err := os.Create(fullPath)
	if err != nil {
		return "", err
	}
	defer dst.Close()

	if _, err := io.Copy(dst, src); err != nil {
		return "", err
	}

	fmt.Println("UPLOAD DIR:", dir)
	fmt.Println("FULL PATH:", fullPath)

	publicURL := s.Url + "/static/" + folder + "/" + filename
	return publicURL, nil
}

func (s *Storage) UploadIconPaymentMethod(ctx context.Context, file *multipart.FileHeader) (string, error) {
	const maxIconSize = 5000 * 1024
	return s.UploadFile(ctx, file, "icon-method", []string{".svg", ".png", ".webp", ".heic"}, maxIconSize)
}

func (s *Storage) UploadCodeQris(ctx context.Context, file *multipart.FileHeader) (string, error) {
	const maxIconSize = 5000 * 1024
	return s.UploadFile(ctx, file, "code-qris", []string{".jpg", ".png", ".webp", ".jpeg", ".heic"}, maxIconSize)
}

func (s *Storage) DeleteFile(
	ctx context.Context,
	fileURL string,
) error {

	if fileURL == "" {
		return nil
	}

	// contoh fileURL:
	// http://localhost:8080/static/order/123456.png

	baseDir, err := os.Getwd()
	if err != nil {
		return err
	}

	// hapus base url
	prefix := fmt.Sprintf("%s/static/", s.Url)
	if !strings.HasPrefix(fileURL, prefix) {
		return nil // atau return error
	}

	relativePath := strings.TrimPrefix(fileURL, prefix)

	// full path ke file
	fullPath := filepath.Join(baseDir, "uploads", relativePath)

	// cek file exists
	if _, err := os.Stat(fullPath); os.IsNotExist(err) {
		return nil // file sudah tidak ada → aman
	}

	// hapus file
	if err := os.Remove(fullPath); err != nil {
		return utils.ErrDeleteFile
	}

	return nil
}
