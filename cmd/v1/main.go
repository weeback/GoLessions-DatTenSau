package main

import (
	"fmt"
	"myapp"
	"myapp/lib/crypto"
	"os"
	"path/filepath"
)

func main() {
	fmt.Printf("Crypto Tools Application\n------------------------------------\n"+
		"\tVersion: %s\n------------------------------------\n", myapp.Version)
	generateRSAKeyPair(2048)
}

func generateRSAKeyPair(keyBits int) {
	var (
		oldName    = filepath.Clean(os.Getenv("HOME") + "/.ssh/id_rsa")
		keyName    string
		passphrase string
	)
	//
	fmt.Println("Generate public/private RSA key pair.")
	keyName = stdReadStr("Enter file in which to save the key ("+oldName+"): ", oldName)
	for i := 1; i <= 3; i++ {
		passphrase = stdReadStr("Enter passphrase (empty for no passphrase): ")
		if passphrase == stdReadStr("Enter same passphrase again: ") {
			break
		}
		if i < 3 {
			fmt.Println("Passphrases do not match.  Try again.")
			continue
		} else {
			fmt.Println("Too many failures")
			os.Exit(1)
		}
	}
	// generate RSA key pair
	rsa := crypto.NewRSA()
	key, pub, err := rsa.GenerateKeyPair(keyBits)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	p1, p2, err := rsa.EncodePKCS8(key)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	// create directory if not exists
	if err = os.MkdirAll(filepath.Dir(keyName), 0755); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	// write private and public key to file
	if err = os.WriteFile(keyName, p1, 0644); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	if err = os.WriteFile(keyName+".pub", p2, 0644); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("Your identification has been saved in %s.\nYour public key has been saved in %s.pub.\n", keyName, keyName)
	//
	fp, art := rsa.GenerateRandomArt(pub)
	fmt.Printf("The key fingerprint is:\nSHA256:%s %s\n", fp, os.Getenv("USERNAME")+"@"+os.Getenv("COMPUTERNAME"))
	fmt.Printf("%s\n", art)
}
