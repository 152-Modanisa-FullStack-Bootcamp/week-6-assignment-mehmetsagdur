package config

import (
	"encoding/json"
	"io"
	"os"
)

type Config struct {
	InitialBalanceAmount int `json:"initialBalanceAmount"`
	MinimumBalanceAmount int `json:"minimumBalanceAmount"`
}

var C = &Config{}

func init() {
	// test ve normal calistirmayi bulamadim. onerilere acigim
	file, err := os.Open(".config/" + env + ".json")
	if err != nil {
		file2, err2 := os.Open("../.config/" + env + ".json")
		if err2 != nil {
			panic(err2)
		}
		defer file2.Close()

		read, err2 := io.ReadAll(file2)
		if err2 != nil {
			panic(err2)
		}
		err2 = json.Unmarshal(read, C)
		if err2 != nil {
			panic(err2)
		}
	} else {
		defer file.Close()

		read, err := io.ReadAll(file)
		if err != nil {
			panic(err)
		}
		err = json.Unmarshal(read, C)
		if err != nil {
			panic(err)
		}
	}
}
