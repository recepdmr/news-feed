package post

import (
	"context"
	"os"

	"github.com/gofiber/fiber/v2"
	post "github.com/recepdmr/apigateways/public/services/post/server"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type handler struct {
}

var postServiceUri = os.Getenv("POST_SERVICE_URI")
var postService *post.PostServiceClient

func (h *handler) GetAll(c *fiber.Ctx) error {
	var result []*post.Post = make([]*post.Post, 0)
	bookList, err := (*postService).Get(context.Background(), &post.GetPostsRequest{})
	if err == nil {
		result = bookList.Data
	}

	c.Status(fiber.StatusOK).JSON(result)
	return nil
}

func (*handler) Create(c *fiber.Ctx) error {
	var req post.InsertPostRequest
	if err := c.BodyParser(&req); err == nil {
		_, err := (*postService).Insert(context.Background(), &req)
		if err == nil {
			return c.Status(fiber.StatusCreated).JSON(nil)
		}
	}
	return c.Status(fiber.StatusBadRequest).JSON(nil)
}

func RegisterPostHandler(app *fiber.App) {
	basePath := "/posts/"
	handler := handler{}
	app.Get(basePath, handler.GetAll)
	app.Post(basePath, handler.Create)
}

func init() {
	configureClient()
}

func configureClient() {
	conn, err := grpc.Dial(postServiceUri, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		panic(err)
	}
	client := post.NewPostServiceClient(conn)
	postService = &client
}
