package store

import (
	"context"
	"gae-test-project/src/model"

	"go.mercari.io/datastore/aedatastore"
	"go.mercari.io/datastore/boom"
)

// UserStore は、UserのCRUDを担保する
type UserStore struct {
	kindName string
	b        *boom.Boom
}

// NewUserStore は、UserStoreを用意する
func NewUserStore(c context.Context) (*UserStore, error) {
	cli, err := aedatastore.FromContext(c) // Memo: 2018/12/29 aedatastore であれば cli.Close() を呼び出す必要はない
	if err != nil {
		return nil, err
	}
	b := boom.FromClient(c, cli)
	return &UserStore{
		kindName: b.Kind(model.User{}),
		b:        b,
	}, nil
}

// Get は、Userを一件取得する
func (store *UserStore) Get(email string) (*model.User, error) {
	user := model.NewUser(email, "", "", "")
	if err := store.b.Get(user); err != nil {
		return nil, err
	}
	return user, nil
}

// List は、Userを複数件取得する
func (store *UserStore) List(firstName, lastName, address string) ([]*model.User, error) {
	q := store.b.NewQuery(store.kindName)
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
	if _, err := store.b.GetAll(q, &userList); err != nil {
		return nil, err
	}
	return userList, nil
}

// InsertOrUpdate は、Userを新規作成または更新する
func (store *UserStore) InsertOrUpdate(email, firstName, lastName, address string) (*model.User, error) {
	// TODO: insertとupdateを分ける必要が出たらコメントアウトする
	//if _, err := store.Get(email); err != nil {
	//	if err != datastore.ErrNoSuchEntity {
	//		return err
	//	}
	//}
	user := model.NewUser(email, firstName, lastName, address)
	if _, err := store.b.Put(user); err != nil {
		return nil, err
	}
	return user, nil
}

// Delete は、Userを一件削除する
func (store *UserStore) Delete(email string) error {
	user := model.NewUser(email, "", "", "")
	return store.b.Delete(user)
}
