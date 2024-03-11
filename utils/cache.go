package utils

import (
	"bufio"
	"log"
	"os"
)

func ReadCache(guid string) (bool, error) {
	// Use temp file to cache already sent alerts.
	// Create it if not already existing

	_, err := os.Stat("published.tmp")
	if err != nil {
		os.Create("published.tmp")
	}
	f, err := os.Open("./published.tmp")
	if err != nil {
		log.Printf("cannot read cache.")
		log.Print(err)
		return false, err
	}
	defer f.Close()

	s := bufio.NewScanner(f)
	s.Split(bufio.ScanLines)
	for s.Scan() {
		if s.Text() == guid {
			return true, nil
		}
	}
	return false, nil
}

func WriteCache(guid string) error {
	f, err := os.OpenFile("./published.tmp", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0664)

	if err != nil {
		log.Print("cannot open cache.")
		log.Print(err)
		return err
	}

	defer f.Close()

	_, err = f.WriteString(guid + "\n")

	if err != nil {
		log.Print("Error: Cannot write in cache.")
		log.Print(err)
		return err
	} else {
		log.Printf("Guid successfully write in cache.")
		log.Printf("Value: %s", guid)
		return nil
	}
}
