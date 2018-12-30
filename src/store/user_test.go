package store

import (
	"gae-test-project/src/testutil"
	"net/http"
	"testing"

	"google.golang.org/appengine"
	"google.golang.org/appengine/aetest"
)

func TestUserStore(t *testing.T) {
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

	uStore, err := NewUserStore(c)
	if err != nil {
		t.Fatal(err)
	}

	var (
		email     = "test@sample.co.jp"
		firstName = "テスト"
		lastName  = "太郎"
		address   = "サンプル"
	)

	insertResult, err := uStore.InsertOrUpdate(email, firstName, lastName, address)
	if err != nil {
		t.Fatal(err)
	}
	testutil.AssertEquals(t, "insertResult.Title", insertResult.Email, email)
	testutil.AssertEquals(t, "insertResult.Detail.FirstName", insertResult.Detail.FirstName, firstName)
	testutil.AssertEquals(t, "insertResult.Detail.LastName", insertResult.Detail.LastName, lastName)
	testutil.AssertEquals(t, "insertResult.Detail.Address", insertResult.Detail.Address, address)

	getResult, err := uStore.Get(email)
	if err != nil {
		t.Fatal(err)
	}
	testutil.AssertEquals(t, "getResult.Title", getResult.Email, email)
	testutil.AssertEquals(t, "getResult.Detail.FirstName", getResult.Detail.FirstName, firstName)
	testutil.AssertEquals(t, "getResult.Detail.LastName", getResult.Detail.LastName, lastName)
	testutil.AssertEquals(t, "getResult.Detail.Address", getResult.Detail.Address, address)

	listResult, err := uStore.List(firstName, lastName, address)
	if err != nil {
		t.Fatal(err)
	}
	testutil.AssertIntEquals(t, "len(listResult)", len(listResult), 1)
	testutil.AssertEquals(t, "listResult[0].Title", listResult[0].Email, email)
	testutil.AssertEquals(t, "listResult[0].Detail.FirstName", listResult[0].Detail.FirstName, firstName)
	testutil.AssertEquals(t, "listResult[0].Detail.LastName", listResult[0].Detail.LastName, lastName)
	testutil.AssertEquals(t, "listResult[0].Detail.Address", listResult[0].Detail.Address, address)

	if err := uStore.Delete(email); err != nil {
		t.Fatal(err)
	}

	if _, err := uStore.Get(email); err == nil {
		t.Fatal(err)
	}
}
