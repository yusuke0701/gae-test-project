package store

import (
	"context"
	"fmt"
	"time"

	"cloud.google.com/go/datastore"

	"github.com/yusuke0701/gae-test-project/backend/db/datastore/model"
	errs "github.com/yusuke0701/goutils/error"
	"github.com/yusuke0701/goutils/manufacture"
)

// Thread は、スレッド情報のDB操作を担保する
type Thread struct{}

// Insert は、スレッドを一件挿入する
func (tStore *Thread) Insert(ctx context.Context, t *model.Thread) error {
	id, err := tStore.newID()
	if err != nil {
		return err
	}
	t.ID = id

	if err := tStore.canInsert(ctx, t); err != nil {
		return err
	}

	t.CreatedAt = time.Now()

	if _, err := datastoreClient.Put(ctx, tStore.newKey(t.ID), t); err != nil {
		return err
	}
	return nil
}

// Get は、スレッドを一件取得する
func (tStore *Thread) Get(ctx context.Context, id string) (t *model.Thread, err error) {
	if err := datastoreClient.Get(ctx, tStore.newKey(id), t); err != nil {
		if err == datastore.ErrNoSuchEntity {
			return nil, &errs.ErrNotFound{Msg: "no such entity"}
		}
		return nil, err
	}
	return
}

// List は、スレッドを一覧取得する
func (tStore *Thread) List(ctx context.Context) (ts []*model.Thread, err error) {
	q := datastore.NewQuery(tStore.kind())
	_, err = datastoreClient.GetAll(ctx, q, &ts)
	return
}

// Update は、スレッドを一件更新する
func (tStore *Thread) Update(ctx context.Context, t *model.Thread) error {
	if err := tStore.canUpdate(ctx, t); err != nil {
		return err
	}

	t.UpdatedAt = time.Now()

	if _, err := datastoreClient.Put(ctx, tStore.newKey(t.ID), t); err != nil {
		return err
	}
	return nil
}

// Delete は、スレッドを一件削除する
func (tStore *Thread) Delete(ctx context.Context, id string) error {
	if err := tStore.canDelete(ctx, id); err != nil {
		return err
	}

	t, err := tStore.Get(ctx, id)
	if err != nil {
		return err
	}

	t.DeletedAt = time.Now()

	if _, err := datastoreClient.Put(ctx, tStore.newKey(t.ID), t); err != nil {
		return err
	}
	return nil
}

// 内部向けメソッド郡

func (tStore *Thread) kind() string {
	return "thread"
}

func (tStore *Thread) newKey(id string) *datastore.Key {
	return datastore.NameKey(tStore.kind(), id, nil)
}

func (tStore *Thread) newID() (string, error) {
	return manufacture.NewUUID()
}

func (tStore *Thread) canInsert(ctx context.Context, t *model.Thread) error {
	if _, err := tStore.Get(ctx, t.ID); err != nil {
		if _, ok := err.(*errs.ErrNotFound); ok {
			// ok
		} else {
			return err
		}
	} else {
		return &errs.ErrConflict{Msg: fmt.Sprintf("invalid id = %s", t.ID)}
	}

	if _, err := (&Tag{}).GetMulti(ctx, t.TagIDs); err != nil {
		return err
	}
	return nil
}

func (tStore *Thread) canUpdate(ctx context.Context, new *model.Thread) error {
	old, err := tStore.Get(ctx, new.ID)
	if err != nil {
		return err
	}

	if old.UpdatedAt.After(new.UpdatedAt) {
		return &errs.ErrConflict{Msg: fmt.Sprintf("invalid updateAt")}
	}

	if _, err := (&Tag{}).GetMulti(ctx, new.TagIDs); err != nil {
		return err
	}
	return nil
}

func (tStore *Thread) canDelete(ctx context.Context, id string) error {
	if _, err := tStore.Get(ctx, id); err != nil {
		return err
	}
	return nil
}
