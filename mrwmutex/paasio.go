package mrwmutex

import (
  "io"
  "sync/atomic"
  "paasio/paasio"
)

type readCounter struct {
  reader io.Reader
  bytesRead atomic.Int64
  callCount atomic.Int64
  mutex MultiRWMutex
}

type writeCounter struct {
  writer io.Writer
  bytesWritten atomic.Int64
  callCount atomic.Int64
  mutex MultiRWMutex
}

type readWriteCounter struct {
  rc paasio.ReadCounter
  wc paasio.WriteCounter
}

func NewWriteCounter(writer io.Writer) paasio.WriteCounter {
  return &writeCounter {
    writer: writer,
  }
}

func NewReadCounter(reader io.Reader) paasio.ReadCounter {
  return &readCounter {
    reader: reader,
  }
}

func NewReadWriteCounter(readwriter io.ReadWriter) paasio.ReadWriteCounter {
  return &readWriteCounter {
    NewReadCounter(readwriter),
    NewWriteCounter(readwriter),
  }
}

func (rc *readCounter) Read(p []byte) (int, error) {
  bytesRead, err := rc.reader.Read(p)
  rc.mutex.WLock()
  defer rc.mutex.WUnlock()

  rc.callCount.Add(1)
  rc.bytesRead.Add(int64(bytesRead))

  return bytesRead, err
}

func (rc *readCounter) ReadCount() (int64, int) {
  rc.mutex.RLock()
  defer rc.mutex.RUnlock()

  return rc.bytesRead.Load(), int(rc.callCount.Load())
}

func (wc *writeCounter) Write(p []byte) (int, error) {
  bytesWritten, err := wc.writer.Write(p)

  wc.mutex.WLock()
  defer wc.mutex.WUnlock()

  wc.callCount.Add(1)
  wc.bytesWritten.Add(int64(bytesWritten))

  return bytesWritten, err
}

func (wc *writeCounter) WriteCount() (int64, int) {
  wc.mutex.RLock()
  defer wc.mutex.RUnlock()

  return wc.bytesWritten.Load(), int(wc.callCount.Load())
}

func (rwc *readWriteCounter) Read(p []byte) (int, error) {
  return rwc.rc.Read(p)
}

func (rwc *readWriteCounter) ReadCount() (int64, int) {
  return rwc.rc.ReadCount()
}

func (rwc *readWriteCounter) Write(p []byte) (int, error) {
  return rwc.wc.Write(p)
}

func (rwc *readWriteCounter) WriteCount() (int64, int) {
  return rwc.wc.WriteCount()
}
