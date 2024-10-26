package repository

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	task "github.com/venil7/func/task"
)

type ItemRepository interface {
	Item(id int32) (Item, error)
	Items(ids []int32) ([]Item, error)
}

type ItemClient struct {
	client http.Client
}

func New() *ItemClient {
	client := http.Client{}
	return &ItemClient{
		client: client,
	}
}

func (r *ItemClient) Item(id int32) (Item, error) {
	body := task.Task[io.ReadCloser](func() (io.ReadCloser, error) {
		url := fmt.Sprintf("https://hacker-news.firebaseio.com/v0/item/%d.json", id)
		resp, err := r.client.Get(url)
		if err != nil {
			return nil, err
		}
		return resp.Body, nil
	})

	item := task.Then(body, func(body io.ReadCloser) (Item, error) {
		defer body.Close()
		decoder := json.NewDecoder(body)
		var item Item
		err := decoder.Decode(&item)
		return item, err
	})

	return item()
}

func (r *ItemClient) Items(ids []int32) ([]Item, error) {
	item := task.From1(r.Item)
	items := task.Traverse(ids, func(id int32) task.Task[Item] {
		return item(id)
	})
	return items()
}
