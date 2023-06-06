package main

func main() {
	// go tem if - switch e select
	// no go não tem else if ou elsif

	// == != < >
	// && e ||

	if 1 < 2 {
		println("1 é menor que 2")
	} else {
		println("1 não é menor que 2")
	}

	a := 1

	switch a {
	case 1:
		println("1")
	default:
		println("0")
	}

}
