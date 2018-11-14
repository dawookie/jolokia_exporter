// Copyright © 2017 Alexander Pinnecke <alexander.pinnecke@googlemail.com>

package main

import "github.com/jaxxstorm/jolokia_exporter/cmd"

var version = "snapshot"

func main() {
	cmd.Execute(version)
}
