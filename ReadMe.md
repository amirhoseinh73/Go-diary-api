# Restful Api With Go

> These features are available in this repo

1. Dockerize Go &check;
2. Registration with a username and password &check;
3. Login with a username and password &check;
4. Create a new diary entry
5. Retrieve all diary entries


## Login API

> POST: /auth/login --> diary_api/controller.Login


## Register API

> POST: /auth/register --> diary_api/controller.Register


## add Entry

> POST: /api/entry --> diary_api/controller.AddEntry


## get all Entries

> GET: /api/entry --> diary_api/controller.GetAllEntries