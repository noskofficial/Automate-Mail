# Automate Mail
Automatically send emails.

## Instruction
This program reads a csv file expecting name and email at the second and third column respectively.

### Set Credentials
To send an email a username and password is also required of the sender account.
The program reada a .env file named creds.env. So create a creds.env file in the root directory of the project
with the following contents:

SENDEREMAIL=senderemailhere
SENDERPASSWORD=senderpasswordhere 

### Running the program 
```bash
go run mailer.go <csv_file>
```
