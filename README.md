# challenge
Challenge

URL: https://hiveapplogin.tk

To register you need to visit https://hiveapplogin.tk/register, after registered you have to visit https://hiveapplogin.tk/login

the backups are being created every hour, to be able to decrypt you can use the cryptortool that was written in go, instructions of how to use it can be seen executing
the tool without any parameters.

e .g **./cryptortool -operation decrypt -s arquivo.txt.zzz -p "password" -d "file.tar.gz"**

extract the tar.gz file **tar -xzvf file.tar.gz** and in the folder check the Users.sql file to verify the users from the database.

to deploy the application locally you first need to create the **local_network** network in docker.
e .g **docker network create local_network**

replace the content of .env file with the content of .env.local file

**docker-compose up**



