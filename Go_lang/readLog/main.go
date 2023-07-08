/*
Reading logfile as it updates
*/
package main

import (
	"bufio"
	"log"
	"os"
	"sync"
)

// SeekInfo represents arguments to `os.Seek`
type SeekInfo struct {
	Offset int64
	Whence int // os.SEEK_*
}

type logger interface {
	Fatal(v ...interface{})
	Fatalf(format string, v ...interface{})
	Fatalln(v ...interface{})
	Panic(v ...interface{})
	Panicf(format string, v ...interface{})
	Panicln(v ...interface{})
	Print(v ...interface{})
	Printf(format string, v ...interface{})
	Println(v ...interface{})
}

type Tail struct {
	Filename string
	Lines    chan *Line
	Config

	file   *os.File
	reader *bufio.Reader

	watcher watch.FileWatcher
	changes *watch.FileChanges

	tomb.Tomb // provides: Done, Kill, Dying

	lk sync.Mutex
}

// Config is used to specify how a file must be tailed.
type Config struct {
	// File-specifc
	Location  *SeekInfo // Seek to this location before tailing
	ReOpen    bool      // Reopen recreated files (tail -F)
	MustExist bool      // Fail early if the file does not exist
	Poll      bool      // Poll for file changes instead of using inotify
	Pipe      bool      // Is a named pipe (mkfifo)
	//RateLimiter *ratelimiter.LeakyBucket

	// Generic IO
	Follow      bool // Continue looking for new lines (tail -f)
	MaxLineSize int  // If non-zero, split longer lines into multiple lines

	// Logger, when nil, is set to tail.DefaultLogger
	// To disable logging: set field to tail.DiscardingLogger
	Logger logger
}

// Tail File begins tailing the file. Outputs the steam.
func TailFile(filelocation string, config Config) (*Tail, error) {
	if config.ReOpen && !config.Follow {
		util.Fatal("cannot set ReOpen without Follow.")
	}

	t := &Tail{
		Filename: filename,
		Lines:    make(chan *Line),
		Config:   config,
	}

	// when Logger was not specified in config, use default logger
	if t.Logger == nil {
		t.Logger = log.New(os.Stderr, "", log.LstdFlags)
	}

	if t.Poll {
		t.watcher = watch.NewPollingFileWatcher(filename)
	} else {
		t.watcher = watch.NewInotifyFileWatcher(filename)
	}

	if t.MustExist {
		var err error
		t.file, err = OpenFile(t.Filename)
		if err != nil {
			return nil, err
		}
	}

	go t.tailFileSync()

	return t, nil
}
