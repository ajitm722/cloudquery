package utilities

import (
	"os"
	"path/filepath"
	"strings"
)

func xdgDefault(envVar, fallback string) string {
	if v := os.Getenv(envVar); v != "" {
		return v
	}
	return fallback
}

func ExpandPath(path string) string {
	if path == "" {
		return path
	}
	if strings.HasPrefix(path, "~/") {
		if home, err := os.UserHomeDir(); err == nil {
			path = filepath.Join(home, path[2:])
		}
	}

	mapping := map[string]string{
		"XDG_STATE_HOME":  xdgDefault("XDG_STATE_HOME", filepath.Join(os.Getenv("HOME"), ".local", "state")),
		"XDG_DATA_HOME":   xdgDefault("XDG_DATA_HOME", filepath.Join(os.Getenv("HOME"), ".local", "share")),
		"XDG_CONFIG_HOME": xdgDefault("XDG_CONFIG_HOME", filepath.Join(os.Getenv("HOME"), ".config")),
		"XDG_CACHE_HOME":  xdgDefault("XDG_CACHE_HOME", filepath.Join(os.Getenv("HOME"), ".cache")),
	}

	expanded := os.Expand(path, func(key string) string {
		if val, ok := mapping[key]; ok {
			return val
		}
		return os.Getenv(key)
	})

	return expanded
}
