package main

import "C"

import (
	"fmt"
	"os"
	"unsafe"

	"github.com/jessevdk/go-flags"
	"github.com/pkt-cash/pktd/btcutil/er"
	"github.com/pkt-cash/pktd/lnd"
	"github.com/pkt-cash/pktd/lnd/signal"
)

func main() {
	// Load the configuration, and parse any command line options. This
	// function will also set up logging properly.
	loadedConfig, err := lnd.LoadConfig()
	if err != nil {
		errr := er.Wrapped(err)
		if e, ok := errr.(*flags.Error); !ok || e.Type != flags.ErrHelp {
			// Print error if not due to help request.
			_, _ = fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

		// Help was requested, exit normally.
		os.Exit(0)
	}

	// Hook interceptor for os signals.
	if err := signal.Intercept(); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	//	we want to disable the use of macaroons so, force that it's turned off
	loadedConfig.NoMacaroons = true

	// Call the "real" main in a nested manner so the defers will properly
	// be executed in the case of a graceful shutdown.
	if err := lnd.Main(
		loadedConfig, lnd.ListenerCfg{}, signal.ShutdownChannel(),
	); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

//export startService
func startService(numParams C.int, params **C.char) {
	fmt.Println("startService()")
	length := int(numParams)
	tmpslice := (*[(1 << 29) - 1]*C.char)(unsafe.Pointer(params))[:length:length]
	parameters := make([]string, length)
	for i, s := range tmpslice {
		parameters[i] = C.GoString(s)
	}

	for i := 0; i < len(parameters); i++ {
		parameter := parameters[i]
		fmt.Println(parameters[i])
		os.Args = append(os.Args, parameter)
	}
	main()
}
