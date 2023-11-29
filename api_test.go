package main

import (
	"encoding/json"
	"io"
	"net/http"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/samvimes01/go-rest1/db"
	"github.com/samvimes01/go-rest1/models"
	"github.com/samvimes01/go-rest1/utils"
	"github.com/steinfletcher/apitest"
)

func newApp() *fiber.App {
	var app *fiber.App = NewFiberApp()

	db.InitDatabase(utils.GetValue("DB_TEST_HOST"), utils.GetValue("DB_TEST_NAME"))

	return app
}

func getItem() models.Item {
	db.InitDatabase(utils.GetValue("DB_TEST_HOST"), utils.GetValue("DB_TEST_NAME"))
	item, err := db.SeedItem()
	if err != nil {
		panic(err)
	}

	return item
}

func cleanup(res *http.Response, req *http.Request, apiTest *apitest.APITest) {
	if http.StatusOK == res.StatusCode || http.StatusCreated == res.StatusCode {
		db.CleanSeeders()
	}
}

func FiberToHandlerFunc(app *fiber.App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		resp, err := app.Test(r)
		if err != nil {
			panic(err)
		}

		for k, vv := range resp.Header {
			for _, v := range vv {
				w.Header().Add(k, v)
			}
		}
		w.WriteHeader(resp.StatusCode)

		if _, err := io.Copy(w, resp.Body); err != nil {
			panic(err)
		}
	}
}

func getJWTToken(t *testing.T) string {
	db.InitDatabase(utils.GetValue("DB_TEST_HOST"), utils.GetValue("DB_TEST_NAME"))
	user, err := db.SeedUser()
	if err != nil {
		panic(err)
	}

	var userRequest *models.UserRequest = &models.UserRequest{
		Email:    user.Email,
		Password: user.Password,
	}

	var resp *http.Response = apitest.New().
		HandlerFunc(FiberToHandlerFunc(newApp())).
		Post("/api/v1/login").
		JSON(userRequest).
		Expect(t).
		Status(http.StatusOK).
		End().Response

	var response *models.Response[string] = &models.Response[string]{}

	json.NewDecoder(resp.Body).Decode(&response)

	var token string = response.Data

	var JWT_TOKEN = "Bearer " + token

	return JWT_TOKEN
}

func TestSignup_Success(t *testing.T) {
	userData, err := utils.CreateFaker[models.User]()

	if err != nil {
		panic(err)
	}

	var userRequest *models.UserRequest = &models.UserRequest{
		Email:    userData.Email,
		Password: userData.Password,
	}

	apitest.New().
		Observe(cleanup).
		HandlerFunc(FiberToHandlerFunc(newApp())).
		Post("/api/v1/signup").
		JSON(userRequest).
		Expect(t).
		Status(http.StatusOK).
		End()
}

func TestSignup_ValidationFailed(t *testing.T) {
	var userRequest *models.UserRequest = &models.UserRequest{
		Email:    "",
		Password: "",
	}

	apitest.New().
	  Observe(cleanup).
		HandlerFunc(FiberToHandlerFunc(newApp())).
		Post("/api/v1/signup").
		JSON(userRequest).
		Expect(t).
		Status(http.StatusBadRequest).
		End()
}

func TestLogin_Success(t *testing.T) {
	db.InitDatabase(utils.GetValue("DB_TEST_HOST"), utils.GetValue("DB_TEST_NAME"))
	user, err := db.SeedUser()
	if err != nil {
		panic(err)
	}

	var userRequest *models.UserRequest = &models.UserRequest{
		Email:    user.Email,
		Password: user.Password,
	}

	apitest.New().
		Observe(cleanup).
		HandlerFunc(FiberToHandlerFunc(newApp())).
		Post("/api/v1/login").
		JSON(userRequest).
		Expect(t).
		Status(http.StatusOK).
		End()
}

func TestLogin_ValidationFailed(t *testing.T) {
	var userRequest *models.UserRequest = &models.UserRequest{
		Email:    "",
		Password: "",
	}

	apitest.New().
	  Observe(cleanup).
		HandlerFunc(FiberToHandlerFunc(newApp())).
		Post("/api/v1/login").
		JSON(userRequest).
		Expect(t).
		Status(http.StatusBadRequest).
		End()
}

func TestLogin_Failed(t *testing.T) {
	var userRequest *models.UserRequest = &models.UserRequest{
		Email:    "notfound@mail.com",
		Password: "123123",
	}

	apitest.New().
	  Observe(cleanup).
		HandlerFunc(FiberToHandlerFunc(newApp())).
		Post("/api/v1/login").
		JSON(userRequest).
		Expect(t).
		Status(http.StatusInternalServerError).
		End()
}

