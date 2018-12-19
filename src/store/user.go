package store

import (
	"context"
	"gae-test-project/src/model"

	"go.mercari.io/datastore"
	"go.mercari.io/datastore/aedatastore"
)

const UserKindName = "User"

// UserStore は、UserのCRUDを担保する
type UserStore struct {
	cli datastore.Client
}

// NewUserStore は、UserStoreを用意する
func NewUserStore(c context.Context) (*UserStore, error) {
	cli, err := aedatastore.FromContext(c) // Memo: 2018/12/29 aedatastore であれば cli.Close() を呼び出す必要はない
	if err != nil {
		return nil, err
	}
	return &UserStore{
		cli: cli,
	}, nil
}

// newKey は、Keyの生成を担当する
func (store *UserStore) newKey(email string) datastore.Key {
	return store.cli.NameKey(UserKindName, email, nil)
}

// Get は、Userを一件取得する
func (store *UserStore) Get(email string) (*model.User, error) {
	user := model.NewUser(email, "", "", "")
	if err := store.cli.Get(store.cli.Context(), store.newKey(email), user); err != nil {
		return nil, err
	}
	return user, nil
}

// List は、Userを複数件取得する
func (store *UserStore) List(firstName, lastName, address string) ([]*model.User, error) {
	q := store.cli.NewQuery(UserKindName)
	if firstName != "" {
		q = q.Filter("Detail.FirstName =", firstName)
	}
	if lastName != "" {
		q = q.Filter("Detail.LastName =", lastName)
	}
	if address != "" {
		q = q.Filter("Detail.Address =", address)
	}

	userList := make([]*model.User, 0)
	if _, err := store.cli.GetAll(store.cli.Context(), q, userList); err != nil {
		return nil, err
	}
	return userList, nil
}

// InsertOrUpdate は、Userを新規作成または更新する
func (store *UserStore) InsertOrUpdate(email, firstName, lastName, address string) (*model.User, error) {
	user := model.NewUser(email, firstName, lastName, address)
	// TODO: insertとupdateを分ける必要が出たらコメントアウトする
	//if _, err := store.Get(user.Email); err != nil {
	//	if err != datastore.ErrNoSuchEntity {
	//		return err
	//	}
	//}
	if _, err := store.cli.Put(store.cli.Context(), store.newKey(email), user); err != nil {
		return nil, err
	}
	return user, nil
}

// Delete は、Userを一件削除する
func (store *UserStore) Delete(email string) error {
	return store.cli.Delete(store.cli.Context(), store.newKey(email))
}
