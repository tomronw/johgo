package logger

import (
	"github.com/labstack/gommon/log"
	"io"
	logging "log"
	"os"
	"time"
)

var (
	ApiWarningLogger *logging.Logger
	ApiInfoLogger    *logging.Logger
	ApiErrorLogger   *logging.Logger
)

func init() {

	now := time.Now()
	f, err := os.OpenFile("internal/core/logs/api/"+now.Format("01-02-2006")+"-johgo-api.log",
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		logging.Fatalf("error opening file: %v", err)
	}

	mw := io.MultiWriter(os.Stdout, f)

	ApiInfoLogger = logging.New(mw, "[INFO] - ", logging.Ldate|logging.Ltime|logging.Lshortfile)
	ApiWarningLogger = logging.New(mw, "[WARNING] - ", logging.Ldate|logging.Ltime|logging.Lshortfile)
	ApiErrorLogger = logging.New(mw, "[ERROR] - ", logging.Ldate|logging.Ltime|logging.Lshortfile)

	log.SetOutput(mw)

}
