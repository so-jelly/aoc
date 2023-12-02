package main

import "testing"

var dirSizeTestData = `$ cd /
$ ls
dir a
14848514 b.txt
8504156 c.dat
dir d
$ cd a
$ ls
dir e
29116 f
2557 g
62596 h.lst
$ cd e
$ ls
584 i
$ cd ..
$ cd ..
$ cd d
$ ls
4060174 j
8033020 d.log
5626152 d.ext
7214296 k`

func TestDirSize(t *testing.T) {

	expect := 95437

	smallDirsSize, delDirSize := DirSize([]byte(dirSizeTestData), 100000)

	if smallDirsSize != expect {
		t.Errorf("have %v, want %v", smallDirsSize, expect)
	}

	expectDel := 24933642

	if delDirSize != expectDel {
		t.Errorf("have %v, want %v", delDirSize, expectDel)
	}

}

var dirSize, delDir int

func BenchmarkDirSize(b *testing.B) {
	for n := 0; n < b.N; n++ {
		dirSize, delDir = DirSize([]byte(dirSizeTestData), 100000)
	}

}
