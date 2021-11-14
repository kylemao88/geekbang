// +build wireinject

package fourmgr

import "github.com/google/wire"

func InitializeUser() user {
	wire.Build(NewUser, NewHead, NewNick)
	return user{}
}
