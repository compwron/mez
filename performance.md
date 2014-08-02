This is an exercise on how to think through the performance implications of an app. 

Use go benchmark against the app

go test -bench . -benchmem -cpu 1,2,4
http://blog.golang.org/profiling-go-programs

Plausible heavy load edge cases:
Many POST /game/koan
Many POST /game/guess
Big json to /game
Big

Plausible attacks:
# Repeatedly POST /game and POST /game/guess with the same thing to keep anyone else from playing


Error cases to look at:
# terrible JSON to every endpoint
# hitting unimplemented endpoints
