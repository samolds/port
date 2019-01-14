// Copyright (C) 2018 - 2019 Sam Olds

package database

import (
	"context"
	"errors"

	"cloud.google.com/go/datastore"
	"google.golang.org/api/option"
)

type DB struct {
	// https://cloud.google.com/datastore/docs/concepts/overview
	gaeDatastore *datastore.Client
}

func New(ctx context.Context, dbID string, credFile string) (db *DB,
	err error) {
	if dbID == "" {
		return nil, errors.New("a database id is required")
	}

	db = &DB{}
	if credFile == "" {
		db.gaeDatastore, err = datastore.NewClient(ctx, dbID)
	} else {
		db.gaeDatastore, err = datastore.NewClient(ctx, dbID,
			option.WithCredentialsFile(credFile))
	}

	if err != nil {
		return nil, err
	}

	return db, nil
}

func (db *DB) Get(ctx context.Context, kind Kind, id string) (
	r interface{}, err error) {
	key := datastore.NameKey(kind.String(), id, nil)
	err = db.gaeDatastore.Get(ctx, key, r)
	if err != nil {
		return nil, err
	}
	return r, nil
}

func (db *DB) GetAllLinks(ctx context.Context) ([]Link, error) {
	var r []Link
	query := datastore.NewQuery(LinkKind.String()).Order("display")
	_, err := db.gaeDatastore.GetAll(ctx, query, &r)
	if err != nil {
		return nil, err
	}
	return r, nil
}

func (db *DB) GetMostRecentNowText(ctx context.Context) (*NowText, error) {
	var r []NowText
	query := datastore.NewQuery(NowTextKind.String()).Order(
		"-creationTime").Limit(1)
	_, err := db.gaeDatastore.GetAll(ctx, query, &r)
	if err != nil {
		return nil, err
	}
	if len(r) != 1 {
		errors.New("more nowText results than expected")
	}
	return &r[0], nil
}

func (db *DB) Put(ctx context.Context, kind Kind, d Link) (id string,
	err error) {
	id, err = GenerateUUID()
	if err != nil {
		return "", err
	}

	key := datastore.NameKey(kind.String(), id, nil)
	_, err = db.gaeDatastore.Put(ctx, key, &d)
	if err != nil {
		return "", err
	}

	return id, nil
}
