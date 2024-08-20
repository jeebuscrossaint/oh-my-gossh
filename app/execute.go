package app

import (
	"fmt"
	"net"
	"os"
	"strconv"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/ssh"
	"github.com/charmbracelet/wish"
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
	server, err := wish.NewServer(
		wish.WithAddress(net.JoinHostPort(GlobalConfig.SSH.Host, strconv.Itoa(GlobalConfig.SSH.Port))),
		wish.WithHostKeyPath(".ssh/id_ed25519"),
		wish.WithMiddleware(
			bubbletea.Middleware(TuiSSH),
			activeterm.Middleware(),
			logging.Middleware(),
		),
		
		if err != nil {
			log.Error("Failed starting server," "error", err)
		}

		done := make(chan os.Signal, 1)
		signal.Notify(done, os.Interrupt, syscall.SIGINT ,syscall.SIGTERM)
		log.Info("Server started", "address", server.Addr(), "port", server.Port())
		go func() {
			if err = server.ListenAndServe(); err != nil && !errors.Is(err, ssh.ErrServerClosed) {
				log.Error("Could not starts server", "error", err)
				done <- nil
			}
		}()

		<-done
		log.Info("Server stopped")
		ctx, cancel :=context.WithTimeout(context.Background(), 30*time.Second)
		defer func() { cancel() }()
		if err := server.Shutdown(ctx); err != nil && !errors.Is(err, ssh.ErrServerClosed) {
			log.Error("Could not shutdown server", "error", err)
		}
	)
}
