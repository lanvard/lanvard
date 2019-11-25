package main

import (
	"lanvard/foundation/http"
	http2 "lanvard/http"
	httpInterface "lanvard/interface/http"
	"lanvard/src/bootstrap"
	"log"
	net "net/http"
)

func main() {
	net.HandleFunc("/", handleKernel)

	log.Println("Server is ready to handle requests")
	if err := net.ListenAndServe(":80", nil); err != nil && err != net.ErrServerClosed {
		log.Fatal("Could not listen", err)
	}

	log.Println("Server stopped")
}

func handleKernel(response net.ResponseWriter, request *net.Request) {

	/*
	   |--------------------------------------------------------------------------
	   | Turn On The Lights
	   |--------------------------------------------------------------------------
	   |
	   | We need to illuminate PHP development, so let us turn on the lights.
	   | This bootstraps the framework and gets it ready for use, then it
	   | will load up this application so that we can run it and send
	   | the responses back to the browser and delight our users.
	   |
	*/
	app := bootstrap.App()

	/*
	   |--------------------------------------------------------------------------
	   | Register the response writer
	   |--------------------------------------------------------------------------
	   |
	   | Lanvard only uses the response writer here in main.go. But we register
	   | the response writer if you need it anyway
	   |
	*/
	app.Container.Singleton(
		(*net.ResponseWriter)(nil),
		response,
	)

	/*
	   |--------------------------------------------------------------------------
	   | Run The Application
	   |--------------------------------------------------------------------------
	   |
	   | Once we have the application, we can handle the incoming request
	   | through the kernel, and send the associated response back to
	   | the client allowing them to enjoy the creative
	   | and wonderful application we have prepared for them.
	   |
	*/
	kernel := app.Make((*httpInterface.Kernel)(nil)).(http.KernelStruct)
	appResponse := kernel.Handle(http2.Request(app, *request))

	println("appResponse")
	println(appResponse)

	// todo convert custom 'buffer' response to default go response
}
