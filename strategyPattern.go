// Implementing Strategy Pattern in Go

package main

type Duck struct {
	IQuackBehaviour
	IFlyBehaviour
	IEatBehaviour
}

type CloudDuck struct {
	NewQuack
	OldFly
	NewEat
}

// RoadDuck inheriting Duck is not needed as it is not using any of the behaviours
type RoadDuck struct {
	Duck
}

type MountainDuck struct {
	OldQuack
	OldFly
	OldEat
}

type IQuackBehaviour interface {
	Quack()
}

type NewQuack struct {
}

func (nq *NewQuack) Quack() {
	println("New Quack")
}

type OldQuack struct {
}

func (nq *OldQuack) Quack() {
	println("Old Quack")
}

type IFlyBehaviour interface {
	Fly()
}

type NewFly struct {
}

func (nf *NewFly) Fly() {
	println("New Fly")
}

type OldFly struct {
}

func (of *OldFly) Fly() {
	println("Old Fly")
}

type IEatBehaviour interface {
	Eat()
}

type NewEat struct {
}

func (ne *NewEat) Eat() {
	println("New Eat")
}

type OldEat struct {
}

func (oe *OldEat) Eat() {
	println("Old Eat")
}

func main() {
	// Cloud Duck should implement New Quack
	// Mountain Duck should implement Old Quack

	var cd CloudDuck
	cd.NewQuack.Quack()
	cd.OldFly.Fly()
	cd.NewEat.Eat()

	var md MountainDuck
	md.OldQuack.Quack()
	md.OldFly.Fly()
	md.OldEat.Eat()

}
