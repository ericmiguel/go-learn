# Creating Go modules

The following guide will ilustrate the Go lang basics using a module creating as study case. All the content in this tutorial can be found at the [Go docs](https://golang.org/doc/tutorial/create-module).

This tutorial's sequence includes seven brief topics that each illustrate a different part of the language.

* Create a module
  * Write a small module with functions you can call from another module.
* Call your code from another module
  * Import and use your new module.
* Return and handle an error
  * Add simple error handling.
* Return a random greeting
  * Handle data in slices (Go's dynamically-sized arrays).
* Return greetings for multiple people
  * Store key/value pairs in a map.
* Add a test
  * Use Go's built-in unit testing features to test your code.
* Compile and install the application
  * Compile and install your code locally.

Go code is grouped into packages, and packages are grouped into modules. Your module specifies dependencies needed to run your code, including the Go version and the set of other modules it requires.

---

## Create and call a Go module

Start by creating a Go module. In a module, you collect one or more related packages for a discrete and useful set of functions.

1. Open a command prompt and cd to your home directory
```
cd
```
2. Create a greetings directory for your Go module source code.

```
mkdir greetings
cd greetings
```

3. Start your module using the [go mod init command](https://golang.org/ref/mod#go-mod-init).

Run the go mod init command, giving it your module path -- here, use example.com/greetings. If you publish a module, this must be a path from which your module can be downloaded by Go tools (Github, for example). That would be your code's repository.

```
go mod init example.com/greetings
```

1. In your text editor, create a file in which to write your code and call it greetings.go. The base code (with some comments) for the current step is above.

```Go
// Declare a greetings package to collect related functions.
package greetings

import "fmt"

// Function structure: Function name(parameter parameter type) return type
// Hello returns a greeting for the named person.
func Hello(name string) string {

	// In Go, the := operator is a shortcut for declaring
	// and initializing a variable in one line
	// (Go uses the value on the right to determine the variable's type).

	// Return a greeting that embeds the name in a message.
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message
}
```

5. Create a Hello module, which will use the Greetings module.

```
mkdir hello
cd hello
go mod init example.com/hello
```

The base code for the current step is above.

```Go
package main

import (
	"fmt"
	"example.com/greetings"
)

func main() {
	message := greetings.Hello("Eric")
	fmt.Println(message)
}
```

6. For production use, you’d publish the example.com/greetings module from its repository (with a module path that reflected its published location), where Go tools could find it to download it. For now, because you haven't published the module yet, you need to adapt the example.com/hello module so it can find the example.com/greetings code on your local file system. Inside our Hello module folder, run:
```
go mod edit -replace=example.com/greetings=../greetings
```

7. From the command prompt in the hello directory, run the go mod tidy command to synchronize the example.com/hello module's dependencies, adding those required by the code, but not yet tracked in the module.

```
go mod tidy
```

8. Run your code to confirm that it works
```
go run .
```
it should output 'Hi, {your name}. Welcome!'.


---

## Return and handle an error

1. In greetings/greetings.go, add the code highlighted below

There's no sense sending a greeting back if you don't know who to greet. Return an error to the caller if the name is empty. Copy the following code into greetings.go and save the file.

```Go
package greetings

import (
    "errors"
    "fmt"
)

// Hello returns a greeting for the named person.
func Hello(name string) (string, error) {
    // If no name was given, return an error with a message.
    if name == "" {
        return "", errors.New("empty name")
    }

    // If a name was received, return a value that embeds the name
    // in a greeting message.
    message := fmt.Sprintf("Hi, %v. Welcome!", name)
    return message, nil
}
```

2. In your hello/hello.go file, handle the error now returned by the Hello function, along with the non-error value.

```Go
package main

import (
    "fmt"
    "log"

    "example.com/greetings"
)

func main() {
    // Set properties of the predefined Logger, including
    // the log entry prefix and a flag to disable printing
    // the time, source file, and line number.
    log.SetPrefix("greetings: ")
    log.SetFlags(0)

    // Request a greeting message.
    message, err := greetings.Hello("")
    // If an error was returned, print it to the console and
    // exit the program.
    if err != nil {
        log.Fatal(err)
    }

    // If no error was returned, print the returned message
    // to the console.
    fmt.Println(message)
}
```

3. At the command line in the hello directory, run hello.go to confirm that the code works.

```
go run .
```

That's common error handling in Go: Return an error as a value so the caller can check for it.

---

## Return a random greeting

To do this, you'll use a Go slice. A slice is like an array, except that its size changes dynamically as you add and remove items. The slice is one of Go's most useful types.

1. In greetings/greetings.go, change your code so it looks like the following

```go
package greetings

import (
    "errors"
    "fmt"
    "math/rand"
    "time"
)

// Hello returns a greeting for the named person.
func Hello(name string) (string, error) {
    // If no name was given, return an error with a message.
    if name == "" {
        return name, errors.New("empty name")
    }
    // Create a message using a random format.
    message := fmt.Sprintf(randomFormat(), name)
    return message, nil
}

// init sets initial values for variables used in the function.
func init() {
    rand.Seed(time.Now().UnixNano())
}

// randomFormat returns one of a set of greeting messages. The returned
// message is selected at random.
func randomFormat() string {
    // A slice of message formats.
    formats := []string{
        "Hi, %v. Welcome!",
        "Great to see you, %v!",
        "Hail, %v! Well met!",
    }

    // Return a randomly selected message format by specifying
    // a random index for the slice of formats.
    return formats[rand.Intn(len(formats))]
}
```

2. In hello/hello.go, change your code so it looks like the following.

```go
package main

import (
    "fmt"
    "log"

    "example.com/greetings"
)

func main() {
    // Set properties of the predefined Logger, including
    // the log entry prefix and a flag to disable printing
    // the time, source file, and line number.
    log.SetPrefix("greetings: ")
    log.SetFlags(0)

    // Request a greeting message.
    message, err := greetings.Hello("Eric")
    // If an error was returned, print it to the console and
    // exit the program.
    if err != nil {
        log.Fatal(err)
    }

    // If no error was returned, print the returned message
    // to the console.
    fmt.Println(message)
}
```

3. At the command line, in the hello directory, run hello.go to confirm that the code works. Run it multiple times, noticing that the greeting changes.

---

## Return greetings for multiple people
   
In the last changes you'll make to your module's code, you'll add support for getting greetings for multiple people in one request. In other words, you'll handle a multiple-value input, then pair values in that input with a multiple-value output. To do this, you'll need to pass a set of names to a function that can return a greeting for each of them.

But there's a hitch. Changing the Hello function's parameter from a single name to a set of names would change the function's signature. If you had already published the example.com/greetings module and users had already written code calling Hello, that change would break their programs.

In this situation, a better choice is to write a new function with a different name. The new function will take multiple parameters. That preserves the old function for backward compatibility.

1. In greetings/greetings.go, change your code so it looks like the following.

```go
package greetings

import (
    "errors"
    "fmt"
    "math/rand"
    "time"
)

// Hello returns a greeting for the named person.
func Hello(name string) (string, error) {
    // If no name was given, return an error with a message.
    if name == "" {
        return name, errors.New("empty name")
    }
    // Create a message using a random format.
    message := fmt.Sprintf(randomFormat(), name)
    return message, nil
}

// Hellos returns a map that associates each of the named people
// with a greeting message.
func Hellos(names []string) (map[string]string, error) {
    // A map to associate names with messages.
    messages := make(map[string]string)
    // Loop through the received slice of names, calling
    // the Hello function to get a message for each name.
    for _, name := range names {
        message, err := Hello(name)
        if err != nil {
            return nil, err
        }
        // In the map, associate the retrieved message with
        // the name.
        messages[name] = message
    }
    return messages, nil
}

// Init sets initial values for variables used in the function.
func init() {
    rand.Seed(time.Now().UnixNano())
}

// randomFormat returns one of a set of greeting messages. The returned
// message is selected at random.
func randomFormat() string {
    // A slice of message formats.
    formats := []string{
        "Hi, %v. Welcome!",
        "Great to see you, %v!",
        "Hail, %v! Well met!",
    }

    // Return one of the message formats selected at random.
    return formats[rand.Intn(len(formats))]
}
```

* Add a Hellos function whose parameter is a slice of names rather than a single name. Also, you change one of its return types from a string to a map so you can return names mapped to greeting messages.
* Have the new Hellos function call the existing Hello function. This helps reduce duplication while also leaving both functions in place.
* Create a messages map to associate each of the received names (as a key) with a generated message (as a value). In Go, you initialize a map with the following syntax: make(map[key-type]value-type). You have the Hellos function return this map to the caller. For more about maps, see Go maps in action on the Go blog.
* Loop through the names your function received, checking that each has a non-empty value, then associate a message with each. In this for loop, range returns two values: the index of the current item in the loop and a copy of the item's value. You don't need the index, so you use the Go blank identifier (an underscore) to ignore it. For more, see The blank identifier in Effective Go.


2. In your hello/hello.go calling code, pass a slice of names, then print the contents of the names/messages map you get back.

```Go
package main

import (
    "fmt"
    "log"

    "example.com/greetings"
)

func main() {
    // Set properties of the predefined Logger, including
    // the log entry prefix and a flag to disable printing
    // the time, source file, and line number.
    log.SetPrefix("greetings: ")
    log.SetFlags(0)

    // A slice of names.
    names := []string{"Gladys", "Samantha", "Darrin"}

    // Request greeting messages for the names.
    messages, err := greetings.Hellos(names)
    if err != nil {
        log.Fatal(err)
    }
    // If no error was returned, print the returned map of
    // messages to the console.
    fmt.Println(messages)
}
```

3. At the command line, change to the directory that contains hello/hello.go, then use go run to confirm that the code works.

```
go run .
```

---

## Add a test

Testing your code during development can expose bugs that find their way in as you make changes. In this topic, you add a test for the Hello function.

Go's built-in support for unit testing makes it easier to test as you go. Specifically, using naming conventions, Go's testing package, and the go test command, you can quickly write and execute tests.

1. In the greetings directory, create a file called greetings_test.go. **Ending a file's name with _test.go tells the go test command that this file contains test functions.**

```Go
package greetings

import (
    "testing"
    "regexp"
)

// TestHelloName calls greetings.Hello with a name, checking
// for a valid return value.
func TestHelloName(t *testing.T) {
    name := "Gladys"
    want := regexp.MustCompile(`\b`+name+`\b`)
    msg, err := Hello("Gladys")
    if !want.MatchString(msg) || err != nil {
        t.Fatalf(`Hello("Gladys") = %q, %v, want match for %#q, nil`, msg, err, want)
    }
}

// TestHelloEmpty calls greetings.Hello with an empty string,
// checking for an error.
func TestHelloEmpty(t *testing.T) {
    msg, err := Hello("")
    if msg != "" || err == nil {
        t.Fatalf(`Hello("") = %q, %v, want "", error`, msg, err)
    }
}
```


- Implement two tests:
  - TestHelloName calls the Hello function, passing a name value with which the function should be able to return a valid response message. If the call returns an error or an unexpected response message (one that doesn't include the name you passed in), you use the t parameter's Fatalf method to print a message to the console and end execution.
  - TestHelloEmpty calls the Hello function with an empty string. This test is designed to confirm that your error handling works. If the call returns a non-empty string or no error, you use the t parameter's Fatalf method to print a message to the console and end execution.

2. At the command line in the greetings directory, run the [go test command](https://golang.org/cmd/go/#hdr-Test_packages) to execute the test. You can add the -v flag to get verbose output that lists all of the tests and their results.

```
go test -v
```

3. Break the greetings.Hello function to view a failing test.

```Go
// A broken version of previous function (for test purposes)
func BrokenHello(name string) (string, error) {
	// If no name was given, return an error with a message.
	if name == "" {
		return name, errors.New("empty name")
	}
	// Create a message using a random format.
	// message := fmt.Sprintf(randomFormat(), name)
	message := fmt.Sprint(randomFormat())
	return message, nil
}
```

and add a test for the above function

```Go
// A broken version of previous function (for test purposes)
func BrokenHello(name string) (string, error) {
	// If no name was given, return an error with a message.
	if name == "" {
		return name, errors.New("empty name")
	}
	// Create a message using a random format.
	// message := fmt.Sprintf(randomFormat(), name)
	message := fmt.Sprint(randomFormat())
	return message, nil
}
```

--- 

## Compile and install the application

This topic introduces two additional commands for building code:

- The [go build command](https://golang.org/cmd/go/#hdr-Compile_packages_and_dependencies) compiles the packages, along with their dependencies, but it doesn't install the results.
- The [go install command](https://golang.org/ref/mod#go-install) compiles and installs the packages.

1. From the command line in the hello directory, run the go build command to compile the code into an executable.

```
go build
```

2. From the command line in the hello directory, run the new hello executable to confirm that the code works.

On Linux:
```
./hello
```

3. Next, you'll install the executable so you can run it without specifying its path. Discover the Go install path, where the go command will install the current package.

```
go list -f '{{.Target}}'
```

4. Add the Go install directory to your system's shell path and run the install command

on Linux:
```
export PATH=$PATH:/path/to/your/install/directory
go install
```

Remember to remove the executable name from the path :laughing:  

You can also edit your rc:
```
nano ~/.bashrc
```

And insert at the end of the file
```
export PATH="/home/$USER/go/bin/:$PATH"
```

After, you can reload your terminal 
```
source ~/.bashrc
```

The ```hello``` command will return the names you coded.

