package main

import (
	_ "embed"
	"flag"
	"fmt"
	"log"
	"os"
	"path"

	. "github.com/quentinlintz/cmdtop/config"
)

var (
	//go:embed LICENSE
	licenseText string
	version     = "<unknown>"
	usage       = `Usage: %s [options]
Print top commands used from your shell history

Options:
`
)

func main() {
	cfg := Config{}

	flag.Usage = func() {
		name := path.Base(os.Args[0])
		fmt.Fprintf(os.Stderr, usage, name)
		flag.PrintDefaults()
	}
	flag.IntVar(&cfg.Top, "top", 5, "how many top commands to return")
	flag.BoolVar(&cfg.ShowLicense, "license", false, "print license and exit")
	flag.BoolVar(&cfg.ShowVersion, "version", false, "print version and exit")
	flag.Parse()

	if cfg.ShowLicense {
		fmt.Print(licenseText)
		os.Exit(0)
	}

	if cfg.ShowVersion {
		name := path.Base(os.Args[0])
		fmt.Printf("%s version %s\n", name, version)
		os.Exit(0)
	}

	if err := ParseEnv(&cfg); err != nil {
		log.Fatalf("error parsing environment: %s", err)
	}

	if err := ValidateConfig(&cfg); err != nil {
		log.Fatalf("error validating config: %s", err)
	}

	fmt.Printf("Config: %+v\n", cfg)
}
