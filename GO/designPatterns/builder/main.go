package main

import (
	"fmt"
)

type Vehicle struct {
	Type string
	Color string
	Wheels string
}

type Builder struct {
	buildProcess BuildProcess
}

func (b *Builder) SetBuildProcess(bp BuildProcess) {
	b.buildProcess = bp
}

func (b *Builder) Construct() Vehicle {
	b.buildProcess.SetType().SetColor().SetWheels()
	return b.buildProcess.GetVehicle()
}

func (b *Builder) PrintVehicle() {
	vehicle := b.buildProcess.GetVehicle()
	fmt.Printf("Type: %s \n", vehicle.Type)
	fmt.Printf("Color: %s \n", vehicle.Color)
	fmt.Printf("Wheels: %s \n", vehicle.Wheels)
	fmt.Printf("===============\n")
}

type BuildProcess interface {
	SetType() BuildProcess
	SetColor() BuildProcess
	SetWheels() BuildProcess
	GetVehicle() Vehicle
}


// Car implements BuildProcess
type Car struct {
	vehicle Vehicle
}

func (c *Car) SetType() BuildProcess {
	c.vehicle.Type = "family"
	return c
}

func (c *Car) SetColor() BuildProcess {
	c.vehicle.Color = "blue"
	return c
}

func (c *Car) SetWheels() BuildProcess {
	c.vehicle.Wheels = "steel"
	return c
}

func (c *Car) GetVehicle() Vehicle {
	return c.vehicle
}

// Truck implements BuildProcess
type Truck struct {
	vehicle Vehicle
}

func (c *Truck) SetType() BuildProcess {
	c.vehicle.Type = "18wheelers"
	return c
}

func (c *Truck) SetColor() BuildProcess {
	c.vehicle.Color = "black"
	return c
}

func (c *Truck) SetWheels() BuildProcess {
	c.vehicle.Wheels = "aluminum"
	return c
}

func (c *Truck) GetVehicle() Vehicle {
	return c.vehicle
}


/*

As we can see, the client only need to do:
Tell the builder of what type of the product
Tell the builder to construct, the step by step will be done in the Builder, the client doesn’t care about the detail process
Tell the builder to print the report when the construction process is done
We can reuse the step by step to create the object and hide the complexity to create it. If there is another new electronic product should be produced, we can just define the new struct and implement the BuildProcess interface, we also achieved the flexibility and scalability in our code.

*/

func main() {
	builder := Builder{}

	car := &Car{}
	builder.SetBuildProcess(car)
	builder.Construct()
	builder.PrintVehicle()

	truck := &Truck{}
	builder.SetBuildProcess(truck)
	builder.Construct()
	builder.PrintVehicle()
}