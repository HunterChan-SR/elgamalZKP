package elgamal

import (
	tedwards "github.com/consensys/gnark-crypto/ecc/twistededwards"
	"github.com/consensys/gnark/frontend"
	twistededwards "github.com/consensys/gnark/std/algebra/native/twistededwards"
)

var g twistededwards.Point = twistededwards.Point{X: "9671717474070082183213120605117400219616337014328744928644933853176787189663", Y: "16950150798460657717958625567821834550301663161624707787222815936182638968203"}

func Encrypt(curve twistededwards.Curve, pk, msg twistededwards.Point, r frontend.Variable) (c1, c2 twistededwards.Point) {
	rPK := curve.ScalarMul(pk, r)
	c1 = curve.Add(msg, rPK)
	c2 = curve.ScalarMul(g, r)
	return c1, c2
}

type ElgamalCircuit struct {
	PK  twistededwards.Point `gnark:",public"`
	MSG twistededwards.Point
	R   frontend.Variable
	C1  twistededwards.Point `gnark:",public"`
	C2  twistededwards.Point `gnark:",public"`
}

func (circuit *ElgamalCircuit) Define(api frontend.API) error {
	curve, _ := twistededwards.NewEdCurve(api, tedwards.BN254)
	c1, c2 := Encrypt(curve, circuit.PK, circuit.MSG, circuit.R)

	cc1 := curve.ScalarMul(circuit.C1, 1)
	cc2 := curve.ScalarMul(circuit.C2, 1)

	curve.API()
	api.AssertIsEqual(c1.X, cc1.X)
	api.AssertIsEqual(c1.Y, cc1.Y)
	api.AssertIsEqual(c2.X, cc2.X)
	api.AssertIsEqual(c2.Y, cc2.Y)
	return nil
}
