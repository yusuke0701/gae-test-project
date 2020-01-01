package store

import (
	"context"
	"fmt"
	"gae-test-project/model"
	"gae-test-project/util"
	"time"

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

	now := time.Now()
	a.CreatedAt = now
	a.UpdatedAt = now

	if _, err := util.DatastoreClient.Put(ctx, aStore.newKey(a.ID), a); err != nil {
		return err
	}
	return nil
}

// Get は、コメントを一件取得する
func (aStore *Account) Get(ctx context.Context, id string) (a *model.Account, err error) {
	if err := util.DatastoreClient.Get(ctx, aStore.newKey(id), a); err != nil {
		if err == datastore.ErrNoSuchEntity {
			return nil, &util.ErrNotFound{Msg: "no such entity"}
		}
		return nil, err
	}
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
	if err := aStore.canUpdate(ctx, a); err != nil {
		return err
	}

	a.UpdatedAt = time.Now()

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

func (aStore *Account) canUpdate(ctx context.Context, new *model.Account) error {
	old, err := aStore.Get(ctx, new.ID)
	if err == datastore.ErrNoSuchEntity {
		return &util.ErrNotFound{Msg: "no such entity"}
	} else if err != nil {
		return err
	}

	if old.UpdatedAt.After(new.UpdatedAt) {
		return &util.ErrConflict{Msg: fmt.Sprintf("invalid updateAt")}
	}
	return nil
}
