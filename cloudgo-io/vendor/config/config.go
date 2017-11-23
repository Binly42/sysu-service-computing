package config

import (
	"convention/codec"
	"log"
	"os"
	"time"
)

var debugMode = true
var logToConsoleMode = true

func DebugMode() bool        { return debugMode }
func LogToConsoleMode() bool { return logToConsoleMode }

// type Config = map[string](interface{})

// Config holds all configure of Cloudgo system.
var Config = make(map[string](interface{}))

func Load(decoder codec.Decoder) {
	cfg := &(Config)
	// CHECK: Need check if have already exactly loaded ALL config (i.e. eof) ?
	if err := decoder.Decode(cfg); err != nil {
		log.Fatal(err)
	}
}

func Save(encoder codec.Encoder) error {
	return encoder.Encode(Config)
}

// ... paths

// WorkingDir for cloudgo.
func WorkingDir() string {
	location, existed := os.LookupEnv("HOME")
	if !existed || DebugMode() {
		location = "."
	}
	ret := location + "/.cloudgo.d/"
	return ret
}

func LogPath() string { return WorkingDir() + "cloudgo_" + time.Now().Format("20060102_15") + ".log" }

var neededFilepaths = []string{}

func NeededFilepaths() []string {
	return neededFilepaths
}

func ensurePathsNeededExist() {
	if err := os.MkdirAll(WorkingDir(), 0777); err != nil {
		log.Fatal(err)
	}

	for _, path := range NeededFilepaths() {
		if _, err := os.Stat(path); os.IsNotExist(err) {
			f, err := os.Create(path)
			defer f.Close()
			if err != nil {
				log.Fatal(err)
			}
		}
	}
}

func init() {
	ensurePathsNeededExist()
}
