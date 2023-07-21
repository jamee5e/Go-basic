package main

import (
	"os"

	"github.com/jamee5e/jame-shop-tutorial/config"
	"github.com/jamee5e/jame-shop-tutorial/modules/servers"
	databases "github.com/jamee5e/jame-shop-tutorial/pkg/databases/migration"
)

func envPath() string {
	if len(os.Args) == 1 {
		return ".env"
	} else {
		return os.Args[1]
	}
}
func main() {
	cfg := config.LoadConfig(envPath())
	db := databases.DbConnect(cfg.Db())
	defer db.Close()

	servers.NewServer(cfg, db).Start()
}
