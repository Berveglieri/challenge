package tools

import (
	"database/sql"
	"flag"
	"fmt"
	"github.com/challenge/hive/compressor"
	"github.com/challenge/hive/cryptor"
	_ "github.com/lib/pq"
	"io/ioutil"
	"log"
	"net/url"
	"os"
	"os/exec"
	"strings"
	"time"
)

// EmptyValue function returns an error if not all args are passed
func EmptyValue() {
	fmt.Printf("Usage: %s [OPTIONS] argument ...\n", os.Args[0])
	flag.PrintDefaults()
}

// ReadPassword function reads an input without echoing it and cast it to string
func ReadPassword(password string) string {

	return password
}

// Exporter function exports an environment variable that was returned by Readpassword function
func Exporter(password string) {
	os.Setenv("PGPASSWORD", ReadPassword(password))
}

func listTables(host string, port string, user string, database string) []string {

	var tables = []string{}

	connStr := "postgres://" + user + ":" + url.QueryEscape(os.Getenv("PGPASSWORD")) + "@" + host + ":" + port + "/" + database + "?sslmode=require"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("connected to server %s\n", host)

	defer db.Close()

	rows, err := db.Query("SELECT c.relname FROM pg_class c WHERE c.relkind = 'S' UNION SELECT table_name FROM information_schema.tables WHERE table_type='BASE TABLE' AND table_schema='public';")
	if err != nil {
		log.Fatal(err)
	}

	cols, err := rows.Columns()
	if err != nil {
		fmt.Println("Failed to get columns", err)
	}

	raw := make([][]byte, len(cols))
	result := make([]string, len(cols))

	dataTables := make([]interface{}, len(cols))
	for i, _ := range raw {
		dataTables[i] = &raw[i]
	}

	for rows.Next() {
		err = rows.Scan(dataTables...)
		if err != nil {
			fmt.Println("Failed to scan row", err)
		}

		for i, j := range raw {
			if j == nil {
				result[i] = "\\N"
			} else {
				result[i] = string(j)
			}

			for _, i := range result {
				tables = append(tables, i)
			}
		}

	}

	return tables
}

func createDir(database string) string{

	var bDirectory  = ""

	fmt.Println("creating backup directory")
	_, err := os.Stat(database)

	if os.IsNotExist(err) {
		errDir := os.MkdirAll(database+"_"+time.Now().Format(time.RFC3339), 0755)
		if errDir != nil {
			log.Fatal(err)
		}
	}

	root, err := os.Getwd()
	if err != nil{
		panic(err)
	}

	files, err := ioutil.ReadDir(root)
	if err != nil {
		panic(err)
	}

	for _, f := range files {
		if strings.HasPrefix(f.Name(), database+"_") {
			bDirectory = f.Name()
			os.Chdir(bDirectory)
		}
	}

	return bDirectory

}

// ExecBackup function executes a system call to execute pg_dumb binary
func ExecBackup(done chan string, host string, port string, user string, database string, password string) {

	Exporter(password)
	dir := createDir(database)

	for _, s := range listTables(host, port, user, database) {
		fmt.Println("Backing up " + s + " table")
		cmdExecuteBackup := exec.Command("pg_dump", "-h", host, "-p", port, "-U", user, "-d", database, "-t", s, "-f", s+".sql")
		output, err := cmdExecuteBackup.CombinedOutput()
		done <- s

		if err != nil {
			fmt.Println(fmt.Sprint(err) + ": " + string(output))
		}
	}

	os.Chdir("../")

	fmt.Println("compressing backup")
	compressor.Compress(dir, dir+".tar.gz")

	if _, err := os.Stat(dir+".tar.gz"); err == nil {
		fmt.Println("File compressed successfully")

		if err != nil{
			panic(err)
		}
	}

	fmt.Println("encrypting backup")
	cryptor.Encrypt(dir+".tar.gz", "IceCream", dir+".zzz")
	if _, err := os.Stat(dir+".zzz"); err == nil {
		fmt.Println("File encrypted successfully")

		if err != nil{
			panic(err)
		}
	}

	os.RemoveAll(dir)
	os.RemoveAll(dir+".tar.gz")

	close(done)

}







