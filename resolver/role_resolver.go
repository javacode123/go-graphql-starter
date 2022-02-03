package resolver

import (
	"github.com/graph-gophers/graphql-go"
	"github.com/javacode123/go-graphql-starter/model"
)

type roleResolver struct {
	role *model.Role
}

func (r *roleResolver) ID() graphql.ID {
	return graphql.ID(r.role.ID)
}

func (r *roleResolver) Name() *string {
	return &r.role.Name
}
