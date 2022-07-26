# When to use Buffered and Unbuffered channels

- By default, channels are unbuffered and they are easy to understand: one goroutine writes and waits for another goroutine to pick up its work.

- While in case of buffered channnels, you have to pick a size, since buffered channels never have unlimited buffers.

- Buffered channels are useful when you
  - know how many goroutines you have launched
  - want to limit the number of goroutines you will launch
  - want to limit the amount of work that is queued up.
