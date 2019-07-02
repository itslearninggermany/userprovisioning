package userprovisioning

import (
	"fmt"
	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
	"io/ioutil"
	"os"
)

type sFTPSetup struct {
	ipadresse string
	username  string
	password  string
}

// Makes a Setup for an SFTPUpload
func SFTPSetup(ipadresse string, username string, password string) *sFTPSetup {
	a := new(sFTPSetup)
	a.ipadresse = ipadresse
	a.password = username
	a.username = password
	return a
}

// Upload a file with the credentials given by  creating the SFTP-Object
func (a *sFTPSetup) UploadSFTPData(input []byte) (string, error) {
	config := &ssh.ClientConfig{
		User:            a.username, //sftpUsername
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		Auth: []ssh.AuthMethod{
			ssh.Password(a.password), //sftpPasswort
		},
	}

	config.SetDefaults()
	sshConn, err := ssh.Dial("tcp", a.ipadresse+":22", config) //sftpServer
	if err != nil {
		return "", err
	}
	defer sshConn.Close()

	c, err := sftp.NewClient(sshConn)
	if err != nil {
		return "", err
	}
	defer c.Close()

	remoteFile, err := c.Create("upload.xml")
	if err != nil {
		return "", err
	}

	_, err = remoteFile.Write(input)
	if err != nil {
		return "", err
	}

	log := ">>>>>>>>>>>>>>Here is the data which was uploaded: <<<<<<<<<<<<<<<<<<< \n" + string(input)

	return log, nil
}

// Upload a file with the credentials given by  creating the SFTP-Object
func (a *sFTPSetup) UploadSFTP(file string, path string, deleteFileAfterUpload bool) {
	config := &ssh.ClientConfig{
		User:            a.username, //sftpUsername
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		Auth: []ssh.AuthMethod{
			ssh.Password(a.password), //sftpPasswort
		},
	}

	config.SetDefaults()
	sshConn, err := ssh.Dial("tcp", a.ipadresse+":22", config) //sftpServer
	if err != nil {
		fmt.Print(err)
		return
	}
	defer sshConn.Close()

	c, err := sftp.NewClient(sshConn)
	if err != nil {
		fmt.Print(err)
		return
	}
	defer c.Close()

	remoteFileName := file
	remoteFile, err := c.Create(remoteFileName)
	if err != nil {
		fmt.Println(err)
		return
	}

	upload, err := ioutil.ReadFile(path + file)
	if err != nil {
		fmt.Println(err)
		return
	}
	_, err = remoteFile.Write(upload)
	if err != nil {
		fmt.Println(err)
		return
	}
	if deleteFileAfterUpload {
		os.Remove(path + file)
	}

	fmt.Println(">>>>>>>>>>>>>>Here is the data which was uploaded: <<<<<<<<<<<<<<<<<<<")
	fmt.Printf("%s", upload)
}
