# Get familiar with Ruby, Python and Go

## Variables and Data types

### Ruby

  ```
  int_num = 10
  float_num = 10.5
  name = "hello"
  num_arr = [10, 20, 30, 40, 50]

  # Accessing items in Array
  num_arr.each do |num|
    puts num
  end

  # output: print 10, 20, 30, 40, 50 with new line

  # -------------------- Hash ---------------
  fruits_hash = { apple: 10, orange: 5, banana: 20 }

  # Accessing items in Hash
  fruits_hash.each do |name, qty|
    puts "We have fruit #{name} of #{qty} quantity"
  end

  # output
  We have fruit apple of 10 quantity
  We have fruit orange of 5 quantity
  We have fruit banana of 20 quantity
  ```

### Python

  ```
  int_num = 10
  float_num = 10.5
  name = "hello"
  num_arr = [10, 20, 30, 40, 50]

  # Accessing items in Array
  for num in num_arr:
    print(num)

  # output: print 10, 20, 30, 40, 50 with new line  

  # -------------------- Dictionary ---------------
  fruits_hash = { "apple": 10, "orange": 5, "banana": 20 }

  # Accessing items in Dictionary
  for name,qty in fruits_hash.items():
    print(f"We have fruit {name} of {qty} quantity")

  # output
  We have fruit apple of 10 quantity
  We have fruit orange of 5 quantity
  We have fruit banana of 20 quantity
  ```

### Go

  ```
  var age int  // there are others like int8, int16, int32, ...
  age = 10

  or

  age := 10

  float_num := 10.5
  fmt.Println(reflect.TypeOf(float_num).String())

  var name string
	name = "foo"

  var num_arr [3]int
  num_arr[0] = 10
  num_arr[1] = 20
  num_arr[2] = 30

  # Accessing items in Array
  for _, num := range num_arr {
    fmt.Println(num)
  }  

  # output: print 10, 20, 30 with new line

  // -------------------- Maps ---------------
  fruits_hash := make(map[string]int)

  fruits_hash["apple"] = 10
  fruits_hash["orange"] = 5
  fruits_hash["banana"] = 20

  # Accessing items in Dictionary
  for name, qty := range fruits_hash {
		fmt.Println(fmt.Sprintf("We have fruit %s of %d quantity", name, qty))
	}

  # output
  We have fruit apple of 10 quantity
  We have fruit orange of 5 quantity
  We have fruit banana of 20 quantity
  ```

## Conditions

### Ruby

  ```
  num1 = 10
  num2 = 15

  if num1 < num2
    puts "#{num1} is less than #{num2}"
  elsif num1 > num2
    puts "#{num1} is greater than #{num2}"
  else
    puts "Both are equal."
  end
  ```

### Python

  ```
  num1 = 10
  num2 = 15

  if num1 < num2:
    print(f"{num1} is less than {num2}")
  elif num1 > num2:
    print(f"{num1} is greater than {num2}")
  else:
    print("Both are equal.")
  ```

### Go

  ```
  num1 := 10
  num2 := 15

  if num1 < num2 {
    fmt.Println(fmt.Sprintf("%d is less than %d", num1, num2))
  } else if num1 > num2 {
    fmt.Println(fmt.Sprintf("%d is greater than %d", num1, num2))
  } else {
    fmt.Println("Both are equal.")
  }
  ```

## Functions or Methods

### Ruby

  ```
  def add(a, b)
    a + b
  end

  puts(add(10, 15))
  ```

### Python

  ```
  def add(a, b):
    return a + b

  print(add(10, 15))
  ```

### Go

  ```
  func add(a int, b int) int {
  	return a + b
  }

  fmt.Println(add(10, 15))
  ```

## Class

### Ruby

  ```
  class Employee
    def initialize(fname, lname, total_leaves, leaves_taken)
      @first_name = fname
      @last_name = lname
      @total_leaves = total_leaves
      @leaves_taken = leaves_taken
    end

    def leaves_remaining
      puts("#{@first_name} #{@last_name} has #{leaves_left} leaves remaining")
    end

    private

    def leaves_left
      @total_leaves - @leaves_taken
    end
  end

  e = Employee.new("Foo", "Bar", 30, 15)
  e.leaves_remaining
  ```

### Python

  ```
  class Employee:
    def __init__(self, fname, lname, total_leaves, leaves_taken):
      self.first_name = fname
      self.last_name = lname
      self.total_leaves = total_leaves
      self.leaves_taken = leaves_taken

    def leaves_remaining(self):
      print(f"{self.first_name} {self.last_name} has {self.total_leaves - self.leaves_taken}")

  e = Employee("Foo", "Bar", 30, 20)
  e.leaves_remaining()
  ```

### Go

  ```
  package main

  import (
  	"fmt"
  )

  // Employee is an employee class
  type Employee struct {
  	FirstName   string
  	LastName    string
  	TotalLeaves int
  	LeavesTaken int
  }

  // LeavesRemaining gives no. of total leaves balance
  func (e Employee) LeavesRemaining() {
  	fmt.Printf("%s %s has %d leaves remaining\n",
  		e.FirstName, e.LastName, (e.TotalLeaves - e.LeavesTaken))
  }

  func main() {
  	e := Employee{
  		FirstName:   "Foo",
  		LastName:    "Bar",
  		TotalLeaves: 30,
  		LeavesTaken: 15,
  	}
  	e.LeavesRemaining()
  }
  ```
