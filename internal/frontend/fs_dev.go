//go:build !prod
// +build !prod

package frontend

import (
	"fmt"
	"io/fs"
	"os"
	"runtime"
)

func GetAssetFS() (fs.FS, error) {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		return nil, fmt.Errorf("can not get caller information")
	}
	basePath := filename[:len(filename)-9]
	path := fmt.Sprintf("%s/%s", basePath, "nerdlocke-frontend/dist")
	return os.DirFS(path), nil
}
