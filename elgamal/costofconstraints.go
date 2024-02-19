package elgamal

import (
	"crypto/rand"
	"fmt"

	"github.com/consensys/gnark-crypto/ecc"
	bn254 "github.com/consensys/gnark-crypto/ecc/bn254/twistededwards"
	tedwards "github.com/consensys/gnark-crypto/ecc/twistededwards"
	"github.com/consensys/gnark/backend/groth16"
	"github.com/consensys/gnark/frontend"
	"github.com/consensys/gnark/frontend/cs/r1cs"
	twistededwards "github.com/consensys/gnark/std/algebra/native/twistededwards"
)

func Elgamal_test() {
	var assignment ElgamalCircuit
	params, _ := twistededwards.GetCurveParams(tedwards.BN254)
	var G bn254.PointAffine
	G.X.SetBigInt(params.Base[0])
	G.Y.SetBigInt(params.Base[1])

	SK, _ := rand.Int(rand.Reader, params.Order)
	var PK bn254.PointAffine
	PK.ScalarMultiplication(&G, SK)
	assignment.PK = twistededwards.Point{X: PK.X, Y: PK.Y}

	msg, _ := rand.Int(rand.Reader, params.Order)
	var MSG bn254.PointAffine
	PK.ScalarMultiplication(&G, msg)
	assignment.MSG = twistededwards.Point{X: MSG.X, Y: MSG.Y}

	R, _ := rand.Int(rand.Reader, params.Order)
	assignment.R = R

	var C1, C2 bn254.PointAffine
	C1.ScalarMultiplication(&PK, R)
	C1.Add(&MSG, &C1)
	C2.ScalarMultiplication(&G, R)
	assignment.C1 = twistededwards.Point{X: C1.X, Y: C1.Y}
	assignment.C2 = twistededwards.Point{X: C2.X, Y: C2.Y}

	//
	var circuit ElgamalCircuit
	ccs, _ := frontend.Compile(ecc.BN254.ScalarField(), r1cs.NewBuilder, &circuit)
	pk, vk, _ := groth16.Setup(ccs)

	witness, _ := frontend.NewWitness(&assignment, ecc.BN254.ScalarField())
	publicWitness, _ := witness.Public()
	proof, _ := groth16.Prove(ccs, pk, witness)
	err := groth16.Verify(proof, vk, publicWitness)
	if err != nil {
		fmt.Println("invalid proof")
	}
}
