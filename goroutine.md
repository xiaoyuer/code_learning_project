# goroutine

Coroutines are simply concurrent subroutines \(functions, closures, or methods in Go\) that are _nonpreemptive_—that is, they cannot be interrupted. Instead, coroutines have multiple points throughout which allow for suspension or reentry.

非抢占式

Goroutines don’t define their own suspension or reentry points;

Go’s runtime observes the runtime behavior of goroutines and automatically suspends them when they block and then resumes them when they become unblocked.





Coroutines, and thus goroutines, are implicitly concurrent constructs, but concurrency is not a property _of_ a coroutine: something must host several coroutines simultaneously and give each an opportunity to execute—otherwise, they wouldn’t be concurrent! 



Go’s mechanism for hosting goroutines is an implementation of what’s called an _M:N scheduler_, which means it maps `M` green threads to `N` OS threads.



