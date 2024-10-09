package config

import "github.com/fsnotify/fsnotify"

type Reloader struct {
	configPath string
	fsnotify.Watcher
}

func (r *Reloader) Reload() error {
	return nil
}
