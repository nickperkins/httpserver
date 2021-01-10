package main

import (
	"fmt"
	"log"
	"os"

	poker "github.com/nickperkins/httpserver"
)

const dbFileName = "game.db.json"

func main() {

	store, closeStore, err := poker.FileSystemPlayerStoreFromFile(dbFileName)
	if err != nil {
		log.Fatal(err)
	}
	defer closeStore()
	fmt.Println("Let's play poker")
	fmt.Println("Type {Name} wins to record a win")
	poker.NewCLI(store, os.Stdin, poker.BlindAlerterFunc(poker.StdOutAlerter)).PlayPoker()

}
