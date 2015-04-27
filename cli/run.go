package cli

import (
	"net/http"

	"github.com/codegangsta/cli"
	"github.com/go-soa/charon/controller/web"
	"github.com/go-soa/charon/service"
	"github.com/julienschmidt/httprouter"
)

var (
	runCommand = cli.Command{
		Name:   "run",
		Usage:  "starts server",
		Action: runCommandAction,
	}
)

func runCommandAction(context *cli.Context) {
	service.InitConfig(context.GlobalString("environment"))
	service.InitLogger(service.Config.Logger)
	service.InitDB(service.Config.DB)
	service.InitRepositoryManager(service.DBPool)
	service.InitPasswordHasher(service.Config.PasswordHasher)
	service.InitTemplates(service.Config.Templates)
	service.InitMailer(service.Config.Mailer, service.MailTemplates)

	router := httprouter.New()

	setupStaticRoutes(router)
	setupWebRoutes(router)

	listenOn := service.Config.Server.Host + ":" + service.Config.Server.Port
	service.Logger.Fatal(http.ListenAndServe(listenOn, router))
}

func setupWebRoutes(router *httprouter.Router) {
	container := web.ServiceContainer{
		Logger:         service.Logger,
		DB:             service.DBPool,
		RM:             service.RepositoryManager,
		PasswordHasher: service.PasswordHasher,
		Mailer:         service.Mail,
		Templates:      service.Templates,
	}

	registrationIndex := web.NewHandler(
		"registration_index",
		web.NewMiddlewares(
			(*web.Handler).RegistrationSuccess,
		),
		container,
	)

	registrationSuccess := web.NewHandler(
		"registration_success",
		web.NewMiddlewares(
			(*web.Handler).RegistrationSuccess,
		),
		container,
	)

	registrationCreate := web.NewHandler(
		"registration_index",
		web.NewMiddlewares(
			(*web.Handler).RegistrationCreate,
		),
		container,
	)

	router.Handler("GET", "/registration", registrationIndex)
	router.Handler("GET", "/registration/success", registrationSuccess)
	router.Handler("POST", "/registration", registrationCreate)
}

func setupStaticRoutes(router *httprouter.Router) {
	router.ServeFiles("/assets/*filepath", http.Dir("assets"))
}