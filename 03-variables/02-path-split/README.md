# Return the file attribute from a path

Write a function to split a file system path into a directory and filename component, return the
file part and discard the other component. Use the `/` symbol as a directory
separator. If the path consists of a single filename with no directory, then return the full path.


For instance, the file part of the path `/usr/bin/go` is
`go`, and the file part of `vmlinuz` is
`vmlinuz`

Implement your code in a function called `splitPath` that has the below signature:

``` go
// splitPath returns the file component of a file path.
func splitPath(fullPath string) string {
    ...
}
```

Insert your code into the file `exercise.go` near the placeholder `// INSERT YOUR CODE HERE`.

HINT: use the [`path.Split()`](https://pkg.go.dev/path#Split) function.
