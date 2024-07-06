// main.go
package main

import (
	"fmt"           // Import the fmt package for formatted I/O
	"os"            // Import the os package for operating system functions
	"os/exec"       // Import the os/exec package for running external commands
	"os/signal"     // Import the os/signal package for signal handling
	"path/filepath" // Import the path/filepath package for file path operations
	"syscall"       // Import the syscall package for system calls
	"time"          // Import the time package for time-related functions

	"github.com/fsnotify/fsnotify" // Import the fsnotify package for file system notifications
)

func main() {
    watchDir := "./" // Set watchDir variable to current directory
    go watchChanges(watchDir) // Start a goroutine to watch for changes in the directory

    // Initial run
    runServer() // Start the server

    // Keep the program running
    sigs := make(chan os.Signal, 1) // Create a channel for receiving signals
    signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM) // Notify the program to handle SIGINT and SIGTERM signals
    <-sigs // Wait for a signal
}

var cmd *exec.Cmd // Declare a global variable cmd of type *exec.Cmd

func runServer() {
    // Build the binary
    buildCmd := exec.Command("go", "build", "-o", "app", "./cmd/main.go") // Create a command to build the binary
    buildCmd.Stdout = os.Stdout // Set the stdout of the build command to os.Stdout
    buildCmd.Stderr = os.Stderr // Set the stderr of the build command to os.Stderr
    if err := buildCmd.Run(); err != nil { // Run the build command and check for errors
        fmt.Println("Build failed:", err) // Print an error message if the build fails
        return
    }

    // Run the binary
    cmd = exec.Command("./app") // Create a command to run the binary
    cmd.Stdout = os.Stdout // Set the stdout of the run command to os.Stdout
    cmd.Stderr = os.Stderr // Set the stderr of the run command to os.Stderr
    if err := cmd.Start(); err != nil { // Start running the binary and check for errors
        fmt.Println("Failed to start server:", err) // Print an error message if the server fails to start
    }
}

func stopServer() {
    if cmd != nil && cmd.Process != nil { // Check if the cmd is not nil and cmd.Process is not nil
        if err := cmd.Process.Kill(); err != nil { // Kill the server process and check for errors
            fmt.Println("Failed to stop server:", err) // Print an error message if stopping the server fails
        }
        cmd.Wait() // Wait for the server to stop
    }
}

func watchChanges(dir string) {
    watcher, err := fsnotify.NewWatcher() // Create a new filesystem watcher
    if err != nil { // Check for errors in creating the watcher
        fmt.Println("Error creating watcher:", err) // Print an error message if creating the watcher fails
        return
    }
    defer watcher.Close() // Close the watcher when the function ends

    done := make(chan bool) // Create a channel for signaling when the function is done

    go func() { // Start a goroutine
        for { // Loop indefinitely
            select { // Listen for channel events
            case event := <-watcher.Events: // Handle events from the watcher
                if event.Op&fsnotify.Write == fsnotify.Write { // Check if the event is a write operation
                    fmt.Println("Modified file:", event.Name) // Print the modified file name
                    stopServer() // Stop the server
                    time.Sleep(1 * time.Second) // Prevent rapid restarting
                    runServer() // Restart the server
                }
            case err := <-watcher.Errors: // Handle errors from the watcher
                fmt.Println("Error:", err) // Print the error
            }
        }
    }()

    err = filepath.Walk(dir, func(path string, info os.FileInfo, err error) error { // Walk through the directory
        if err != nil { // Check for errors
            return err // Return the error
        }
        if info.IsDir() { // Check if the path is a directory
            return watcher.Add(path) // Add the directory to the watcher
        }
        return nil
    })
    if err != nil { // Check for errors in adding directories
        fmt.Println("Error adding directories:", err) // Print an error message
    }


    <-done
}

