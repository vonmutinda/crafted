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
	m := http.NewServeMux() 

	s := &http.Server{
		Addr			: p,
		Handler  		: m,
		IdleTimeout		: 1*time.Second,
		ReadTimeout		: 1*time.Second,
		WriteTimeout	: 120*time.Second, 
	}
	
	go func(){ 
		fmt.Println(fmt.Sprintf("Server up and running on http://localhost%s", config.PORT))
		log.Fatal("Go run go!", s.ListenAndServe())
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, os.Kill)

	sig := <-c 
	log.Println("Got a signal :", sig)

	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	
	s.Shutdown(ctx)
}

