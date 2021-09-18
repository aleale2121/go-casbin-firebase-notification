package initiator

import (
	"log"
	"os"
	routing "template/internal/adapter/glue/routing"
	compHandler "template/internal/adapter/http/rest/server/company"
	usrHandler "template/internal/adapter/http/rest/server/user"
	"template/internal/adapter/repository"
	"template/internal/adapter/storage/persistence/company"
	"template/internal/adapter/storage/persistence/user"
	"template/internal/constant/model"
	compUsecase "template/internal/module/company"
	usrUsecase "template/internal/module/user"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	en_translations "github.com/go-playground/validator/v10/translations/en"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// global validator instance
var (
	uni      *ut.UniversalTranslator
	validate *validator.Validate
	trans    ut.Translator
)

func Initialize() {

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		en := en.New()
		uni := ut.New(en, en)
		// this is usually know or extracted from http 'Accept-Language' header
		// also see uni.FindTranslator(...)
		trans, _ = uni.GetTranslator("en")
		en_translations.RegisterDefaultTranslations(v, trans)
	}

	DATABASE_URL := "postgres://postgres:admin@localhost:5432/demo?sslmode=disable"

	conn, err := gorm.Open(postgres.Open(DATABASE_URL), &gorm.Config{
		SkipDefaultTransaction: true,
	})

	// conn.AutoMigrate migrates gorm models
	conn.AutoMigrate(&model.User{}, &model.Role{}, &model.UserCompanyRole{}, &model.PushedNotification{}, &model.Notification{}, &model.Company{})

	if err != nil {
		log.Printf("Error when Opening database connection: %v", err)
		os.Exit(1)
	}

	usrPersistence := user.UserInit(conn)
	compPersistence := company.CompanyInit(conn)

	usrRepo := repository.UserInit()
	usrUsecase := usrUsecase.Initialize(usrRepo, usrPersistence)
	usrHandler := usrHandler.UserInit(usrUsecase, trans)

	compUsecase := compUsecase.Initialize(compPersistence)
	compHandler := compHandler.CompanyInit(compUsecase, trans)

	router := gin.Default()

	//  group: v1
	v1 := router.Group("/v1")
	routing.UserRoutes(v1, usrHandler)
	routing.CompanyRoutes(v1, compHandler)
	router.Run()
	log.Println(usrHandler)
}
