package middleware

import (
	"github.com/dgrijalva/jwt-go"
	jailer "github.com/iris-contrib/middleware/jwt"
	"server/bootstrap"
)

type Authentication struct {
}

func (s Authentication) Attach(bootstrap *bootstrap.Bootstrapper) {

}

func Auth() *jailer.Middleware {

	return jailer.New(jailer.Config{
		ValidationKeyGetter: func(token *jwt.Token) (i interface{}, e error) {
			return []byte("secret"), nil
		},
		SigningMethod: jwt.SigningMethodHS256,
	})
}
