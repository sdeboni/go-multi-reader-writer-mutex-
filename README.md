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
ReadMutex-16            	     382	   3148276 ns/op
ReadRWMutex-16          	     338	   3504227 ns/op
ReadAtomic-16           	     396	   2975656 ns/op
ReadMultiRWMutex-16     	     402	   2942809 ns/op

WriteMutex-16           	     367	   3218048 ns/op
WriteRWMutex-16         	     337	   3493808 ns/op
WriteAtomic-16          	     387	   3035166 ns/op
WriteMultiRWMutex-16    	     392	   2997623 ns/op
