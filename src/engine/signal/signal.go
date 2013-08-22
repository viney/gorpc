package signal

import (
	"os"
	"os/signal"
	"syscall"
)

const sigChanLen = 16

func Serve() {
	in := make(chan os.Signal, sigChanLen)
	signal.Notify(in, syscall.SIGQUIT, syscall.SIGINT, syscall.SIGTERM, syscall.SIGTSTP, syscall.SIGKILL)
	for sig := range in {
		switch sig {
		case syscall.SIGQUIT, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL:
			return
		case syscall.SIGTSTP:
			syscall.Kill(syscall.Getpid(), syscall.SIGSTOP)
		case syscall.SIGHUP:
			// reload config, etc
		case syscall.SIGUSR1:
			// miscalleaneous signals for own use
		default:
			// warn about unhandled signal, etc
		}
	}
}
