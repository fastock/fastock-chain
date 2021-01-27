package cases

import (
	"os"
)

// GetBackendDBDir return the path "$GOPATH/src/../../x/backend/cases"
func GetBackendDBDir() string {
	gopath := os.Getenv("GOPATH")
	dir := gopath + "/src/../../x/backend/cases"
	return dir
}
