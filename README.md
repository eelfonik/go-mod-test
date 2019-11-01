# Go for Javascript devs
## What's (might be) new
### `slice`
A data structure built on good old fixed size `array`, allows you to do more things like `append` (you can basically see it as `array` in js).

### `pointer`
Use `&` & `*` wisely, you got the memory address in your hand

### `range` & `for` loop
Nope, no `map` no `filter` no `reduce`, Go is mostly imperative on this subject.

  ```go
  vals := map[string]int{"a": 1, "b": 2}
  // Hello python 👀
  for key, value := range vals {
    fmt.Printf(key + ": " + strconv.Itoa(value)+ "\n")
  }
  // b: 2
  // a: 1
  ```

### `struct`
You can treat it as some sort of type def, but you can `new` one and get back its pointer:

  ```go
  type Woo struct {
    sound string
    animalId int
  }

  func main() {
    wolf := new(Woo) // it's a pointer

    // pointers are automatically dereferenced. so we can directly use dot notation

    wolf.sound = "woo"
    wolf.animalId = 1
    println(wolf) // 0x41a784
    fmt.Println(wolf) // &{woo 1}

    println(wolf.sound) // woo 

    realWolf := Woo{"woo", 1}
    fmt.Println(realWolf) // {woo 1}
  }
  ```

### `method`
It's a plain function with `receiver`. 

- `receiver` can be either a pointer or a value. But for any build-in type already has a pointer pointed to the underlying data (ex: `map`, `slice`, `array`), normally you don't need a pointer receiver.
 
- `receiver` could not be basic or unamed type, it allows you want to write things like `p.isAllowed()` instead of `isAllowed(p)`, consider the following:

  ```go
  type conf struct {
    isDog bool
  }

  //function
  func isAllowed(animal conf) bool {
    return animal.isDog
  }

  //method
  func (animal conf) isAllowed() bool {
    return animal.isDog
  }

  // if you want to mutate the struct, use a pointer receiver
  func (animal *conf) changeNature() {
    animal.isDog = false
  }

  // usage
  func main() {
    dog := conf{isDog: true}
    // function call
    fmt.Println(isAllowed(dog)) //true
    // method call
    fmt.Println(dog.isAllowed()) //true
    
    // we changed the underlying data struct value
    dog.changeNature()
    fmt.Println(dog.isDog) //false
    fmt.Println(dog.isAllowed()) //false
  }
  ```

### `interface`
Nope, it's not the `interface` you know in typescript. It's a type defined use the signatures of some `method`s.  If you try to define some data fields inside interface, there will errors (but if you really want, you can do some trick with `embedded struct` like this one : https://stackoverflow.com/questions/26027350/go-interface-fields).

Any type that implements all its methods is an implementation of this `interface`. Go prefers tiny interface. It helps if you want to do some kind of `polymorphism`.







