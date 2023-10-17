package core

import (
	"github.com/labstack/gommon/log"
	"io"
	logging "log"
	"os"
	"time"
)

var (
	WarningLogger *logging.Logger
	InfoLogger    *logging.Logger
	ErrorLogger   *logging.Logger
)

func init() {
	// create logging files and init
	now := time.Now()
	f, err := os.OpenFile("internal/core/logs/engine/"+now.Format("01-02-2006")+"-johgo-engine.log",
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		logging.Fatalf("error opening file: %v", err)
	}

	mw := io.MultiWriter(os.Stdout, f)

	InfoLogger = logging.New(mw, "[INFO] - ", logging.Ldate|logging.Ltime|logging.Lshortfile)
	WarningLogger = logging.New(mw, "[WARNING] - ", logging.Ldate|logging.Ltime|logging.Lshortfile)
	ErrorLogger = logging.New(mw, "[ERROR] - ", logging.Ldate|logging.Ltime|logging.Lshortfile)

	log.SetOutput(mw)

}
