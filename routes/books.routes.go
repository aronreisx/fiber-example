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

	router.Put("/:id", func(c *fiber.Ctx) error {
		id, err := c.ParamsInt("id")

		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Request error",
			})
		}

		type Request struct {
			Title string
		}

		var body Request

		c.BodyParser(&body)

		for _, book := range books {
			if book.Id == id {
				book.Title = body.Title
			}
		}

		return c.JSON(fiber.Map{
			"books": books,
		})
	})

	router.Delete("/:id", func(c *fiber.Ctx) error {
		id, err := c.ParamsInt("id")

		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Request error",
			})
		}

		for index, book := range books {
			if book.Id == id {
				books = append(books[0:index], books[index + 1:]...)
			}
		}

		return c.JSON(fiber.Map{
			"books": books,
		})
	})
}