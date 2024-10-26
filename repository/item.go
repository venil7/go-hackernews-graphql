package repository

import (
	"errors"
)

type Item struct {
	Id          int32
	Deleted     bool
	Type        string
	By          string
	Time        int32
	Text        string
	Dead        bool
	Parent      int32
	Poll        int32
	Kids        []int32
	Url         string
	Score       int32
	Title       string
	Parts       []int32
	Descendants int32
	repo        *ItemRepository
}

func (item *Item) SetRepo(repo *ItemRepository) {
	item.repo = repo
}

func (item Item) Items() ([]Item, error) {
	if item.repo == nil {
		return nil, errors.New("item.client not set")
	}
	return (*item.repo).Items(item.Kids)
}
