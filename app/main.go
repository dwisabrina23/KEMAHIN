package main

import (
	_driverFactory "kemahin/drivers"

	_userService "kemahin/businesses/users"
	_userController "kemahin/controllers/users"
	_userRepo "kemahin/drivers/databases/users"

	_eventService "kemahin/businesses/events"
	_eventController "kemahin/controllers/events"
	_eventRepo "kemahin/drivers/databases/events"

	_orgService "kemahin/businesses/organizers"
	_orgController "kemahin/controllers/organizer"
	_orgRepo "kemahin/drivers/databases/organizers"

	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
	"gorm.io/gorm"

	_middleware "kemahin/app/middlewares"
	_routes "kemahin/app/routes"
	_dbDriver "kemahin/drivers/mysql"
	"log"
)

func init() {
	viper.SetConfigFile(`app/config/config.json`)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	if viper.GetBool(`debug`) {
		log.Println("Service RUN on DEBUG mode")
	}

}

func dbMigrate(db *gorm.DB) {
	db.AutoMigrate(
		&_userRepo.Users{},
		&_eventRepo.Events{},
		&_orgRepo.Organizer{},
	)
}

func main() {
	configDB := _dbDriver.ConfigDB{
		DB_Username: viper.GetString(`database.user`),
		DB_Password: viper.GetString(`database.pass`),
		DB_Host:     viper.GetString(`database.host`),
		DB_Port:     viper.GetString(`database.port`),
		DB_Database: viper.GetString(`database.name`),
	}

	db := configDB.InitialDB()
	dbMigrate(db)

	configJWT := _middleware.ConfigJWT{
		SecretJWT:       viper.GetString(`jwt.secret`),
		ExpiresDuration: viper.GetInt(`jwt.expired`),
	}

	e := echo.New()

	userRepo := _driverFactory.NewUserRepository(db)
	userService := _userService.NewService(userRepo, &configJWT)
	userCtrl := _userController.NewUserController(userService)

	eventRepo := _driverFactory.NewEventRepository(db)
	eventService := _eventService.NewService(eventRepo)
	eventCtrl := _eventController.NewEventController(eventService)

	orgRepo := _driverFactory.NewOrgRepository(db)
	orgService := _orgService.NewService(orgRepo, &configJWT)
	orgCtrl := _orgController.NewOrgController(orgService)

	routesInit := _routes.ControllerList{
		JWTMiddleware:       configJWT.Init(),
		UserController:      *userCtrl,
		EventController:     *eventCtrl,
		OrganizerController: *orgCtrl,
	}
	routesInit.RouteRegister(e)

	log.Fatal(e.Start(viper.GetString("server.address")))
}
