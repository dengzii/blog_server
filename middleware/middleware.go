package middleware

import "server/bootstrap"

type Middleware interface {
	Attach(bootstrap *bootstrap.Bootstrapper)
}

func WithBootstrap(bootstrap *bootstrap.Bootstrapper) {

}
