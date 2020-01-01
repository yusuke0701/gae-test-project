package store

import (
	"context"
	"fmt"
	"gae-test-project/model"
	"gae-test-project/util"
	"time"

	"cloud.google.com/go/datastore"
)

// Comment は、コメントのDB操作を担保する
type Comment struct{}

func (cStore *Comment) kind() string {
	return "comment"
}

func (cStore *Comment) newKey(id string) *datastore.Key {
	return datastore.NameKey(cStore.kind(), id, nil)
}

// Insert は、コメントを一件挿入する
func (cStore *Comment) Insert(ctx context.Context, c *model.Comment) error {
	if err := cStore.canInsert(ctx, c.ID); err != nil {
		return err
	}

	now := time.Now()
	c.CreatedAt = now
	c.UpdatedAt = now

	if _, err := util.DatastoreClient.Put(ctx, cStore.newKey(c.ID), c); err != nil {
		return err
	}
	return nil
}

// Get は、コメントを一件取得する
func (cStore *Comment) Get(ctx context.Context, id string) (c *model.Comment, err error) {
	if err := util.DatastoreClient.Get(ctx, cStore.newKey(id), c); err != nil {
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
	_, err = util.DatastoreClient.GetAll(ctx, q, &cs)
	return
}

// Update は、コメントを一件更新する
func (cStore *Comment) Update(ctx context.Context, c *model.Comment) error {
	if err := cStore.canUpdate(ctx, c); err != nil {
		return nil
	}

	c.UpdatedAt = time.Now()

	if _, err := util.DatastoreClient.Put(ctx, cStore.newKey(c.ID), c); err != nil {
		return err
	}
	return nil
}

func (cStore *Comment) canInsert(ctx context.Context, id string) error {
	if _, err := cStore.Get(ctx, id); err != nil {
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

func (cStore *Comment) canUpdate(ctx context.Context, new *model.Comment) error {
	old, err := cStore.Get(ctx, new.ID)
	if err != nil {
		if err == datastore.ErrNoSuchEntity {
			return &util.ErrNotFound{Msg: "no such entity"}
		}
		return err
	}

	if old.UpdatedAt.After(new.UpdatedAt) {
		return &util.ErrConflict{Msg: fmt.Sprintf("invalid updateAt")}
	}
	return nil
}
