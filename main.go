package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"time"
)

type example struct{
	q string
	a string
}
func main() {
	csvFile:=flag.String("fileName", "problems.csv", "A problems to solve")
	gameTime:=flag.Int("gameTime", 20, "Value of time to play the game.")
	flag.Parse()
	file, err := os.Open(*csvFile)
	if err!= nil{
		fmt.Println("Couldn't open a file.")
	}
	r := csv.NewReader(file)
	lines, err := r.ReadAll()
	if err!=nil{
		fmt.Println("There's no lines.")
	}

	problems :=ParseLines(lines)
	timer :=time.NewTimer(time.Duration(*gameTime)*time.Second)
	correct:=0
	for i, p := range problems{
		fmt.Printf("Problem number %d: %s=", i+1, p.q)
		answerCh :=make(chan string)
		go func() {
			var answer string
			fmt.Scanf("%s\n", &answer)
			answerCh<-answer
		}()
		select{
			 case <-timer.C:
				 fmt.Printf("\nOkay, time's up, you've scored %d of %d points. ",correct, len(lines))
				 return
			 case answer := <-answerCh:
			 	if answer==p.a{
			 		correct++
				}
		}
	}
	fmt.Printf("\nOkay, you've scored %d of %d points. ",correct, len(lines))
}
func ParseLines(lines [][]string)[]example{
	ret := make([]example, len(lines))
	for i, line := range lines{
		ret[i]=example{
			q:line[0],
			a:line[1],
		}
	}
	return ret
}

