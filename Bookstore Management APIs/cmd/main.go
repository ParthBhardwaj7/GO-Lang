package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/parthbhardwaj93/Bookstore-Management-APIs/package/routes"
)

func main() {

