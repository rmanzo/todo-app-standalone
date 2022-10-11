package reprogramable

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
	"time"
)

type Weight struct {
	Name          string  `json:"name"`
	Date		  time.Time `json:"date"`
	CurrentWeight float64 `json:"currentWight"`
}

type Weights []Weight

func TouchFile(name string) error {
	file, err := os.OpenFile(name, os.O_RDONLY|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	return file.Close()
}

func (w *Weights) AddToSlice(weight float64) {
	entry := Weight{
		Name:	"rino",
		Date:	time.Now(),
		CurrentWeight: weight,
	}

	*w = append(*w, entry)

}

func (w *Weights) SaveToFile(filename string) error {
	byteArray, err := json.Marshal(w)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(filename, byteArray, 0644)
	if err != nil {
		return err
	}
	return nil
}

func (w *Weights) LoadFromFile(filename string) error {
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			TouchFile(filename)
		}

	}
	json.Unmarshal(file, w)
	return nil

}
