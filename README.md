# Automate Mail
Automatically send emails.

## Instruction
This program reads a csv file expecting `name` and `email` at the second and third column respectively.

### Set Credentials
To send an email a username and password of the sender account is required.
The program reada a `.env` file named `cred.env`. So create a `cred.env` file in the root directory of the project
with the following contents:
```
SenderEmail=senderemailhere
SenderPassword=senderpasswordhere 
```
### Running the program 
```bash
go run mailer.go <csv_file>
```
