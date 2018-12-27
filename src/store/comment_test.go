package store

import (
	"gae-test-project/src/testutil"
	"net/http"
	"testing"

	"google.golang.org/appengine"
	"google.golang.org/appengine/aetest"
)

func TestCommentStore(t *testing.T) {
	inst, err := aetest.NewInstance(nil)
	if err != nil {
		t.Fatal(err)
	}
	defer inst.Close()

	r, err := inst.NewRequest(http.MethodGet, "/", nil)
	if err != nil {
		t.Fatal(err)
	}
	c := appengine.NewContext(r)

	cStore, err := NewCommentStore(c)
	if err != nil {
		t.Fatal(err)
	}

	var (
		email = "test@sample.co.jp"
		title = "hoge"
		body  = "fuga"
	)

	insertResult, err := cStore.Insert(email, title, body)
	if err != nil {
		t.Fatal(err)
	}
	testutil.AssertEquals(t, "insertResult.Title", insertResult.Title, title)
	testutil.AssertEquals(t, "insertResult.Body", insertResult.Body, body)

	getResult, err := cStore.Get(email, insertResult.ID)
	if err != nil {
		t.Fatal(err)
	}
	testutil.AssertEquals(t, "getResult.Title", getResult.Title, title)
	testutil.AssertEquals(t, "getResult.Body", getResult.Body, body)

	listByEmailResult, err := cStore.ListByEmail(email)
	if err != nil {
		t.Fatal(err)
	}
	testutil.AssertIntEquals(t, "len(listByEmailResult)", len(listByEmailResult), 1)
	testutil.AssertEquals(t, "listByEmailResult[0].Title", listByEmailResult[0].Title, title)
	testutil.AssertEquals(t, "listByEmailResult[0].Body", listByEmailResult[0].Body, body)
}
