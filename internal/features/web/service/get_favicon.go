package web_service

import (
	"fmt"
	"os"
	"path"
)

func (s *WebService) GetFavicon() ([]byte, error) {
	faviconFilePath := path.Join(
		os.Getenv("PROJECT_ROOT"),
		"/public/favicon.svg",
	)

	favicon, err := s.webRepository.GetFile(faviconFilePath)
	if err != nil {
		return nil, fmt.Errorf("get file from repository: %w", err)
	}

	return favicon, nil
}
