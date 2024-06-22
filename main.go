package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "ping..."})
	})
	r.Run("localhost:10008") // listen and serve on 0.0.0.0:8080
}

// harish $ go run main.go
// [GIN-debug] [WARNING] Creating an Engine instance with the Logger and Recovery middleware already attached.

// [GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
//  - using env:   export GIN_MODE=release
//  - using code:  gin.SetMode(gin.ReleaseMode)

// [GIN-debug] GET    /ping                     --> main.main.func1 (3 handlers)
// [GIN-debug] [WARNING] You trusted all proxies, this is NOT safe. We recommend you to set a value.
// Please check https://pkg.go.dev/github.com/gin-gonic/gin#readme-don-t-trust-all-proxies for details.
// [GIN-debug] Listening and serving HTTP on localhost:10008
// [GIN] 2024/06/23 - 00:04:42 | 200 |        85.5µs |       127.0.0.1 | GET      "/ping"
// [GIN] 2024/06/23 - 00:04:42 | 404 |         625ns |       127.0.0.1 | GET      "/favicon.ico"
// [GIN] 2024/06/23 - 00:04:47 | 200 |      30.042µs |       127.0.0.1 | GET      "/ping"
// [GIN] 2024/06/23 - 00:04:48 | 200 |      16.834µs |       127.0.0.1 | GET      "/ping"



