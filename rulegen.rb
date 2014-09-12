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