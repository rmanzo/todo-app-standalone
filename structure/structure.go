package structure

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"time"

	"github.com/alexeyco/simpletable"
)

type Todo struct {
	Task          string  `json:"task"`
	Done		  bool `json:"done"`
	CreatedAt	 time.Time `json:"createdAt"`
	CompletedAt	  time.Time	`json:"comepletedAt"`
}

type Todos []Todo

func TouchFile(name string) error {
	file, err := os.OpenFile(name, os.O_RDONLY|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	return file.Close()
}

func (t *Todos) AddToSlice(task string) {
	entry := Todo{
		Task:	task,
		Done:	false,
		CreatedAt:	time.Now(),
		CompletedAt: time.Time{},
	}

	*t = append(*t, entry)

}

func (t *Todos) SaveToFile(filename string) error {
	byteArray, err := json.Marshal(t)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(filename, byteArray, 0644)
	if err != nil {
		return err
	}
	return nil
}

func (t *Todos) LoadFromFile(filename string) error {
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			TouchFile(filename)
		}
	}
	json.Unmarshal(file, t)
	return nil

}
func (t *Todos) Complete(index int) {
	slice := *t
	if index > 0 || index < len(slice) {
		slice[index - 1].Done = true
		slice[index - 1].CompletedAt = time.Now()
	}
}

func (t *Todos) Print() {
	table := simpletable.New()
	table.Header = &simpletable.Header{
		Cells: []*simpletable.Cell{
			{Align: simpletable.AlignCenter, Text: "#"},
			{Align: simpletable.AlignCenter, Text: "Task"},
			{Align: simpletable.AlignCenter, Text: "Done"},
			{Align: simpletable.AlignRight, Text: "CreatedAt"},
			{Align: simpletable.AlignRight, Text: "CompletedAt"},
		},
	}
	var cells [][]*simpletable.Cell

	for index, item := range *t {
		index++
		task := Blue(item.Task)
		done := Blue("no")
		complete := "-"
		if item.Done {
			task = green(fmt.Sprintf("\u2705 %s", item.Task))
			done = green("yes")
			complete = item.CompletedAt.Format(time.RFC822)
		}
		cells = append(cells, *&[]*simpletable.Cell{
			{Text: fmt.Sprintf("%d", index)},
			{Text: task},
			{Text: done},
			{Text: item.CreatedAt.Format(time.RFC822)},
			{Text: complete},
		})
	}

	table.Body = &simpletable.Body{Cells: cells}
	table.Footer = &simpletable.Footer{Cells: []*simpletable.Cell{
		{Align: simpletable.AlignCenter, Span: 5, Text: red(fmt.Sprintf("Your have %d pending todos", t.CountPending()))},
	}}

	table.SetStyle(simpletable.StyleUnicode)

	table.Println()
}
func (t *Todos) Delete(index int) {
	index--
	s := *t
	if index <= 0 || index > len(s) {
		//return error.New("invalid index")
	} 
	*t = append(s[:index], s[index+1:]...)

}

func (t *Todos) CountPending() int{
	total := 0
	for _, item := range *t {
		if !item.Done {
			total++
		}
	}
	return total
}
