package resolvers

import (
	"gql-test/repository"
)

type Resolver struct {
	repo repository.ItemRepository
}

func New(repo repository.ItemRepository) *Resolver {
	return &Resolver{
		repo: repo,
	}
}

func (res *Resolver) Item(args struct{ Id int32 }) (repository.Item, error) {
	item, err := res.repo.Item(args.Id)
	item.SetRepo(&res.repo)
	return item, err
}
