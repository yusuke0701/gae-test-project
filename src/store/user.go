package store

import (
	"context"
	"gae-test-project/src/model"

	"google.golang.org/appengine/datastore"
)

const UserKindName = "user"

type UserStore struct {
	Context context.Context // appengineのContextを渡す
}

func (store *UserStore) Get(email string) (*model.User, error) {
	user := &model.User{}
	key := datastore.NewKey(store.Context, UserKindName, email, 0, nil)
	if err := datastore.Get(store.Context, key, user); err != nil {
		return nil, err
	}
	return user, nil
}

func (store *UserStore) List(firstName, lastName, address string) ([]*model.User, error) {
	q := datastore.NewQuery(UserKindName)
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
	if _, err := q.GetAll(store.Context, &userList); err != nil {
		return nil, err
	}
	return userList, nil
}

func (store *UserStore) InsertOrUpdate(user *model.User) error {
	// TODO: insertとupdateを分ける必要が出たらコメントアウトする
	//if _, err := store.Get(user.Email); err != nil {
	//	if err != datastore.ErrNoSuchEntity {
	//		return err
	//	}
	//}
	key := datastore.NewKey(store.Context, UserKindName, user.Email, 0, nil)
	if _, err := datastore.Put(store.Context, key, user); err != nil {
		return err
	}
	return nil
}

func (store *UserStore) Delete(email string) error {
	key := datastore.NewKey(store.Context, UserKindName, email, 0, nil)
	return datastore.Delete(store.Context, key)
}
