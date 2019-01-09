package api

import (
	"gae-test-project/src/store"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
)

// InitCommentAPI は、CommentAPIをルーティングに設定する
func InitCommentAPI(g *gin.RouterGroup) {
	api := &commentAPI{}
	g.GET(":email/:id", api.get)
	g.GET(":email", api.listByEmail)
	g.PUT("", api.insert)
	g.POST(":id", api.update)
}

type commentAPI struct{}

func (api *commentAPI) get(ginC *gin.Context) {
	appC := appengine.NewContext(ginC.Request)

	email := ginC.Param("email")
	log.Debugf(appC, "email = %s", email)

	stringID := ginC.Param("id")
	id, err := strconv.ParseInt(stringID, 10, 64)
	if err != nil {
		ginC.String(http.StatusInternalServerError, err.Error())
		return
	}
	log.Debugf(appC, "id = %s", id)

	cStore, err := store.NewCommentStore(appC)
	if err != nil {
		ginC.String(http.StatusInternalServerError, err.Error())
		return
	}
	user, err := cStore.Get(email, id)
	if err != nil {
		ginC.String(http.StatusInternalServerError, err.Error())
		return
	}
	ginC.JSON(http.StatusOK, user)
}

func (api *commentAPI) listByEmail(ginC *gin.Context) {
	appC := appengine.NewContext(ginC.Request)

	email := ginC.Param("email")
	log.Debugf(appC, "email = %s", email)

	cStore, err := store.NewCommentStore(appC)
	if err != nil {
		ginC.String(http.StatusInternalServerError, err.Error())
		return
	}
	user, err := cStore.ListByEmail(email)
	if err != nil {
		ginC.String(http.StatusInternalServerError, err.Error())
		return
	}
	ginC.JSON(http.StatusOK, user)
}

func (api *commentAPI) insert(ginC *gin.Context) {
	appC := appengine.NewContext(ginC.Request)

	email, ok := ginC.GetQuery("email")
	if !ok {
		ginC.String(http.StatusBadRequest, "email is required")
		return
	}
	title, ok := ginC.GetQuery("title")
	if !ok {
		ginC.String(http.StatusBadRequest, "title is required")
		return
	}
	body, ok := ginC.GetQuery("body")
	if !ok {
		ginC.String(http.StatusBadRequest, "body is required")
		return
	}

	cStore, err := store.NewCommentStore(appC)
	if err != nil {
		ginC.String(http.StatusInternalServerError, err.Error())
		return
	}
	user, err := cStore.Insert(email, title, body)
	if err != nil {
		log.Errorf(appC, err.Error())
		ginC.String(http.StatusInternalServerError, err.Error())
		return
	}
	ginC.JSON(http.StatusOK, user)
}

func (api *commentAPI) update(ginC *gin.Context) {
	appC := appengine.NewContext(ginC.Request)

	stringID := ginC.Param("id")
	id, err := strconv.ParseInt(stringID, 10, 64)
	if err != nil {
		ginC.String(http.StatusInternalServerError, err.Error())
		return
	}
	email, ok := ginC.GetQuery("email")
	if !ok {
		ginC.String(http.StatusBadRequest, "email is required")
		return
	}
	title, ok := ginC.GetQuery("title")
	if !ok {
		ginC.String(http.StatusBadRequest, "title is required")
		return
	}
	body, ok := ginC.GetQuery("body")
	if !ok {
		ginC.String(http.StatusBadRequest, "body is required")
		return
	}

	cStore, err := store.NewCommentStore(appC)
	if err != nil {
		ginC.String(http.StatusInternalServerError, err.Error())
		return
	}
	user, err := cStore.Update(id, email, title, body)
	if err != nil {
		log.Errorf(appC, err.Error())
		ginC.String(http.StatusInternalServerError, err.Error())
		return
	}
	ginC.JSON(http.StatusOK, user)
}
