package router

var (
	routers map[string]Router
)

type Router interface {
	GetUir() string
	GetController() interface{}
	GetNeedAuth() bool
}

func Routers() map[string]Router {
	return routers
}

func register(router Router) {
	routers[router.GetUir()] = router
}

func unRegister(party string) {
	delete(routers, party)
}

func init(){
	routers = make(map[string]Router)
}
