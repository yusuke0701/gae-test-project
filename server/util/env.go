package util

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"cloud.google.com/go/compute/metadata"
)

var (
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

	if err := checkEnv(); err != nil {
		log.Fatal(err)
	}
}

// SetTraceID トレースIDをセットする
func SetTraceID(r *http.Request) {
	TraceID = strings.SplitN(r.Header.Get("X-Cloud-Trace-Context"), "/", 2)[0]
}

func checkEnv() error {
	if ProjectID == "" {
		return errors.New("プロジェクトID が空です")
	}
	if ServiceAccountName == "" {
		return errors.New("サービスアカウント名 が空です")
	}
	if ServiceAccountID == "" {
		return errors.New("サービスアカウントID が空です")
	}
	return nil
}
