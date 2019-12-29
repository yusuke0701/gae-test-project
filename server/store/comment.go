package store

import (
	"context"
	"gae-test-project/connection"
	"gae-test-project/model"

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

// InsertOrUpadte は、コメントを一件保存する
func (cStore *Comment) InsertOrUpadte(ctx context.Context, c *model.Comment) error {
	if _, err := connection.DatastoreClient.Put(ctx, cStore.newKey(c.ID), c); err != nil {
		return err
	}
	return nil
}

// Get は、コメントを一件取得する
func (cStore *Comment) Get(ctx context.Context, id string) (c *model.Comment, err error) {
	err = connection.DatastoreClient.Get(ctx, cStore.newKey(id), c)
	return
}

// List は、コメントを一覧取得する
func (cStore *Comment) List(ctx context.Context) (cs []*model.Comment, err error) {
	q := datastore.NewQuery(cStore.kind())
	_, err = connection.DatastoreClient.GetAll(ctx, q, &cs)
	return
}
