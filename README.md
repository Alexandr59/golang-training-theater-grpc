# golang-training-Theater

## Description

##### Backend Application for storing theater

|Path|Method|Description|
|:---:|:---:|:---:|
|```/api/v1/tickets```|```GET```|get all tickets|
|```/api/v1/posters```|```GET```|get all posters|
|```/api/v1/users/{id}```|```GET```|get all users by ```account id```|
|```/api/v1/account/{id}```|```GET```|get account by ```id```|
|```/api/v1/genre/{id}```|```GET```|get genre by ```id```|
|```/api/v1/hall/{id}```|```GET```|get hall by ```id```|
|```/api/v1/location/{id}```|```GET```|get location by ```id```|
|```/api/v1/performance/{id}```|```GET```|get performance by ```id```|
|```/api/v1/place/{id}```|```GET```|get place by ```id```|
|```/api/v1/poster/{id}```|```GET```|get poster by ```id```|
|```/api/v1/price/{id}```|```GET```|get price by ```id```|
|```/api/v1/role/{id}```|```GET```|get role by ```id```|
|```/api/v1/schedule/{id}```|```GET```|get schedule by ```id```|
|```/api/v1/sector/{id}```|```GET```|get sector by ```id```|
|```/api/v1/ticket/{id}```|```GET```|get ticket by ```id```|
|```/api/v1/user/{id}```|```GET```|get user by ```id```|
|```/api/v1/account/{id}```|```DELETE```|delete account by ```id```|
|```/api/v1/genre/{id}```|```DELETE```|delete genre by ```id```|
|```/api/v1/hall/{id}```|```DELETE```|delete hall by ```id```|
|```/api/v1/location/{id}```|```DELETE```|delete location by ```id```|
|```/api/v1/performance/{id}```|```DELETE```|delete performance by ```id```|
|```/api/v1/place/{id}```|```DELETE```|delete place by ```id```|
|```/api/v1/poster/{id}```|```DELETE```|delete poster by ```id```|
|```/api/v1/price/{id}```|```DELETE```|delete price by ```id```|
|```/api/v1/role/{id}```|```DELETE```|delete role by ```id```|
|```/api/v1/schedule/{id}```|```DELETE```|delete schedule by ```id```|
|```/api/v1/sector/{id}```|```DELETE```|delete sector by ```id```|
|```/api/v1/ticket/{id}```|```DELETE```|delete ticket by ```id```|
|```/api/v1/user/{id}```|```DELETE```|delete user by ```id```|

|Path|Method|Description|Body example|
|:---:|:---:|:---:|:---|
|```/api/v1/account```|```POST```|create new account|```{"FirstName":"ExampleFirstName","LastName":"ExampleLastName","PhoneNumber":"ExamplePhoneNumber","Email":"Example@gmail.com"}```|
|```/api/v1/genre```|```POST```|create new genre|```{"Name":"ExampleName"}```|
|```/api/v1/hall```|```POST```|create new hall|```{"AccountId":"1","Name":"ExampleName","Capacity":"9999","LocationId":"1"}```|
|```/api/v1/location```|```POST```|create new location|```{"AccountId":"1","Address":"ExampleAddress","PhoneNumber":"+777777777777"}```|
|```/api/v1/performance```|```POST```|create new performance|```{"AccountId":"1","Name":"ExampleName","GenreId":"1","Duration":"1:00"}```|
|```/api/v1/place```|```POST```|create new place|```{"SectorId":"10","Name":"1"}```|
|```/api/v1/poster```|```POST```|create new poster|```{"AccountId":"1","ScheduleId":"6","Comment":"ExampleCommit"}```|
|```/api/v1/price```|```POST```|create new price|```{"AccountId":"1","SectorId":"9","PerformanceId":"1","Price":"99"}```|
|```/api/v1/role```|```POST```|create new role|```{"Name":"ExampleName"}```|
|```/api/v1/schedule```|```POST```|create new schedule|```{"AccountId":"1","PerformanceId":"1","Date":"2021-04-25 13:00","HallId":"1"}```|
|```/api/v1/sector```|```POST```|create new sector|```{"Name":"ExampleName"}```|
|```/api/v1/ticket```|```POST```|create new ticket|```{"AccountId":"1","ScheduleId":"6","PlaceId":"1","DateOfIssue":"now()","paid":"true","reservation":"true","destroyed":"true"}```|
|```/api/v1/user```|```POST```|create new user|```{"AccountId":"1","FirstName":"ExampleFirstName","LastName":"ExampleLastName","RoleId":"1","LocationId":"1","PhoneNumber":"+777777777777"}```|
|```/api/v1/account```|```PUT```|update account|```{"Id":"18","FirstName":"ExampleFirstName","LastName":"ExampleLastName","PhoneNumber":"ExamplePhoneNumber","Email":"Example@gmail.com"}```|
|```/api/v1/genre```|```PUT```|update genre|```{"Id":"12","Name":"ExampleName"}```|
|```/api/v1/hall```|```PUT```|update hall|```{"Id":"6","AccountId":"1","Name":"ExampleName","Capacity":"9999","LocationId":"1"}```|
|```/api/v1/location```|```PUT```|update location|```{"Id":"6","AccountId":"1","Address":"ExampleAddress","PhoneNumber":"+777777777777"}```|
|```/api/v1/performance```|```PUT```|update performance|```{"Id":"6","AccountId":"1","Name":"ExampleName","GenreId":"1","Duration":"1:00"}```|
|```/api/v1/place```|```PUT```|update place|```{"Id":"11","SectorId":"10","Name":"2"}```|
|```/api/v1/poster```|```PUT```|update poster|```{"Id":"7","AccountId":"1","ScheduleId":"6","Comment":"ExampleCommit"}```|
|```/api/v1/price```|```PUT```|update price|```{"Id":"36",AccountId":"1","SectorId":"9","PerformanceId":"1","Price":"99"}```|
|```/api/v1/role```|```PUT```|update role|```{"Id":"8","Name":"ExampleName"}```|
|```/api/v1/schedule```|```PUT```|update schedule|```{"Id":"11","AccountId":"1","PerformanceId":"1","Date":"2021-04-25 13:00","HallId":"1"}```|
|```/api/v1/sector```|```PUT```|update sector|```{"Id":"19","Name":"ExampleName"}```|
|```/api/v1/ticket```|```PUT```|update ticket|```{"Id":"26","AccountId":"1","ScheduleId":"6","PlaceId":"1","DateOfIssue":"now()","paid":"true","reservation":"true","destroyed":"true"}```|
|```/api/v1/user```|```PUT```|update user|```{"Id":"8","AccountId":"1","FirstName":"ExampleFirstName","LastName":"ExampleLastName","RoleId":"1","LocationId":"1","PhoneNumber":"+777777777777"}```|

## Usage

1. Run server on port ```8080```

> ```go run ./theater-gorm/cmd/main.go```

2. Open URL ```http://localhost:8080```
