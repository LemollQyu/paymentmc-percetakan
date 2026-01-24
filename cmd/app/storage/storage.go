package storage

type Storage struct {
	BaseDir string
	Url     string
}

func NewStorage(baseDir, url string) *Storage {
	return &Storage{
		BaseDir: baseDir,
		Url:     url,
	}
}
