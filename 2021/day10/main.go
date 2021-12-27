package main

func main() {
	t := []string{
		"{([(<{}[<>[]}>{[]{[(<()>",
		"[[<[([]))<([[{}[[()]]]",
		"[{[{({}]{}}([{[{{{}}([]",
		"[<(<(<(<{}))><([]([]()",
		"<{([([[(<>()){}]>(<<{{",
	}



/*

    {([(<{}[<>[]}>{[]{[(<()> - Expected ], but found } instead.
    [[<[([]))<([[{}[[()]]] - Expected ], but found ) instead.
    [{[{({}]{}}([{[{{{}}([] - Expected ), but found ] instead.
    [<(<(<(<{}))><([]([]() - Expected >, but found ) instead.
    <{([([[(<>()){}]>(<<{{ - Expected ], but found > instead.

	
    ): 3 points.
    ]: 57 points.
    }: 1197 points.
    >: 25137 points.

		*/
