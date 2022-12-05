package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	blogdashboardcontroller "blogging/BlogDashBoard/Controller"
	blogdashboardservice "blogging/BlogDashBoard/Service"

	"github.com/gorilla/mux"
)

var (
	DashboardController *blogdashboardcontroller.DashboardController
	DashBoardService    *blogdashboardservice.DashBoardService
)

func init() {
	fmt.Println("==================== Service Init ================================")
	fmt.Println("DashBoardService Init")
	DashBoardService = blogdashboardservice.NewDashBoardService()
	fmt.Println("DashboardControler Init")
	DashboardController = blogdashboardcontroller.NewDashBoardController(DashBoardService)
	fmt.Println("==================================================================")
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/api/write_blog", DashboardController.WriteBlog).Methods("POST")
	router.HandleFunc("/api/read_blog", DashboardController.ReadBlog).Methods("GET")
	router.HandleFunc("/api/delete_blog", DashboardController.DeleteBlog).Methods("DELETE")
	router.HandleFunc("/api/health", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(map[string]bool{"ok": true})
	}).Methods("GET")

	srv := &http.Server{
		Handler: router,
		Addr:    "localhost:8000",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())
}
