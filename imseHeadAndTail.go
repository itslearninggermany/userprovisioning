// Name: Thomas Nordmann (thomas.nordmann@itslearning.com)
// Date: 13.01.2019
// Function of the package: Userprovisioning for IMS-E
// Function of the file:

package userprovisioning

import (
	"strconv"
	"time"
)

const xMLHeadPart1 string = "<?xml version=\"1.0\" encoding=\"ISO-8859-1\" standalone=\"no\"?><!DOCTYPE enterprise PUBLIC \"IMS Enterprise/LMS Interoperability DTD\" \"http://www.fs.usit.uio.no/DTD/ims_epv1p1.dtd\"><enterprise><properties lang=\"NO\"><datasource>univention</datasource><target>LMS</target><datetime>"
const xMLHeadPart2 string = "2018-10-24</datetime></properties>"
const xMLTail string = "</enterprise>"

// The beginning of an IMSE-File is created
func CreateXMLHead() []byte {
	now := time.Now()
	year, month, day := now.Date()
	tmp := xMLHeadPart1 + strconv.Itoa(day) + "." + month.String() + "." + strconv.Itoa(year) + xMLHeadPart2
	return []byte(tmp)
}

// The ending of an IMSE-File is created
func CreateXMLTail() []byte {
	return []byte(xMLTail)
}
