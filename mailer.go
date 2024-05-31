package main

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"log"
	"net/smtp"
	"os"
	"text/template"

	"github.com/joho/godotenv"
)

func main() {
	if len(os.Args) != 2 {
		log.Printf("Error: Usage: %s <file_path>", os.Args[0])
		return
	}
	// Specify your CSV file path
	csvFilePath := os.Args[1]

	// Read CSV file
	records, err := readCSV(csvFilePath)
	if err != nil {
		fmt.Println("Error reading CSV:", err)
		return
	}

	err = godotenv.Load("cred.env")
	if err != nil {
		log.Fatalf("Some error occured while reading the Env file. Err: %s", err)
	}
	// SMTP server configuration
	smtpServer := "smtp.gmail.com"
	smtpPort := 587
	senderEmail := os.Getenv("SenderEmail")
	senderPassword := os.Getenv("SenderPassword")

	// Email template
	subject := "3 Days Linux Workshop: Congratulations on Completing the Linux Workshop!"

	// Iterate through records and send emails
	for _, record := range records {
		// Extract necessary information from the CSV record
		email := record[2]
		name := record[1]
		templateContent, err := readTemplate("./templates/thankyouforparticipation.txt")
		tmpl, err := template.New("emailTemplate").Parse(templateContent)
		if err != nil {
			log.Fatal("Something went wrong")
		}

		var body bytes.Buffer
		err = tmpl.Execute(&body, struct{ Name string }{name})
		if err != nil {
			log.Fatal("Something went wrong")
		}
		// Compose the email body

		// Send email
		err = sendEmail(smtpServer, smtpPort, senderEmail, senderPassword, email, subject, string(body.Bytes()))
		if err != nil {
			fmt.Printf("Error sending email to %s: %v\n", email, err)
		} else {
			fmt.Printf("Email sent successfully to %s\n", email)
		}
	}
}
func readTemplate(filePath string) (string, error) {
	content, err := os.ReadFile(filePath)
	if err != nil {
		return "", err
	}
	return string(content), nil
}
func readCSV(filePath string) ([][]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	return records, nil
}

func sendEmail(server string, port int, from, password, to, subject, body string) error {
	auth := smtp.PlainAuth("", from, password, server)
	msg := fmt.Sprintf("To: %s\r\n"+
		"Subject: %s\r\n"+
		"\r\n"+
		"%s", to, subject, body)

	err := smtp.SendMail(fmt.Sprintf("%s:%d", server, port), auth, from, []string{to}, []byte(msg))
	if err != nil {
		return err
	}

	return nil
}
