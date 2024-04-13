# MarkasBali Go Final Project

This is a repo for markasbali ft kominfo final project for hacker class go <br>
first thing you must do to run it is running this command <br><br>
<code>go mod tidy</code> <br><br>
then you have to run all these go services with command (1 terminal tab for 1 service) <br><br>
<code>go run main.go</code> <br><br>
for frontend (mobile app) it use ionic framework
so you have to run <br><br>
<code>npm ci</code> <br><br>
and then run it locally <br><br>
<code>ionic serve</code>

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
- get specific database with latest backup information
- restore all databases with latest backup file

<br><br>
this repo maintained by <b>Komang Mardika</b> and <b>Dewa Antasena</b>