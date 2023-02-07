package main

//global variabel för priserna, en lista av heltal där index motsvarar längden garn
var h []uint

//struct av typ res som sparar maxpriset för en längd garn samt längden garn som heltal, och huruvida den har beräknats som en boolsk variabel
type res struct {
	price    uint
	computed bool
	list     []res
}

//hjälpfunktion för att hitta maximala värdet av två, tar in två heltal och returnerar ett av dem, T(n) = O(1)
func findMax(x, y res) res {
	if x.price > y.price {
		return x
	} else {
		return y
	}
}

//rekursiv funktion, tar in meter garn som heltal och returnerar det maximala värdet för just den mängden garn som ett heltal, värdet sparas i en lista med strukturer av typen res
//T(n) är O(n^2), visas i g)
func P(i uint, r []res) res {
	var v1, v2 res
	if i >= uint(len(h))-1 { //om meter garn överstiger index läggs nollor till i listan
		h = append(h, make([]uint, i-uint(len(h))+1)...)
		r = append(r, make([]res, i-uint(len(r))+1)...)
	}
	r[i] = res{}
	if i == 0 {
		r[i].price = 0
		r[i].computed = true
		return r[i]
	} else {
		v1 = res{}
		for j := uint(1); j <= i; j++ {
			if r[i-j].computed == true {
				v2.price = h[j] + r[i-j].price
			} else {
				vtemp := P(i-uint(j), r)
				v2.price = h[j] + vtemp.price
				v2.list = append(r[i].list, vtemp)
			}
			v1 = findMax(v1, v2)
		}
	}
	r[i].price = v1.price
	r[i].computed = true
	r[i].list = append(r[i].list, v1)
	return r[i]
}

//testfunktion som testar för givna kostnader, ger panik om summorna inte stämmer
func tp() {
	h = []uint{0, 2, 5, 6, 9}
	r := make([]res, len(h))
	if P(0, r).price != 0 {
		panic("FEL")
	}
	if P(1, r).price != 2 {
		panic("FEL")
	}
	if P(2, r).price != 5 {
		panic("FEL")
	}
	if P(3, r).price != 7 {
		panic("FEL")
	}
	if P(4, r).price != 10 {
		panic("FEL")
	}
	if P(5, r).price != 12 {
		panic("FEL")
	}
}

func main() {
	tp()
}
