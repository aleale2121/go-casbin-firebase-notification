package initiator

import (
	"log"
	"net/http"
	"os"
	routing "template/internal/adapter/glue/routing"
	authHandler "template/internal/adapter/http/rest/server/auth"
	compHandler "template/internal/adapter/http/rest/server/company"
	permHandler "template/internal/adapter/http/rest/server/permission"
	rlHandler "template/internal/adapter/http/rest/server/role"
	usrHandler "template/internal/adapter/http/rest/server/user"
	"template/internal/adapter/repository"
	"template/internal/adapter/storage/persistence/company"
	"template/internal/adapter/storage/persistence/role"
	"template/internal/adapter/storage/persistence/user"
	"template/internal/constant/model"
	authUsecase "template/internal/module/auth"
	compUsecase "template/internal/module/company"
	roleUsecase "template/internal/module/role"
	usrUsecase "template/internal/module/user"
    casAuth "template/platform/casbin"
	// "github.com/casbin/casbin/v2"
	// gormadapter "github.com/casbin/gorm-adapter/v3"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"

	"github.com/gin-gonic/gin"
	en_translations "github.com/go-playground/validator/v10/translations/en"

	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// global validator instance
var (
	// uni      *ut.UniversalTranslator
	validate *validator.Validate
	trans    ut.Translator
)

func Initialize() {

	// if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
	en := en.New()
	uni := ut.New(en, en)
	trans, _ = uni.GetTranslator("en")
	validate = validator.New()
	en_translations.RegisterDefaultTranslations(validate, trans)
	// }

	DATABASE_URL := "postgres://postgres:admin@localhost:5432/demo?sslmode=disable"

	conn, err := gorm.Open(postgres.Open(DATABASE_URL), &gorm.Config{
		SkipDefaultTransaction: true,
	})

	if err != nil {
		log.Printf("Error when Opening database connection: %v", err)
		os.Exit(1)
	}

	// conn.AutoMigrate migrates gorm models

	conn.AutoMigrate(&model.Role{}, &model.User{}, &model.UserCompanyRole{},  &model.Company{})

	a, _ := gormadapter.NewAdapterByDBWithCustomTable(conn, &model.CasbinRule{})
	e ,err:= casbin.NewEnforcer("../../rbac_model.conf", a)
	if err != nil {
		panic(err)
	}

	usrPersistence := user.UserInit(conn)
	compPersistence := company.CompanyInit(conn)
	rolePersistent := role.RoleInit(conn)

	roleUsecase := roleUsecase.RoleInitialize(rolePersistent)
	roleHandler := rlHandler.NewRoleHandler(roleUsecase, trans)

	usrRepo := repository.UserInit()
	usrUsecase := usrUsecase.Initialize(usrRepo, usrPersistence, validate, trans)
	usrHandler := usrHandler.UserInit(usrUsecase, trans)

	jwtManager := authUsecase.NewJWTManager("secret")
	authUsecases := authUsecase.Initialize(usrPersistence,*jwtManager)
	authHandlers := authHandler.NewAuthHandler(authUsecases)

	compUsecase := compUsecase.Initialize(compPersistence, validate, trans)
	compHandler := compHandler.CompanyInit(compUsecase, trans)

	casbinAuth:=casAuth.NewCasbin(e)
	permHandler := permHandler.PermissionInit(casbinAuth, trans)

	router := gin.Default()

	router.Use(authHandlers.Authorizer(e))
	router.Use(corsMW())

	v1 := router.Group("/v1")
	routing.UserRoutes(v1, usrHandler)
	routing.CompanyRoutes(v1, compHandler)
	routing.RoleRoutes(v1, roleHandler)
	routing.PermissionRoutes(v1, permHandler)
	routing.AuthRoutes(v1, authHandlers)

	logrus.WithFields(logrus.Fields{
		"host": "127.0.0.1",
		"port": ":8080",
	}).Info("Starts Serving on HTTP")

	log.Fatal(http.ListenAndServe("127.0.0.1"+":"+"8080", router))
}
func corsMW() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, PUT, PATCH, OPTIONS, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}

		c.Next()
	}
}
