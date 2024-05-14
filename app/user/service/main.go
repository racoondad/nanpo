package main

import (
	"github.com/racoondad/nanpo"
)

func main() {
	app := nanpo.NewApplication()
	installConfig(app)
	installDatabase(app)
	installRedis(app)
	installMiddleware(app)
	installEventBus(app)
	installRouter(app)
	installWorkPool(app)
	installCronPool(app)
	installServer(app)
	app.Run()
}

func installConfig(app nanpo.Application) {
}

func installDatabase(app nanpo.Application) {
}

func installRedis(app nanpo.Application) {
}

func installMiddleware(app nanpo.Application) {
}

func installEventBus(app nanpo.Application) {
}

func installRouter(app nanpo.Application) {
}

func installWorkPool(app nanpo.Application) {
}

func installCronPool(app nanpo.Application) {
}

func installServer(app nanpo.Application) {
}
