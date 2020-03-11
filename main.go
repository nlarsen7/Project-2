package main

import (
  "fmt"
  "math/rand"
  "time"
)

func main() {
rand.Seed(time.Now().UnixNano())
var player int
var computer int
playerScore:=0 
computerScore:=0
var roll int
fmt.Println("Welcome to the game of Pig. The rules are simple, roll a die and get to 100. 1's don't give you any points and end your turn with a 0 for that round.")
for computerScore<100||playerScore<100{
for playerScore<100{
  roll=rand.Intn(6)+1
  if roll==1{
    playerScore=0
    break
  }else{
    playerScore=playerScore+roll
  }
//ask user to roll or hold
fmt.Println("enter 1 to hold or 0 to roll again.")
fmt.Println(playerScore)
fmt.Scanln(&roll)
if roll==1{
  fmt.Println("Player Score is",playerScore+player,"and your turn is over")
  break
} 
}
for computerScore<100{
  roll=rand.Intn(6)+1
  if roll==1{
    computerScore=0
    break
  }else{
    computerScore=computerScore+roll
  }
  fmt.Println("enter 1 to hold or 0 to roll again.")
  fmt.Println(computerScore)
  fmt.Scanln(&roll)
  if roll==1{
    fmt.Println("Computer Score is",computer+computerScore,"and your turn is over")
    break
  }
}
player=playerScore
computer=computerScore
}  
}
