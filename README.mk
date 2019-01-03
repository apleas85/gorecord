# GORECORD

Gorecord is a golang command line tool used to read records from file.  Files must have the fields; LastName,FirstName,Gender,FavoriteColor,DateOfBirth and be seperated in one of three ways 
LastName | FirstName | Gender | FavoriteColor | DateOfBirth
LastName, FirstName, Gender, FavoriteColor, DateOfBirth
LastName FirstName Gender FavoriteColor DateOfBirth

There are 3 test files provided located in the project root directory; spaceRecords.txt, commaRecords.txt and pipeRecords.txt.  You can also provide your own files.
The user has the option to sort that data in three ways
1. By gender
2. By last name
3. by birthday

# Installation
1. clone into workspace and cd into project directory
2. install all dependencies rungo dep ensure
3. to build run go build

# Command Line 
- There are three available flags for read -byGender -byBirthday -byLastName, if multiple flags are set the precedents is byGender, byBirthday, and then byLastName
- ./gorecords read --files={file1},{file2},{file3} {flag}```
- To start a server run ./gorecords server 
- For help run ./gorecords --help

# Server Mode
The server will start listen on port 9010.  There are 4 paths; 
1. POST /records - Post a single data line in any of the 3 formats supported by your existing code
2, GET /records/gender - returns records sorted by gender
3. GET /records/birthdate - returns records sorted by birthdate
4. GET /records/name - returns records sorted by name


