package main

import (
	"log"

	"github.com/wley3337/learning/tree/main/go/go_bank/internal/api"
	"github.com/wley3337/learning/tree/main/go/go_bank/internal/store"
)

func main() {
	db, err := store.NewConnection()
	if err != nil {
		log.Fatal(err)
	}

	if err := store.Setup(db); err != nil {
		log.Fatal(err)
	}

	err = api.RunServer(":3000", db)
	if err != nil {
		_ = db.Close()
		log.Fatal(err)
	}
}

// package main
//
// import (
// 	"github.com/wley3337/learning/tree/main/go/go_bank/internals/api"
// 	st "github.com/wley3337/learning/tree/main/go/go_bank/internals/store"
// 	"log"
// )
//
// func main() {
// 	store, err := st.NewPostgresStore()
// 	if err != nil {
// 		log.Fatal(err)
// 	}
//
// 	if err := store.Init(); err != nil {
// 		log.Fatal(err)
// 	}
//
// 	server := api.NewAPIServer(":3000", store)
// 	server.Run()
// }
