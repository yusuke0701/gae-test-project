package store

import (
	"context"
	"fmt"
	"gae-test-project/model"
	"gae-test-project/util"

	"cloud.google.com/go/datastore"
)

// Account は、アカウント情報のDB操作を担保する
type Account struct{}

func (aStore *Account) kind() string {
	return "account"
}

func (aStore *Account) newKey(id string) *datastore.Key {
	return datastore.NameKey(aStore.kind(), id, nil)
}

// Insert は、コメントを一件挿入する
func (aStore *Account) Insert(ctx context.Context, a *model.Account) error {
	if err := aStore.canInsert(ctx, a.ID); err != nil {
		return err
	}

	if _, err := util.DatastoreClient.Put(ctx, aStore.newKey(a.ID), a); err != nil {
		return err
	}
	return nil
}

// Get は、コメントを一件取得する
func (aStore *Account) Get(ctx context.Context, id string) (a *model.Account, err error) {
	err = util.DatastoreClient.Get(ctx, aStore.newKey(id), a)
	return
}

// List は、コメントを一覧取得する
func (aStore *Account) List(ctx context.Context) (as []*model.Account, err error) {
	q := datastore.NewQuery(aStore.kind())
	_, err = util.DatastoreClient.GetAll(ctx, q, &as)
	return
}

// Update は、コメントを一件更新する
func (aStore *Account) Update(ctx context.Context, a *model.Account) error {
	if err := aStore.canUpdate(ctx, a.ID); err != nil {
		return err
	}

	if _, err := util.DatastoreClient.Put(ctx, aStore.newKey(a.ID), a); err != nil {
		return err
	}
	return nil
}

func (aStore *Account) canInsert(ctx context.Context, id string) error {
	if _, err := aStore.Get(ctx, id); err != nil {
		if err == datastore.ErrNoSuchEntity {
			// ok
		} else {
			return err
		}
	} else {
		return &util.ErrConflict{Msg: fmt.Sprintf("invalid id = %s", id)}
	}
	return nil
}

func (aStore *Account) canUpdate(ctx context.Context, id string) error {
	if _, err := aStore.Get(ctx, id); err != nil {
		if err == datastore.ErrNoSuchEntity {
			return &util.ErrNotFound{Msg: "no such entity"}
		}
		return err
	}
	return nil
}
