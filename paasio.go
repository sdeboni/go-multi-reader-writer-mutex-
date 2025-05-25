package paasio

import "io"

type readCounter struct {
  reader io.Reader
  bytesRead int64
  callCount int
}
type writeCounter struct {
  writer io.Writer
  bytesWritten int64
  callCount int
}

type readWriteCounter struct {
  rc ReadCounter
  wc WriteCounter
}

func NewWriteCounter(writer io.Writer) WriteCounter {
  return &writeCounter {
    writer: writer,
  }
}

func NewReadCounter(reader io.Reader) ReadCounter {
  return &readCounter {
    reader: reader,
  }
}

func NewReadWriteCounter(readwriter io.ReadWriter) ReadWriteCounter {
  return &readWriteCounter {
    NewReadCounter(readwriter),
    NewWriteCounter(readwriter),
  }
}

func (rc *readCounter) Read(p []byte) (int, error) {
  bytesRead, err := rc.reader.Read(p)
  rc.callCount++
  rc.bytesRead += int64(bytesRead)
  return bytesRead, err
}

func (rc *readCounter) ReadCount() (int64, int) {
  return rc.bytesRead, rc.callCount
}

func (wc *writeCounter) Write(p []byte) (int, error) {
  bytesWritten, err := wc.writer.Write(p)
  wc.callCount++
  wc.bytesWritten += int64(bytesWritten)
  return bytesWritten, err
}

func (wc *writeCounter) WriteCount() (int64, int) {
  return wc.bytesWritten, wc.callCount
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
