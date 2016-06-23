package handler

import (
	"log"

	cat "github.com/sunfmin/myou/categories/proto/categories"
	"golang.org/x/net/context"
)

type CatHandler struct{}

func (e *CatHandler) List(ctx context.Context, filter *cat.Filter, cats *cat.CategoryList) error {
	log.Print("Received CatHandler.Call request")

	return nil
}
