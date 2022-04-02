package DataHandle

import (
	"bufio"
	"io"
	"os"
)

func ReadConf(fileName string) ([]string, error) {
	f, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	var result []string
	defer f.Close()
	br := bufio.NewReader(f)
	for {
		a, _, c := br.ReadLine()
		if c == io.EOF {
			break
		}
		result = append(result, string(a))
	}
	return result, nil
}
