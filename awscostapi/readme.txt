for this app...

we need to install golang in our system
need code editer (i'm prefer VS code)

install AWS CLI by given commands if used ubuntu os (snap apt -get aws cli)
configure credential
            -- open terminal and hit commond: aws configure
            -- set credential like, access key, secret key, region, output format etc.
            -- choose region ap-south-1 because service charged is less for this region.

in the AWSCOSTAPI app...
we have used several function each and every function used through main

The function in the file is getDate(), it returns range of date in format YY-MM-DD

The function GetAwsCost() is where we request the information to AWS, we give the details related to what is needed,

we set the level of granularity (daily) or Monthly or yearly as per our requirement.

create a file then store in CSV.
CSV was chosen for ease of use, almost every spreadsheet program can open those and if not the we processed with a simple text file


The main function is where we use the CLI library, for executing commands, and calling functions

after all..
hit command go run main.go
it will genrate a csv file and you simply open it in spreadsheet