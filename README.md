# GO backup and restore databases

This is a repo for <br><b>markasbali ft kominfo final project for hacker class: programming with golang</b> <br>
The app features such as backup and restore multiple database, it implements some go great features such as:
- concurrency with channeling and goroutine 
- gorm orm
- go fiber a restful api framework for go
- microservice
- message exchange using gRPC
- websocket using gorilla websocket
- unit testing
<br><br>
So what is first thing you must do to run it? <br><br>
1. run this command <br><br>
<code>go mod tidy</code> <br><br>
2. then you must create empty MySQL / MariaDB databases see <br><br><code>cli-service/config.json</code> for databases list<br><br>
3. then you have to run all these go services with command (1 terminal tab for 1 service) there are 3 service atm cli/web/webscoket<br><br>
<code>go run main.go</code> <br><br>
4. for frontend (mobile app) it use ionic framework
so you have to run <br><br>
<code>npm i</code> <br><br>
You need to use node lts (version 20 atm) and then run it locally <br><br>
<code>ionic serve</code>
<br><br>
5. create tmp folder in root of cli-service if not exists<br>
6. and create storage folder in root of web-service if also not exists

### project structure
 
- CLI Service (cli-service) / backend
- Web Service (web-service) / backend
- Web Socket Service (ws-service) / backend
- Mobile App (client/tabs) / frontend
- gRPC (proto) / backend

### what you can do within the app: <br>
- reset maintained databases (drop all tables)
- migrate and seeding databases with table books and cars with csv data
- run backup all databases
- get all list of every backed up databases with latest backup file information
- get specific database with all backup information
- restore all databases with latest backup file
- catch error using websocket

<br>
this repo maintained by <b>Komang Mardika</b> with contributor <b>Dewa Antasena</b>