package main

import (
	"flag"
	"fmt"
	"os"
	"github.com/challenge/hive/cryptor"
)

func main() {

	//cryptor.Encrypt("hive_2020-04-04T23:05:11+02:00.tar.gz","uyara", "hive_2020-04-04T23:05:11+02:00.tar.gz.dyx")
	//cryptor.Decrypt("hive_2020-04-04T23:05:11+02:00.tar.gz.dyx", "uyara", "hive_2020-04-04T23:05:11+02:00.tar.gz")
	//

	var (
		//operation to be executed in the file/directory e. g. encrypt|decrypt
		operation string
		//source file flag
		sfile string
		//passhphrase flag
		passphrase string
		//destination file flag
		dfile string
	)

	flag.StringVar(&operation, "operation", "", "This flag defines the operation to be executed in the file")
	flag.StringVar(&sfile, "s", "", "source file name")
	flag.StringVar(&passphrase, "p", "", "string to generate the hash")
	flag.StringVar(&dfile, "d", "", "destination file")

	flag.Parse()
	flag.Usage = emptyValue

	if operation == "" || sfile == "" || passphrase == "" || dfile == "" {
		emptyValue()
		os.Exit(1)
	} else if operation == "encrypt" {
		cryptor.Encrypt(sfile, passphrase, dfile)
	} else if operation == "decrypt" {
		cryptor.Decrypt(sfile, passphrase, dfile)
	}
}

func emptyValue() {
	fmt.Printf("Usage: %s [OPTIONS] argument ...\n", os.Args[0])
	flag.PrintDefaults()
}
