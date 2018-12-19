package api

import (
	"net/http"

	"gae-test-project/src/store"

	"github.com/gin-gonic/gin"
	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
)

// InitUserAPI は、UserAPIをルーティングに設定する
func InitUserAPI(g *gin.RouterGroup) {
	api := &UserAPI{}
	g.GET(":email", api.Get)
	g.GET("", api.List)
	g.PUT("", api.InsertOrUpdate)
	g.DELETE("", api.Delete)
}

// UserAPI はUserAPIをまとめる
type UserAPI struct{}

// Get は、User取得API
func (api *UserAPI) Get(ginC *gin.Context) {
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

// List は、User一覧取得API
func (api *UserAPI) List(ginC *gin.Context) {
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

// InsertOrUpdate は、User新規作成または更新API
func (api *UserAPI) InsertOrUpdate(ginC *gin.Context) {
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

// Delete は、User削除API
func (api *UserAPI) Delete(ginC *gin.Context) {
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
