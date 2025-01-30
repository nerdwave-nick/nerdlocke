//go:build prod
// +build prod

package frontend

import (
	"embed"
	"io/fs"
)

//go:embed nerdlocke-frontend/dist
var embedFrontend embed.FS

func GetAssetFS() (fs.FS, error) {
	f, err := fs.Sub(embedFrontend, "nerdlocke-frontend/dist")
	if err != nil {
		return nil, error
	}

	return f, nil
}
