# Interfaces with Structs

In this exercise, you'll build a game store along with a set of functions to query the store.

1. Define a `Shape` interface with the following functions:
   - `Area()` calculates the shape's area and returns `float64`
   - `Perimeter()` calculates the shape's perimeter and returns `float64`

2. Declare the following structs with appropriate elements:
   - `Circle`
   - `Rectangle`

3. Write a constructor functions:
   - `NewCircle(Radius float64) Circle`
   - `NewRectangle(Width,Height float64) Rectangle`

3. Implement interface `Shape` for `Circle` and `Rectangle`

Insert your code into the file `exercise.go`.

Hint: read about how to use [interfaces](https://go.dev/tour/methods/9).
