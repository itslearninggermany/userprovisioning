// Name: Thomas Nordmann (thomas.nordmann@itslearning.com)
// Date: 13.01.2019
// Function of the package: Userprovisioning for IMS-E
// Function of the file:

package userprovisioning

import (
	"encoding/xml"
	"errors"
	"fmt"
	"io/ioutil"
	//"math/rand"
	//"time"
)

type person struct {
	XMLName   xml.Name `xml:"person"`
	Recstatus string   `xml:"recstatus,attr"`
	Sourcedid struct {
		Source string `xml:"source"`
		ID     string `xml:"id"` // Synchronisationkey
	} `xml:"sourcedid"`
	Userid string `xml:"userid"` // USERNAME
	Name   struct {
		Fn string `xml:"fn"`
		N  struct {
			Family string `xml:"family"`
			Given  string `xml:"given"`
		} `xml:"n"`
	} `xml:"name"`
	Demographics struct {
		Bday string `xml:"bday"`
	} `xml:"demographics"`
	Email string    `xml:"email"`
	Tel   []telefon `xml:"tel"`
	Adr   struct {
		Street   string `xml:"street"`
		Locality string `xml:"locality"`
		Pcode    string `xml:"pcode"`
	} `xml:"adr"`
	Institutionrole struct {
		Primaryrole         string `xml:"primaryrole,attr"`
		Institutionroletype string `xml:"institutionroletype,attr"`
	} `xml:"institutionrole"`
	Extension struct {
		Relationship []childParent `xml:"relationship"`
	} `xml:"extension"`
}

// Creates a new Person
func Person(institution, syncID, username, givenname, familyname, birthday, email, street, locality, pcode, kindOfPersonStudentOrStaffOrParent string, tel []telefon, childOrParentIds []string, parent bool) person {
	a := new(person)
	a.Recstatus = "1"
	a.Sourcedid.Source = institution
	a.Sourcedid.ID = syncID
	a.Userid = username
	a.Name.Fn = givenname + " " + familyname
	a.Name.N.Family = familyname
	a.Name.N.Given = givenname
	a.Email = email
	a.Demographics.Bday = birthday
	a.Tel = tel
	a.Adr.Street = street
	a.Adr.Locality = locality
	a.Adr.Pcode = pcode
	a.Institutionrole.Primaryrole = "Yes"

	a.Institutionrole.Institutionroletype = kindOfPersonStudentOrStaffOrParent
	a.Extension.Relationship = MakeAChildSlice(institution, childOrParentIds, parent)
	return *a
}

// Parse the Person to a XML-File.
// It gives the XML-Content as []byte and all Person that were parsed as String and an error
// when something went wrong.
func (a *person) ParseToXML() ([]byte, string, error) {
	output, err := xml.Marshal(a)
	if err != nil {
		err2 := errors.New(err.Error() + " Person: " + a.Userid + " was not(!!!) imported. \n")
		return output, "", err2
	}
	log := "Person: " + a.Userid + " is imported. \n"

	return output, log, err
}

/*

// The filename is the name of the UserID; The format is XML.
// If there is already a personfile, it will be deletet.
func (a *person) Save2File (pathname string) (oldPersonDeletet bool) {
oldPersonDeletet = false;
if (IsFileInTheDirectory (pathname, a.Userid + ".xml")) {
	os.Remove(pathname+a.Userid+".xml");
	oldPersonDeletet = true;
}

file, err := os.Create(pathname+a.Userid+".xml")
_, err = file.Write(a.ParseToXML())
if err != nil {
	fmt.Println(err)
	return
}
return

}
*/

// Read file which was stored via SavePersonFiles
func ReadPersonFile(filename string) person {
	var a person
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
		return a
	}
	xml.Unmarshal(content, &a)
	return a
}

/*
func (a *person) Inmport2Database (nameOfTheDatabase string) error{
db := ConnectionZuDB(nameOfTheDatabase);
defer db.Close()

	stmtOut, err := db.Prepare("SELECT COUNT(idInstitution) FROM Institution WHERE name = ?;")
	if err != nil {
		return err
	}
	defer stmtOut.Close()
	var institutionid int

// First get the institutionid, if there is no, a new one will be created in the database.
	rows, err = stmtOut.QueryRow(a.Sourcedid.Source).Scan(&institutionid)
	if err != nil {
		return err
	}


	stmt , err := db.Prepare("INSERT INTO Institution (name) VALUE (?)")
		if err != nil {
			return err
		}
		_,err = stmt.Exec(a.Sourcedid.Source)

		return err
		stmt.Close()

		stmtOut2, err := db.Prepare("SELECT idInstitution FROM Institution WHERE name = ?;")
		if err != nil {
			return err
		}
		defer stmtOut2.Close()

		err = stmtOut2.QueryRow(a.Sourcedid.Source).Scan(&institutionid)
		if err != nil {
			return err
		}
	}

// Second: Create a new idPerson, which one is unique
rand.Seed(time.Now().UTC().UnixNano())


// Third: Import the person to the database
stmt , err := db.Prepare("INSERT INTO Person (idPerson, idInstitution, username, givenname, familyname, birthday, email, street, locality, pcode, kindOfPersonStudentOrStaffOrParent,synckey) VALUE (?,?,?,?,?,?,?,?,?,?,?,?)")
if err != nil {
	return err
}
defer stmt.Close()

idPerson :=  rand.Intn(2147483647)


_,err = stmt.Exec(idPerson,institutionid, a.Userid, a.Name.N.Given, a.Name.N.Family, a.Demographics.Bday, a.Email, a.Adr.Street,a.Adr.Locality, a.Adr.Pcode, a.Institutionrole.Institutionroletype,a.Sourcedid.ID)

if (err != nil) {
	return err
}

fmt.Println("Person with username" + a.Userid + "was imported to the database")
return nil
}

/*

func GetPersonFromDatabase (UserID string) person {

	rows , err := db.Query("SELECT * FROM Person;")
	//_, err = db.Query("INSERT INTO Person (idPerson, idInstitution, idUser, givenname, familyname, birthday, email, street, locality, pcode, kindOfPersonStudentOrStaffOrParent) VALUE  (266333,1,123499956,'Felicitas','Nordmann','17.01.2015','cecilia@mail.de','Keithstraße 29','Berlin','10787','student');")
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan( , &name)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(id, name)
	}


}


func GetAllPersonsFromDatabase () []person {

}

func DeletePersonFromDatabase (UserID string) (

)



*/
