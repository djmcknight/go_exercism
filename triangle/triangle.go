// Package triangle should have a package comment that summarizes what it's about.
package triangle

import "math"

// Kind is the type for the kind of triangle
type Kind string

// Triangle Types
const (
	NaT = "NaT" // not a triangle
	Equ = "Equ" // equilateral
	Iso = "Iso" // isosceles
	Sca = "Sca" // scalene
)

// isTriangle determines if the floats you're passing are triangles
func isTriangle(a, b, c float64) bool {
	if a+b+c > 0 {
		if a <= (b+c) && b <= (a+c) && c <= (a+b) {
			return true
		}
	}
	return false
}

// isEquilateral checks if your inputs are an Equilateral Triangle
func isEquilateral(a, b, c float64) bool {
	if a == b && a == c {
		return true
	}
	return false
}

// isIsosceles checks if your inputs are an Isosceles Triangle
func isIsosceles(a, b, c float64) bool {
	if a == b || a == c || b == c {
		return true
	}
	return false
}

// isScalene checks if your inputs are an Scalene Triangle
func isScalene(a, b, c float64) bool {
	if a != b && a != c {
		return true
	}
	return false

}

// KindFromSides identifies what type of triangle you're using as a parameter.
func KindFromSides(a, b, c float64) Kind {
	if math.IsInf(a, 0) || math.IsInf(b, 0) || math.IsInf(c, 0) {
		return NaT
	}
	if !isTriangle(a, b, c) {
		return NaT
	}
	if isEquilateral(a, b, c) {
		return Equ
	}
	if isIsosceles(a, b, c) {
		return Iso
	}
	if isScalene(a, b, c) {
		return Sca
	}
	return NaT

}
