package config

import (
	"os"

	"github.com/vaughan0/go-ini"
)

type env struct {
	path string
	data ini.Section
}

func newEnv(path string) env {
	return env{path: path}
}

func (e *env) Parse() error {
	f, err := ini.LoadFile(e.path)
	if err != nil {
		return err
	}
	e.data = f.Section("")

	return nil
}

func (e env) Set() {
	for k, v := range e.data {
		os.Setenv(k, v)
	}
}
