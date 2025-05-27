package value

import (
  "io"
  "sync/atomic"
  "paasio/paasio"
)

type stats struct {
  byteCount int64
  callCount int
}

type readCounter struct {
  reader io.Reader
  stats atomic.Value
}
type writeCounter struct {
  writer io.Writer
  stats atomic.Value
}

type readWriteCounter struct {
  rc paasio.ReadCounter
  wc paasio.WriteCounter
}

func NewWriteCounter(writer io.Writer) paasio.WriteCounter {
  w := writeCounter {
    writer: writer,
  }
  w.stats.Store(&stats{0, 0})
  return &w
}

func NewReadCounter(reader io.Reader) paasio.ReadCounter {
  r := readCounter {
    reader: reader,
  }
  r.stats.Store(&stats{0, 0})
  return &r
}

func NewReadWriteCounter(readwriter io.ReadWriter) paasio.ReadWriteCounter {
  return &readWriteCounter {
    NewReadCounter(readwriter),
    NewWriteCounter(readwriter),
  }
}

func (rc *readCounter) Read(p []byte) (int, error) {
  bytesRead, err := rc.reader.Read(p)

  success := false;

  for !success {
    oldStats := rc.stats.Load().(*stats)
    newStats := stats { 
      byteCount: oldStats.byteCount + int64(bytesRead),
      callCount: oldStats.callCount + 1,
    }
    success = rc.stats.CompareAndSwap(oldStats, &newStats)
  }

  return bytesRead, err
}

func (rc *readCounter) ReadCount() (int64, int) {
  stats := rc.stats.Load().(*stats)
  return stats.byteCount, stats.callCount
}

func (wc *writeCounter) Write(p []byte) (int, error) {
  bytesWritten, err := wc.writer.Write(p)
  
  success := false
  for !success {
    oldStats := wc.stats.Load().(*stats)
    newStats := stats {
      byteCount: oldStats.byteCount + int64(bytesWritten),
      callCount: oldStats.callCount + 1,
    }
    success = wc.stats.CompareAndSwap(oldStats, &newStats)
  }

  return bytesWritten, err
}

func (wc *writeCounter) WriteCount() (int64, int) {
  stats := wc.stats.Load().(*stats);
  return stats.byteCount, stats.callCount
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
