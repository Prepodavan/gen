package typewriter

import "os"

type Config struct {
	Filter                func(os.FileInfo) bool
	IgnoreTypeCheckErrors bool
	RequestedPackages []string
	RequestedPackagesPaths map[string]string
}

var DefaultConfig = &Config{}
