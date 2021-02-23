package middleware

import (
	"bytes"
	"io/ioutil"
	"log"
	"os"
	"time"

	"github.com/getsentry/sentry-go"
	"github.com/gin-gonic/gin"
)

// LogTerminal func
func LogTerminal(c *gin.Context) {
	reqMethod := c.Request.Method
	reqPath := c.Request.URL.Path
	log.Println(reqMethod, " -> ", reqPath)
}

// LogSentry func
func LogSentry(c *gin.Context) {
	reqMethod := c.Request.Method
	reqURL := c.Request.URL.Path
	buf, _ := ioutil.ReadAll(c.Request.Body)
	// fmt.Println(string(buf))
	c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(buf))
	sentryMessage := string(reqMethod) + " -> " + reqURL + "\n" + string(buf)
	Sentry(sentryMessage)
}

// Sentry func to log with sentry.io
func Sentry(data string) {
	dsn := os.Getenv("SENTRY_DSN")

	err := sentry.Init(sentry.ClientOptions{
		// Either set your DSN here or set the SENTRY_DSN environment variable.
		Dsn: dsn,
	})
	if err != nil {
		sentry.CaptureException(err)
		log.Fatalf("sentry.Init: %s", err)
	}
	// Flush buffered events before the program terminates.
	// Set the timeout to the maximum duration the program can afford to wait.
	defer sentry.Flush(2 * time.Second)
	sentry.CaptureMessage(data)
}
