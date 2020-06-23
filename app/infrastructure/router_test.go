import (
	"github.com/gin-gonic/gin"
	"github.com/hiroshimashu/ei-rest/app/controllers"
)

var Router *gin.Engine

func TestUserRouter(t *testing.T) {

}

func init() {
    router := gin.Default()

    userController := controllers.NewMockUserController()

    router.POST("/users", func(c *gin.Context) { userController.Create(c) })
    router.GET("/users", func(c *gin.Context) { userController.Index(c) })

    Router = router
}