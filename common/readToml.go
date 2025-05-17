package common



import (
	"log"

	"github.com/BurntSushi/toml"
)

// --------------------------------------------------------------------
// function reads the constants from the config.toml file
// --------------------------------------------------------------------
func ReadTomlConfig(pFilename string) interface{} {
	var f interface{}
	if _, err := toml.DecodeFile(pFilename, &f); err != nil {
		log.Println(err)
	}
	return f
}