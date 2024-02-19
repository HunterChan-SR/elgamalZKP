package elgamal

import (
	tedwards "github.com/consensys/gnark-crypto/ecc/twistededwards"
	"github.com/consensys/gnark/frontend"
	twistededwards "github.com/consensys/gnark/std/algebra/native/twistededwards"
)

type ElgamalCircuit struct {
	PK  twistededwards.Point `gnark:",public"`
	MSG twistededwards.Point
	R   frontend.Variable
	C1  twistededwards.Point `gnark:",public"`
	C2  twistededwards.Point `gnark:",public"`
}

func (circuit *ElgamalCircuit) Define(api frontend.API) error {
	curve, _ := twistededwards.NewEdCurve(api, tedwards.BN254)

	params, _ := twistededwards.GetCurveParams(tedwards.BN254)
	g := twistededwards.Point{X: params.Base[0], Y: params.Base[1]}
	c1 := curve.Add(circuit.MSG, curve.ScalarMul(circuit.PK, circuit.R))
	c2 := curve.ScalarMul(g, circuit.R)
	api.AssertIsEqual(c1.X, circuit.C1.X)
	api.AssertIsEqual(c2.X, circuit.C2.X)
	return nil
}
