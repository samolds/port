// Copyright (C) 2018 - 2020 Sam Olds

package database

import (
	"time"
)

type Kind string

func (m Kind) String() string { return string(m) }

const (
	LinkKind    Kind = "link"
	NowTextKind Kind = "nowtext"
)

// https://godoc.org/cloud.google.com/go/datastore
type Link struct {
	CreationTime time.Time `datastore:"creationTime"`
	Href         string    `datastore:"href"`
	Display      string    `datastore:"display"`
}

type NowText struct {
	CreationTime  time.Time `datastore:"creationTime"`
	ProfileImgSrc string    `datastore:"profileImgSrc"`
	HTMLText      string    `datastore:"htmlText,noindex"`
}
