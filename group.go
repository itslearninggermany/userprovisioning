package userprovisioning

import (
	"encoding/xml"
	"errors"
	"strconv"
)

type group struct {
	XMLName   xml.Name `xml:"group"`
	Sourcedid struct {
		Source string `xml:"source"` // always igslengede
		ID     string `xml:"id"`
	} `xml:"sourcedid"`
	Grouptype struct {
		Scheme    string `xml:"scheme"` // always igslengede
		Typevalue struct {
			Level string `xml:"level,attr"` // Level 0: Root; level 1: School; Level 2: Groups
			Value string `xml:",chardata"`
		} `xml:"typevalue"`
	} `xml:"grouptype"`
	Description struct {
		Short string `xml:"short"` // Name the same like ID!
	} `xml:"description"`
	Relationship struct {
		Relation  string `xml:"relation,attr"` // always 1
		Sourcedid struct {
			Source string `xml:"source"` // alway igslengede
			ID     string `xml:"id"`     // When it is rool --> root; When it is the School --> root; When it is a Group of the School --> Schoolname
		} `xml:"sourcedid"`
		Label string `xml:"label"` // root level; school level; groups level
	} `xml:"relationship"`
}

func Group(institution string, scheme string, grouptypevalue string, groupname string, groupID int, level int, parentgroupID int, parentlevel int) *group {
	a := new(group)
	a.Sourcedid.ID = strconv.Itoa(groupID)
	a.Sourcedid.Source = institution
	a.Grouptype.Scheme = scheme
	a.Grouptype.Typevalue.Level = strconv.Itoa(level)
	a.Grouptype.Typevalue.Value = grouptypevalue
	a.Description.Short = groupname
	a.Relationship.Relation = "1"
	a.Relationship.Sourcedid.Source = institution
	a.Relationship.Sourcedid.ID = strconv.Itoa(parentgroupID)
	a.Relationship.Label = strconv.Itoa(parentlevel)
	return a
}

func (a *group) Parse2XML() ([]byte, string, error) {
	x, err := xml.Marshal(a)
	if err != nil {
		err2 := errors.New(err.Error() + " Group " + a.Description.Short + " with ID " + a.Sourcedid.ID + " was not (!!!!!) created \n")
		return nil, "", err2
	}
	log := "Group " + a.Description.Short + " with ID " + a.Sourcedid.ID + " was  created \n"
	return x, log, nil
}

/*

func (a *group) GetGroupID () int {
	x,_ := strconv.Atoi (a.Sourcedid.ID)
	return x
}

func (a *group) GetGroupName () string {
	return a.Description.Short
}

func (a *group) GetGroupLevel () int {
	x , _ := strconv.Atoi(a.Grouptype.Typevalue.Level)
	return  x
}

func (a *group) GetParentInformation () (parentGroupID, parentLevel int){
	parentGroupID, _  = strconv.Atoi(a.Relationship.Sourcedid.ID)
	parentLevel, _ = strconv.Atoi(a.Relationship.Label)
	return
}
*/
