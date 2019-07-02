package userprovisioning

// Name: Thomas Nordmann (thomas.nordmann@itslearning.com)
// Date: 13.01.2019
// Function of the package: Userprovisioning for IMS-E
// Function of the file: the Child-Struc for Parents/Child Relation

type child struct {
	Sourcedid struct {
		Source string `xml:"source"`
		ID     string `xml:"id"`
	} `xml:"sourcedid"`
	Label string `xml:"label"`
}

// Creates a new Child.
func Child(institution, childID string) *child {
	a := new(child)
	a.Sourcedid.Source = institution
	a.Sourcedid.ID = childID
	a.Label = "child"
	return a
}

// GetChild gibt die Institution und die ChildID zurück
func (a *child) GetChild() (institution, childID string) {
	return a.Sourcedid.Source, a.Sourcedid.ID
}

// MakeAChildSlice erstellt ein Array mit Childs.
func MakeAChildSlice(institution string, childIDs []string) (res []child) {
	for i := 0; i < len(childIDs); i++ {
		a := Child(institution, childIDs[i])
		res = append(res, *a)
	}
	return
}
