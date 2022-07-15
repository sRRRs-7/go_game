package load

import "flag"

var (
	ConfigFile = flag.String("config-file", "config.json", "path to custom configuration file")
	TextFile = flag.String("field-file", "field.txt", "path to custom text file")
)