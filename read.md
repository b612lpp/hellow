08.03.2025
Hello project has been started.
No concreet functionality just testing features. 

Now app calculates and returns incoming integer from json
curl -H "Content-Type: application/json" -d '{"name":22}' http://localhost:8080/calc/


curl -X GET -H "Content-Type: application/json" -d '{"value":22, "cost":30}' http://localhost:8080/calc/


Next step is logging. Added slog events: server start, main page, API happy and wrong request.