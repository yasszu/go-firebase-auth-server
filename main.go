package main

import (
	"context"
	"flag"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/mux"

	"go-firebase-auth-server/application/usecase"
	"go-firebase-auth-server/infrastructure/db"
	"go-firebase-auth-server/infrastructure/firebase"
	"go-firebase-auth-server/infrastructure/persistence"
	"go-firebase-auth-server/interfaces/handler"
	_middleware "go-firebase-auth-server/interfaces/middleware"
	"go-firebase-auth-server/util"
)

func main() {
	var wait time.Duration
	flag.DurationVar(&wait, "graceful-timeout", time.Second*30, "the duration for which the server gracefully wait for existing connections to finish - e.g. 15s or 1m")
	flag.Parse()

	// Load conf
	cnf := util.NewConf()

	// Establish DB connection
	conn, err := db.NewConn(cnf)
	if err != nil {
		panic(err.Error())
	}

	r := mux.NewRouter()

	authenticationService := firebase.NewAuthenticationService()
	userRepository := persistence.NewUserRepository(conn)
	userUsecase := usecase.NewUserUsecase(userRepository, authenticationService)
	middleware := _middleware.NewMiddleware(userUsecase)
	indexHandler := handler.NewIndexHandler(conn)
	authHandler := handler.NewAuthHandler(conn, userUsecase)
	userHandler := handler.NewUserHandler(conn, userUsecase)

	r.Use(middleware.Logging)
	r.Use(middleware.CORS)
	root := r.PathPrefix("").Subrouter()
	v1 := r.PathPrefix("/v1").Subrouter()
	v1.Use(middleware.FirebaseAuth)

	indexHandler.Register(root)
	authHandler.Register(root)
	userHandler.Register(v1)

	srv := &http.Server{
		Addr:         cnf.Server.Addr(),
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      r,
	}

	// Run our server in a goroutine so that it doesn't block.
	go func() {
		log.Printf(" ⇨ graceful timeout: %s", wait)
		log.Printf(" ⇨ http server started on %s", cnf.Server.Addr())
		if err = srv.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()

	c := make(chan os.Signal, 1)
	// We'll accept graceful shutdowns when quit via SIGINT (Ctrl+C)`
	// SIGKILL, SIGQUIT or SIGTERM (Ctrl+/) will not be caught.
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	// Block until we receive our signal.
	<-c
	log.Println("received stop signal")

	// Create a deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer func() {
		log.Println("cancel")
		cancel()
	}()
	// Doesn't block if no connections, but will otherwise wait
	// until the timeout deadline.
	_ = srv.Shutdown(ctx)
	// Optionally, you could run srv.Shutdown in a goroutine and block on
	// <-ctx.Done() if your application should wait for other services
	// to finalize based on context cancellation.
	log.Println("shutting down")
}
