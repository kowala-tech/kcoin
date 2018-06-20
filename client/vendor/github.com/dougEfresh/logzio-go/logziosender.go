// Copyright © 2017 Douglas Chimento <dchimento@gmail.com>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package logzio

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/beeker1121/goque"
	"go.uber.org/atomic"
)

const (
	maxSize              = 3 * 1024 * 1024 // 3 mb
	defaultHost          = "https://listener.logz.io:8071"
	defaultDrainDuration = 5 * time.Second
)

var defaultBuffDir string
var newLine = byte(10)

// LogzioSender instance of the
type LogzioSender struct {
	queue         *goque.Queue
	drainDuration time.Duration
	buf           *bytes.Buffer
	draining      atomic.Bool
	mux           sync.Mutex
	token         string
	url           string
	debug         io.Writer
	tls           bool
}

// SenderOptionFunc options for logz
type SenderOptionFunc func(*LogzioSender) error

// New creates a new Logzio sender with a token and options
func New(token string, options ...SenderOptionFunc) (*LogzioSender, error) {
	l := &LogzioSender{
		buf:           bytes.NewBuffer(make([]byte, maxSize)),
		drainDuration: defaultDrainDuration,
		url:           fmt.Sprintf("%s/?token=%s", defaultHost, token),
		token:         token,
	}
	q, err := goque.OpenQueue(fmt.Sprintf("%s%s%s%s%d", os.TempDir(), string(os.PathSeparator), "logzio-buffer", string(os.PathSeparator), time.Now().UnixNano()))
	if err != nil {
		return nil, err
	}
	l.queue = q
	for _, option := range options {
		if err := option(l); err != nil {
			return nil, err
		}
	}
	go l.start()
	return l, nil
}

// SetTempDirectory Use this temporary dir
func SetTempDirectory(dir string) SenderOptionFunc {
	return func(l *LogzioSender) error {
		l.queue.Drop()
		q, err := goque.OpenQueue(dir)
		if err != nil {
			return err
		}
		l.queue = q
		return nil
	}
}

// SetUrl set the url which maybe different from the defaultUrl
func SetUrl(url string) SenderOptionFunc {
	return func(l *LogzioSender) error {
		l.url = fmt.Sprintf("%s/?token=%s", url, l.token)
		l.debugLog("logziosender.go: Setting url to %s\n", l.url)
		return nil
	}
}

// SetDebug mode and send logs to this writer
func SetDebug(debug io.Writer) SenderOptionFunc {
	return func(l *LogzioSender) error {
		l.debug = debug
		return nil
	}
}

// SetDrainDuration to change the interval between drains
func SetDrainDuration(duration time.Duration) SenderOptionFunc {
	return func(l *LogzioSender) error {
		l.drainDuration = duration
		return nil
	}
}

// Send the payload to logz.io
func (l *LogzioSender) Send(payload []byte) error {
	_, err := l.queue.Enqueue(payload)
	return err
}

func (l *LogzioSender) start() {
	l.drainTimer()
}

// Stop will close the LevelDB queue and do a final drain
func (l *LogzioSender) Stop() {
	defer l.queue.Close()
	l.Drain()
}

func (l *LogzioSender) drainTimer() {
	for {
		time.Sleep(l.drainDuration)
		l.Drain()
	}
}

// Drain - Send remaining logs
func (l *LogzioSender) Drain() {
	l.mux.Lock()
	if l.draining.Load() {
		l.mux.Unlock()
		l.debugLog("logziosender.go: Already draining\n")
		return
	}
	l.debugLog("logziosender.go: draining queue\n")
	l.mux.Unlock()
	l.draining.Toggle()
	defer l.draining.Toggle()
	var (
		err     error
		item    *goque.Item
		bufSize int
	)
	l.buf.Reset()
	for bufSize < maxSize && err == nil {
		item, err = l.queue.Peek()
		if item != nil {
			// NewLine is appended tp item.Value
			if len(item.Value)+bufSize+1 > maxSize {
				break
			}
			bufSize += len(item.Value)
			l.debugLog("logziosender.go: Adding item %d with size %d (total buffSize: %d)\n", item.ID, len(item.Value), bufSize)
			item, _ = l.queue.Dequeue()
			l.buf.Write(append(item.Value, newLine))
		}
	}
	if bufSize > 0 {
		//l.debugLog("logziosender.go: Sending %s (%d) to %s\n", l.buf.String(), l.buf.Len(), l.url)
		resp, err := http.Post(l.url, "text/plain", l.buf)
		if err != nil {
			l.debugLog("logziosender.go: Error sending logs to %s\n", l.url)
			l.requeue()
			return
		}
		if resp.StatusCode == http.StatusOK {
			l.debugLog("logziosender.go: Accepted payload\n")
			return
		}
		if resp.StatusCode == http.StatusUnauthorized {
			l.errorLog("logziosender.go: Unauthorized access to %s\n", l.url)
			l.requeue()
			return
		}
		b, _ := ioutil.ReadAll(resp.Body)
		l.errorLog("logziosender.go: Error sending %s to %s\ncode: %s response:%s\n", l.buf.String(), l.url, resp.StatusCode, string(b))
	}
}

func (l *LogzioSender) Sync() error {
	l.Drain()
	return nil
}

func (l *LogzioSender) requeue() {
	l.debugLog("logziosender.go: Requeue %s", l.buf.String())
	l.Send(l.buf.Bytes())
}

func (l *LogzioSender) debugLog(format string, a ...interface{}) {
	if l.debug != nil {
		fmt.Fprintf(l.debug, format, a...)
	}
}

func (l *LogzioSender) errorLog(format string, a ...interface{}) {
	fmt.Fprintf(os.Stderr, format, a...)
}

func (l *LogzioSender) Write(p []byte) (n int, err error) {
	return len(p), l.Send(p)
}
