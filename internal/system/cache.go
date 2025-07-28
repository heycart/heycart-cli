package system

import (
	"os"
	"path"
)

func GetHeyCartCliCacheDir() string {
	cacheDir, _ := os.UserCacheDir()

	return path.Join(cacheDir, "heycart-cli")
}
