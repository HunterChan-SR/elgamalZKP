package main

import (
	"elgamalZKP/elgamal"
)

func main() {
	elgamal.Elgamal_test()
}

// 21:18:01 INF compiling circuit
// 21:18:01 INF parsed circuit inputs nbPublic=6 nbSecret=3
// 21:18:01 INF building constraint builder nbConstraints=6083
// 21:18:02 DBG constraint system solver done nbConstraints=6083 took=3.496625
// 21:18:02 DBG prover done backend=groth16 curve=bn254 nbConstraints=6083 took=43.540041
// 21:18:02 DBG verifier done backend=groth16 curve=bn254 took=1.095583
