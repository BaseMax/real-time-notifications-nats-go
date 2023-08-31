package main

func main() {
	// Configuration setup

	// DB setup

	// Models setup

	// Start application
	r := InitRoutes()
	r.Logger.Fatal(r.Start(":8000"))
}
