# implement size match rule
# implement ^/>/< rule
# implement pip math rules (with config file mapping pips to sizes)
# write out the rules for each of the example rules on the website
# Make all beginning rules work http://www.koryheath.com/zendo/rules-for-beginners/
# Deploy somewhere publically accessible
# Save past games and view old games at GET /games/old
# Make chars chars instead of strings
# Persistence across server restarts, possibly using Bolt https://github.com/boltdb/bolt

Rules to make work

DONE Number of pieces in rule is satisfied
DONE Number of pieces excluded in rule is avoided
DONE Multiple rules (must be at least one but not 3 pieces) can be parsed

all its pieces are the same color.
all its pieces are the same size.
all its pieces are flat.
it contains at least one red piece.
it contains at least one small piece.
it contains at least one piece of each of the four colors.
it contains no green pieces.
it contains no large pieces.
it contains at least one medium yellow piece.
it contains exactly two pieces.
it contains two or more upright pieces.
it contains a piece pointing at another piece.
it contains an ungrounded piece.
it contains at least one green piece and at least one blue piece.
it contains at least two pieces that are touching each other.
