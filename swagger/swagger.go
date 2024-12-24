package swagger

import (
	"embed"
	"net/http"
)

//go:embed swagger-ui/*
var swaggerFiles embed.FS

func ServeSwaggerUI() http.Handler {
	return http.FileServer(http.FS(swaggerFiles))
}
