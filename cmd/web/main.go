package main

import (
	"flag"
	"log/slog"
	"net/http"
	"os"
)

type application struct {
	logger *slog.Logger
}

func main() {
	addr := flag.String("addr", ":4000", "HTTP network address")
	flag.Parse()
	// creating a custom logger
	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		AddSource: true,
	}))

	app := &application{
		logger: logger,
	}

	// Initialize a new servemux
	mux := http.NewServeMux()
	fileServer := http.FileServer(http.Dir(("./ui/static")))
	mux.Handle("GET /static/", http.StripPrefix("/static", fileServer))
	mux.HandleFunc("GET /{$}", app.home)
	mux.HandleFunc("GET /snippet/veiw/{id}", app.snippetView)
	mux.HandleFunc("GET /snippet/create", app.snippetCreate)
	mux.HandleFunc("POST /{$}", app.snippetCreatePost)

	logger.Info("Starting server", "addr", *addr)

	err := http.ListenAndServe(*addr, mux)
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}
}
