package main

import (
	"elgamal/elgamal"

	"github.com/consensys/gnark-crypto/ecc"
	"github.com/consensys/gnark/backend/groth16"
	"github.com/consensys/gnark/frontend"
	"github.com/consensys/gnark/frontend/cs/r1cs"

	twistededwards "github.com/consensys/gnark/std/algebra/native/twistededwards"
)

func main() {
	var circuit elgamal.ElgamalCircuit
	ccs, _ := frontend.Compile(ecc.BN254.ScalarField(), r1cs.NewBuilder, &circuit)
	pk, vk, _ := groth16.Setup(ccs)

	PK := twistededwards.Point{X: "10850314672982369172672003967451632033159284953813776714598433980631397319933", Y: "8965398796740420932285406259591821562477691583691085976005966743474079100246"}
	MSG := twistededwards.Point{X: "17584003342443288212611502702720852971326889038729647608488525487636419282607", Y: "2125049271648756652653196499038894612649967510935307762984437955535699116622"}
	C1 := twistededwards.Point{X: "16840994858898498356996151304195349526959402024690888114512301638848133108243", Y: "1906440559241814703106077917483893328233037979587993449544693732315371263125"}
	C2 := twistededwards.Point{X: "15378195395970144197460528272196928189431405947459526788349558588114049237546", Y: "13825167311208359590171659020021831249552434905319228210383512419763619111810"}
	assignment := elgamal.ElgamalCircuit{PK: PK, MSG: MSG, R: 211248238708487, C1: C1, C2: C2}

	witness, _ := frontend.NewWitness(&assignment, ecc.BN254.ScalarField())
	publicWitness, _ := witness.Public()
	proof, _ := groth16.Prove(ccs, pk, witness)
	groth16.Verify(proof, vk, publicWitness)
}

/*
sk 1913791588
pk {X: 10850314672982369172672003967451632033159284953813776714598433980631397319933, Y: 8965398796740420932285406259591821562477691583691085976005966743474079100246}
msg 114514
MSG {X: 17584003342443288212611502702720852971326889038729647608488525487636419282607, Y: 2125049271648756652653196499038894612649967510935307762984437955535699116622}
r 211248238708487
rpk {X: 14007696396215044731334306980895292497843112922881865103344294727961969124948, Y: 19572170633557472093401501526937270044032112965223295387483807592888056488403}
c1 {X: 16840994858898498356996151304195349526959402024690888114512301638848133108243, Y: 1906440559241814703106077917483893328233037979587993449544693732315371263125}
c2 {X: 15378195395970144197460528272196928189431405947459526788349558588114049237546, Y: 13825167311208359590171659020021831249552434905319228210383512419763619111810}
*/

/*
21:46:23 INF compiling circuit
21:46:23 INF parsed circuit inputs nbPublic=6 nbSecret=3
21:46:23 INF building constraint builder nbConstraints=12679
21:46:25 DBG constraint system solver done nbConstraints=12679 took=11.107541
21:46:25 DBG prover done backend=groth16 curve=bn254 nbConstraints=12679 took=69.601583
21:46:25 DBG verifier done backend=groth16 curve=bn254 took=1.980917
*/
