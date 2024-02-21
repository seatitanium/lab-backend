package lab_backend

import "github.com/gin-gonic/gin"

func main() {
	router := gin.New()
	router.Use(middlewares())

}
