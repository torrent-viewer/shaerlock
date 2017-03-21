package shaerlock

import (
	"path/filepath"
	"strings"
)

// SleuthMatroska is the function parsing files conforming to the Matroska
// Container Specification.
func SleuthMatroska(path string) (Media, error) {
	fullPath, err := filepath.Abs(path)
	if err != nil {
		return Media{}, err
	}
	return Media{
		Name: strings.TrimSuffix(filepath.Base(path), filepath.Ext(path)),
		Extension: strings.TrimPrefix(filepath.Ext(path), "."),
		FullPath: fullPath,
	}, nil
}
