package store

import (
	"context"
	"fmt"
	"time"

	"cloud.google.com/go/datastore"

	"github.com/yusuke0701/gae-test-project/model"
	"github.com/yusuke0701/gae-test-project/util"
)

// Account は、アカウント情報のDB操作を担保する
type Account struct{}

// Insert は、アカウントを一件挿入する
func (aStore *Account) Insert(ctx context.Context, a *model.Account) error {
	id, err := aStore.newID()
	if err != nil {
		return err
	}
	a.ID = id

	if err := aStore.canInsert(ctx, a.ID); err != nil {
		return err
	}

	a.CreatedAt = time.Now()

	if _, err := util.DatastoreClient.Put(ctx, aStore.newKey(a.ID), a); err != nil {
		return err
	}
	return nil
}

// Get は、アカウントを一件取得する
func (aStore *Account) Get(ctx context.Context, id string) (a *model.Account, err error) {
	if err := util.DatastoreClient.Get(ctx, aStore.newKey(id), a); err != nil {
		if err == datastore.ErrNoSuchEntity {
			return nil, &util.ErrNotFound{Msg: "no such entity"}
		}
		return nil, err
	}
	return
}

// List は、アカウントを一覧取得する
func (aStore *Account) List(ctx context.Context) (as []*model.Account, err error) {
	q := datastore.NewQuery(aStore.kind())
	_, err = util.DatastoreClient.GetAll(ctx, q, &as)
	return
}

// Update は、アカウントを一件更新する
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

// Delete は、アカウントを一件削除する
func (aStore *Account) Delete(ctx context.Context, id string) error {
	if err := aStore.canDelete(ctx, id); err != nil {
		return err
	}

	a, err := aStore.Get(ctx, id)
	if err != nil {
		return err
	}

	a.DeletedAt = time.Now()

	if _, err := util.DatastoreClient.Put(ctx, aStore.newKey(a.ID), a); err != nil {
		return err
	}
	return nil
}

// 内部向けメソッド郡

func (aStore *Account) kind() string {
	return "account"
}

func (aStore *Account) newKey(id string) *datastore.Key {
	return datastore.NameKey(aStore.kind(), id, nil)
}

func (aStore *Account) newID() (string, error) {
	return util.NewUUID()
}

func (aStore *Account) canInsert(ctx context.Context, id string) error {
	if _, err := aStore.Get(ctx, id); err != nil {
		if _, ok := err.(*util.ErrNotFound); ok {
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
	if err != nil {
		return err
	}

	if old.UpdatedAt.After(new.UpdatedAt) {
		return &util.ErrConflict{Msg: fmt.Sprintf("invalid updateAt")}
	}
	return nil
}

func (aStore *Account) canDelete(ctx context.Context, id string) error {
	if _, err := aStore.Get(ctx, id); err != nil {
		return err
	}
	return nil
}
