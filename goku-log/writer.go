package goku_log

import (
	"bufio"
	"bytes"
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"
)

//MaxBufferd buffer最大值
const MaxBufferd = 1024 * 500

var (
	bufferPool = &sync.Pool{
		New: func() interface{} {
			return new(bytes.Buffer)
		},
	}
)

//FileWriterByPeriod 文件周期写入
type FileWriterByPeriod struct {
	wC         chan *bytes.Buffer
	dir        string
	file       string
	period     LogPeriod
	enable     bool
	cancelFunc context.CancelFunc
	locker     sync.Mutex
	wg         sync.WaitGroup
	expire     time.Duration
}

//NewFileWriteBytePeriod 获取新的FileWriterByPeriod
func NewFileWriteBytePeriod() *FileWriterByPeriod {
	w := &FileWriterByPeriod{
		locker: sync.Mutex{},
		wg:     sync.WaitGroup{},
		enable: false,
	}

	return w
}
func (w *FileWriterByPeriod) getExpire() time.Duration {
	w.locker.Lock()
	expire := w.expire
	w.locker.Unlock()
	return expire
}

//Set 设置
func (w *FileWriterByPeriod) Set(dir, file string, period LogPeriod, expire time.Duration) {
	fileName := strings.TrimSuffix(file, ".log")

	w.locker.Lock()
	w.file = fileName
	w.dir = dir
	w.period = period
	w.expire = expire
	w.locker.Unlock()
}

//Open 打开
func (w *FileWriterByPeriod) Open() {
	w.locker.Lock()
	defer w.locker.Unlock()

	if w.enable {
		return
	}

	ctx, cancel := context.WithCancel(context.Background())
	w.cancelFunc = cancel
	w.wC = make(chan *bytes.Buffer, 100)
	w.wg.Add(1)
	w.enable = true
	go w.do(ctx)
}

//Close 关闭
func (w *FileWriterByPeriod) Close() {

	isClose := false
	w.locker.Lock()
	if !w.enable {
		w.locker.Unlock()
		return
	}

	if w.cancelFunc != nil {
		isClose = true
		w.cancelFunc()
		w.cancelFunc = nil
	}
	w.enable = false
	w.locker.Unlock()
	if isClose {
		w.wg.Wait()
	}

}
func (w *FileWriterByPeriod) Write(p []byte) (n int, err error) {

	l := len(p)
	if !w.enable {
		return l, nil
	}
	buffer := bufferPool.Get().(*bytes.Buffer)
	buffer.Reset()
	buffer.Write(p)

	w.wC <- buffer
	return l, nil
}

func (w *FileWriterByPeriod) do(ctx context.Context) {
	w.initFile()
	f, lastTag, e := w.openFile()
	if e != nil {
		fmt.Printf("open log file:%s\n", e.Error())
		return
	}

	buf := bufio.NewWriter(f)
	t := time.NewTicker(time.Second * 5)
	defer t.Stop()
	tflusth := time.NewTimer(time.Second)

	for {
		select {
		case <-ctx.Done():
			{
				for len(w.wC) > 0 {
					p := <-w.wC
					buf.Write(p.Bytes())
					bufferPool.Put(p)
				}
				buf.Flush()
				f.Close()
				t.Stop()
				w.wg.Done()
				return
			}

		case <-t.C:
			{
				if buf.Buffered() > 0 {
					buf.Flush()
					tflusth.Reset(time.Second)
				}
				if lastTag != w.timeTag(time.Now()) {

					f.Close()
					w.history(lastTag)
					fnew, tag, err := w.openFile()
					if err != nil {
						return
					}
					lastTag = tag
					f = fnew
					buf.Reset(f)

					go w.dropHistory()
				}

			}
		case <-tflusth.C:
			{
				if buf.Buffered() > 0 {
					buf.Flush()
				}
				tflusth.Reset(time.Second)
			}
		case p := <-w.wC:
			{
				buf.Write(p.Bytes())
				bufferPool.Put(p)
				if buf.Buffered() > MaxBufferd {
					buf.Flush()
				}
				tflusth.Reset(time.Second)
			}
		}
	}
}
func (w *FileWriterByPeriod) timeTag(t time.Time) string {

	w.locker.Lock()
	tag := t.Format(w.period.FormatLayout())
	w.locker.Unlock()
	return tag
}
func (w *FileWriterByPeriod) history(tag string) {

	path := filepath.Join(w.dir, fmt.Sprintf("%s.log", w.file))
	histroy := filepath.Join(w.dir, fmt.Sprintf("%s-%s.log", w.file, tag))
	_ = os.Rename(path, histroy)

}
func (w *FileWriterByPeriod) dropHistory() {
	expire := w.getExpire()
	expireTime := time.Now().Add(-expire)
	pathPatten := filepath.Join(w.dir, fmt.Sprintf("%s-*", w.file))
	files, err := filepath.Glob(pathPatten)
	if err == nil {
		for _, f := range files {
			if info, e := os.Stat(f); e == nil {

				if expireTime.After(info.ModTime()) {
					_ = os.Remove(f)
				}
			}
		}
	}
}
func (w *FileWriterByPeriod) initFile() {
	_ = os.MkdirAll(w.dir, os.ModePerm)
	path := filepath.Join(w.dir, fmt.Sprintf("%s.log", w.file))
	nowTag := w.timeTag(time.Now())
	if info, e := os.Stat(path); e == nil {

		timeTag := w.timeTag(info.ModTime())
		if timeTag != nowTag {
			w.history(timeTag)
		}
	}

	w.dropHistory()

}

func (w *FileWriterByPeriod) openFile() (*os.File, string, error) {
	path := filepath.Join(w.dir, fmt.Sprintf("%s.log", w.file))
	nowTag := w.timeTag(time.Now())
	f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		return nil, "", err
	}
	return f, nowTag, err

}
