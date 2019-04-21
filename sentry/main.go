package main

import (
	"fmt"
	"github.com/getsentry/raven-go"
	"log"
	"time"
)

func main() {
	l1, err := newLogger("development", "my-app", "1.0.0")
	if err != nil {
		log.Fatal(err)
	}
	l2, err := newLogger("staging", "my-app", "1.0.0")
	if err != nil {
		log.Fatal(err)
	}

	for i := 0; i < 5; i++ {
		logID := "6B6ACF38-A90B-4A5D-AA8D-8CFF0737F2CA"

		l1.error(logID, "Hello ^ ^")
		l2.error(logID, "Hello ^ ^")

		time.Sleep(1 * time.Second)
	}

	l1.close()
	l2.close()

	time.Sleep(3 * time.Second)
}

type Logger struct {
	env string
	app string
	ver string
	c   *raven.Client
}

func newLogger(env string, app string, ver string) (*Logger, error) {
	raven.SetMaxQueueBuffer(1000)

	c, err := raven.New("https://14f621b116a141729c430854f08b0597:40315af41f2342189a662e5b7a33191b@sentry.io/1443202")
	if err != nil {
		return nil, err
	}
	c.SetEnvironment(env)
	c.SetDefaultLoggerName(fmt.Sprintf("%s@%s", app, ver))

	return &Logger{
		env: env,
		app: app,
		ver: ver,
		c:   c,
	}, nil
}

func (l *Logger) close() {
	l.c.Close()
}

func (l *Logger) error(logID string, message string) {
	tags := map[string]string{
		"app":     l.app,
		"version": l.ver,
		"log_id":  logID,
	}
	e := raven.Extra{}
	p := raven.NewPacketWithExtra(fmt.Sprintf("[%s] %s", time.Now().Format(time.RFC3339), message), e)
	p.Level = raven.ERROR
	p.Fingerprint = []string{l.env, l.app, l.ver, logID}
	eventID, _ := l.c.Capture(p, tags)
	log.Print(eventID)
}
