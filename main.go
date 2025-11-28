package main

import (
	"blog_Agregator/internal/config"
	"blog_Agregator/internal/database"
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

type state struct {
	cfg *config.Config
	db  *database.Queries
}

func main() {
	conf, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}
	poinconf := &conf

	db, err := sql.Open("postgres", conf.Db_url)
	if err != nil {
		fmt.Printf("%v", err)
		os.Exit(1)
	}

	dbQueries := database.New(db)

	s := &state{
		db:  dbQueries,
		cfg: poinconf,
	}

	cmds := commands{
		registeredCommands: make(map[string]func(*state, command) error),
	}

	cmds.register("login", handlerLogin)
	cmds.register("register", handlerRegister)
	cmds.register("reset", handlerReset)
	cmds.register("users", handlerUsers)
	cmds.register("agg", handlerAgg)
	cmds.register("addfeed", handlerAddFeed)
	cmds.register("feeds", handlerFeeds)
	cmds.register("follow", handlerFollow)
	cmds.register("following", handlerFollowing)

	if len(os.Args) < 2 {
		fmt.Println("error: Not enough arguments given")
		os.Exit(1)
	}

	cmdName := os.Args[1]
	cmdArgs := os.Args[2:]

	err = cmds.run(s, command{Name: cmdName, Args: cmdArgs})
	if err != nil {
		fmt.Printf("%v\n", err)
		os.Exit(1)
	}

}
