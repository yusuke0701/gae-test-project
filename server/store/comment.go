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

// Comment は、コメントのDB操作を担保する
type Comment struct{}

// Insert は、コメントを一件挿入する
func (cStore *Comment) Insert(ctx context.Context, c *model.Comment) error {
	id, err := cStore.newID(ctx)
	if err != nil {
		return err
	}
	c.ID = id

	if err := cStore.canInsert(ctx, c); err != nil {
		return err
	}

	c.CreatedAt = time.Now()

	if _, err := datastoreClient.Put(ctx, cStore.newKey(c.ID), c); err != nil {
		return err
	}
	return nil
}

// Get は、コメントを一件取得する
func (cStore *Comment) Get(ctx context.Context, id int64) (c *model.Comment, err error) {
	if err := datastoreClient.Get(ctx, cStore.newKey(id), c); err != nil {
		if err == datastore.ErrNoSuchEntity {
			return nil, &util.ErrNotFound{Msg: "no such entity"}
		}
		return nil, err
	}
	return
}

// List は、コメントを一覧取得する
func (cStore *Comment) List(ctx context.Context) (cs []*model.Comment, err error) {
	q := datastore.NewQuery(cStore.kind())
	_, err = datastoreClient.GetAll(ctx, q, &cs)
	return
}

// Update は、コメントを一件更新する
func (cStore *Comment) Update(ctx context.Context, c *model.Comment) error {
	if err := cStore.canUpdate(ctx, c); err != nil {
		return err
	}

	c.UpdatedAt = time.Now()

	if _, err := datastoreClient.Put(ctx, cStore.newKey(c.ID), c); err != nil {
		return err
	}
	return nil
}

// Delete は、コメントを一件削除する
func (cStore *Comment) Delete(ctx context.Context, id int64) error {
	if err := cStore.canDelete(ctx, id); err != nil {
		return err
	}

	c, err := cStore.Get(ctx, id)
	if err != nil {
		return err
	}

	c.DeletedAt = time.Now()

	if _, err := datastoreClient.Put(ctx, cStore.newKey(c.ID), c); err != nil {
		return err
	}
	return nil
}

// 内部向けメソッド郡

func (cStore *Comment) kind() string {
	return "comment"
}

func (cStore *Comment) newKey(id int64) *datastore.Key {
	return datastore.IDKey(cStore.kind(), id, nil)
}

func (cStore *Comment) newID(ctx context.Context) (int64, error) {
	latest := new(model.Comment)
	{
		q := datastore.NewQuery(cStore.kind())
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

func (cStore *Comment) canInsert(ctx context.Context, c *model.Comment) error {
	if _, err := cStore.Get(ctx, c.ID); err != nil {
		if _, ok := err.(*util.ErrNotFound); ok {
			// ok
		} else {
			return err
		}
	} else {
		return &util.ErrConflict{Msg: fmt.Sprintf("invalid id = %d", c.ID)}
	}
	if _, err := (&Thread{}).Get(ctx, c.ThreadID); err != nil {
		return err
	}
	return nil
}

func (cStore *Comment) canUpdate(ctx context.Context, new *model.Comment) error {
	old, err := cStore.Get(ctx, new.ID)
	if err != nil {
		return err
	}

	if old.UpdatedAt.After(new.UpdatedAt) {
		return &util.ErrConflict{Msg: fmt.Sprintf("invalid updateAt")}
	}
	return nil
}

func (cStore *Comment) canDelete(ctx context.Context, id int64) error {
	if _, err := cStore.Get(ctx, id); err != nil {
		return err
	}
	return nil
}
