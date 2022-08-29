# Factory pattern

- Delegating the creation of new instances of structures to a different part of the program.
- Working at the interface level instead of with concrete implementations.
- Grouping families of objects to obtain a family object creator.

## Approach

First of all, we need to create the interface that all the structs will have in common

```go
package factory

// Noiser is the interface that all the animals in this factory will share
type Noiser interface {
  // Noise will return the noise that the animal is capable of making.
  Noise() string
}
```

Next step, is to create an enum (it is not necessary, but it is recommeded) representing each
of the animals

```go
// the previous code goes here

// this is the type that represents each of the animals
type Animal int

const (
  Dog Animal = iota
  Cat
  Duck
)

type Dog struct {
  Name string
  Legs uint8
}

type Cat struct {
  Name string
  Legs uint8
}

type Duck struct {
  Name string
  Legs uint8
}

func (d *Dog) Noise() string {
  return "Guau"
}

func (c *Cat) Noise() string {
  return "Miau"
}

func (c *Duck) Noise() string {
  return "Quack"
}
```

Finally, we have to create the factory method to generate this previous structs in demand

```go
// Here the previous code goes

// ErrNoiseAnimalDoNotExists is used when an incorrect animal is passed as a parameter
var ErrNoiseAnimalDoNotExists = errors.New("noiser animal do not exists.")

// GetNoiserAnimal generates the noiser animal depending on the animal passed as a parameter
func GetNoiserAnimal(animal Animal) Noiser {
  var (
    noiser Noiser
    err error
  )

  switch animal {
  case Dog:
    noiser = Dog{}
  case Cat:
    noiser = Cat{}
  case Duck:
    noiser = Duck{}
  default:
    err = ErrNoiserAnimalDoNotExists
  }
  
  return noiser, err
}
```
