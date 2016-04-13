package routers

import (
  "github.com/codegangsta/negroni"
  "github.com/gorilla/mux"
  "github.com/shijuvar/go-web/taskmanager/common"
  "github.com/shijuvar/go-web/taskmanager/controllers"
)

func SetItemRoutes(router *mux.Router) *mux.Router {
  itemRouter := mux.NewRouter()
  itemRouter.HandleFunc("/items", controllers.CreateItem).Methods("POST")
  itemRouter.HandleFunc("/items/{id}", controllers.UpdateItem).Methods("PUT")
  itemRouter.HandleFunc("/items", controllers.GetItems).Methods("GET")
  itemRouter.HandleFunc("/items/{id}", controllers.GetItemById).Methods("GET")
  itemRouter.HandleFunc("/items/user/{id}", controllers.GetItemsByUser).Methods("GET")
  itemRouter.HandleFunc("/items/{id}", controllers.DeleteItem).Methods("DELETE")
  router.PathPrefix("/items").Handler(negroni.New(
    negroni.HandlerFunc(common.authorize),
    negroni.Wrap(itemRouter),
  ))
  return router
}
