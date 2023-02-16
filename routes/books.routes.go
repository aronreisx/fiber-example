package routes
import "github.com/gofiber/fiber/v2"

type Book struct {
	Title string `json:"title"`
	Id int `json:"id"`
}

func UseBooksRoute(router fiber.Router) {
	books := []*Book{
		{
			Title: "Clean code",
			Id: 1,
		},
		{
			Title: "Clean architecture",
			Id: 2,
		},
	}

	router.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"books": books,
		})
	})

	router.Get("/:id", func(c *fiber.Ctx) error {
		id, _ := c.ParamsInt("id")
		var bookFound Book

		for _, book := range books {
			if book.Id == id {
				bookFound = *book
			}
		}

		return c.JSON(fiber.Map{
			"book": bookFound,
		})
	})

	router.Post("/", func(c *fiber.Ctx) error {
		type Request struct {
			Title string
			Id int
		}

		var body Request

		c.BodyParser(&body)

		newBook := Book {
			Title: body.Title,
			Id: len(books) + 1,
		}

		books = append(books, &newBook)

		return c.JSON(fiber.Map{
			"books" books,
		})
	})
}