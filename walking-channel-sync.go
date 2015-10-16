package main

import (
        "fmt"
        "math/rand"
        "time"
    )

/*
Every morning, Alice and Bob go for a walk, and being creatures of habit, they follow the same routine every day.

First, they both prepare, grabbing sunglasses, perhaps a belt, closing open windows, turning off ceiling fans, and pocketing their phones and keys.

Once they’re both ready, which typically takes each of them between 60 and 90 seconds, they arm the alarm, which has a 60 second delay.

While the alarm is counting down, they both put on their shoes, a process which tends to take each of them between 35 and 45 seconds.

Then they leave the house together and lock the door, before the alarm has finished its countdown.

Write a program to simulate Alice and Bob’s morning routine.

Here’s some sample output from running a solution to this problem.
*/


type Task struct {
    name string
    task string
    minTime int
    maxTime int
}

func init() {
    rand.Seed(time.Now().UTC().UnixNano())
}

func delayed_task(c chan string, done chan bool, t Task) {
    start := time.Now()
    res := t.name + " started " + t.task
    c <- res
    time.Sleep(time.Duration(t.minTime) * time.Second)
    time.Sleep(time.Duration(rand.Intn(t.maxTime - t.minTime)) * time.Second)
    res = t.name + " finished " + t.task + " - " + time.Duration(time.Since(start)).String()
    c <- res
    done <- true
}


func main() {
    tasks := make(chan string)
    done := make(chan bool)
    
    fmt.Println("Let's go for a walk!")
    go delayed_task(tasks, done, Task{"Bob", "getting ready", 5, 10})
    go delayed_task(tasks, done, Task{"Alice", "getting ready", 5, 10})

    go func() {
        for i := range tasks {
            fmt.Println(i)
        }
    }()
    
    for i := 0; i < 2; i++ {
        <-done
    }    
    
    go delayed_task(tasks, done, Task{"Alarm", "counting down", 20, 21})
    
    go delayed_task(tasks, done, Task{"Alice", "putting on shoes", 15, 18})
    go delayed_task(tasks, done, Task{"Bob", "putting on shoes", 15, 18})
    
    for i := 0; i < 3; i++ {
        <-done
    }    
    
    fmt.Println("Alice and Bob Left!")
}