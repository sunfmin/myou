package tests

import (
	"fmt"
	"testing"

	"golang.org/x/net/context"

	"github.com/micro/go-micro"
	proto "github.com/sunfmin/myou/categories/proto/categories"
)

func getCats() (cc proto.CategoriesClient) {
	serv := micro.NewService(
		micro.Name("myou.srv.categories"),
	)
	cc = proto.NewCategoriesClient("myou.srv.categories", serv.Client())
	return
}

func BenchmarkCategoryList(b *testing.B) {
	// fmt.Println(n)

	cats := getCats()

	// run the Fib function b.N times
	for n := 0; n < b.N; n++ {

		res, err := cats.List(context.TODO(), &proto.Filter{Name: "H", Project: &proto.Project{Id: "1"}})
		if err != nil {
			fmt.Println("err: ", err)
			continue
		}
		if len(res.Categories) == 0 {
			fmt.Println("categories not found.")
		}
	}
}

func BenchmarkCategoryPut(b *testing.B) {
	cats := getCats()

	// run the Fib function b.N times
	for n := 0; n < b.N; n++ {

		res, err := cats.Put(context.TODO(), &proto.Input{
			Project:  &proto.Project{Id: "1"},
			Category: &proto.Category{Id: fmt.Sprintf("%d", n+1), Name: fmt.Sprintf("Hello %d", n+1)},
		})

		if err != nil {
			fmt.Println("err: ", err)
		}
		_ = res
	}
}
