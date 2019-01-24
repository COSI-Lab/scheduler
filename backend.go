package main

import (
	"fmt" //Printing and writing
	"net/http" //Actually handles http
	"encoding/csv" //Read the csv
	"os" //Open the file
	"encoding/json" //Make json
	"time" //Parse the time of a class
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
	//Kind of thing that the lab is, "lec" for lecture, "lab" for a lab, "dis" for discussion, and "UNK" for unkown
	//UNK should never reach the front end
	Kind string
}

//Global Variable (Used for computational efficency)
var cList []Course

func handleErr(e error){
	//This function takes an error and panics if it exists, used so frequently it makes sense to be a function
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
		//Gets each row of the file and parses into tmp
		tmp := Course{ ClassNbr: rec[0], Subject: rec[1], Catalog: rec[2], Section: rec[3], Description:rec[4], Capacity:rec[5], Enrolled:rec[6], Seats:rec[7], Waitlist:rec[8], Cart:rec[9], Overflow:rec[10], StartTime:rec[11], EndTime:rec[12], MeetingDays:rec[12], Room:rec[14], Prof:rec[15], StartDate:rec[16], EndDate:rec[17] , Kind:"UNK" }

		//Guesses at what kind of class the class is (lecture, lab or discussion)
		if len([]rune(tmp.MeetingDays)) > 1 {
			tmp.Kind = "lec"
		} else {
				//Figure out how long the current class takes
				startTime,_ := time.Parse("3:04PM",tmp.StartTime)
				endTime, _ := time.Parse("3:04PM",tmp.EndTime)
				dur :=startTime.Sub(endTime )
				if (dur.Hours() > 1){
					tmp.Kind = "lab"
				} else {
					tmp.Kind = "dis"
				}
		}

		ret = append(ret,tmp)
	}

	return ret
}

func classHandler(w http.ResponseWriter, r *http.Request){
	//This function handles sending the list of classes as json objects. 
	classes := cList //parseClases()
	for _,course := range classes {
		jason,err := json.Marshal(course)
		handleErr(err)
		fmt.Fprintf(w,"%s",jason)
	}
}

func main(){
	port := ":8080"
	cList = parseClases()
	http.HandleFunc("/classes",classHandler)
	http.Handle("/",http.FileServer(http.Dir("./static")))
	fmt.Printf("Listening on port %s\n",port)
	parseClases()
	http.ListenAndServe(port,nil)
}
