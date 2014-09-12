options = {:count_negation => ["!", ""],
			:piece_count => [1, 2, 3, 4, 5],
			:piece_orientation => "^",
			:color_negation => ["!", ""],
			:color_count => [1, 2, 3, 4, 5],
			:color => ["R", "G", "B"]
		}

p options

rule_piece_count = 2 # it might be too mean to do do multi=part rules... this game is hard
p (0..rand(rule_piece_count)).map { 
	options.map {|option, values|
		values[rand(values.size)]
	}.join("")
}.join(",")

puts "Example rule explanation: \n"
	+ "\n 1^!3B means there must be one or more pieces, but there may not be exactly three blue pieces."
	+ "\n !4^!4G means that there must not be exactly four pieces (but ther can be more or less), and there may not be exactly four green pieces (rule is kind of redundant)"
	+ "\n 3^!4R,!1^3G means that there must be 3 or more pieces but not 4 red pieces. Also, there may not be exactly one piece and there must be at least 3 green pieces."