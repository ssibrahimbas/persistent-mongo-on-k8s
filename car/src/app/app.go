package app

import (
	"log"

	cnf "github.com/ssibrahimbas/ssi-core/pkg/config"
	"github.com/ssibrahimbas/ssi-core/pkg/db"
	"github.com/ssibrahimbas/ssi-core/pkg/helper"
	"github.com/ssibrahimbas/ssi-core/pkg/http"
	"github.com/ssibrahimbas/ssi-core/pkg/i18n"
	"github.com/ssibrahimbas/ssi-core/pkg/validator"
	"ssibrahimbas.com/persistent-mongo/car/src/config"
	"ssibrahimbas.com/persistent-mongo/car/src/internal"
)

type App struct {
	Cnf  *config.App
	Db   *db.MongoDB
	I18n *i18n.I18n
	Http *http.Client
	Val  *validator.Validator
	Repo *internal.Repo
	Srv  *internal.Service
	Hnd  *internal.Handler
}

func New() *App {
	return &App{}
}

func (a *App) Init() *App {
	a.loadConfig()
	a.loadDb()
	a.loadI18n()
	a.loadHttp()
	a.loadValidator()
	a.loadInternal()
	return a
}

func (a *App) Serve() {
	log.Fatal(a.Http.Listen(a.Cnf.Host + a.Cnf.Port))
}

func (a *App) loadDb() {
	uri := db.CalcMongoUri(db.UriParams{
		Host: a.Cnf.Db.Host,
		Port: a.Cnf.Db.Port,
		User: a.Cnf.Db.Usr,
		Pass: a.Cnf.Db.Pw,
		Db:   a.Cnf.Db.Name,
	})
	d, err := db.NewMongo(uri, a.Cnf.Db.Name)
	helper.CheckErr(err)
	a.Db = d
}

func (a *App) loadConfig() {
	cnfg := config.App{}
	cnf.LoadConfig(&cnfg)
	a.Cnf = &cnfg
}

func (a *App) loadI18n() {
	i := i18n.New(a.Cnf.I18n.Fallback)
	i.LoadLanguages(a.Cnf.I18n.LocaleDir, a.Cnf.I18n.Locales...)
	a.I18n = i
}

func (a *App) loadHttp() {
	a.Http = http.New(a.I18n)
}

func (a *App) loadValidator() {
	val := validator.New(a.I18n)
	val.ConnectCustom()
	val.RegisterTagName()
	a.Val = val
}

func (a *App) loadInternal() {
	a.Repo = internal.NewRepo(&internal.RepoConfig{
		MongoDB: a.Db,
	})
	a.Srv = internal.NewService(&internal.ServiceConfig{
		Repo: a.Repo,
		I18n: a.I18n,
	})
	a.Hnd = internal.NewHandler(&internal.HandlerConfig{
		Service:    a.Srv,
		Config:     a.Cnf,
		I18n:       a.I18n,
		Validator:  a.Val,
		HttpClient: a.Http,
	})
	a.Hnd.InitAllVersions()
}
