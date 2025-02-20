package module

import (
	"github.com/vsrtferrum/OzonIntro/internal/errors"
	"github.com/vsrtferrum/OzonIntro/internal/model"
)

type ModuleActions interface {
	AddComment(model.Comments) (uint64, error)
	AddPost(model.Post) (uint64, error)
	GetPosts() (*[]model.PostList, error)
	GetPost(uint64) (*model.Post, *[]model.Comments, error)
}

func (module *Module) AddComment(data *model.WriteComment) (uint64, error) {
	if len(data.Text) > 2000 {
		return 0, errors.ErrSizeComment
	}
	res, err := module.storage.WriteComment(data)
	return res, err
}

func (module *Module) AddPost(data *model.WritePost) (uint64, error) {
	return module.storage.WritePost(data)
}

func (module *Module) GetPosts() (*[]model.PostList, error) {
	return module.storage.GetPostsList()
}

func (module *Module) GetPost(id uint64) (*model.Post, *[]model.Comments, error) {
	return module.storage.GetPostById(id)
}
