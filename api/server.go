package api

import (
	"log"
	"fmt"
	"context"
	"os/signal"
	"os" 
	"time" 
	"net/http" 
	"github.com/vonmutinda/crafted/config"
	"github.com/vonmutinda/crafted/api/database"
	"github.com/vonmutinda/crafted/api/router" 
)

func Run(){  
	config.Load()

	_, err := database.Connect()

	if err != nil{
		log.Println(err)
	}
	
	Listen( config.PORT)
}

func Listen(p string){
	
	// set routes 
	r := router.New()

	// server configurations
	s := &http.Server{
		Addr			: p,
		Handler  		: r,
		IdleTimeout		: 1*time.Second,
		ReadTimeout		: 1*time.Second,
		WriteTimeout	: 120*time.Second, 
	}
	
	// server goroutine
	go func(){ 
		fmt.Println(fmt.Sprintf("Server up and running on http://localhost%s", config.PORT))
		log.Fatal("Go run go!", s.ListenAndServe())
	}()
	
	// for graceful shutdown - channel recieves signal
	// CTRL+C or SIGTERM
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, os.Kill)

	// blocks until a signal is recieved
	sig := <-c 
	log.Println("Signal :",sig)
	
	// context. cancel func complains if ignored.
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second) 
	defer cancel() 
	
	log.Println("Server shutting down...")
	s.Shutdown(ctx)
	os.Exit(0)
}

