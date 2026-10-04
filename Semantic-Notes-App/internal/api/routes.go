package api

import "github.com/gin-gonic/gin"

func RegisterRoutes(
	router *gin.Engine,
	handler *Handler,
) {

	router.GET("/", func(c *gin.Context) {

		c.JSON(
			200,
			gin.H{
				"message": "Semantic Notes API",
			},
		)

	})

	router.POST(
		"/notes",
		handler.CreateNote,
	)

	router.GET(
		"/notes",
		handler.GetNotes,
	)

	router.POST(
		"/notes/search",
		handler.SearchNotes,
	)

}
