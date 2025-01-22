package server

import (
	// "log"
	// "net/http"

	"net/http"
	"time"

	"github.com/gorilla/websocket"
	"github.com/levelstudio/payroll-4ta-crud/pkg/models"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"github.com/levelstudio/payroll-4ta-crud/pkg/db"
	"github.com/levelstudio/payroll-4ta-crud/pkg/graph"
)

func graphqlHandler() gin.HandlerFunc {
	// NewExecutableSchema and Config are in the generated.go file
	// Resolver is in the resolver.go file
	srv := handler.New(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{
		UsersRepo: db.UsersRepo{
			DB:                 db.DB,
			CreateUserObserver: map[string]chan *models.User{},
			UpdateUserObserver: map[string]chan *models.User{},
			DeleteUserObserver: map[string]chan *models.User{},
		},
	}}))

	srv.AddTransport(transport.POST{})
	srv.AddTransport(transport.Websocket{
		KeepAlivePingInterval: 10 * time.Second,
		Upgrader: websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
		},
		/* InitFunc: func(ctx context.Context, initPayload transport.InitPayload) (context.Context, error) {
			return webSocketInit(ctx, initPayload)
		}, */
	})
	var mb int64 = 1 << 20
	srv.AddTransport(transport.MultipartForm{
		MaxMemory:     2 * mb,
		MaxUploadSize: 5 * mb,
	})
	srv.Use(extension.Introspection{})

	return func(c *gin.Context) {
		srv.ServeHTTP(c.Writer, c.Request)
	}
}

func playgroundHandler() gin.HandlerFunc {
	h := playground.Handler("GraphQL", "/query")

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func HttpServer(addr string) error {

	r := gin.Default()
	/* r.Use(cors.New(cors.Config{
			AllowOrigins:     []string{"http://localhost:4200"},
			AllowMethods:     []string{"GET", "POST", "DELETE", "PUT", "PATCH"},
			AllowHeaders:     []string{"Origin"},
			ExposeHeaders:    []string{"Content-Length"},
			AllowCredentials: true,
			AllowOriginFunc: func(origin string) bool {
				// return origin == "http://localhost:4200"
	      return true
			},
		})) */
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"}
	r.Use(cors.New(config))
	r.MaxMultipartMemory = 8 << 20
	r.Static("/files", "./files")
	r.Any("/query", graphqlHandler())
	r.Any("/subs", graphqlHandler())
	r.GET("/", playgroundHandler())
	return r.Run(addr)
}
