package cmd

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/spf13/cobra"
	"github.com/vonmutinda/crafted/api/auto" 
	"github.com/vonmutinda/crafted/api/router"
	"github.com/vonmutinda/crafted/config"
)

var serverCmd = &cobra.Command{
	Use: "crafted",
	Aliases: []string{"runserver", "cm"},
	Short: "Crafted Microservices Server",
	Long: `Start Crafted Service...`,
	Run: func(cmd *cobra.Command, args []string){
		run()
	},
}

func init(){
	rootCmd.AddCommand(serverCmd)
}


func run(){ 
	auto.Load() 

	// set routes 
	r := router.LoadCORS(router.New())

	// server configurations
	s := &http.Server{
		Addr			: string(config.PORT),
		Handler  		: r,
		IdleTimeout		: 1*time.Second,
		ReadTimeout		: 1*time.Second,
		WriteTimeout	: 120*time.Second, 
		MaxHeaderBytes	: 1 << 20,
	}
	
	// server goroutine
	go func(){ 
		log.Println(fmt.Sprintf("Server up and running on http://localhost%s", config.PORT))
		log.Fatal("Go run go! ", s.ListenAndServe())
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

