package middlewares

import (
	"net/http"

	"github.com/HoloPirates/mogupantsu/templates"
	"github.com/gin-gonic/gin"
)

// NotFoundHandler : Controller for displaying 404 error page
func NotFoundHandler(c *gin.Context) {
	templates.HttpError(c, http.StatusNotFound)
}
