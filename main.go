package main

import (
    "log"
    "net/http"
    "school_management_app/routers"
    "school_management_app/controllers"
)

func main() {
    controllers.InitDB()
    defer controllers.DB.Close()

    routers.SetupRoutes()

    log.Println("Server listening on port 8080...")
    log.Fatal(http.ListenAndServe(":8080", nil))
}
