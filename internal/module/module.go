package module

import (
	"github.com/vsrtferrum/OzonIntro/internal/storage"
)

type Module struct {
	storage storage.StorageAtions
}

func NewModule(storage storage.StorageAtions) *Module {
	return &Module{storage: storage}
}
