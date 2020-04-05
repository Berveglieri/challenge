package main

import (

"flag"
"fmt"
"os"
"github.com/challenge/hive/tools"
)

var (

	//operation to be executed in the database e. g. backup
	operation string
	//host flag
	host string
	//port flag
	port string
	//user flag
	user string
	//user password
	password string
	//database name
	database string
)

func main() {

	flag.StringVar(&operation, "operation", "", "This flag defines the operation to be executed in the database")
	flag.StringVar(&host, "h", "", "host IP or FQDN")
	flag.StringVar(&port, "p", "", "port database is listening for requests")
	flag.StringVar(&user, "u", "", "user to connect to database")
	flag.StringVar(&password, "P", "", "User password")
	flag.StringVar(&database, "d", "", "database to connect")

	flag.Parse()
	flag.Usage = tools.EmptyValue

	if operation == "" || host == "" || port == "" || user == "" || database == "" || password == "" {
		tools.EmptyValue()
		os.Exit(1)
	} else if operation == "backup" {

		done := make(chan string)
		go tools.ExecBackup(done, host, port, user, database, password)
		for {
			v, ok := <-done
			if ok == false {
				break
			} else {
				fmt.Println("done ", v)
			}
		}
	}
}

