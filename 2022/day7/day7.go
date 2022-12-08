package day7

import (
	"math"
	"strconv"
	"strings"
)

/*
cd /

ls
dir a
14848514 b.txt
8504156 c.dat
dir d

cd a
ls
dir e
29116 f
2557 g
62596 h.lst
cd e
ls
584 i
cd ..
cd ..
cd d
ls
4060174 j
8033020 d.log
5626152 d.ext
7214296 k


*/

func TaskA(commandBlocks []string) int {
	path := Stack{}

	children := make(map[string][]string)
	dirSizes := make(map[string]int)

	for _, block := range commandBlocks {
		parseBlock(block, &path, &children, &dirSizes)

	}

	ans := 0

	for absPath := range dirSizes {
		var size = dfs(absPath, &children, &dirSizes)
		if size < 100000 {
			ans += size
		}
	}

	return ans
}

func dfs(absPath string, children *map[string][]string, dirSize *map[string]int) int {
	size := (*dirSize)[absPath]
	for _, child := range (*children)[absPath] {
		size += dfs(child, children, dirSize)
	}
	return size
}

func parseBlock(block string, path *Stack, children *map[string][]string, dirSize *map[string]int) {
	blockLines := strings.Split(block, "\n")
	command := blockLines[0]
	outputs := blockLines[1:]

	commandParts := strings.Split(command, " ")

	if commandParts[0] == "cd" {
		if commandParts[1] == ".." {
			path.Pop()
		} else {
			path.Push(commandParts[1])
		}
		return
	}
	absPath := strings.Join(*path, "/")

	filesSize := 0

	for _, line := range outputs {
		if !strings.HasPrefix(line, "dir") {
			intVar, _ := strconv.Atoi(strings.Split(line, " ")[0])
			filesSize += intVar
		} else {
			dirName := strings.Split(line, " ")[1]
			(*children)[absPath] = append((*children)[absPath], absPath+"/"+dirName)
		}
	}

	(*dirSize)[absPath] = filesSize

}

func TaskB(commandBlocks []string) int {
	path := Stack{}

	children := make(map[string][]string)
	dirSizes := make(map[string]int)

	for _, block := range commandBlocks {
		parseBlock(block, &path, &children, &dirSizes)

	}
	unused := 70000000 - dfs("/", &children, &dirSizes)
	requiredSpace := 30000000 - unused
	ans := math.MaxUint32
	for absPath := range dirSizes {
		var size = dfs(absPath, &children, &dirSizes)
		if size >= requiredSpace {
			ans = minValue(ans, size)
		}
	}

	return ans
}

func minValue(val1 int, val2 int) int {
	if val1 < val2 {
		return val1
	}
	return val2
}
