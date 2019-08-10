package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"time"

	"github.com/linuxkit/virtsock/pkg/vsock"
	"github.com/spf13/cobra"

	"github.com/mutagen-io/mutagen/cmd"
	"github.com/mutagen-io/mutagen/pkg/agent"
	"github.com/mutagen-io/mutagen/pkg/housekeeping"
	"github.com/mutagen-io/mutagen/pkg/logging"
	"github.com/mutagen-io/mutagen/pkg/synchronization/endpoint/remote"
)

const (
	// housekeepingInterval is the interval at which housekeeping will be
	// invoked by the agent.
	housekeepingInterval = 24 * time.Hour
)

func housekeepRegularly(context context.Context, logger *logging.Logger) {
	// Perform an initial housekeeping operation since the ticker won't fire
	// straight away.
	logger.Println("Performing initial housekeeping")
	housekeeping.Housekeep()

	// Create a ticker to regulate housekeeping and defer its shutdown.
	ticker := time.NewTicker(housekeepingInterval)
	defer ticker.Stop()

	// Loop and wait for the ticker or cancellation.
	for {
		select {
		case <-context.Done():
			return
		case <-ticker.C:
			logger.Println("Performing regular housekeeping")
			housekeeping.Housekeep()
		}
	}
}

func endpointMain(command *cobra.Command, arguments []string) error {
	// Create a channel to track termination signals. We do this before creating
	// and starting other infrastructure so that we can ensure things terminate
	// smoothly, not mid-initialization.
	signalTermination := make(chan os.Signal, 1)
	signal.Notify(signalTermination, cmd.TerminationSignals...)

	// Set up regular housekeeping and defer its shutdown.
	housekeepingContext, housekeepingCancel := context.WithCancel(context.Background())
	defer housekeepingCancel()
	go housekeepRegularly(housekeepingContext, logging.RootLogger.Sublogger("housekeeping"))

	fmt.Println("VSOCK LISTEN")
	ln, err := vsock.Listen(uint32(vsock.CIDAny), uint32(4099))
	if err != nil {
		return err
	}

	for {
		conn, err := ln.Accept()
		if err != nil {
			return err
		}
		go func() {
			remote.ServeEndpoint(logging.RootLogger, conn)
		}()
	}
}

var endpointCommand = &cobra.Command{
	Use:          agent.ModeEndpoint,
	Short:        "Run the agent in endpoint mode",
	RunE:         endpointMain,
	SilenceUsage: true,
}

var endpointConfiguration struct {
	// help indicates whether or not help information should be shown for the
	// command.
	help bool
}

func init() {
	// Grab a handle for the command line flags.
	flags := endpointCommand.Flags()

	// Manually add a help flag to override the default message. Cobra will
	// still implement its logic automatically.
	flags.BoolVarP(&endpointConfiguration.help, "help", "h", false, "Show help information")
}
