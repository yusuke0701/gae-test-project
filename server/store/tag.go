package store

import (
	"context"
	"fmt"
	"time"

	"cloud.google.com/go/datastore"
	"google.golang.org/api/iterator"

	"github.com/yusuke0701/gae-test-project/model"
	"github.com/yusuke0701/gae-test-project/util"
)

// Tag は、タグのDB操作を担保する
type Tag struct{}

// Insert は、タグを一件挿入する
func (tStore *Tag) Insert(ctx context.Context, t *model.Tag) error {
	id, err := tStore.newID(ctx)
	if err != nil {
		return err
	}
	t.ID = id

	if err := tStore.canInsert(ctx, t.ID); err != nil {
		return err
	}

	t.CreatedAt = time.Now()

	if _, err := datastoreClient.Put(ctx, tStore.newKey(t.ID), t); err != nil {
		return err
	}
	return nil
}

// Get は、タグを一件取得する
func (tStore *Tag) Get(ctx context.Context, id int64) (t *model.Tag, err error) {
	if err := datastoreClient.Get(ctx, tStore.newKey(id), t); err != nil {
		if err == datastore.ErrNoSuchEntity {
			return nil, &util.ErrNotFound{Msg: "no such entity"}
		}
		return nil, err
	}
	return
}

// GetMulti は、タグを一括取得する
func (tStore *Tag) GetMulti(ctx context.Context, ids []int64) (ts []*model.Tag, err error) {
	keys := make([]*datastore.Key, 0, len(ts))
	for _, t := range ts {
		keys = append(keys, tStore.newKey(t.ID))
	}
	if err := datastoreClient.GetMulti(ctx, keys, ts); err != nil {
		if err == datastore.ErrNoSuchEntity {
			return nil, &util.ErrNotFound{Msg: "no such entity"}
		}
		return nil, err
	}
	return
}

// List は、タグを一覧取得する
func (tStore *Tag) List(ctx context.Context) (ts []*model.Tag, err error) {
	q := datastore.NewQuery(tStore.kind())
	_, err = datastoreClient.GetAll(ctx, q, &ts)
	return
}

// Update は、タグを一件更新する
func (tStore *Tag) Update(ctx context.Context, t *model.Tag) error {
	if err := tStore.canUpdate(ctx, t); err != nil {
		return err
	}

	t.UpdatedAt = time.Now()

	if _, err := datastoreClient.Put(ctx, tStore.newKey(t.ID), t); err != nil {
		return err
	}
	return nil
}

// Delete は、タグを一件削除する
func (tStore *Tag) Delete(ctx context.Context, id int64) error {
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

func (tStore *Tag) kind() string {
	return "tag"
}

func (tStore *Tag) newKey(id int64) *datastore.Key {
	return datastore.IDKey(tStore.kind(), id, nil)
}

func (tStore *Tag) newID(ctx context.Context) (int64, error) {
	latest := new(model.Tag)
	{
		q := datastore.NewQuery(tStore.kind()).KeysOnly()
		q = q.Order("-ID")

		iter := datastoreClient.Run(ctx, q)
		if _, err := iter.Next(latest); err != nil {
			if err == iterator.Done {
				util.LogInfof(ctx, "start a new thread")
				return 1, nil
			}
			return 0, err
		}
	}
	return latest.ID + 1, nil
}

func (tStore *Tag) canInsert(ctx context.Context, id int64) error {
	if _, err := tStore.Get(ctx, id); err != nil {
		if _, ok := err.(*util.ErrNotFound); ok {
			// ok
		} else {
			return err
		}
	} else {
		return &util.ErrConflict{Msg: fmt.Sprintf("invalid id = %d", id)}
	}
	return nil
}

func (tStore *Tag) canUpdate(ctx context.Context, new *model.Tag) error {
	old, err := tStore.Get(ctx, new.ID)
	if err != nil {
		return err
	}

	if old.UpdatedAt.After(new.UpdatedAt) {
		return &util.ErrConflict{Msg: fmt.Sprintf("invalid updateAt")}
	}
	return nil
}

func (tStore *Tag) canDelete(ctx context.Context, id int64) error {
	if _, err := tStore.Get(ctx, id); err != nil {
		return err
	}
	return nil
}
