package main

import (
	"fmt"
	"log"
	"path/filepath"
	"sort"
	"strconv"
	"strings"

	"github.com/j18e/adventofcode/2022/common"
)

const (
	diskSize      = 70000000
	requiredSpace = 30000000
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	lines := common.ReadInput("input.txt")
	root := populate(lines)
	part1 := 0
	dirList := root.List()
	for _, dir := range dirList {
		size := dir.Size()
		if size < 100000 {
			part1 += size
		}
	}

	sort.Slice(dirList, func(i, j int) bool {
		return dirList[i].Size() < dirList[j].Size()
	})

	needed := requiredSpace - (diskSize - root.Size())

	var part2 int
	for _, dir := range dirList {
		if dir.Size() >= needed {
			part2 = dir.Size()
			break
		}
	}
	fmt.Println("part 1:", part1)
	fmt.Println("part 2:", part2)
	return nil
}

func (d *Dir) List() []*Dir {
	res := []*Dir{d}
	for _, dir := range d.Dirs {
		res = append(res, dir.List()...)
	}
	return res
}

func populate(lines []string) *Dir {
	root := NewDir("/", nil)
	cur := root

	for i, ln := range lines {
		split := strings.Split(ln, " ")
		if split[0] != "$" {
			continue
		}
		switch split[1] {
		case "cd":
			switch split[2] {
			case "/":
				cur = root
			case "..":
				cur = cur.Parent
			default:
				name := split[2]
				cur.AddDir(name)
				cur = cur.Dirs[name]
			}
		case "ls":
		loop:
			for _, ln := range lines[i+1:] {
				split := strings.Split(ln, " ")
				name := split[1]
				switch split[0] {
				case "$":
					break loop
				case "dir":
					cur.AddDir(name)
				default:
					size, err := strconv.Atoi(split[0])
					if err != nil {
						panic(err)
					}
					cur.AddFile(name, size)
				}
			}
		default:
			panic(fmt.Sprintf("unrecognized cmd at line %d: %s", i, split[1]))
		}
	}
	return root
}

func NewDir(name string, parent *Dir) *Dir {
	dir := &Dir{
		Name:   name,
		Parent: parent,
		Dirs:   make(map[string]*Dir),
		Files:  make(map[string]int),
	}
	return dir
}

type Dir struct {
	Name   string
	Parent *Dir
	Dirs   map[string]*Dir
	Files  map[string]int
	size   *int
}

func (d *Dir) AddDir(name string) *Dir {
	if d.Dirs[name] != nil {
		return d.Dirs[name]
	}
	dir := NewDir(name, d)
	d.Dirs[name] = dir
	return dir
}

func (d *Dir) AddFile(name string, size int) {
	d.Files[name] = size
}

func (d *Dir) Pwd() string {
	stack := []string{d.Name}
	for p := d.Parent; p != nil; p = p.Parent {
		stack = append([]string{p.Name}, stack...)
	}
	return filepath.Join(stack...)
}

func (d *Dir) Size() int {
	if d.size != nil {
		return *d.size
	}
	total := 0
	for _, d := range d.Dirs {
		total += d.Size()
	}
	for _, f := range d.Files {
		total += f
	}
	d.size = &total
	return total
}
