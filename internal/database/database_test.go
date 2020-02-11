package database

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"testing"
)

func Test_insert(t *testing.T) {
	db, err := sql.Open("sqlite3", "./files/users.db")
	if err != nil {
		log.Panic(err)
	}
	defer db.Close()
	_, err = db.Exec("CREATE TABLE users (Id INTEGER PRIMARY KEY NOT NULL, UserName STRING, FullName STRING, City STRING, BirthDate DATE, Department STRING, Gender STRING, ExperienceYears INTEGER);")
	_, err = db.Exec("INSERT INTO users (Id, UserName, FullName, City, BirthDate, Department, Gender, ExperienceYears) VALUES (0, 'Zebralead', 'Benjamin Thompson', 'Ransom Canyon', '1990-10-20', 'IT', 'male', 3);")
	_, err = db.Exec("INSERT INTO users (Id, UserName, FullName, City, BirthDate, Department, Gender, ExperienceYears) VALUES (1, 'Howlerwool', 'James Jackson', 'Cienega Springs', '1985-04-10', 'IT', 'male', 3);")
	_, err = db.Exec("INSERT INTO users (Id, UserName, FullName, City, BirthDate, Department, Gender, ExperienceYears) VALUES (2, 'Jaguarband', 'Lily Brown', 'Ransom Canyon', '1980-04-10', 'IT', 'female', 3);")
	_, err = db.Exec("INSERT INTO users (Id, UserName, FullName, City, BirthDate, Department, Gender, ExperienceYears) VALUES (3, 'Servantbald', 'Ella Harris', 'Lucien', '2001-11-15', 'IT', 'female', 3);")
	_, err = db.Exec("INSERT INTO users (Id, UserName, FullName, City, BirthDate, Department, Gender, ExperienceYears) VALUES (4, 'Edgesmall', 'David Moore', 'Campden', '1999-03-23', 'IT', 'male', 3);")
	_, err = db.Exec("INSERT INTO users (Id, UserName, FullName, City, BirthDate, Department, Gender, ExperienceYears) VALUES (5, 'Serpentnavy', 'Elijah Johnson', 'Kingsbridge', '1945-04-03', 'IT', 'male', 3);")
	_, err = db.Exec("INSERT INTO users (Id, UserName, FullName, City, BirthDate, Department, Gender, ExperienceYears) VALUES (6, 'Hissershallow', 'James Williams', 'Buffalo City', '1992-07-01', 'IT', 'male', 3);")
	_, err = db.Exec("INSERT INTO users (Id, UserName, FullName, City, BirthDate, Department, Gender, ExperienceYears) VALUES (7, 'Chatterdog', 'Anthony Brown', 'Ransom Canyon', '1999-03-07', 'IT', 'male', 3);")
	_, err = db.Exec("INSERT INTO users (Id, UserName, FullName, City, BirthDate, Department, Gender, ExperienceYears) VALUES (8, 'Headmad', 'Mia Martinez', 'Lucien', '1997-06-28', 'IT', 'female', 3);")
	_, err = db.Exec("INSERT INTO users (Id, UserName, FullName, City, BirthDate, Department, Gender, ExperienceYears) VALUES (9, 'Lighterleaf', 'Mia Garcia', 'Burrton', '2007-05-09', 'IT', 'female', 3);")
	_, err = db.Exec("INSERT INTO users (Id, UserName, FullName, City, BirthDate, Department, Gender, ExperienceYears) VALUES (10, 'Gargoylewild', 'Anthony Jackson', 'Skidaway Island', '2003-09-01', 'IT', 'male', 3);")
	_, err = db.Exec("INSERT INTO users (Id, UserName, FullName, City, BirthDate, Department, Gender, ExperienceYears) VALUES (11, 'Kittenspice', 'Alexander Martinez', 'Baldock', '2000-01-25', 'IT', 'male', 3);")
	_, err = db.Exec("INSERT INTO users (Id, UserName, FullName, City, BirthDate, Department, Gender, ExperienceYears) VALUES (12, 'Beaklead', 'Aiden Smith', 'Newstead', '1918-10-26', 'IT', 'male', 3);")
	_, err = db.Exec("INSERT INTO users (Id, UserName, FullName, City, BirthDate, Department, Gender, ExperienceYears) VALUES (13, 'Howlervolcano', 'Sophia Harris', 'Campden', '1998-01-07', 'IT', 'female', 3);")
	_, err = db.Exec("INSERT INTO users (Id, UserName, FullName, City, BirthDate, Department, Gender, ExperienceYears) VALUES (14, 'Apenoble', 'Charlotte Jones', 'San Martin', '1977-08-12', 'IT', 'female', 3);")
	if err != nil {
		panic(err)
	}
}
