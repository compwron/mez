This is intended to grow into an implementation of the [game Zendo](https://en.wikipedia.org/wiki/Zendo_(game)) playable via GET and POST requests and hosted on a central server. So far only two simple types of rules (color and count) have been implemented. 

Pull requests are welcome, along with other suggestions. I do not claim to be fluent in golang. 

[![Build Status](https://travis-ci.org/compwron/mez.svg)](https://travis-ci.org/compwron/mez)

To play with an existing game server:
````
fill this in with commandline curl stuff
````

To run the game server locally:

````
export GOHOME=<your go home path>
cd $GOHOME/src
git clone <this repo> && cd <this repo>
go install 
./../../bin/mez
````



Game flow:
====

* A player sets the rule with two koans (one true, one false) (POST /game {"rule": "1^", "true":"1^SG", "false":"1>SG"})
* Other players submit koans (POST /game/koan {"koan":"1>SG"})
* Players view koans and results (GET /game)
* Players guess rule (POST /game/guess {"rule":"3^G"})
* If guess is correct, game resets


Rules Guide
====

Colors:

* "G" At least one piece must be green
* "1G" The same as "G"
* "2G" At least two pieces must be green
* "!G" No pieces may be green
* "!2G" One or three etc pieces may be green, but not two pieces. 
* "GR" Nonsensical rule 
* "G,R" At least one piece must be green. At least one piece must be red. No single-piece koans will be true.

Development:
====

How to run tests and see coverage percentage:
````
go tool cover -html=coverage.out
````

Run just one test:
````
go test -run <testname>  # <- actually this is a regex running against test names
````

Before pushing code:
````
go fmt 
go test 
````

Internal Koan nomenclature: 
====
(for development, not for playing the game)

* koan "1^SG"
* koan "1^SG,2^MB"
* koanChunks ["1^SG","2^MB"]
* koanPieces ["1", "^", "S", "G"] # should never have "," in it
