package controllers

import (
	"go-auth/database"
	"go-auth/helpers"
	"go-auth/models"
	"log"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)


func Login(c *fiber.Ctx) error {
	// data validation
	reqData := new(ReqLogin)
	if err := c.BodyParser(&reqData); err != nil {
		return err
	}
	errors := helpers.ValidateStruct(*reqData)
	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errors)
	}

	var user = models.User{}

	// checking if user exists
	if err := database.
		DB.Where("email = ?", reqData.Email).First(&user).Error; 
		err != nil {
			c.Status(fiber.StatusNotFound)
			return c.JSON(fiber.Map{
				"message": "user not found",
			})
	}
	if user.Id == 0 {
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"message": "user not found",
		})
	}

	// checking if password matches user
	if err := bcrypt.CompareHashAndPassword(user.Password, []byte(reqData.Password)); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "inccorect password",
		})
	}

	// generating jwt token
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer: strconv.Itoa(int(user.Id)),
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
	})
	token, err := claims.SignedString([]byte(SecretKey))
	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": err,
		})
	}

	// saving jwt to cookie
	cookie := fiber.Cookie{
		Name: "jwt",
		Value: token,
		Expires: time.Now().Add(time.Hour * 24),
		HTTPOnly: true,
	}

	c.Cookie(&cookie)

	log.Println("Successfully logged user in")
	return c.JSON(fiber.Map{
		"message": "Successfully logged in",
	})
}