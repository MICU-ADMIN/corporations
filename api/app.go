package main

import (
	"net/http"

	"mosque.icu.corporations.api/services/hadith_service"

	"github.com/gin-gonic/gin"
)




func main() {

router := gin.Default()

router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "hello")
	})

// hadiths
	router.GET("/hadith", hadith_service.GetHadith)
	router.GET("/ahadith", hadith_service.GetAhadith)
	router.GET("/ahadith/filter", hadith_service.GetAhadithFiltered)



	router.Run()

	
}


