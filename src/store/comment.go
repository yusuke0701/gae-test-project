package store

import (
	"context"
	"gae-test-project/src/model"

	"go.mercari.io/datastore/aedatastore"
	"go.mercari.io/datastore/boom"
)

// CommentStore は、CommentのCRUDを担保する
type CommentStore struct {
	kindName string
	b        *boom.Boom
}

// NewCommentStore は、CommentStoreを用意する
func NewCommentStore(c context.Context) (*CommentStore, error) {
	cli, err := aedatastore.FromContext(c) // Memo: 2018/12/29 aedatastore であれば cli.Close() を呼び出す必要はない
	if err != nil {
		return nil, err
	}
	b := boom.FromClient(c, cli)
	return &CommentStore{
		kindName: b.Kind(model.Comment{}),
		b:        b,
	}, nil
}

// Get は、Commentを一件取得する
func (store *CommentStore) Get(email string, id int64) (*model.Comment, error) {
	comment, err := model.NewComment(store.b, email, id, "", "")
	if err != nil {
		return nil, err
	}
	if err := store.b.Get(comment); err != nil {
		return nil, err
	}
	return comment, nil
}

// ListByEmail は、emailを元にCommentを複数件取得する
func (store *CommentStore) ListByEmail(email string) ([]*model.Comment, error) {
	user := model.NewUser(email, "", "", "")
	key, err := store.b.KeyError(user)
	if err != nil {
		return nil, err
	}

	q := store.b.NewQuery(store.kindName).Ancestor(key)
	comments := make([]*model.Comment, 0)
	if _, err := store.b.GetAll(q, &comments); err != nil {
		return nil, err
	}
	return comments, nil
}

// Insert は、Commentを新規作成する
func (store *CommentStore) Insert(email, title, body string) (*model.Comment, error) {
	comment, err := model.NewComment(store.b, email, 0, title, body) // idに0を入れることで、Datastore側でよしなにやってくれる
	if err != nil {
		return nil, err
	}
	if _, err := store.b.Put(comment); err != nil {
		return nil, err
	}
	return comment, nil
}

// Update は、Commentを更新する
func (store *CommentStore) Update(id int64, email, title, body string) (*model.Comment, error) {
	comment, err := model.NewComment(store.b, email, id, title, body)
	if err != nil {
		return nil, err
	}
	if _, err := store.b.Put(comment); err != nil {
		return nil, err
	}
	return comment, nil
}

// Delete は、Commentを削除する
func (store *CommentStore) Delete(id int64) error {
	comment, err := model.NewComment(store.b, "", id, "", "")
	if err != nil {
		return err
	}
	return store.b.Delete(comment)
}