func TestGetItems_Success(t *testing.T) {
	apitest.New().
	  Observe(cleanup).
		HandlerFunc(FiberToHandlerFunc(newApp())).
		Get("/api/v1/items").
		Expect(t).
		Status(http.StatusOK).
		End()
}

func TestGetItem_Success(t *testing.T) {
	var item models.Item = getItem()

	apitest.New().
		Observe(cleanup).
		HandlerFunc(FiberToHandlerFunc(newApp())).
		Get("/api/v1/items/" + item.ID).
		Expect(t).
		Status(http.StatusOK).
		End()
}

func TestGetItem_NotFound(t *testing.T) {
	apitest.New().
	  Observe(cleanup).
		HandlerFunc(FiberToHandlerFunc(newApp())).
		Get("/api/v1/items/0").
		Expect(t).
		Status(http.StatusNotFound).
		End()
}

func TestCreateItem_Success(t *testing.T) {
	itemData, err := utils.CreateFaker[models.Item]()
	if err != nil {
		panic(err)
	}

	var itemRequest *models.ItemRequest = &models.ItemRequest{
		Name:     itemData.Name,
		Price:    itemData.Price,
		Quantity: itemData.Quantity,
	}

	var token string = getJWTToken(t)

	apitest.New().
		Observe(cleanup).
		HandlerFunc(FiberToHandlerFunc(newApp())).
		Post("/api/v1/items").
		Header("Authorization", token).
		JSON(itemRequest).
		Expect(t).
		Status(http.StatusCreated).
		End()
}

func TestCreateItem_ValidationFailed(t *testing.T) {
	var itemRequest *models.ItemRequest = &models.ItemRequest{
		Name:     "",
		Price:    0,
		Quantity: 0,
	}

	var token string = getJWTToken(t)

	apitest.New().
	  Observe(cleanup).
		HandlerFunc(FiberToHandlerFunc(newApp())).
		Post("/api/v1/items").
		Header("Authorization", token).
		JSON(itemRequest).
		Expect(t).
		Status(http.StatusBadRequest).
		End()
}

func TestUpdateItem_Success(t *testing.T) {
	var item models.Item = getItem()

	var itemRequest *models.ItemRequest = &models.ItemRequest{
		Name:     item.Name,
		Price:    item.Price,
		Quantity: item.Quantity,
	}

	var token string = getJWTToken(t)

	apitest.New().
		Observe(cleanup).
		HandlerFunc(FiberToHandlerFunc(newApp())).
		Put("/api/v1/items/"+item.ID).
		Header("Authorization", token).
		JSON(itemRequest).
		Expect(t).
		Status(http.StatusOK).
		End()
}

func TestUpdateItem_ValidationFailed(t *testing.T) {
	var item models.Item = getItem()

	var itemRequest *models.ItemRequest = &models.ItemRequest{
		Name:     "",
		Price:    0,
		Quantity: 0,
	}

	var token string = getJWTToken(t)

	apitest.New().
	  Observe(cleanup).
		HandlerFunc(FiberToHandlerFunc(newApp())).
		Put("/api/v1/items/"+item.ID).
		Header("Authorization", token).
		JSON(itemRequest).
		Expect(t).
		Status(http.StatusBadRequest).
		End()
}

func TestUpdateItem_Failed(t *testing.T) {
	var itemRequest *models.ItemRequest = &models.ItemRequest{
		Name:     "changed",
		Price:    10,
		Quantity: 10,
	}

	var token string = getJWTToken(t)

	apitest.New().
		Observe(cleanup).
		HandlerFunc(FiberToHandlerFunc(newApp())).
		Put("/api/v1/items/0").
		Header("Authorization", token).
		JSON(itemRequest).
		Expect(t).
		Status(http.StatusNotFound).
		End()
}

func TestDeleteItem_Success(t *testing.T) {
	var item models.Item = getItem()

	var token string = getJWTToken(t)

	apitest.New().
	  Observe(cleanup).
		HandlerFunc(FiberToHandlerFunc(newApp())).
		Delete("/api/v1/items/"+item.ID).
		Header("Authorization", token).
		Expect(t).
		Status(http.StatusOK).
		End()
}

func TestDeleteItem_Failed(t *testing.T) {
	var token string = getJWTToken(t)

	apitest.New().
	  Observe(cleanup).
		HandlerFunc(FiberToHandlerFunc(newApp())).
		Delete("/api/v1/items/0").
		Header("Authorization", token).
		Expect(t).
		Status(http.StatusNotFound).
		End()
}
