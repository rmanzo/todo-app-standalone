package main

import (
	"flag"
	"fmt"

	"github.com/rmanzo/reprogramable"
)

const filename = ".db"

func main() {

	add := flag.Float64("add", 0.0, "Enter your weight")
	list := flag.Bool("list", false, "List your weights")
	cleft := flag.Bool("cleft", false, "Calories left for the week")
	flag.Parse()

	weightTable := &reprogramable.Weights{}
	err := weightTable.LoadFromFile(filename)
	if err != nil {
		fmt.Println(err)
	}

	switch {
	case *add != 0:
		weightTable.AddToSlice(*add)
		fmt.Println("Just testing")
		fmt.Println(*weightTable)
		weightTable.SaveToFile(filename)

	case *list:
		for index, item :=  range(*weightTable) {
			fmt.Println(index,item.Date,item.Name,item.CurrentWeight)
		}
	case *cleft:
		fmt.Println("1000 Calories left")
	}

	

	
}
