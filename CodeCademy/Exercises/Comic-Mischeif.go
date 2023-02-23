/*
Comic Mischief
Use variables and types in the Go programming language to maintain the catalog of items in your comic book store Comic Mischief. */

package main

import "fmt"

func main() {
	var publisher, writer, title, finalThoughts, artist, paperquality string
	var year, pageNumber, costyear, costPage, cost int
	//The Above Variabled defined as "int" can be replaced with "float32" for more accurate costs.
	var grade, maxGrade float32

	fmt.Println("")

	title = "Mr. GoToSleep"
	writer = "Tracey Hatchet"
	artist = "Jewel Tampson"
	publisher = "DizzyBooks Publishing Inc."
	year = 1997
	pageNumber = 14
	grade = 6.5
	maxGrade = 10.0
	finalThoughts = "This book was fantastic to Read and also I like the author a lot. But, I gave it a low rating due to the book being quite old"
	paperquality = "HIGH"

	costyear = year / 1000
	costPage = pageNumber / 10
	cost = costyear + costPage

	if paperquality == "HIGH" {
		cost = cost + 6
	} else if paperquality == "MEDIUM" {
		cost = cost / 2
	} else if paperquality == "LOW" {
		cost = cost / 10
	}

	fmt.Println("This is a rating of a book called:", title, ". It was written by", writer, "and its publisher is", publisher, "And its Artist is", artist, "and its was published in the year", year, ". It has a total of", pageNumber, "Pages. I Really like this book. I will rate this book", grade, " out of", maxGrade, ". My Final Thoughts on the book are:", finalThoughts, "The Overall cost of the book was", cost, "$.")

	fmt.Println("")

	title = "Epic Vol. 1"
	writer = "Ryan N. Shawn"
	artist = "The Spiegeler-winning 'Phoebe Paperclips'"
	publisher = "DizzyBooks Publishing Inc."
	year = 2013
	pageNumber = 160
	grade = 4.8
	maxGrade = 5.0
	finalThoughts = "This Book was quite good, just a minor defect while shipping, the Hard back of the book got ruined. Honestly, It was a good book to read, And overall amazing."
	paperquality = "MEDIUM"

	costyear = year / 1000
	costPage = pageNumber / 10
	cost = costyear + costPage

	if paperquality == "HIGH" {
		cost = cost + 6
	} else if paperquality == "MEDIUM" {
		cost = cost / 2
	} else if paperquality == "LOW" {
		cost = cost / 10
	}

	fmt.Println("This is a rating of a book called:", title, ". It was written by", writer, "and its publisher is", publisher, "And its Artist is", artist, "and its was published in the year", year, ". It has a total of", pageNumber, "Pages. I Really like this book. I will rate this book", grade, " out of", maxGrade, ". My Final Thoughts on the book are:", finalThoughts, "The Overall cost of the book was", cost, "$.")

	fmt.Println("")

	title = "Harry Potter and the Goblet of Fire"
	writer = "J.K. Rowling"
	artist = "Patrick Doyle"
	publisher = "BloomsBury"
	year = 2005
	pageNumber = 617
	grade = 9.6
	maxGrade = 10.0
	finalThoughts = "This book was a really good addition to the harry potter universe."
	paperquality = "LOW"

	costyear = year / 1000
	costPage = pageNumber / 10
	cost = costyear + costPage

	if paperquality == "HIGH" {
		cost = cost + 6
	} else if paperquality == "MEDIUM" {
		cost = cost / 2
	} else if paperquality == "LOW" {
		cost = cost / 10
	}

	fmt.Println("This is a rating of a book called:", title, ". It was written by", writer, "and its publisher is", publisher, "And its Artist is", artist, "and its was published in the year", year, ". It has a total of", pageNumber, "Pages. I Really like this book. I will rate this book", grade, " out of", maxGrade, ". My Final Thoughts on the book are:", finalThoughts, "The Overall cost of the book was", cost, "$.")

	fmt.Println("")

}
