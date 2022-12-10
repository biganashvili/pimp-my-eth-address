# pimp-my-eth-address
This tool uses a brute-force approach to generate an address that matches the given criteria.

Available flags:
-p prefix
-c contains
-s suffix
## Usage
```
//generate address that starts with 1987
go run main.go -p="1987"

//generate address that ends with 1987
go run main.go -s="1987"


//generate address that starts whith 0, contains C0DE, and ends with 9
go run main.go -p="0" -c="C0DE" -s="9"
```
