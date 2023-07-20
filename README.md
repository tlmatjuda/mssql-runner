# MS SQL Server Runner

This is a small CLI app meant to integrate with a database of your choice to execute ".sql" files against it.<br/>
It the current integration is for Microsoft SQL Server.<br/>

At the moment I am using my own internal modules so it's not reusable, until I can publish those modules as well.<br/>
It will run SQL files in no particular order, so make sure you run sql chagnes that are not dependant on each other.<br/><br/>

## TECH STACK

For interest sake I used :

* Google's **[Go](https://go.dev/)** v.1.20.x
* IntelliJ IDEA with Go Lang Plugin<br/><br/>

## DRIVING THE APP

Once you have the binary / executable based on the OS you can do the following to drive.<br/>
You will need a json configuration : **"mssql-env-conf.json"** inside a sub-folder named **"configs"**<br/>
So you you will end up having something like the folliwinig in your file system :<br/><br/>

```plaintext
configs/mssql-env-conf.json
mssql-runner
mssql-runner.exe
```
<br/>

The json config structure is as follows :

```json
[
  {
    "environment": "dev",
    "username": "<dev-user>",
    "password": "<dev-pass>",
    "port": 1234,
    "host": "<dev-host>"
  }
]
```
<br/><br/>
You drive this tool by running the executable file.

When doing so you need to supply two parameters

* **Environment Name** : Which will match the **"environment"** field in the json file.
* **Directory Path** : This is the folder where your **".sql"** files are

<br/>

Type something like the following on Unix Based systems :

```bash
./mssql-runner dev /<wherever>/<your>/<sql-files>/<are>
```

Windows OS systems

```powershell
mssql-runner.exe dev /<wherever>/<your>/<sql-files>/<are>
```




You will then be presented with the following on your terminal

```bash
INFO  2023/07/20 09:51:28 You are about to execute SQL files in no particular order towards the : [ dev ] environment
INFO  2023/07/20 09:51:28
INFO  2023/07/20 09:51:28 The database details are as follows :
INFO  2023/07/20 09:51:28 HOST : <dev-host>
INFO  2023/07/20 09:51:28 PORT : 1234
INFO  2023/07/20 09:51:28 USER : <dev-user>
INFO  2023/07/20 09:51:28 PASSWORD : ( Yea right :) )
INFO  2023/07/20 09:51:28
INFO  2023/07/20 09:51:28 If this is correct, please type either : "Yes" to continue or "No" to stop the process.
```





Then you should type either **"Yes"** to continue or **"No"** to terminate

When you have confirmed then you will see something like :

```bash
INFO  2023/07/20 09:49:09 You are about to execute SQL files in no particular order towards the : [ dev ] environment
INFO  2023/07/20 09:49:09
INFO  2023/07/20 09:49:09 The database details are as follows :
INFO  2023/07/20 09:49:09 HOST : <dev-host>
INFO  2023/07/20 09:49:09 PORT : 1433
INFO  2023/07/20 09:49:09 USER : <dev-user>
INFO  2023/07/20 09:49:09 PASSWORD : ( Yea right :) )
INFO  2023/07/20 09:49:09
INFO  2023/07/20 09:49:09 If this is correct, please type either : "Yes" to continue or "No" to stop the process.
Yes
INFO  2023/07/20 09:49:11 Found 1 SQL files
INFO  2023/07/20 09:49:11 Connected to the database successfully!
INFO  2023/07/20 09:49:11 Running SQL File : /<wherever>/<your>/<sql-files>/create-person-tables.sql
INFO  2023/07/20 09:49:11 Process complete
```


Finally you may go and verify your data on the database service to make sure things went fine.<br/><br/>

## CONCLUSION
This cli app is a used to execute SQL files towards your target database. 