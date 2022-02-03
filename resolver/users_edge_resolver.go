package resolver

import (
	"github.com/graph-gophers/graphql-go"
	"github.com/javacode123/go-graphql-starter/model"
)

type usersEdgeResolver struct {
	cursor graphql.ID
	model  *model.User
}

func (r *usersEdgeResolver) Cursor() graphql.ID {
	return r.cursor
}

func (r *usersEdgeResolver) Node() *userResolver {
	return &userResolver{u: r.model}
}
