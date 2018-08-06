package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"time"
)
var ar [101]interface{}

func set() {
	for i:=1;i<=100;i++ {
		ar[i]=i
	}
	ar[39]="SNAKE"
	ar[5]="LADDER"
	ar[77]="LADDER"
	ar[25]="SNAKE"
	ar[33]="LADDER"
	ar[45]="SNAKE"
	ar[53]="SNAKE"
	ar[62]="LADDER"
	ar[97]="SNAKE"
	ar[72]="SNAKE"
}
func random(min, max int) int {
    rand.Seed(time.Now().Unix())
    return rand.Intn(max - min) + min
}

func flow(){
	
	in:=""
	in1:=""
	cur:=-1
	cur1:=-1
	//t:=0
	t1:=0
	count:=0
	count1:=0
	ch:=""
	ch1:=""
	g:=0
	g1:=0
	var m interface{}
	var m1 interface{}
	Loop:
	fmt.Println("Press y to dice or n to quit ## USER--1")
	if(cur==-1) {
		cur=0
	}
	if(cur!=-1) {
		ar[cur]=m
	}
	fmt.Scanln(&in)
	h:=exec.Command("clear")
	h.Stdout=os.Stdout
	h.Run()
	if(in=="n") {
		fmt.Println("PLAYER-1 QUIT")
		fmt.Println("PLAYER-2 WON !!!!")
		return
	}
	if(in=="y") {
		
		di:=random(1,6)
		count++
		g=0
		cur=cur+di
		if(cur>100){
			cur=cur-di
		}
		if(ar[cur]=="LADDER"){
			ch="PLAYER-1 GOT A LADDER"
			g=1
			if(cur==5) {
				cur=27
			}
			if(cur==77) {
				cur=95
			}
			if(cur==33) {
				cur=55
			}
			if(cur==62) {
				cur=96
			}
		}
		if(ar[cur]=="SNAKE") {
			ch="PLAYER-1 GOT A SNAKE BITE"
			g=1
			if(cur==39) {
				cur=17
			}
			if(cur==25){
				cur=7
			}
			if(cur==45){
				cur=35
			}
			if(cur==53) {
				cur=31
			}
			if(cur==97) {
				cur=67
			}
			if(cur==72) {
				cur=18
			}
		}
		if(g!=1) {
			ch="PLAYER-1 GOT NEITHER SNAKE NOR LADDER"
		}
		fmt.Println("DICE ROLLED :",di)
		fmt.Print(ch)
		if(ar[cur]==100){
			fmt.Println("\t\t\t#######  PLAYER-1 WON ########")
			fmt.Println("\t\tTOTAL MOVES TAKEN",count)
			t1=1
			return 
		} 
		m=ar[cur]
		if(ar[cur]=="*P2*") {
			ar[cur]="*P1*,*P2*"
		}
		if(ar[cur]!="*P2*") {
			
			ar[cur]="*P1*"
		}
		board()
		
		
	}
	fmt.Println("Press y to dice or n to quit ## USER--2")
	if(cur1==-1) {
		cur1=0
	}
	if(cur1!=-1) {
		ar[cur1]=m1
	}
	fmt.Scanln(&in1)
	h1:=exec.Command("clear")
	h1.Stdout=os.Stdout
	h1.Run()
	if(in1=="n") {
	fmt.Println("PLAYER-2 QUIT")
	fmt.Println("PLAYER-1 WON !!!!")
		return
	}
	if(in1=="y") {
		
		di:=random(1,6)
		count1++
		g1=0
		cur1=cur1+di
		if(cur1>100){
			cur1=cur1-di
		}
		if(ar[cur1]=="LADDER"){
			ch1="PLAYER-2 GOT A LADDER"
			g1=1
			if(cur1==5) {
				cur1=27
			}
			if(cur1==77) {
				cur1=95
			}
			if(cur1==33) {
				cur1=55
			}
			if(cur1==62) {
				cur1=96
			}
		}
		if(ar[cur1]=="SNAKE") {
			ch1="PLAYEr-2 GOT A SNAKE BITE"
			g1=1
			if(cur1==39) {
				cur1=17
			}
			if(cur1==25){
				cur1=7
			}
			if(cur1==45){
				cur1=35
			}
			if(cur1==53) {
				cur1=31
			}
			if(cur1==97) {
				cur1=67
			}
			if(cur1==72) {
				cur1=18
			}
		}
		if(g1!=1) {
			ch1="PLAYER-2 GOT NEITHER SNAKE NOR LADDER"
		}
		fmt.Println("DICE ROLLED :",di)
		fmt.Print(ch1)
		if(ar[cur1]==100){
			fmt.Println("\t\t\t#######  PLAYER-2 WON ########")
			fmt.Println("\t\tTOTAL MOVES TAKEN",count1)
			t1=1
			return 
		} 
		m1=ar[cur1]
		if(ar[cur1]=="*P1*") {
			ar[cur1]="*P1,P2*"
		}
		if(ar[cur1]!="*P1*") {
			ar[cur1]="*P2*"
		}
		board()
		
		goto Loop
	}
	if(t1==0) {	
		goto Loop
	}
}
func board() {
	fmt.Println()
	for i:=100;i>=1;i-- {
		j:=0
		fmt.Print("\t\t\t")
		for j<10 {
			fmt.Print(ar[i],"\t")
			i--
			j++
		}
		fmt.Print("\n\n")
		k:=0
		i=i-9
		fmt.Print("\t\t\t")
		for k<10 {
			fmt.Print(ar[i],"\t")
			i++
			k++
		}
		fmt.Print("\n\n")
		i=i-9
		i--
	}
}

func main() {
	set()	
	board()
	flow()
}