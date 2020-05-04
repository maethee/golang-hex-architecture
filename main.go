package main

import (
	"fmt"
	"log"

	server "github.axa.com/maethee-chakkuchantorn/psp-improvement/implem/gin.server"
	"github.axa.com/maethee-chakkuchantorn/psp-improvement/infra"
	"github.axa.com/maethee-chakkuchantorn/psp-improvement/internal/app/uc"
	"github.com/asticode/go-astikit"
	"github.com/asticode/go-astilectron"
)

func main() {
	//start gin server
	run()

	//start with astictron
	startAstilectron()
}

func startAstilectron() {
	// Set logger
	l := log.New(log.Writer(), log.Prefix(), log.Flags())

	// Create astilectron
	a, err := astilectron.New(l, astilectron.Options{
		AppName:           "Test",
		BaseDirectoryPath: "example",
	})
	if err != nil {
		l.Fatal(fmt.Errorf("main: creating astilectron failed: %w", err))
	}
	defer a.Close()

	// Handle signals
	a.HandleSignals()

	// Start
	if err = a.Start(); err != nil {
		l.Fatal(fmt.Errorf("main: starting astilectron failed: %w", err))
	}

	// New window
	var w *astilectron.Window
	if w, err = a.NewWindow("website/index.html", &astilectron.WindowOptions{
		Center: astikit.BoolPtr(true),
		Height: astikit.IntPtr(700),
		Width:  astikit.IntPtr(700),
	}); err != nil {
		l.Fatal(fmt.Errorf("main: new window failed: %w", err))
	}

	// Create windows
	if err = w.Create(); err != nil {
		l.Fatal(fmt.Errorf("main: creating window failed: %w", err))
	}

	//Open debug mode
	w.OpenDevTools()
	
	// Blocking pattern
	a.Wait()
}

func run() {
	ginServer := infra.NewServer(5000, 0)
	//mapping route
	server.NewRouter(uc.HandlerConstructor{}.New()).SetRoutes(ginServer.Router)
	ginServer.Start()
	return
}
