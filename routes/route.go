package routes

import (
	"Backend-berkah/config"
	"Backend-berkah/controller"
	"Backend-berkah/helper"
	"log"
	"net/http"
	// Middleware untuk autentikasi dan otorisasi JWT
)

func URL(w http.ResponseWriter, r *http.Request) {
    // Log request method dan path
    log.Printf("Incoming request: %s %s", r.Method, r.URL.Path)

    // Set Access Control Headers
    if config.SetAccessControlHeaders(w, r) {
        return
    }

    // Load environment variables
    config.LoadEnv()

    // Ambil metode dan path dari request
    method := r.Method
    path := r.URL.Path

    // Routing berdasarkan method dan path
    switch {
    // User authentication routes
    case method == "POST" && path == "/register":
        controller.Register(w, r)
    case method == "POST" && path == "/login":
        controller.Login(w, r)

    // Google OAuth routes
    case method == "GET" && path == "/auth/google/login":
        controller.HandleGoogleLogin(w, r) // Menangani login dengan Google
    case method == "GET" && path == "/auth/callback":
        controller.HandleGoogleCallback(w, r) // Menangani callback dari Google

    // Admin route untuk manage CRUD
	case method == "GET" && path == "/retreive/data":
		controller.GetLocation(w, r)	
    case method == "GET" && path == "/getlocation":
        controller.GetAllLocation(w, r)
    case method == "POST" && path == "/createlocation":
        controller.CreateLocation(w, r)
    case method == "PUT" && path == "/updatelocation":
        controller.UpdateLocation(w, r)
    case method == "DELETE" && path == "/deletelocation":
        controller.DeleteLocation(w, r)	

    // Default route
    default:
        helper.NotFound(w, r)
    }
}
