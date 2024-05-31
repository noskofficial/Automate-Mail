# Automate Mail
Automatically send emails.

## Instruction
This program reads a csv file expecting `name` and `email` at the second and third column respectively.

### Configuration
To send an email a username and password of the sender account is required.
The program reads a `.env` file named `config.env`. So create a `config.env` file in the root directory of the project
with the following contents:
```
SenderEmail=senderemailhere
SenderPassword=senderpasswordhere
EmailSubject="This is a test email"
```

### Adding Email Content
It is recommended to add a `.txt` file in the template directory with the email contents.

To access the name of the reciever use `{{.Name}}` anywhere within the template file.

### Running the program 
```bash
go run mailer.go <csv_file> --template <template_file_path>
```
