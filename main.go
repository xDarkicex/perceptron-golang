package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/appliedgo/perceptron/draw"
)

// Perceptron Holds our simple Perceptron,
// with weights and a bias inputs.
type Perceptron struct {
	weights []float32
	bias    float32
}

// Heaviside function is a mathmatic function too return a number from 0-1
func (p *Perceptron) heaviside(f float32) int32 {
	if f < 0 {
		return 0
	}
	return 1
}

// NewPerceptron Creates a perceptron with n inputs.
// Weights and bias are init with random numbers between -1 - 1
func NewPerceptron(n int32) *Perceptron {
	var i int32
	w := make([]float32, n, n)
	for i = 0; i < n; i++ {
		w[i] = ((rand.Float32() * 2) - 1)
	}
	return &Perceptron{
		weights: w,
		bias:    ((rand.Float32() * 2) - 1),
	}
}

func (p *Perceptron) Process(inputs []int32) int32 {
	sum := p.bias
	for i, input := range inputs {
		sum += float32(input) * p.weights[i]
	}
	return p.heaviside(sum)
}

func (p *Perceptron) Adjust(inputs []int32, delta int32, learningRate float32) {
	for i, input := range inputs {
		p.weights[i] += float32(input) * float32(delta) * learningRate
	}
	p.bias += float32(delta) * learningRate
}

var m, b int32

func f(x int32) int32 {
	return m*x + b
}

func isAboveLine(point []int32, f func(int32) int32) int32 {
	x := point[0]
	y := point[1]
	if y > f(x) {
		return 1
	}
	return 0
}

func train(p *Perceptron, iters int, rate float32) {
	for i := 0; i < iters; i++ {
		point := []int32{
			rand.Int31n(201) - 101,
			rand.Int31n(201) - 101,
		}
		actual := p.Process(point)
		expected := isAboveLine(point, f)
		delta := expected - actual
		p.Adjust(point, delta, rate)
	}
}

func verify(p *Perceptron) int32 {
	var correctAnswers int32 = 0
	c := draw.NewCanvas()
	for i := 0; i < 100; i++ {
		point := []int32{
			rand.Int31n(201) - 101,
			rand.Int31n(201) - 101,
		}
		result := p.Process(point)
		if result == isAboveLine(point, f) {
			correctAnswers++
		}
		c.DrawPoint(point[0], point[1], result == 1)
	}
	c.DrawLinearFunction(m, b)
	c.Save()
	return correctAnswers
}

func main() {
	rand.Seed(time.Now().UnixNano())
	m = rand.Int31n(11) - 5
	b = rand.Int31n(101) - 51
	p := NewPerceptron(2)
	iterations := 100000
	var learningRate float32 = 0.1 // can be anyway from 0 - 1
	train(p, iterations, learningRate)
	successRate := verify(p)
	fmt.Printf("%d%% of the answers were correct. \n", successRate)
}
