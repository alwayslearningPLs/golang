# Singleton pattern

As a general rule, we consider using the Singleton pattern when the following applies:

- We need a single, shared value, of some particular type.
- We need to restrict object creation of some type to a single unit along the entire program.

## How to create a Singleton

### Mutex approach

1. Create a mutex to lock access to shared data (in this case the singleton struct)
2. Check if it is nil and create a new one if the previous statement is true.
3. Return the singleton you have created.

```go
package singleton

import "sync"

type singleton struct {
  // Here the values of the structure singleton
}

var (
  s *singleton
  m sync.Mutex
) 

// GetInstance is used to create a singleton struct being thread-safe
func GetInstance() *singleton {
  if s == nil {
    m.Lock()
    defer m.Unlock()
    if s == nil {
      s = new(singleton)
    }
  }
}
```

### init() approach

We can also create a init() function, which will be executed at the beginning of the application.
This function is executed one time per file per package.

```go
package singleton

type singleton struct {
  // Here the code goes
}

var s *singleton

func init() {
  s = new(singleton)
}

// GetInstance is used to just return the singleton instance
func GetInstance() *singleton {
  return s
}
```

### sync.Do approach

The last option is to use the function `Do` from the package `sync`. In this case, we ensure that the singleton
will be instantiated only one time.

```go
package singleton

import "sync"

type singleton struct {
  // Here the code goes
}

var (
  s *singleton

  sOnce sync.Once
)

// GetInstance is used to create a singleton struct being thread-safe
func GetInstance() *singleton {
  sOnce.Do(func() {
    s = new(singleton)
  })
  return s
}
```

Notice that we haven't checked if the singleton is nil or not. This is why `sync.Do` is doing the hard work for us.
The function inside `sync.Do` will be called just one time, so we really don't care if singleton is nil or not, because it will.
