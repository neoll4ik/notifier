package signal

import (
	"os"
	"os/signal"
	"syscall"
        . "github.com/aavzz/notifier/setup/syslog"
	. "github.com/aavzz/notifier/setup/pidfile"
	. "github.com/aavzz/notifier/setup/cfgfile"
)

func SignalHandling() {

	sigusr1 := make(chan os.Signal, 1)
	sighup := make(chan os.Signal, 1)
	sigint := make(chan os.Signal, 1)
	sigquit := make(chan os.Signal, 1)
	sigterm := make(chan os.Signal, 1)

	signal.Notify(sigusr1, syscall.SIGUSR1)
	signal.Notify(sighup, syscall.SIGHUP)
	signal.Notify(sigint, syscall.SIGINT)
	signal.Notify(sigquit, syscall.SIGQUIT)
	signal.Notify(sigterm, syscall.SIGTERM)

	go func() {
		<-sigusr1
		SysLog.Info("SIGUSR1 received")
		os.Exit(-1)
	}()

	go func() {
		for {
			<-sighup
			SysLog.Info("SIGHUP received, re-reading configuration file")
			ReadConfig()
		}
	}()

	go func() {
		<-sigint
		SysLog.Info("SIGINT received, exitting")
		RemovePidFile()
		os.Exit(0)
	}()

	go func() {
		<-sigquit
		SysLog.Info("SIGQUIT received, exitting")
		RemovePidFile()
		os.Exit(0)
	}()

	go func() {
		<-sigterm
		SysLog.Info("SIGTERM received, exitting")
		RemovePidFile()
		os.Exit(0)
	}()

}
