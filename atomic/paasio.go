package atomic

import (
  "io"
  "sync/atomic"
  "time"
  "math/rand"
  "math"
  "paasio/paasio"
)

type readCounter struct {
  reader io.Reader
  counter atomic.Uint64
}
type writeCounter struct {
  writer io.Writer
  counter atomic.Uint64
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

func (rc *readCounter) updateCounter(bytes int) {
  success := false
  patience := 1000

  for !success && patience > 0 {
    patience--

    oldCounter := rc.counter.Load()

    bytesRead := oldCounter >> 32
    callCount := oldCounter & uint64(math.MaxUint32)

    bytesRead += uint64(bytes)

    if (bytesRead > math.MaxUint32) {
      panic("overflow")
    }
    callCount++

    newCounter := (bytesRead << 32) + callCount

    success = rc.counter.CompareAndSwap(oldCounter, newCounter) 
    if !success {
      n := rand.Intn(10)
      time.Sleep(time.Duration(n)*time.Nanosecond)
    }
  } 
  if patience == 0 {
    panic("too many update conflicts")
  }
}

func (rc *readCounter) Read(p []byte) (int, error) {
  bytesRead, err := rc.reader.Read(p)

  rc.updateCounter(bytesRead) 

  return bytesRead, err
}

func (rc *readCounter) ReadCount() (int64, int) {
  counter := rc.counter.Load();

  bytesRead := int64(counter >> 32)
  callCount := int(counter & math.MaxUint32)

  return bytesRead, callCount
}


func (wc *writeCounter) updateCounter(bytes int) {
  success := false
  patience := 1000

  for !success && patience > 0 {
    patience--

    oldCounter := wc.counter.Load()

    bytesWritten := oldCounter >> 32
    callCount := oldCounter & uint64(math.MaxUint32)

    bytesWritten += uint64(bytes)

    if (bytesWritten > math.MaxUint32) {
      panic("overflow")
    }
    callCount++

    newCounter := (bytesWritten << 32) + callCount

    success = wc.counter.CompareAndSwap(oldCounter, newCounter) 
    if !success {
      n := rand.Intn(10)
      time.Sleep(time.Duration(n)*time.Nanosecond)
    }
  } 
  if patience == 0 {
    panic("too many update conflicts")
  }
}

func (wc *writeCounter) Write(p []byte) (int, error) {
  bytesWritten, err := wc.writer.Write(p)

  wc.updateCounter(bytesWritten)

  return bytesWritten, err
}

func (wc *writeCounter) WriteCount() (int64, int) {
  counter := wc.counter.Load();

  bytesWritten := int64(counter >> 32)
  callCount := int(counter & math.MaxUint32)

  return bytesWritten, callCount
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
