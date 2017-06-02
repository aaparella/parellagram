Title: The Power of Ad-Hoc Interfaces
Date: June 1, 2017

Much like many other langugaes, Go has support for interfaces. They define a set of behaviors and allow us to not worry about differing implementation details behind those behaviors.

```
type UserSource struct {
    FetchUser(int) User    
    CreateUser(...) error
    ...
}

func emailUser(id int, source UserSource) {
    user, err := source.FetchUser(id)
    ...
}
```

In the above example we abstract away the source of user information. It could be a database, an API call, a local file on disk, doesn't matter. All we need to know is that source is a struct that implements the methods required of it in the `UserSource` interface.

*Unlike* most other languages, Go does not require a struct to explicitly declare that it implements an interface. This means that if we define an interface, any struct that has the appropriate methods may automatically be used anywhere that that interface is accepted.

```
type Stringer interface {
    ToString() string
}

type User struct {
    Email, First, Last string
}

type (u *User) ToString() string {
    return fmt.Sprintf("%s %s, %s", u.First, u.Last, u.Email)
}
```

We can now use the `User` struct anywhere that a `Stringer` is accepted, even though we never explicitly declared that `User` implements the `Stringer` interface.

```
func Display(s Stringer) {
    fmt.Printf(s.ToString())
}

u := &User{
    First: "Alex", 
    Last: "Parella",
    Email: "aaparella@gmail.com",
}
Display(u) // "Alex Parella, aaparella@gmail.com"
```

The obvious benefit of this is that it helps to reduce noise. As Go is heavily oriented around using lots of small interfaces, having each implementing struct declare that it implements each interface explicitly would get very old, very fast. Another, less obvious benefit is that it allows us to define interfaces that are implemented by structs *that we did not write*.

Take this example, adapted from a Java project I worked on recently at work.

In this function, a number of objects are used to create a list of search criterion that are used to search for an item.

```
List<String> criteria = new ArrayList<>();

// One of many
YearRange yearRange = details.getYearRange(); // Years manufactured
if (yearRange != null) {
    criteria.add(rangeCriteria(yearRange.getMin(), yearRange.getMax()));
}
```

If there are many different criteria (in the example this is being drawn from, there were twenty four) this blows up very quickly. We would like to be able to make this behavior generic.

```
public String rangeCriteria(List<String> criteria, Range range) { 
    if (range != null) {
        crieria.add(String.format(...));
    }
}

// Could also be defined with a closure to avoid passing criteria as an argument
rangeCriteria(criteria, yearRange);
```

The issue with this, in our particular case, was that the different range objects did not share a superclass, nor did they all implement the same interface that exposed this behavior, and we could not modify the class definitions to implement such an interface. Unfortunately this meant we could not easily create a generic method without having to create a wrapper function that accepted each type of object separately. This sort of situation is where the ability of structs in Go to implement interfaces without explicitly declaring that they do so would come into play.

In this particular case, we could define such an interface as follows.

```
type Range interface {
    getMin() string
    getMax() string
}

func addRange(r Range) string {
    return fmt.Sprintf(...)
}
```

Because interfaces are defined implicitly, all of the objects that we need to use can be accessed through this generic interface without modification. This allows us to make a function that is generic over multiple struct types that could be from entirely different packages altogether. Were the structs required to explicitly implement interfaces, this would not be possible.

I initially decided not to write this post as I thought the example situation would be too rare, but as I was just asked earlier today how to deal with this exact problem, perhaps it's not as rare as I thought.
