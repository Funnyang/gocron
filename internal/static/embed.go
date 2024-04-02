package static

import (
	"embed"
	"io/fs"
)

//go:embed web
var webFs embed.FS
var WebServerRoot fs.FS

func init() {
	WebServerRoot, _ = fs.Sub(webFs, "web")
}
