package initiator

import (
	"log"
	"os"
	routing "template/internal/adapter/glue/routing"
	usrHandler "template/internal/adapter/http/rest/server/user"
	rlHandler "template/internal/adapter/http/rest/server/role"
	"template/internal/adapter/repository"
	"template/internal/adapter/storage/persistence/user"
	"template/internal/adapter/storage/persistence/role"
	"template/internal/constant/model"
	usrUsecase "template/internal/module/user"
	roleUsecase "template/internal/module/role"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	en_translations "github.com/go-playground/validator/v10/translations/en"
	// "github.com/casbin/casbin/v2"
	// gormadapter "github.com/casbin/gorm-adapter/v3"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// global validator instance
var (
	// uni      *ut.UniversalTranslator
	// validate *validator.Validate
	trans    ut.Translator
)

func Initialize() {

	// en := en.New()
	// uni = ut.New(en, en)
	// trans, _ := uni.GetTranslator("en")
	// validate = validator.New()
	// en_translations.RegisterDefaultTranslations(validate, trans)
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

	if err != nil {
		log.Printf("Error when Opening database connection: %v", err)
		os.Exit(1)
	}

	// conn.AutoMigrate migrates gorm models
	conn.AutoMigrate(&model.Role{},&model.User{}, &model.UserCompanyRole{}, &model.PushedNotification{}, &model.Notification{}, &model.Company{})

	// a, _ := gormadapter.NewAdapterByDBWithCustomTable(conn, &model.CasbinRule{})
	// e, _ := casbin.NewEnforcer("../rbac_model.conf", a)

	// Load the policy from DB.
	// e.LoadPolicy()

	// Check the permission.
	// e.Enforce("alice", "data1", "read")

	// Modify the policy.
	// e.AddPolicy(...)
	// e.RemovePolicy(...)

	// Save the policy back to DB.
	// e.SavePolicy()
	usrPersistence := user.UserInit(conn)
	rolePersistent:=persistence.RoleInit(conn)
    roleUsecase:=roleUsecase.RoleInitialize(rolePersistent)
    roleHandler:=rlHandler.NewRoleHandler(roleUsecase,trans)

	usrRepo := repository.UserInit()
	usrUsecase := usrUsecase.Initialize(usrRepo, usrPersistence)
	usrHandler := usrHandler.UserInit(usrUsecase, trans)

	router := gin.Default()

	//  group: v1
	v1 := router.Group("/v1")
	routing.UserRoutes(v1, usrHandler)

	roleGroup:=v1.Group("/roles")
	routing.RoleRoutes(roleGroup,roleHandler)

	router.Run()
	log.Println(usrHandler)
}
