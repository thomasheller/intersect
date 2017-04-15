# intersect

intersect prints the lines found in all given files, sorted.

```sh
$ go get github.com/thomasheller/intersect
$ cat a.txt
123
foo
bar
baz
$ cat b.txt
baz
bar
123
456
$ cat c.txt
456
foo
bar
baz
$ intersect a.txt b.txt c.txt
bar
baz
$ intersect a.txt b.txt
123
bar
baz
$ intersect a.txt c.txt
bar
baz
foo
$ intersect b.txt c.txt
456
bar
baz
```
