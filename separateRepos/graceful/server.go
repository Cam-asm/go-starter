// Package graceful implements wrappers which gracefully shutdown when given interrupt signals.
package graceful

// TODO move graceful into a separate Git repository
import (
	"context"
	"net/http"
	"os"
	"os/signal"

	"github.com/sirupsen/logrus"
)

// Server gracefully shuts down srv HTTP server without interrupting any active connections.
// See pkg.go.dev/net/http#Server.Shutdown for more details.
func Server(srv *http.Server) {
	idleConnectionsClosed := make(chan struct{})
	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt)
		<-sigint

		// When an interrupt signal is received then shut down.
		logrus.Info("received interrupt signal. Attempting to shutdown HTTP server", srv.Addr)
		if err := srv.Shutdown(context.Background()); err != nil {
			// Error from closing listeners, or context timeout:
			logrus.Error("HTTP server Shutdown:", err)
		}
		close(idleConnectionsClosed)
	}()

	logrus.Info("serving on address", srv.Addr)
	if err := srv.ListenAndServe(); err != http.ErrServerClosed {
		// Error starting or closing listener:
		logrus.Fatalln("HTTP server ListenAndServe:", err)
	}

	<-idleConnectionsClosed
}
