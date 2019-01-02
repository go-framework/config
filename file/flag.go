package file

import "flag"

// Register config file flag.
func RegisterConfigFileFlag(filename string) *string {
	return flag.String("config-file", filename, "config file")
}
