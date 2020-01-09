
package controllers

import "github.com/mohit/todo-app-backend/api/middlewares"

func (s *Server) initializeRoutes() {

	// Home Route
	s.Router.HandleFunc("/", middlewares.SetMiddlewareAuthentication(s.Home)).Methods("GET")

	// Login Route
	s.Router.HandleFunc("/login", middlewares.SetMiddlewareJSON(s.Login)).Methods("POST")
	
	//Users Route
	s.Router.HandleFunc("/users", middlewares.SetMiddlewareAuthentication(s.GetUsers)).Methods("GET")
	s.Router.HandleFunc("/users", middlewares.SetMiddlewareAuthentication(s.CreateUser)).Methods("POST")
	s.Router.HandleFunc("/users/{id}", middlewares.SetMiddlewareAuthentication(s.DeleteUser)).Methods("DELETE")
	s.Router.HandleFunc("/users/{id}", middlewares.SetMiddlewareAuthentication(s.GetUser)).Methods("GET")
	s.Router.HandleFunc("/users/{id}", middlewares.SetMiddlewareAuthentication(s.UpdateUser)).Methods("PUT")

	//Tasks Route
	s.Router.HandleFunc("/tasks", middlewares.SetMiddlewareAuthentication(s.GetTasks)).Methods("GET")
	s.Router.HandleFunc("/tasks", middlewares.SetMiddlewareAuthentication(s.CreateTask)).Methods("POST")
	s.Router.HandleFunc("/tasks/{id}", middlewares.SetMiddlewareAuthentication(s.DeleteTask)).Methods("DELETE")
	s.Router.HandleFunc("/tasks/{id}", middlewares.SetMiddlewareAuthentication(s.GetTask)).Methods("GET")
	s.Router.HandleFunc("/tasks/{id}", middlewares.SetMiddlewareAuthentication(s.UpdateTask)).Methods("PUT")

}