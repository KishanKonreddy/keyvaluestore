package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/dheerajkumaralla/distributed-kv-store/config"
	"github.com/dheerajkumaralla/distributed-kv-store/db"
	"github.com/dheerajkumaralla/distributed-kv-store/web"
)

var (
	dbLocation  = flag.String("db-location", "", "Path to the bolt db")
	httpAddr    = flag.String("http-addr", "localhost:8080", "Host and port for the HTTP server")
	shardConfig = flag.String("shard-config", "shard-config.toml", "Path to the (static) shard config file")
	shard       = flag.String("shard", "", "shard name")
)

func parseFlags() {
	flag.Parse()
	if *dbLocation == "" {
		log.Fatal("db-location is required.")
	}

	if *shard == "" {
		log.Fatal("shard is required.")
	}
}

func main() {
	parseFlags()

	var c *config.Config
	c, err := config.LoadConfigurationFromFile(*shardConfig)

	if err != nil {
		log.Fatalf("Failed to parse shard config file: %v", err)
	}

	shards, err := config.InitializeShardTopology(c.Shards, *shard)

	if err != nil {
		log.Fatalf("Failed to parse shard config: %v", err)
	}

	log.Println("Total shards: ", shards.Count, " Shard index: ", shards.CurrIdx)

	db, close, err := db.CreateDatabaseInstance(*dbLocation)
	if err != nil {
		log.Fatalf("Failed to open database: %v", err)
	}
	defer close()

	server := web.CreateWebServer(db, shards)

	http.HandleFunc("/retrieve", server.HandleValueRetrieval)
	http.HandleFunc("/store", server.HandleValueStorage)

	log.Fatal(http.ListenAndServe(*httpAddr, nil))
}
