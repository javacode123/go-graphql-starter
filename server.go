package main

import (
	gcontext "github.com/javacode123/go-graphql-starter/context"
	h "github.com/javacode123/go-graphql-starter/handler"
	"github.com/javacode123/go-graphql-starter/resolver"
	"github.com/javacode123/go-graphql-starter/schema"
	"github.com/javacode123/go-graphql-starter/service"
	"net/http"

	graphql "github.com/graph-gophers/graphql-go"
	"github.com/javacode123/go-graphql-starter/loader"
	"golang.org/x/net/context"
)

func main() {
	config := gcontext.LoadConfig(".")

	//db, err := gcontext.OpenDB(config)
	//if err != nil {
	//	log.Fatalf("Unable to connect to db: %s \n", err)
	//}
	ctx := context.Background()
	log := service.NewLogger(config)
	roleService := service.NewRoleService(nil, log)
	userService := service.NewUserService(nil, roleService, log)
	authService := service.NewAuthService(config, log)

	ctx = context.WithValue(ctx, "config", config)
	ctx = context.WithValue(ctx, "log", log)
	ctx = context.WithValue(ctx, "roleService", roleService)
	ctx = context.WithValue(ctx, "userService", userService)
	ctx = context.WithValue(ctx, "authService", authService)

	graphqlSchema := graphql.MustParseSchema(schema.GetRootSchema(), &resolver.Resolver{})

	http.Handle("/login", h.AddContext(ctx, h.Login()))

	loggerHandler := &h.LoggerHandler{config.DebugMode}
	http.Handle("/query", h.AddContext(ctx, loggerHandler.Logging(&h.GraphQL{Schema: graphqlSchema, Loaders: loader.NewLoaderCollection()})))

	http.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "graphiql.html")
	}))

	log.Fatal(http.ListenAndServe(":3000", nil))
}
