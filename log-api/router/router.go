package router

var (
	routers map[string]Router
)

type Router interface {
	GetName() string
	GetController() interface{}
}

func Routers() map[string]Router {
	return routers
}

func register(router Router) {
	routers[router.GetName()] = router
}

func unRegister(party string) {
	delete(routers, party)
}

func init(){
	routers = make(map[string]Router)
	register(NewHelloworld())
}
