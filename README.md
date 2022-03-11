# Homework | Week 2

In this project, a list of books is created. We first list the list we have created and then find out whether there are books searched in the list.

## Purpose of application

We have a list of books. This application has 2 tasks.
1. To show all the books in the application as output.
2. If there is a book name given as an input, printing it on the screen, if not, printing an output stating that the book does not exist.

## Used Commands

### list command
```
go run main.go list
```
This command allows us to see the objects we define in the application as a list.

### search command 
```
go run main.go search <bookName>
go run main.go search Harry Potter
```
This command allows us to scan the objects we have defined in the application to check if they exist inside the application.

## Used Functions

### json.Unmarshal()
'''
json.Unmarshal([]byte(bookJson), &books)
'''
Golang provides multiple APIs to work with JSON including to and from built-in and custom data types using the encoding/json package. To parse JSON, we use the Unmarshal() function in package encoding/json to unpack or decode the data from JSON to a struct.

### strings.Join()
'''
strings.Join(searchArg, " ")
'''
strings.Join() Function in Golang concatenates all the elements present in the slice of string into a single string. This function is available in the string package.

### strings.EqualFold()
'''
strings.EqualFold(books[i].Names, bookName)
'''
strings.EqualFold() Function in Golang reports whether s and t, interpreted as UTF-8 strings, are equal under Unicode case-folding, which is a more general form of case-insensitivity.

## Requirements

* Go Language
* Git
* Go Module

## Acknowledgments

* Hat tip to anyone whose code was used
* Inspiration
* etc
