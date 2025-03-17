package main

import (
	"database/sql"
	"encoding/csv"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

type Data struct {
	Gender               string
	Caste                string
	Coaching             string
	ClassTenEducation    string
	ClassTwelveEducation string
	Medium               string
	ClassXPercentage     string
	ClassXIIPercentage   string
	FatherOccupation     string
	MotherOccupation     string
	Time                 string
	Performance          string
}

func getMySQL() *sql.DB {
	db, err := sql.Open("mysql", "root:root@(127.0.0.1:3306)/Having_db?parseTime=true")

	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	return db
}

func main() {
	db := getMySQL()
	defer db.Close()

	file, err := os.Open("D:/go/src/.go/HAVING/Student_Performance_on_an_Entrance_Examination.csv")
	if err != nil {
		log.Fatal("Failed to open file:", err)
	}
	defer file.Close()

	df := csv.NewReader(file)
	data, err := df.ReadAll()
	if err != nil {
		log.Fatal("Failed to read CSV file:", err)
	}

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS student_performance (
		Gender VARCHAR(10),
		Caste VARCHAR(50),
		Coaching VARCHAR(50),
		ClassTenEducation VARCHAR(50),
		ClassTwelveEducation VARCHAR(50),
		Medium VARCHAR(50),
		ClassXPercentage VARCHAR(10),
		ClassXIIPercentage VARCHAR(10),
		FatherOccupation VARCHAR(50),
		MotherOccupation VARCHAR(50),
		Time VARCHAR(50),
		Performance VARCHAR(50)
	)`)
	if err != nil {
		log.Fatal("Failed to create table:", err)
	}

	
	for i := 1; i < len(data); i++ {
		_, err := db.Exec(`INSERT INTO student_performance (
			Gender, Caste, Coaching, ClassTenEducation, ClassTwelveEducation, Medium,
			ClassXPercentage, ClassXIIPercentage, FatherOccupation, MotherOccupation,
			Time, Performance) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`,
			data[i][0], data[i][1], data[i][2], data[i][3], data[i][4], data[i][5],
			data[i][6], data[i][7], data[i][8], data[i][9], data[i][10], data[i][11],
		)
		if err != nil {
			log.Println("Failed to insert record:", err)
		}
	}

	fmt.Println("Data Inserted Successfully")
}
