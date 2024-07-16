package databaseminer

import (
	"fmt"
	"regexp"
)

type DatabaseMiner interface {
	GetSchema() (*Schema, error)
}

type Schema struct {
	Databases []Database
}

type Database struct {
	Name string
	Tables []Table
}

type Table struct {
	Name string
	Columns []string
}

func Search(m DatabaseMiner) error {
	s,err := m.GetSchema()
	if err != nil {
		return err
	}

	re := getRegex()
	for _,db := range s.Databases {
		for _,t := range db.Tables {
			for _,c := range t.Columns {
				for _,r := range re {
					if r.MatchString(c) {
						fmt.Println(db)
						fmt.Printf("[+] HIT: %s\n", c)
					}
				}
			}
		}
	}
	
	return nil
}

func getRegex() []*regexp.Regexp {
	return []*regexp.Regexp{
		regexp.MustCompile(`(?i)social`),
		regexp.MustCompile(`(?i)ssn`),
		regexp.MustCompile(`(?i)pass(word)?`),
		regexp.MustCompile(`(?i)hash`),
		regexp.MustCompile(`(?i)ccnum`),
		regexp.MustCompile(`(?i)card`),
		regexp.MustCompile(`(?i)security`),
		regexp.MustCompile(`(?i)key`),
	}
}
