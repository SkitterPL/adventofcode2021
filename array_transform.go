package main

func linesToColumnsArray(lines []string) []string {
	var columns = make([]string, len(lines[0]))
	for i := 0; i < len(lines[0]); i++ {
		for _, line := range lines {
			columns[i] += string(line[i])
		}
	}
	return columns
}
