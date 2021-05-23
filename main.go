package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"parse_photo_links/app/parsing"
	"parse_photo_links/cfg"

	"gopkg.in/yaml.v2"
)

func main() {
	var (
		exitCode = 1
		cfg      cfg.Config
	)

	flg := flag.NewFlagSet(os.Args[0], flag.ContinueOnError)
	cfgFile := flg.String("c", "", "-c <path to сonfig file>")
	flg.StringVar(cfgFile, "cfg", "", "--config <path to сonfig file>")
	logFile := flg.String("l", "", "-l <path to log file>")
	flg.StringVar(logFile, "log", "", "--log <path to log file>")
	helpFlag := flg.Bool("h", false, "help flag usage")
	flg.BoolVar(helpFlag, "help", false, "help flag usage")
	flg.Parse(os.Args[1:])

	// exitCode = 1
	// load a log file
	if *logFile != "" {
		log.Printf("Log file is: %s", *logFile)
		lf, err := os.OpenFile(*logFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0660)
		if err != nil {
			log.Printf("Error. Cannot open logfile: %s", err)
			_, usage := flag.UnquoteUsage(flg.Lookup("o"))
			log.Printf("Usage: %v", usage)
			os.Exit(exitCode)
		}
		defer lf.Close()

		log.SetOutput(lf)
	}

	exitCode++

	// exitCode = 2
	// works when helps flag is activated
	if *helpFlag {
		flg.PrintDefaults()
		os.Exit(exitCode)
	}

	exitCode++

	//exitCode = 3
	// load config file
	if err := loadCfg(cfgFile, &cfg); err != nil {
		log.Printf("Config file unmarshal error: %s", err)
		_, usage := flag.UnquoteUsage(flg.Lookup("c"))
		log.Printf("Usage: %v", usage)
		os.Exit(exitCode)
	}
	log.Printf("Config: %v", cfg)

	exitCode++

	// TODO: Тут сделать майн функцию

	// fmt.Println()
	PagesContent, err := parsing.ParseAll(&cfg)
	fmt.Println("error from main:", err)

	fmt.Sprintln(PagesContent)
}

// loadCfg - open config file and put config to cfg.Config struct
func loadCfg(path *string, cfg *cfg.Config) error {
	log.Printf("Start loading config")
	defer log.Printf("Config loaded")

	cfgPath := *path
	log.Printf("Config file: %s", cfgPath)
	cfgData, err := ioutil.ReadFile(cfgPath)
	if err != nil {
		return err
	}

	err = yaml.Unmarshal(cfgData, &cfg)
	if err != nil {
		return err
	}
	return nil
}
