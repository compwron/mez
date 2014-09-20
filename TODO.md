Fix golint and govet errors, then add to checkin process
go get code.google.com/p/go.tools/cmd/vet
go vet
go get github.com/golang/lint/golint
go lint

// What part of the rule does a negative apply to? Gotta be the whole thing. So !3^SG means nothing that is 3 AND Small AND G

- t.Skipf("skipping koan chunk addition for now") <- fix this
- disallow multiple ! in rule
- implement pip math rules (with config file mapping pips to sizes)
- write out the rules for each of the example rules on the website
- Make all beginning rules work http://www.koryheath.com/zendo/rules-for-beginners/
- Deploy somewhere publically accessible
- Save past games and view old games at GET /games/old
- Make chars chars instead of strings
- Persistence across server restarts, possibly using Bolt https://github.com/boltdb/bolt

Rules to make work

- DONE Number of pieces in rule is satisfied [Example: "3"]
- DONE Number of pieces excluded in rule is avoided [Example: "!3"] // There can be 1, 2, 4, etc pieces
- DONE Multiple rules (must be at least one but not 3 pieces) can be parsed
- DONE all its pieces are the same color. [Example: "!G, !R, B"] // Therefore all pieces must be the one remaining color, blue
- DONE all its pieces are the same size. [Example: "L"] // All pieces must be L
- DONE all its pieces are flat. [Example: ">"] <- right-lying flat
- DONE it contains at least one red piece. [Example: "1R"]
- DONE it contains at least one small piece. [Example: "1S"]
- DONE it contains at least one piece of each of the four colors. [Example: "1R,1B,1G,1Y"] 
- DONE it contains no green pieces. [Example: "!G"]
- DONE it contains no large pieces. [Example: "!L"]
- DONE it contains at least one medium yellow piece. [Example: "1MY"]
- DONE it contains exactly two pieces. [Example: "2,!3,!4,!5"] # sorry, the specification is not advanced enough to make this a one-chunk rule
- DONE it contains two or more upright pieces. [Example: "2^"] # wait does this actually work
- it contains a piece pointing at another piece.
- it contains an ungrounded piece.
- it contains at least one green piece and at least one blue piece.
- it contains at least two pieces that are touching each other.

More advanced rules:
- exactly half of the pips are from large pieces
- pick some from here http://www.playagaingames.com/games/zendo_some_rules/
