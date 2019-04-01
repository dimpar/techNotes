## How to run tests with coverage?

Navigate to a package with tests. Create c.out if not there. Run:
```
go test -coverprofile=c.out 
go tool cover -html=c.out -o coverage.html
open coverage.html
```
Your browser should open coverage.html