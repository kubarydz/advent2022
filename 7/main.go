package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/kubarydz/advent2022/utils"
)

func main() {
	input := utils.ReadInput("sample")

	fmt.Printf("samples\n")
	resp := calculateV1(input)
	fmt.Printf("sample 1 solution: %d\n", resp)
	if resp != 95437 {
		fmt.Printf("test nr 1 failed\n")
		return
	}

	resp = calculateV2(input)
	fmt.Printf("sample 2 solution: %d\n", resp)
	if resp != 24933642 {
		fmt.Printf("test nr 2 failed\n")
		return
	}

	fmt.Printf("inputs\n")
	input = utils.ReadInput("input")
	resp = calculateV1(input)
	fmt.Printf("input 1 solution: %d\n", resp)

	resp = calculateV2(input)
	fmt.Printf("input 2 solution: %d\n", resp)
}

func calculateV1(input []byte) int {
	lines := utils.ChunkInput(input)

	fileSystem := buildFileSystem(lines)
	fileSystem.calculateInnerDirSizes()

	sumSmallDirs := 0
	calculateSmallDirsTotal(fileSystem, &sumSmallDirs)

	return sumSmallDirs
}

func calculateV2(input []byte) int {
	lines := utils.ChunkInput(input)

	fileSystem := buildFileSystem(lines)
	fileSystem.calculateInnerDirSizes()

	smallestPossible := 0
	wantedSize := 30000000 - (70000000 - fileSystem.getTotalDirSize())
	fmt.Printf("wanted size: %d\n", wantedSize)
	findBigEnoughDirSize(fileSystem, &smallestPossible, wantedSize)

	return smallestPossible
}

func isCommand(line []byte) bool {
	return line[0] == '$'
}

func isChangeDir(line []byte) bool {
	return isCommand(line) && line[2] == 'c' && line[3] == 'd'
}

func isLs(line []byte) bool {
	return isCommand(line) && line[2] == 'l' && line[3] == 's'
}

func isDir(line []byte) bool {
	return line[0] == 'd' && line[1] == 'i' && line[2] == 'r'
}

func isFile(line []byte) bool {
	return !isCommand(line) && !isDir(line)
}

func getDirName(line []byte) string {
	return string(line[4:])
}

func getChangeDirDestination(line []byte) string {
	return string(line[5:])
}

func getFileSize(line []byte) int {
	s := string(line)
	number := strings.Split(s, " ")[0]
	fileSize, _ := strconv.Atoi(number)
	return fileSize
}

type dir struct {
	name          string
	filesSize     int
	innerDirs     []*dir
	previousDir   *dir
	innerDirsSize int
}

func (d dir) getTotalDirSize() int {
	return d.filesSize + d.innerDirsSize
}

func (d *dir) calculateInnerDirSizes() int {
	sum := 0
	for _, innerDir := range d.innerDirs {
		sum += innerDir.calculateInnerDirSizes()
	}
	d.innerDirsSize = sum
	return d.filesSize + d.innerDirsSize
}

func calculateSmallDirsTotal(d *dir, sum *int) {
	if (d.filesSize + d.innerDirsSize) <= 100000 {
		*sum += d.filesSize + d.innerDirsSize
	}
	for _, inner := range d.innerDirs {
		calculateSmallDirsTotal(inner, sum)
	}
}

func buildFileSystem(lines [][]byte) *dir {
	topDir := &dir{name: "/", innerDirs: []*dir{}}

	currentDir := topDir
	for i := 1; i < len(lines); i++ {
		if isChangeDir(lines[i]) {
			destName := getChangeDirDestination(lines[i])
			if destName == "/" {
				currentDir = topDir
				continue
			}
			if destName == ".." {
				currentDir = currentDir.previousDir
				continue
			}

			for _, innerDir := range currentDir.innerDirs {
				if destName == innerDir.name {
					currentDir = innerDir
					break
				}
			}
			continue
		}

		if isLs(lines[i]) {
			for i++; ; i++ {
				if i >= len(lines) || isCommand(lines[i]) {
					break
				}
				if isDir(lines[i]) {
					currentDir.innerDirs = append(currentDir.innerDirs, &dir{name: getDirName(lines[i]), previousDir: currentDir})
				}
				if isFile(lines[i]) {
					currentDir.filesSize += getFileSize(lines[i])
				}
			}
			i--
			continue
		}
	}

	return topDir
}

func findBigEnoughDirSize(d *dir, size *int, wantedSize int) {
	if d.getTotalDirSize() >= wantedSize && (*size == 0 || *size > d.getTotalDirSize()) {
		*size = d.getTotalDirSize()
	}
	if d.innerDirsSize > *size {
		return
	}
	for _, inner := range d.innerDirs {
		findBigEnoughDirSize(inner, size, wantedSize)
	}

}
