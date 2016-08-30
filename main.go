package main

import (
	"fmt"
	"path/filepath"
	"os/exec"
	"os"

	"github.com/BurntSushi/toml"
)

type Tools struct {
	Prefix string
	Compiler string
	Cflags []string
	Linker string
	Ldflags []string
	Archiver string
}

type Config struct {
	Target string
	Toolchain Tools
}

var cfg Config

func ParseConfig() {
	if _, err := toml.DecodeFile("dai.toml", &cfg); err != nil {
		// handle error
		fmt.Println("Error parsing Toml")
	}

	fmt.Printf("%v", cfg)
}

func compile(path string, info os.FileInfo, err error) error {
	cmd := exec.Command(cfg.Toolchain.Compiler, "-c", path)
	//cmd.Env = os.Environ()
	cmd.Run()

	return err
}

func main() {
	println("Hello Dai")
	ParseConfig()

	//cfg.Toolchain.Cflags = append(cfg.Toolchain.Cflags, strings.Fields(os.Getenv("CFLAGS"))...)
	//cfg.Toolchain.Ldflags = append(cfg.Toolchain.Ldflags, strings.Fields(os.Getenv("LDFLAGS"))...)



	err := filepath.Walk("./src", compile)
	if err != nil {
		println("Error collecting files: %v", err)
	}




}
