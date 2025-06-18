# logger
### Simple log to terminal

This simple library prints, based on set LogLevel, nice formatted to terminal

![Screenshot](https://github.com/altafino/logger/blob/master/screenshot.png)

#### usage

The logger supports several levels. When a particular log level is set using `LoggerSettings.Level`, messages will be printed if their own level is numerically less than or equal to `LoggerSettings.Level`. For example, if `LoggerSettings.Level` is `ErrorLevel` (value 2), then messages logged with `ErrorLevel` (2), `HttpLevel` (1), and `InfoLevel` (0) will be printed. The `DebugLevel` has the highest numerical value among standard levels and thus enables logging for all standard message types (Debug, Critical, Error, Http, Info).

Available levels for configuration (e.g., via command-line flags as shown in the example):
*   `info`: When set, logs messages of `InfoLevel`. (Prints: Info)
*   `http`: When set, logs messages of `HttpLevel` and `InfoLevel`. (Prints: Http, Info)
*   `error`: When set, logs messages of `ErrorLevel`, `HttpLevel`, and `InfoLevel`. (Prints: Error, Http, Info)
*   `critical`: When set, logs messages of `CriticalLevel`, `ErrorLevel`, `HttpLevel`, and `InfoLevel`. (Prints: Critical, Error, Http, Info)
*   `debug`: When set, logs messages of `DebugLevel`, `CriticalLevel`, `ErrorLevel`, `HttpLevel`, and `InfoLevel`. (Prints: Debug, Critical, Error, Http, Info)
*   `only`: A special level. When the logger is set to `only` (using `logger.OnlyLevel`), it will *only* print messages logged using the `logger.Only()` function. All other log messages (Info, Debug, Error, etc.) are suppressed. Conversely, calls to `logger.Only()` will *not* be printed if the logger level is set to anything other than `only`.
*   `disabled`: Disables all logging.

The `logger.Info()`, `logger.Http()`, `logger.Error()`, `logger.Critical()`, `logger.Debug()`, and `logger.Only()` functions are used to log messages at their respective levels.

```` go
import (
    "github.com/altafino/logger"
    "flag"
)

func main() {
    level:=flag.String("log","info","sets log level")
    terminalStyle:=flag.String("style","flat","sets terminal style")
    flag.Parse()
    
    InitLog(*level,*terminalStyle)
}

// InitLog is an example helper function demonstrating how to initialize the logger
// by converting string inputs (e.g., from command-line flags or config files)
// into the required logger.Settings struct.
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
	case "only":
		lsettings = logger.Settings{
			Level: logger.OnlyLevel,
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
import (
    "github.com/altafino/logger/middleware"
)
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
