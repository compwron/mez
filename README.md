This is supposed to be a game, played via GET and POST requests. 

Although it is possible for multiple games to run at the same time, let's start with one at a time. Besides, it's more fun that way. 

````
export GOHOME=<your go home path>
cd $GOHOME/src
git clone <this repo> && cd <this repo>
alias gb="go fmt && go test && go install && ./../../bin/mez"
gb
````

How to run tests and see coverage percentage:
````
go tool cover -html=coverage.out
````

Game flow:
A player sets the rule with two koans (one true, one false) (POST /game {"rule": "1^", "true":"1^SG", "false":"1>SG"})
Other players submit koans (POST /game/koan {"koan":"1>SG"})
Players view koans and results (GET /game)
Players guess rule (POST /game/guess {"rule":"3^G"})
If guess is correct, game resets

