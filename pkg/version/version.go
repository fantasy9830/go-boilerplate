package version

import (
	"fmt"
	"net/http"
	"runtime"

	"github.com/gin-gonic/gin"
)

var (
	Version   string
	BuildDate string
)

// PrintCLIVersion print server info
func PrintCLIVersion() string {
	return fmt.Sprintf(
		"version %s, built on %s, %s",
		Version,
		BuildDate,
		runtime.Version(),
	)
}

func APIVersion(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, map[string]string{
		"version":         Version,
		"build_date":      BuildDate,
		"runtime_version": runtime.Version(),
	})
}
