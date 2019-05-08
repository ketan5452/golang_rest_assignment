Clone the repository:

git clone https://github.com/ketan5452/golang_rest_assignment.git

Run below command to run the project:

go run main.go

How to test?

1. Valid test

Hit URL: 
http://localhost:8081/golang?q=bye

Expected Response:
{"Status":200,"Response":"Hello"}

1. Not Valid test

Hit URL: 
http://localhost:8081/golang?q=bye

Expected Response:
{"Status":400,"Response":"Not Found"}
