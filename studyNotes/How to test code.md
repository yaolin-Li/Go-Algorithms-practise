# How to test code

In Go project, we can import "testing" to verify code.([Go test Document](https://pkg.go.dev/testing))

- First, we add "testing" in import
- Use func TestXxx(t *testing.T) as verification function to run the code which being test.
- Example functions may include a concluding line comment that begins with "Output:" and is compared with the standard output of the function when the tests are run. (The comparison ignores leading and trailing space.)
- You can use go test -v xxx.go xxx_test.go to run those test functions. It will tell you "OK" or "FAIL".