package main

import (
	app "gofiber-boilerplate/src"
	"log"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// http://patorjk.com/software/taag/#p=display&c=c&f=Graceful&t=Cozy
var signature = `
---------------------------------
Code by:
        ___  __  ____  _  _ 
       / __)/  \(__  )( \/ )
      ( (__(  O )/ _/  )  / 
       \___)\__/(____)(__/  


📫 Email: devcosmas@gmail.com
---------------------------------
`

// Run starts the app
// @title Gofiber Boilerplate API
// @version 1.0
// @description This is my gofiber boilerplate api server.
// @termsOfService http://swagger.io/terms/
// @contact.name Cozy
// @contact.url https://github.com/ItsCosmas
// @contact.email devcosmas@gmail.com
// @license.name MIT
// @license.url https://github.com/ItsCosmas/gofiber-boilerplate/blob/master/LICENSE
// @host  localhost:8000
// @BasePath /api/v1
func main() {
	loggerConfig := zap.NewProductionConfig()
	loggerConfig.EncoderConfig.TimeKey = "timestamp"
	loggerConfig.EncoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout(time.RFC3339)
	logger, _ := loggerConfig.Build()
	defer logger.Sync() // flushes buffer, if any
	sugar := logger.Sugar()

	sugar.Infof("Initialized App Log test: %s", "L34")

	log.Println(signature)
	app.Run()
}
