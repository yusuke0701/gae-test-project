package store

import (
	"context"
	"gae-test-project/model"
	"gae-test-project/util"

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

// Insert は、コメントを一件保存する
func (cStore *Comment) Insert(ctx context.Context, c *model.Comment) error {
	if _, err := util.DatastoreClient.Put(ctx, cStore.newKey(c.ID), c); err != nil {
		return err
	}
	return nil
}

// Get は、コメントを一件取得する
func (cStore *Comment) Get(ctx context.Context, id string) (c *model.Comment, err error) {
	err = util.DatastoreClient.Get(ctx, cStore.newKey(id), c)
	return
}
