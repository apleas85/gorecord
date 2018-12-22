# GORECORD

Gorecord is a golang command line tool used to read records from file.  Files must have the fields; LastName,FirstName,Gender,FavoriteColor,DateOfBirth and be seperated in one of three ways 
LastName | FirstName | Gender | FavoriteColor | DateOfBirth
LastName, FirstName, Gender, FavoriteColor, DateOfBirth
LastName FirstName Gender FavoriteColor DateOfBirth

The user has the option to sort that data in three ways
1. By gender
2. By last name
3. by birthday

# Installation
1. clone into workspace and cd into project directory
2. install all dependencies ```go dep ensure```
3. build ```go build```

# Command Line 
- There are three available flags for read -byGender -byBirthday -byLastName
- ```./gorecords read --files={file1},{file2},{file3} {flag}```
- To start a server ```./gorecords server``` 
- For help ```./gorecords --help```

# Server Mode
The server will start listen on port 9010.  There are 4 paths; 
1. POST /records - Post a single data line in any of the 3 formats supported by your existing code
2, GET /records/gender - returns records sorted by gender
3. GET /records/birthdate - returns records sorted by birthdate
4. GET /records/name - returns records sorted by name

