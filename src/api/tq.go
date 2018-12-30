package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
	"google.golang.org/appengine/taskqueue"
)

// InitTQAPI は、TaskQueueAPIをルーティングに設定する
func InitTQAPI(g *gin.RouterGroup) {
	api := &tqAPI{}
	g.GET("", api.tqSample)
	g.POST("", api.tqSample)
}

type tqAPI struct{}

func (api *tqAPI) tqSample(ginC *gin.Context) {
	appC := appengine.NewContext(ginC.Request)

	// URLが直叩きされた場合はTaskQueueに投げ込む。
	if ginC.Request.Header.Get("X-AppEngine-QueueName") == "" {
		t := taskqueue.NewPOSTTask(ginC.Request.URL.Path, nil)
		if _, err := taskqueue.Add(appC, t, "tqSample"); err != nil {
			ginC.String(http.StatusInternalServerError, err.Error())
			return
		}
		return
	}

	log.Debugf(appC, "TQの実行に成功")
	ginC.String(http.StatusOK, "TQの実行に成功したよー")
	return
}
