package tests

import (
	"fmt"
	"testing"

	bnats "github.com/micro/go-plugins/broker/nats"
	tnats "github.com/micro/go-plugins/transport/nats"

	"golang.org/x/net/context"

	"github.com/micro/go-micro"
	proto "github.com/sunfmin/myou/categories/proto/categories"
)

func BenchmarkCategoryList(b *testing.B) {
	// fmt.Println(n)
	service := micro.NewService(
		micro.Name("myou.srv.categories"),
		micro.Transport(tnats.NewTransport()),
		// micro.Registry(rnats.NewRegistry()),
		micro.Broker(bnats.NewBroker()),
	)
	cats := proto.NewCategoriesClient("myou.srv.categories", service.Client())

	// run the Fib function b.N times
	for n := 0; n < b.N; n++ {

		res, err := cats.List(context.TODO(), &proto.Filter{Name: "Felix"})
		if err != nil {
			fmt.Println("err: ", err)
		}
		_ = res
	}
}
