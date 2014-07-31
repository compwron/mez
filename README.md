This is supposed to be a game, played via GET and POST requests. 

Although it is possible for multiple games to run at the same time, let's start with one at a time. Besides, it's more fun that way. 

````
export GOHOME=<your go home path>
cd $GOHOME/src
git clone <this repo> && cd <this repo>
alias gb="go fmt && go test && go install && ./../../bin/mez"
gb
````

````
{"a":"b"} 
````
