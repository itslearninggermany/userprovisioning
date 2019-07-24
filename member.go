// Name: Thomas Nordmann (thomas.nordmann@itslearning.com)
// Date: 13.01.2019
// Function of the package: Userprovisioning for IMS-E
// Function of the file:

package userprovisioning

import (
	"encoding/xml"
	"errors"
)

type member struct {
	Sourcedid struct {
		Source string `xml:"source"` // Univention
		ID     string `xml:"id"`     // UID
	} `xml:"sourcedid"`
	Idtype string `xml:"idtype"` // always 1
	Role   struct {
		Recstatus string `xml:"recstatus,attr"` // always 1   // When Mentor then 6
		Roletype  string `xml:"roletype,attr"`  // 01 Student; 02 Staff; 04 Guest
		Subrole   string `xml:"subrole"`        // Student, Staff, Guest
		Status    string `xml:"status"`         // always 1
	} `xml:"role"`
}

func Member(institution string, personID string, role string, mentor bool) member {
	a := new(member)
	a.Sourcedid.Source = institution
	a.Sourcedid.ID = personID
	a.Idtype = "1"
	a.Role.Recstatus = "1"

	var rolenumber string
	if role == "Student" {
		rolenumber = "01"
	} else if role == "Staff" {
		rolenumber = "02"
	} else {
		rolenumber = "04"
	}
	a.Role.Roletype = rolenumber
	a.Role.Subrole = role
	a.Role.Status = "1"
	if mentor {
		a.Role.Roletype = "06"
	}
	return *a
}

// Parses the Memberobject to an xml-bytestream
func (a *member) ParseToXML() ([]byte, string, error) {
	x, err := xml.Marshal(a)
	if err != nil {
		err2 := errors.New(err.Error() + " Member: " + a.Sourcedid.ID + " was not(!!!) added. \n")
		return nil, "", err2
	}
	log := "Member: " + a.Sourcedid.ID + " was not(!!!) added. \n"
	return x, log, nil
}

/*
// The filename is mebmer.[the name of the UserID].xml; The format is XML.
// If there is already a memberfile, it will be deletet.
func (a *member) Save2File (pathname string) (oldMemberDeletet bool) {
	oldMemberDeletet = false;
	if (IsFileInTheDirectory (pathname, "member."+ a.Userid + ".xml")) {
		os.Remove(pathname+a.Userid+".xml");
		oldMemberDeletet = true;
	}

	file, err := os.Create(pathname+"member."+a.Userid+".xml")
	_, err = file.Write(a.ParseToXML())
	if err != nil {
		fmt.Println(err)
		return
	}
	return
}


// Read file which was stored via SaveMemberFiles
func ReadMemberFile (filename string) (member) {
	var a member;
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
		return a
	}
	xml.Unmarshal(content, &a)
	return a
*/
