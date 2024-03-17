package main

import (
	"database/sql"
	"log"

	supa "github.com/nedpals/supabase-go"
)

func configDBClient(sbClient SupabaseClient) *supa.Client {
	supabaseClient := supa.CreateClient(sbClient.Key, sbClient.Url)

	return supabaseClient
}

func openDB(config DatabaseConfig) *sql.DB {
	connStr := "user=" + config.User +
		" password=" + config.Password +
		" host=" + config.Host +
		" port=" + config.Port +
		" dbname=" + config.DBName

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		db.Close()
		log.Fatal(err)
	}

	return db
}
