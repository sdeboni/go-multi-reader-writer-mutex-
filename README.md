# PaaS I/O

Welcome to PaaS I/O on Exercism's Go Track.
If you need help running the tests or submitting your code, check out `HELP.md`.

## Instructions

Report network IO statistics.

You are writing a [PaaS][paas], and you need a way to bill customers based on network and filesystem usage.

Create a wrapper for network connections and files that can report IO statistics.
The wrapper must report:

- The total number of bytes read/written.
- The total number of read/write operations.

[paas]: https://en.wikipedia.org/wiki/Platform_as_a_service

## Source

### Created by

- @bmatsuo

### Contributed to by

- @alebaffa
- @bitfield
- @ekingery
- @eraserix
- @ferhatelmas
- @hilary
- @kytrinyx
- @leenipper
- @petertseng
- @robphoenix
- @sebito91
- @soniakeys
- @tleen

### Based on

Brian Matsuo - https://github.com/bmatsuo


Benchmark comparison between implementations: Mutex, RWMutex, Atomic, MultiRWMutex

ReadTotalMutex-16       	     356	   3395499 ns/op
ReadTotalRWMutex-16     	     333	   3515038 ns/op
ReadTotalAtomic-16      	     385	   3027096 ns/op
ReadTotalMultiRWMutex-16   	     391	   2981718 ns/op

WriteTotalMutex-16      	     373	   3200174 ns/op
WriteTotalRWMutex-16    	     332	   3530359 ns/op
WriteTotalAtomic-16     	     394	   2979828 ns/op
WriteTotalMultiRWMutex-16 	     390	   2988764 ns/op
