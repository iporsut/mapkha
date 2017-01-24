package mapkha

import (
	"bufio"
	"bytes"
	"io/ioutil"
	"log"
	"os"
	"path"
	"runtime"
)

// LoadDict is for loading a word list from file
func LoadDict(path string) (PrefixTree, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	b, err := ioutil.ReadAll(f)
	if err != nil {
		log.Fatal("could not read input:", err)
	}
	scanner := bufio.NewScanner(bytes.NewReader(b))
	wordWithPayloads := make([]string, 0)
	for scanner.Scan() {
		if line := scanner.Text(); len(line) != 0 {
			wordWithPayloads = append(wordWithPayloads, line)
		}
	}
	return MakePrefixTree(wordWithPayloads), nil
}

func MakeDict(words []string) PrefixTree {
	return MakePrefixTree(words)
}

// LoadDefaultDict - loading default Thai dictionary
func LoadDefaultDict() (PrefixTree, error) {
	_, filename, _, _ := runtime.Caller(0)
	return LoadDict(path.Join(path.Dir(filename), "tdict-std.txt"))
}
