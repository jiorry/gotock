package tdx

import (
	"os"
	"path/filepath"
	"strings"
)

func ScanDirGetAllFileName(folder, ext string) ([]string, error) {
	dir, err := os.Open(folder)
	if err != nil {
		return nil, err
	}

	defer dir.Close()

	var arr []string
	var name string

	flist, err := dir.Readdir(0)
	result := make([]string, len(flist))

	for i, info := range flist {
		if info.IsDir() {
			continue
		}

		name = info.Name()

		if filepath.Ext(name) != ext {
			continue
		}

		arr = strings.Split(name, ".")
		result[i] = arr[0][2:]
	}

	return result, nil
}
