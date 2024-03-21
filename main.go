package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"math/rand"
	"os"
	"strconv"
)

type Question struct {
	Question string
	Answer   int
}

type QuestionList struct {
	Questions []Question
}

func (q *Question) CreateQuestion(question string, answer int) {
	q.Question = question
	q.Answer = answer
}

// picks question at random returns true if answer is correct
func (q *QuestionList) Ask() bool {
	idx := rand.Intn(len(q.Questions))
	question := q.Questions[idx].Question
	fmt.Printf("%s = ", question)
	userResponse := getUserInput()
	return q.Questions[idx].CheckAnswer(userResponse)
}

// gets user input and returns it as an int
func getUserInput() int {
	var answer int
	_, err := fmt.Scanf("%d", &answer)
	if err != nil {
		log.Fatal(err)
	}
	return answer
}

// create question list from csv
func (q *Question) CreateQuestionList() QuestionList {
	var questionList QuestionList
	var question Question

	//open file and assign to variable
	file, err := os.Open("data.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	//create file reader
	r := csv.NewReader(file)

	//skip header
	_, err = r.Read()
	if err != nil {
		log.Fatal(err)
	}

	//create question list slice
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		answer, err := strconv.Atoi(record[1])
		if err != nil {
			log.Fatal(err)
		}
		question.CreateQuestion(record[0], answer)
		questionList.Questions = append(questionList.Questions, question)
	}
	return questionList
}

// compares csv and user input
func (q *Question) CheckAnswer(answer int) bool {
	return q.Answer == answer
}

func gameLoop(quizLength int) int {
	var question Question
	questionList := question.CreateQuestionList()
	var correctAnswers int
	var i int
	for i < quizLength {
		correct := questionList.Ask()
		if correct {
			correctAnswers++
		}
		i++
	}
	return correctAnswers
}

func main() {
	if len(os.Args) < 2 {
		fmt.Print("Please provide two arguments:\n  arg[1] = # of questions{int}\n  arg[2] = (s/question){int}\n")
		return
	}
	//var count int
	quizLength, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	if err != nil {
		log.Fatal(err)
	}
	correctAnswers := gameLoop(quizLength)
	fmt.Println("You got", correctAnswers, "correct out of", quizLength)
}
