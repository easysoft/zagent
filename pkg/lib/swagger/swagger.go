package swaggerUtils

import (
	"fmt"
	"github.com/iris-contrib/swagger"
	"github.com/iris-contrib/swagger/swaggerFiles"
	"github.com/kataras/iris/v12"
)

func InitSwaggerDocs(port int, app *iris.Application) {
	config := swagger.Config{
		URL:          fmt.Sprintf("http://localhost:%d/swagger/doc.json", port),
		DeepLinking:  true,
		DocExpansion: "list",
		DomID:        "#swagger-ui",
		Prefix:       "/swagger",
	}

	swaggerUI := swagger.Handler(swaggerFiles.Handler, config)
	app.Get("/swagger", swaggerUI)
	app.Get("/swagger/{any:path}", swaggerUI)
}
