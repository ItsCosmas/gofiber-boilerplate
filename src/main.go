package main

import (
	app "gofiber-boilerplate/api"
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
