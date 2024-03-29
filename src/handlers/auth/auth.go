package auth

import (
	"time"

	"github.com/tmunongo/linkkeep/api/src/models"
	"golang.org/x/crypto/bcrypt"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type JwtCustomClaims struct {
	Name  string `json:"name"`
	ID    uint    `json:"id"`
	Admin bool   `json:"admin"`
	jwt.RegisteredClaims
}

// Valid implements jwt.Claims.
func (j *JwtCustomClaims) Valid() error {
	panic("unimplemented")
}

func Login(c echo.Context) (err error) {
	username := c.FormValue("username")
	password := c.FormValue("password")

	// get user from db by username, if not found return error
	user := &models.User{
		Username: username,
	}

	u, err := user.GetUser()

	if err != nil {
		return c.JSON(500, map[string]string{
            "error": err.Error(),
        })
	}

	// check password
	
	err = bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	if err != nil {
		return c.JSON(401, map[string]string{
			"error": "Invalid credentials",
		})
	}

	claims := &JwtCustomClaims{
		u.Username,
		u.ID,
		true,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 72)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return c.JSON(500, map[string]string{
            "error": err.Error(),
        })
	}

	return c.JSON(200, map[string]string{
		"token": t,
	})
}


func Register(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 4)

	if err != nil {
		return err
	}

	user := &models.User{
		Username: username,
		Password: string(hashedPassword),
	}

	u, err := user.CreateUser()

	if err != nil {
		return err
	}

	claims := &JwtCustomClaims{
		u.Username,
		u.ID,
		false,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 72)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return err
	}

	return c.JSON(200, map[string]string{
		"token": t,
	})
}