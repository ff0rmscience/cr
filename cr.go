package main

import "fmt"
import "math/rand"

const S = 27

type row [S]int
type key [S]row
type machine struct {
	k key
	s row
	a string
	r int
}

func main() {
	var m machine
	m.init()
	m.r = 5
	m.show()
	for i:= 0; i < 10; i++ {
		p := m.randomPlaintext(50)
		c := m.encrypt(p)
		u := m.decrypt(c)
		m.printWord(p)
		m.printWord(c)
		m.printWord(u)
		fmt.Println()
	}
}

func (m machine) comp() [S]int {
	var c [S]int
	for i := 0; i < S; i++ {
		c[i] = m.f(i)
	}
	return c
}

func (m machine) report() {
	fmt.Println("machine key = ")
	for i := 0; i < S; i++ {
		fmt.Println(m.k[i])
	}
	fmt.Println("machine spin = ")
	fmt.Println(m.s)
}
func (m machine) show() {
	for i := 0; i < S; i++ {
		for j := 0; j < S; j++ {
			fmt.Printf("%c", m.a[m.k[i][(j+m.s[i])%S]])
		}
		fmt.Println()
	}
	fmt.Println()
}

func (m *machine) init() {
	for i := 0; i < S; i++ {
		copy(m.k[i][:], rand.Perm(S))
	}
	m.a = "abcdefghijklmnopqrstuvwxyz_"
	m.r = 1
}

func (m machine) f(x int) int {
	for i := 0; i < S; i++ {
		x = m.k[i][(x+m.s[i])%S]
	}
	return x
}

func (m *machine) autospin(r int) {
	var s [S]int
	m.reset()
	for j:=0; j < r; j++ {
		for i := 0; i < S; i++ {
			s[i] = (m.s[i] + m.f(i)) % S
		}
		copy(m.s[:], s[:])
	}
}

func (m *machine) encryptOnce(p []int) []int {
	var c []int
	var d int
	for i := 0; i < len(p); i++ {
		d = m.f(p[i])
		c = append(c, d)
		m.spinOn(m.k[d])
	}
	return c
}

func (m *machine) encrypt(p []int) []int {
	for r := 0; r < m.r; r++ {
		m.autospin(r)
		p = m.encryptOnce(p)
		p = m.reverse(p)
	}
	return p
}

func (m *machine) decrypt(c []int) []int {
	for r:= 0; r < m.r; r++ {
		m.autospin(m.r - 1 - r)
		c = m.reverse(c)
		c = m.decryptOnce(c)
	}
	return c
}

func (m *machine) decryptOnce(c []int) []int {
	var p []int
	var d int
	for i := 0; i < len(c); i++ {
		for d = 0; m.f(d) != c[i]; d++ {
		}
		p = append(p, d)
		m.spinOn(m.k[m.f(d)])
	}
	return p
}

func (m *machine) spinOn(r row) {
	for i := 0; i < S; i++ {
		m.s[i] = (m.s[i] + m.k[r[i]][i])%S
	}
}

func (m machine) printWord(w []int) {
	for i:= 0; i < len(w); i++ {
		fmt.Printf("%c", m.a[w[i]])
	}
	fmt.Printf("\n")
}

func (m machine) randomPlaintext(n int) []int {
	var p []int
	for i := 0; i < n; i++ {
		p = append(p, rand.Intn(S))
	}
	return p
}

func (m *machine) reset() {
	for i:=0;i < S;i++ {
		m.s[i] = 0
	}
}


func (m machine) reverse(n []int) []int {
	for i := 0; i < len(n)/2; i++ {
		j := len(n) - i - 1
		n[i], n[j] = n[j], n[i]
	}
	return n
}
