package main

import (
	"cashmachine/pkg/driver"
	"cashmachine/pkg/usecase"
	"cashmachine/pkg/utils"
	"fmt"
	"os"
	"strconv"

	"github.com/gofiber/fiber/v2"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.SetFormatter(&log.TextFormatter{
		FullTimestamp: true,
	})

	log.Info(">>>>>>>> Server Start, connecting to database <<<<<<<<")
	repo, err := utils.BuildRepository()
	if err != nil {
		log.Fatal(err)
	}

	app := fiber.New()
	app.Post("/new/", func(c *fiber.Ctx) error {
		request := new(usecase.RequestCreate)

		if err := c.BodyParser(request); err != nil {
			c.SendStatus(400)
			c.JSON(usecase.ResponseGeneral{Msg: err.Error()})
			log.WithField("error", err).Info("Invalid User input")

		}

		if request.Value <= 0 {
			c.SendStatus(400)
			c.JSON(usecase.ResponseGeneral{Msg: "Invalid deposit value"})
			log.WithField("error", request).Info("Invalid User input")
			return nil
		}

		response, err := usecase.CreateUsecase(*request, repo)
		if err != nil {
			c.SendStatus(503)
			c.JSON(usecase.ResponseGeneral{Msg: "Failed to create account"})
			log.WithField("error", err).Warn("Failed to create account in database")
			return nil
		}

		c.SendStatus(200)
		c.JSON(response)
		return nil
	})

	app.Get("/balance/:id", func(c *fiber.Ctx) error {
		id, err := strconv.Atoi(c.Params("id"))
		if err != nil {
			c.SendStatus(400)
			c.JSON(usecase.ResponseGeneral{Msg: "Invalid account id format"})
			return nil
		}

		request := usecase.RequestBalance{AccID: id}
		response, err := usecase.BalanceUsecase(request, repo)
		switch {
		case err == driver.ErrInvalidAccount:
			c.SendStatus(404)
			c.JSON(usecase.ResponseGeneral{Msg: "Invalid account id"})
			return nil
		case err != nil:
			c.SendStatus(503)
			c.JSON(usecase.ResponseGeneral{Msg: "Failed to fetch account"})
			log.WithField("error", err).Warn("Failed to fetch account balance in database")
			return nil
		}

		c.SendStatus(200)
		c.JSON(response)
		return nil
	})

	app.Put("/deposit/:id", func(c *fiber.Ctx) error {
		request := new(usecase.RequestDeposit)

		if err := c.BodyParser(request); err != nil {
			c.SendStatus(400)
			c.JSON(usecase.ResponseGeneral{Msg: "Invalid value format"})
			log.WithField("error", err).Info("Invalid User input")
			return nil
		}

		id, err := strconv.Atoi(c.Params("id"))
		if err != nil {
			c.SendStatus(400)
			c.JSON(usecase.ResponseGeneral{Msg: "Invalid account id format"})
			return nil
		}

		request.AccID = id

		response, err := usecase.DepositUsecase(*request, repo)
		switch {
		case err == driver.ErrInvalidAccount:
			c.SendStatus(404)
			c.JSON(usecase.ResponseGeneral{Msg: "Invalid account id"})
			return nil
		case err != nil:
			c.SendStatus(503)
			c.JSON(usecase.ResponseGeneral{Msg: "Failed to deposit value"})
			log.WithField("error", err).Warn("Failed in deposit process")
			return nil
		}

		c.SendStatus(200)
		c.JSON(response)
		return nil
	})

	app.Put("/withdraw/:id", func(c *fiber.Ctx) error {
		request := new(usecase.RequestWithdraw)

		if err := c.BodyParser(request); err != nil {
			c.SendStatus(400)
			c.JSON(usecase.ResponseGeneral{Msg: "Invalid value format"})
			log.WithField("error", err).Info("Invalid User input")
			return nil
		}

		id, err := strconv.Atoi(c.Params("id"))
		if err != nil {
			c.SendStatus(400)
			c.JSON(usecase.ResponseGeneral{Msg: "Invalid account id format"})
			return nil
		}

		request.AccID = id

		response, err := usecase.WithdrawUsecase(*request, repo)
		switch {
		case err == driver.ErrInvalidAccount:
			c.SendStatus(404)
			c.JSON(usecase.ResponseGeneral{Msg: "Invalid account id"})
			return nil
		case err != nil:
			c.SendStatus(503)
			c.JSON(usecase.ResponseGeneral{Msg: "Failed to deposit value"})
			log.WithField("error", err).Warn("Failed in deposit process")
			return nil
		}

		c.SendStatus(200)
		c.JSON(response)
		return nil
	})

	app.Listen(fmt.Sprintf(":%s", os.Getenv("PORT")))
}
