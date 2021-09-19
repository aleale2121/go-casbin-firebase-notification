package initiator

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
	"github.com/joho/godotenv"
	gomail "gopkg.in/mail.v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"template/internal/adapter/glue/routing"
	compHandler "template/internal/adapter/http/rest/server/company"
	email3 "template/internal/adapter/http/rest/server/notification/email"
	publisher3 "template/internal/adapter/http/rest/server/notification/publisher"
	sms3 "template/internal/adapter/http/rest/server/notification/sms"
	rlHandler "template/internal/adapter/http/rest/server/role"
	usrHandler "template/internal/adapter/http/rest/server/user"
	"template/internal/adapter/repository"
	"template/internal/adapter/storage/persistence/company"
	"template/internal/adapter/storage/persistence/notification/email"
	"template/internal/adapter/storage/persistence/notification/publisher"
	"template/internal/adapter/storage/persistence/notification/sms"
	persistence "template/internal/adapter/storage/persistence/role"
	"template/internal/adapter/storage/persistence/user"
	"template/internal/constant"
	"template/internal/constant/errors"
	compUsecase "template/internal/module/company"
	email2 "template/internal/module/notification/email"
	publisher2 "template/internal/module/notification/publisher"
	sms2 "template/internal/module/notification/sms"
	roleUsecase "template/internal/module/role"
	usrUsecase "template/internal/module/user"
)

//global validator instance
var (
	uni      *ut.UniversalTranslator
	validate *validator.Validate
	trans    ut.Translator
)

func Initialize() {
	en := en.New()
	uni := ut.New(en, en)
	trans, _ = uni.GetTranslator("en")
	validate = validator.New()
	en_translations.RegisterDefaultTranslations(validate, trans)

	//DATABASE_URL := "postgres://postgres:yideg2378@localhost:5432/demo?sslmode=disable"
	//conn, err := gorm.Open(postgres.Open(DATABASE_URL), &gorm.Config{
	//	SkipDefaultTransaction: true,
	//})
	//if err != nil {
	//	log.Printf("Error when Opening database connection: %v", err)
	//	os.Exit(1)
	//}
	connStr, _ := constant.DbConnectionString()
	dbConn, err := gorm.Open(postgres.Open(connStr), &gorm.Config{})
	if err != nil {
         log.Fatal("error ",errors.ErrDatabaseConnection)
	}
	x, _ := dbConn.DB()
	defer x.Close()
	dbConn = dbConn.Debug()
	//loading environmental variables
	err = godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	//user and role
	usrPersistence := user.UserInit(dbConn)
	compPersistence := company.CompanyInit(dbConn)
	rolePersistent := persistence.RoleInit(dbConn)

	//notification persistence
	emailPersistent := email.EmailInit(dbConn)
	smsPersistent := sms.SmsInit(dbConn)
	publisherPersistent := publisher.NotificationInit(dbConn)
	err = emailPersistent.MigrateEmail()
	if err != nil {
		log.Fatal("error while creating models %s",err)
	}
	err = smsPersistent.MigrateSms()
	if err != nil {
		log.Fatal("error while creating models %s",err)
	}
	err = publisherPersistent.MigrateNotification()
	if err != nil {
		log.Fatal("error while creating models %s",err)
	}

	//notification services
	emailUsecase := email2.InitializeEmail(emailPersistent)
	smsUsecase := sms2.InitializeSms(smsPersistent)
	publisherUsecase := publisher2.InitializePublisher(publisherPersistent)

	//notification handlers
	m := gomail.NewMessage()
	v := validator.New()
	emailHandler :=email3.NewEmailHandler(emailUsecase,v,m)
	smsHandler :=sms3.NewSmsHandler(smsUsecase,v)
	publisherHandler :=publisher3.NewNotificationHandler(publisherUsecase,v)


	roleUsecase := roleUsecase.RoleInitialize(rolePersistent)
	roleHandler := rlHandler.NewRoleHandler(roleUsecase, trans)

	usrRepo := repository.UserInit()
	usrUsecase := usrUsecase.Initialize(usrRepo, usrPersistence, validate, trans)
	usrHandler := usrHandler.UserInit(usrUsecase, trans)

	compUsecase := compUsecase.Initialize(compPersistence, validate, trans)
	compHandler := compHandler.CompanyInit(compUsecase, trans)

	router := gin.Default()
	//  group: v1
	v1 := router.Group("/v1")
	routing.UserRoutes(v1, usrHandler)
	routing.CompanyRoutes(v1, compHandler)
	routing.RoleRoutes(v1, roleHandler)
	//notification
	routing.EmailRoutes(v1, emailHandler)
	routing.SmsRoutes(v1, smsHandler)
	routing.PublisherRoutes(v1, publisherHandler)
	router.Run()
}
