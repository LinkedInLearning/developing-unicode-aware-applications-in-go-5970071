package graph

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

// ParseGraph parses a graph containing lines like "A → B"
// and returns map of src → []dest.
func ParseGraph(r io.Reader) (map[string][]string, error) {
	graph := make(map[string][]string)
	s := bufio.NewScanner(r)
	lnum := 0
	for s.Scan() {
		lnum++
		line := strings.TrimSpace(s.Text())
		if len(line) == 0 {
			continue
		}

		src, dest, err := parseLine(line)
		if err != nil {
			return nil, fmt.Errorf("%d: %w", lnum, err)
		}

		graph[src] = append(graph[src], dest)
	}

	if err := s.Err(); err != nil {
		return nil, err
	}

	return graph, nil
}

func parseLine(line string) (string, string, error) {
	/* Option 1: Index + len
	i := strings.Index(line, "→")
	if i == -1 {
		return "", "", fmt.Errorf("%q: no → in line", line)
	}

	src := strings.TrimSpace(line[:i])
	dest := strings.TrimSpace(line[i+len("→"):])
	return src, dest, nil
	*/

	src, dest, found := strings.Cut(line, "→")
		if !found {
		return "", "", fmt.Errorf("%q: no → in line", line)
	}

	return strings.TrimSpace(src), strings.TrimSpace(dest),  nil
}
