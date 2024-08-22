package mail

import (
	"bufio"
	"os"
	"strings"
)

type CSVReader struct {
}

func NewCSVReader() CSVReader {
	return CSVReader{}
}

func (c *CSVReader) GetEmailsFromFile(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return []string{}, err
	}

	scanner := bufio.NewScanner(file)

	var emailList []string

	for scanner.Scan() {
		currentEmail := scanner.Text()
		validEmail := strings.Contains(currentEmail, "@")
		if validEmail {
			trimmedEmail := strings.TrimSpace(currentEmail)
			emailList = append(emailList, trimmedEmail)
		}
	}

	return emailList, nil
}
