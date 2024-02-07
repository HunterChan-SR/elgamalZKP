package main

import (
	"github.com/consensys/gnark-crypto/ecc"
	tedwards "github.com/consensys/gnark-crypto/ecc/twistededwards"
	"github.com/consensys/gnark/backend/groth16"
	"github.com/consensys/gnark/frontend"
	"github.com/consensys/gnark/frontend/cs/r1cs"
	"github.com/consensys/gnark/std/algebra/native/twistededwards"
)

type MyCircuit struct {
	P twistededwards.Point `gnark:",public"`
	M frontend.Variable    `gnark:",public"`
}

func (circuit *MyCircuit) Define(api frontend.API) error {
	curve, _ := twistededwards.NewEdCurve(api, tedwards.BN254)
	curve.ScalarMul(circuit.P, circuit.M)

	return nil
}

func test1() {
	var circuit MyCircuit
	ccs, _ := frontend.Compile(ecc.BN254.ScalarField(), r1cs.NewBuilder, &circuit)
	pk, vk, _ := groth16.Setup(ccs)
	p := twistededwards.Point{X: 2, Y: 5}
	assignment := MyCircuit{P: p, M: 7}
	witness, _ := frontend.NewWitness(&assignment, ecc.BN254.ScalarField())
	publicWitness, _ := witness.Public()
	proof, _ := groth16.Prove(ccs, pk, witness)
	groth16.Verify(proof, vk, publicWitness)
}
