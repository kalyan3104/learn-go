### Channel
* A type in Go for transferring data between goroutines
* One or more goroutines write to a channel
* One or more goroutines read back from channel

* Data on channel are typed
* By default channel reads and writes are synchronous.
* Putting a value in channel is just like passing a value to a function
* __Be careful when passing reference types over a channel__

* Zero value for channel is nil

||unbuffered|buffered|nil| closed  |
|---|---|---|---|---|
|__read__|Pause until something is written|Pause if buffer is empty|Hang Forever|Returns immediately with zero value (use comma ok to see if it is closed)|
|__write__|Pause until something is read|Pause only if buffer is full|Hang Forever|PANIC|
|__close__|Works|Works|PANIC|PANIC|