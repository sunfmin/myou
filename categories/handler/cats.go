package handler

import (
	"errors"
	"log"
	"strings"

	"github.com/sunfmin/myou/categories/db"
	proto "github.com/sunfmin/myou/categories/proto/categories"
	"golang.org/x/net/context"
)

type CatHandler struct{}

func (e *CatHandler) List(ctx context.Context, filter *proto.Filter, cats *proto.CategoryList) (err error) {
	log.Print("Received CatHandler.List request")
	if filter.Project == nil || len(filter.Project.Id) == 0 {
		err = errors.New("project not provided.")
		return
	}

	var cs []*db.Category

	wh := db.DB.Where(db.Category{ProjectId: filter.Project.Id})
	if len(strings.TrimSpace(filter.Name)) == 0 {
		wh = wh.Where("name LIKE ?", "%"+filter.Name+"%")
	}
	err = wh.Find(&cs).Error
	for _, c := range cs {
		cats.Categories = append(cats.Categories, &proto.Category{
			Id:   c.Id,
			Name: c.Name,
		})
	}
	return
}

func (e *CatHandler) Put(context context.Context, input *proto.Input, cat *proto.Category) (err error) {
	log.Print("Received CatHandler.Put request")
	c := db.Category{
		Id:        input.Category.Id,
		ProjectId: input.Project.Id,
		Name:      input.Category.Name,
	}
	err = db.DB.Where(db.Category{Id: input.Category.Id, ProjectId: input.Project.Id}).Assign(c).FirstOrCreate(&c).Error
	return
}

func (e *CatHandler) Delete(context context.Context, req *proto.DeleteRequest, cat *proto.Category) error {
	log.Print("Received CatHandler.Delete request")

	return nil
}
