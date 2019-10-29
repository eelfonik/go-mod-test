# Go for Javascript devs
## What's (might be) new
- `slice`: a data structure built on good old fixed size `array`, allows you to do more things like `append` (you can basically see it as `array` in js)
- `pointer`: use `&` & `*` wisely, you got the memory address in your hand
- `range` & `for` loop: yep, no `map` no `filter` no `reduce`, Go is mostly imperative on this subject.
```go
vals := map[string]int{"a": 1, "b": 2}
// Hello python 👀
for key, value := range vals {
  fmt.Printf(key + ": " + strconv.Itoa(value)+ "\n")
}
// b: 2
// a: 1
```

- `struct`: you can treat it as some sort of type def, but you can `new` one and get back its pointer:

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

- `method` is a plain function with `receiver`. 
