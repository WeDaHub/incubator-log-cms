package context

import "github.com/nats-io/jwt"

type Claims struct {
	jwt.AccountClaims
	UserId	uint
}

type Context struct {

}
