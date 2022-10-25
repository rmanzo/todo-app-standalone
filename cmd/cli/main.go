package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/rmanzo/todo-app-standalone/structure"
)

const filename = ".db"

func main() {

	add := flag.String("add", "", "Add a task")
	complete := flag.Int("complete", 0, "Complete a task")
	delete := flag.Int("delete", 0, "Delete a task")
	list := flag.Bool("list", false, "List tasks")
	flag.Parse()

	weightTable := &structure.Todos{}
	err := weightTable.LoadFromFile(filename)
	if err != nil {
		fmt.Println(err)
	}

	switch {
	case *add != "":
		weightTable.AddToSlice(*add)
		weightTable.SaveToFile(filename)

	case *list:
		weightTable.Print()

	case *delete > 0:
		weightTable.Delete(*delete)
		weightTable.SaveToFile(filename)
		fmt.Println("Delete a task")

	case *complete > 0:
		weightTable.Complete(*complete)
		weightTable.SaveToFile(filename)

	default: 
		fmt.Fprintln(os.Stdout, "invalid command")
		os.Exit(0)
	}
}
