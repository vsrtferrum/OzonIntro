package storage

import "github.com/vsrtferrum/OzonIntro/internal/model"

type StorageAtions interface {
	GetPostsList() (*[]model.PostList, error)
	GetPostById(uint64) (*model.Post, *[]model.Comments, error)
	WriteComment(*model.WriteComment) (uint64, error)
	WritePost(*model.WritePost) (uint64, error)
}
