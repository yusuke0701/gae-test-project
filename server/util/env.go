package util

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"cloud.google.com/go/compute/metadata"
)

var (
	// IsLocal ローカル環境の場合は true になる
	IsLocal bool
	// ProjectID GCPプロジェクトID
	ProjectID string
	// ServiceName GAEのサービス名
	ServiceName string
	// Version GAEのバージョン名
	Version string
	// TraceID トレースID
	TraceID string
	// ServiceAccountName GAEのデフォルトサービスアカウント名
	ServiceAccountName string
	// ServiceAccountID GAEのデフォルトサービスアカウントID
	ServiceAccountID string
)

func init() {
	var err error

	ProjectID = os.Getenv("GOOGLE_CLOUD_PROJECT")
	IsLocal = ProjectID == ""
	if !IsLocal {
		ServiceName = os.Getenv("GAE_SERVICE")
		Version = os.Getenv("GAE_VERSION")

		// The account may be empty or the string "default" to use the instance's main account.
		ServiceAccountName, err = metadata.Email("")
		if err != nil {
			log.Fatal(err)
		}

		ServiceAccountID = fmt.Sprintf(
			"projects/%s/serviceAccounts/%s",
			ProjectID,
			ServiceAccountName,
		)
	}
}

// SetTraceID トレースIDをセットする
func SetTraceID(r *http.Request) {
	TraceID = strings.SplitN(r.Header.Get("X-Cloud-Trace-Context"), "/", 2)[0]
}
