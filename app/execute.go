package app

import (
	"context"
	"errors"
	"fmt"
	"net"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/log"
	"github.com/charmbracelet/ssh"
	"github.com/charmbracelet/wish"
	"github.com/charmbracelet/wish/activeterm"
	"github.com/charmbracelet/wish/logging"
)

func Exec() {
	model, options := Tui()
	p := tea.NewProgram(model, options...)
	if _, err := p.Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}

func SSHExec() {
	s, err := wish.NewServer(
		wish.WithAddress(net.JoinHostPort(GlobalConfig.SSH.Host, strconv.Itoa(GlobalConfig.SSH.Port))),
		wish.WithHostKeyPath(".ssh/id_ed25519"),
		wish.WithMiddleware(
			activeterm.Middleware(),
			logging.Middleware(),
			func(sh ssh.Handler) ssh.Handler {
				return func(s ssh.Session) {
					model, options := TuiSSH(s)
					p := tea.NewProgram(model, options...)
					if _, err := p.Run(); err != nil {
						log.Error("Error running program", "error", err)
					}
				}
			},
		),
	)
	if err != nil {
		log.Error("Failed starting server", "error", err)
		return
	}

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	log.Info("Server started", "address", s.Addr, "port", GlobalConfig.SSH.Port)

	go func() {
		if err = s.ListenAndServe(); err != nil && !errors.Is(err, ssh.ErrServerClosed) {
			log.Error("Could not start server", "error", err)
			done <- syscall.SIGTERM
		}
	}()

	<-done
	log.Info("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := s.Shutdown(ctx); err != nil && !errors.Is(err, ssh.ErrServerClosed) {
		log.Error("Could not shutdown server", "error", err)
	}
}