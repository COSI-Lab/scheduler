package main

import (
	"fmt"
	"net/http"
	"encoding/csv"
	"os"
)

type Course struct {
	//Number for the peoplesoft database
	ClassNbr string
	//Subject (ex MA, CS, STAT)
	Subject	string
	//The course number most people think about (ex the 131 in MA131)
	Catalog	string
	//The class section
	Section	string
	//THe course name (This is what is in the spread sheet)
	Description	string
	//How many people can take the course
	Capacity	string
	//How man people are enrolled in the course
	Enrolled	string
	//Availiable seats
	Seats	string
	//How many people are on a waitlist
	Waitlist	string
	//How many people have the course in their shopping cart
	Cart	string
	//How many more people are taking the class, on waitlist or have the class n their shopping cart then the class actuall has seats for
	Overflow	string
	//Time when the class starts
	StartTime	string
	//Time when the class ends
	EndTime	string
	//What days does the calss meet
	MeetingDays	string
	//the room the class is in
	Room	string
	//Who teaches the course
	Prof	string
	//Date that the class starts
	StartDate	string
	//Date the class end
	EndDate	string
}

func rootHandler(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w,"<p> Hello World </p>")
}

func handleErr(e error){
	if e != nil {
		panic(e)
	}
}

func parseClases() (ret []Course ) {
	//This function reads from ./ps.csv and returns a slice of courses
	//Grab all files into memory
	fileData,err := os.Open("./ps.csv")
	handleErr(err)
	reader := csv.NewReader(fileData)
	for rec,err := reader.Read(); err == nil;rec,err = reader.Read() {
		tmp := Course{ ClassNbr: rec[0], Subject: rec[1], Catalog: rec[2], Section: rec[3], Description:rec[4], Capacity:rec[5], Enrolled:rec[6], Seats:rec[7], Waitlist:rec[8], Cart:rec[9], Overflow:rec[10], StartTime:rec[11], EndTime:rec[12], MeetingDays:rec[12], Room:rec[14], Prof:rec[15], StartDate:rec[16], EndDate:rec[17] }

		ret = append(ret,tmp)
	}

	return ret
}

func main(){
	port := ":8080"
	http.Handle("/",http.FileServer(http.Dir("./static")))
	fmt.Printf("Listening on port %s\n",port)
	parseClases()
	http.ListenAndServe(port,nil)
}
