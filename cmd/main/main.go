package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/alim7007/go-bookstore/pkg/routes"
	"github.com/gorilla/mux"
)

func main(){
	r:=mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/",r)
	fmt.Println("server listening on port:8010")
	if err := http.ListenAndServe("localhost:8010", nil); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

// docker run --name pgsql-dev --rm -e POSTGRES_PASSWORD=olim123 -v ${PWD}/postgres-docker:/var/lib/postgresql/data -p 5432:5432 postgres
// docker run -p 8070:70 -e 'PGADMIN_DEFAULT_EMAIL=user@domain.com' -e 'PGADMIN_DEFAULT_PASSWORD=SuperSecret' -d dpage/pgadmin4
// 2137  docker run --name=mysql-dev MYSQL_ROOT_PASSWORD="olim123" --publish 6603:6603 --volume=/root/docker/mysql-dev/conf.d:/etc/mysql/conf.d -d 

// docker exec -it pgsql-dev bin/bash
// psql -h localhost -p 5432 -U postgres

// \l list of databases
// \c Switching Databases
// \dt Listing Tables