package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var sc = bufio.NewScanner(os.Stdin)
var wr = bufio.NewWriter(os.Stdout)

const MAX_QUEUE_LENGTH = 100000

type Process struct {
	name      string
	initTime  int
	remaining int
	endAt     int
}

type Queue struct {
	baseList [MAX_QUEUE_LENGTH]*Process
	head     int
	length   int
}

var queue = Queue{head: 0, length: 0}

func (q *Queue) IsEmpty() bool {
	return q.length == 0
}

func (q *Queue) lookup(i int) *Process {
	idx := (q.head + i) % MAX_QUEUE_LENGTH
	return q.baseList[idx]
}

func (q *Queue) Find(i int) *Process {
	if i >= q.length {
		panic("index out of range")
	}
	return q.lookup(i)
}

func (q *Queue) Enqueue(p *Process) {

	if q.length >= MAX_QUEUE_LENGTH {
		panic("Too many elements. Cannot enque any further")
	}

	idx := (q.head + q.length) % MAX_QUEUE_LENGTH
	q.length += 1
	q.baseList[idx] = p
}

func (q *Queue) Dequeue() *Process {
	idx := q.head
	q.head = (idx + 1) % MAX_QUEUE_LENGTH
	q.length -= 1
	return q.baseList[idx]
}

func newProcess(name string, initTime int) Process {
	return Process{name: name, initTime: initTime, remaining: initTime}
}

func (p *Process) Consume(unitQ int, clock int) int {
	if p.remaining > unitQ {
		p.remaining -= unitQ
		return unitQ
	} else {
		consumed := p.remaining
		p.remaining = 0
		p.endAt = clock + consumed
		return consumed
	}
}

func (p Process) String() string {
	return fmt.Sprintf("Process \"%s\": {initTime: %d, remaining: %d, endAt: %d}", p.name, p.initTime, p.remaining, p.endAt)
}

func (q Queue) String() string {
	elems := make([]string, q.length)

	for i := 0; i < q.length; i += 1 {
		elems[i] = fmt.Sprint(q.lookup(i))
	}
	return "(" + strings.Join(elems, ", ") + ")"
}

func (p *Process) IsRemaining() bool {
	return p.remaining > 0
}

func scanStr() string {
	sc.Scan()
	return sc.Text()
}

func scanInt() int {
	sc.Scan()
	i, err := strconv.Atoi(sc.Text())
	if err != nil {
		panic(err)
	}

	return i
}

func schedule(unitQ int) {
	for clock := 0; !queue.IsEmpty(); {
		p := queue.Dequeue()
		clock += p.Consume(unitQ, clock)
		if p.IsRemaining() {
			queue.Enqueue(p)
		} else {
			fmt.Fprintf(wr, "%s %d\n", p.name, p.endAt)
		}
	}
}

func fetchCond() (n int, unitQ int) {
	// return 5, 100
	return scanInt(), scanInt()
}

func buildProcess(i int) Process {
	// p := [][2]string{
	// 	{"p1", "150"}, {"p2", "80"}, {"p3", "200"}, {"p4", "350"}, {"p5", "20"},
	// }[i]

	// n, _ := strconv.Atoi(p[1])
	// return newProcess(p[0], n)
	return newProcess(scanStr(), scanInt())
}

func main() {
	sc.Split(bufio.ScanWords)
	defer wr.Flush()

	n, unitQ := fetchCond()

	for i := 0; i < n; i += 1 {
		p := buildProcess(i)
		queue.Enqueue(&p)
	}

	schedule(unitQ)
}
