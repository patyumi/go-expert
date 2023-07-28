package main

import (
	"fmt"
	"time"
)

// ---
// Conceitos básicos
// usando o Go o SO consegue separar os processos gerados idenependetemente
// cada processo pode ter THREADS (momentos de processamento)
// quando threads compartilham memória (estão no mesmo processo) conseguem alterar a mesma variável

// ---
// Compartilhamento de memória
// pode gerar RACE CONDITION (threads tentando alterar valor ao mesmo tempo)
// para resolver esse problema, o Go tem o Mutex (mutual exclusion) -> dá um lock na variável quando ela tá sendo usada

// ---
// Concorrência x Paralelismo x Go
// concorrencia ativa a possibilidade de usar o paralelismo

// CONCORRÊNCIA ocorre quando uma tarefa acontece depois da outra de forma sequencial
// PARALELISMO ocorre quando eu consigo executar duas tarefas ao mesmo tempo, independente do término de cada uma
// Go -> faz o paralelismo de acordo com a quantidade de cpus que sua máquina tem
// se tiver uma única CPU ele só vai rodar concorrente

// ---
// Multithreading antigamente
// Processo -> (syscall) -> SO -> threads
// toda vez que uma nova thread vai ser gerada ela começa com 1 mb de memória alocada/padrão
// trabalhar com isso vc usa bastante memória e isso precisa ser observado, além disso, alocações e logs acrescentam nisso

// Scheduler
// gerencia as tarefas/threads
// preemptivo  -> tem tempo determinado por tarefa, a cada x tempo ele troca de tarefa
// cooperativo -> espera uma tarefa terminar para iniciar outra

// o go não utiliza esse scheduler padrão

// ---
// Multithreading no Golang
// para gerar novas threads o go não faz a syscall, ele tem o próprio sistema e suas próprias threads
// elas são green threads ou threads in user land -> elas usam 2kb de memória inicialmente
// Go também tem seu próprio scheduler que trabalha de maneira cooperativa

func task(name string) {
	for i := 0; i < 10; i++ {
		fmt.Printf("%d : Task %s is running\n", i+1, name)
	}
}

// thread main já é 1 thread
func main() {

	go task("A")
	go task("B")

	go func() {
		for i := 0; i < 10; i++ {
			fmt.Printf("%d : Task %s is running\n", i+1, "anonymous")
		}
	}()

	time.Sleep(15 * time.Second)
}
