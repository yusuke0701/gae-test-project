package util

import (
	"errors"
	"fmt"
	"log"
	"os"

	"cloud.google.com/go/compute/metadata"
)

var (
	ProjectID          string
	ServiceAccountName string
	ServiceAccountID   string
)

func init() {
	var err error

	ProjectID = os.Getenv("GOOGLE_CLOUD_PROJECT")

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
