# logger
### Simple log to terminal

This simple library prints, based on set LogLevel, nice formatted to terminal

![Screenshot](https://github.com/altafino/logger/blob/master/screenshot.png)

#### usage

```` go
import {
    "github.com/altafino/logger"
    "flag"
}

func main {
    level:=flag.String("log","info","sets log level")
    terminalStyle:=flag.String("style","flat","sets terminal style")
    flag.Parse()
    
    InitLog(*level,*terminalStyle)
}

func InitLog(level string, style string) {

	fmt.Println("level:", level, "style:", style)

	lsettings := logger.Settings{}

	switch level {
	case "debug":
		lsettings = logger.Settings{
			Level: logger.DebugLevel,
		}
	case "info":
		lsettings = logger.Settings{
			Level: logger.InfoLevel,
		}
	case "http":
		lsettings = logger.Settings{
			Level: logger.HttpLevel,
		}
	case "critical":
		lsettings = logger.Settings{
			Level: logger.CriticalLevel,
		}
	case "disabled":
		lsettings = logger.Settings{
			Level: logger.Disabled,
		}
	default:
		lsettings = logger.Settings{
			Level: logger.InfoLevel,
		}
	}

	lsettings.Output = logger.Terminal

	switch style {
	case "flat":
		lsettings.TerminalStyle = logger.FlatStyle
	case "json":
		lsettings.TerminalStyle = logger.JsonStyle
	default:
		lsettings.TerminalStyle = logger.FlatStyle
	}

	logger.InitLogger(lsettings)
	logger.Info("Logger Started", lsettings)

}


````

#### use as http middleware
##### example with go-chi

```` go
import {
    "github.com/altafino/logger/middleware"
}
// ....
router.Use(
		render.SetContentType(render.ContentTypeJSON), // Set content-Type headers as application/json
		middleware.Logger,             // Log API request calls
		chiMiddleware.DefaultCompress, // Compress results, mostly gzipping assets and json
		chiMiddleware.RedirectSlashes, // Redirect slashes to no slash URL versions
		chiMiddleware.Recoverer,       // Recover from panics without crashing server
		cors.Handler,
	)
// ....
````
