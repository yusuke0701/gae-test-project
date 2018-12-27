package api

import (
	"gae-test-project/src/store"
	"net/http"

	"github.com/gin-gonic/gin"
	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
)

// InitUserAPI は、UserAPIをルーティングに設定する
func InitUserAPI(g *gin.RouterGroup) {
	api := &UserAPI{}
	g.GET(":email", api.get)
	g.GET("", api.list)
	g.PUT("", api.insertOrUpdate)
	g.DELETE("", api.delete)
}

// UserAPI はUserAPIをまとめる
type UserAPI struct{}

func (api *UserAPI) get(ginC *gin.Context) {
	appC := appengine.NewContext(ginC.Request)

	email := ginC.Param("email")
	log.Debugf(appC, "email = %s", email)

	uStore, err := store.NewUserStore(appC)
	if err != nil {
		ginC.String(http.StatusInternalServerError, err.Error())
		return
	}
	user, err := uStore.Get(email)
	if err != nil {
		ginC.String(http.StatusInternalServerError, err.Error())
		return
	}
	ginC.JSON(http.StatusOK, user)
}

func (api *UserAPI) list(ginC *gin.Context) {
	appC := appengine.NewContext(ginC.Request)

	firstName := ginC.Query("firstName")
	lastName := ginC.Query("lastName")
	address := ginC.Query("address")

	uStore, err := store.NewUserStore(appC)
	if err != nil {
		ginC.String(http.StatusInternalServerError, err.Error())
		return
	}
	userList, err := uStore.List(firstName, lastName, address)
	if err != nil {
		ginC.String(http.StatusInternalServerError, err.Error())
		return
	}
	ginC.JSON(http.StatusOK, userList)
}

func (api *UserAPI) insertOrUpdate(ginC *gin.Context) {
	appC := appengine.NewContext(ginC.Request)

	email, ok := ginC.GetQuery("email")
	if !ok {
		ginC.String(http.StatusBadRequest, "email is required")
		return
	}
	firstName, ok := ginC.GetQuery("firstName")
	if !ok {
		ginC.String(http.StatusBadRequest, "firstName is required")
		return
	}
	lastName, ok := ginC.GetQuery("lastName")
	if !ok {
		ginC.String(http.StatusBadRequest, "lastName is required")
		return
	}
	address, ok := ginC.GetQuery("address")
	if !ok {
		ginC.String(http.StatusBadRequest, "address is required")
		return
	}

	uStore, err := store.NewUserStore(appC)
	if err != nil {
		ginC.String(http.StatusInternalServerError, err.Error())
		return
	}
	user, err := uStore.InsertOrUpdate(email, firstName, lastName, address)
	if err != nil {
		log.Errorf(appC, err.Error())
		ginC.String(http.StatusInternalServerError, err.Error())
		return
	}
	ginC.JSON(http.StatusOK, user)
}

func (api *UserAPI) delete(ginC *gin.Context) {
	appC := appengine.NewContext(ginC.Request)

	email := ginC.Param("email")

	uStore, err := store.NewUserStore(appC)
	if err != nil {
		ginC.String(http.StatusInternalServerError, err.Error())
		return
	}
	if err := uStore.Delete(email); err != nil {
		ginC.String(http.StatusInternalServerError, err.Error())
		return
	}
	ginC.String(http.StatusOK, "deleted")
}
