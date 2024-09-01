package main

import (
	"context"
	"database/sql"
	"fmt"
	"real-time-forum/backend"
)

var db *sql.DB

func main() {

	db = backend.OpenDB()
	defer db.Close()
	fmt.Println("DB opened")
	backend.CreateTables()
	//Backround ja mux on ikka veel segased. mõlemad on just kui http requestid aga saan aru
	//et wsi jaoks on backgroundi vaja?
	rootCtx := context.Background()
	ctx, cancel := context.WithCancel(rootCtx)
	backend.SetCtx(ctx)

	backend.SetUpServer()

	defer cancel()

}
