package main

type Vector2 struct {
  x float64
  y float64
}

type Vector3 struct {
  x float64
  y float64
  z float64
}

func Vector2Create(x, y float64) Vector2 {
  newvec := Vector2 {x, y}
  return newvec
}

func Vector2Add(a, b Vector2) Vector2 {
  newvec := Vector2 {a.x + b.x, a.y + b.y}
  return newvec
}

func Vector2Subtract(a, b Vector2) Vector2 {
  newvec := Vector2 {a.x - b.x, a.y - b.y}
  return newvec
}

func Vector2Multiply(a Vector2, b float64) Vector2 {
  newvec := Vector2 {a.x * b, a.y * b}
  return newvec
}

func Vector2Divide(a Vector2, b float64) Vector2 {
  newvec := Vector2 {a.x / b, a.y / b}
  return newvec
}

func Vector2Dot(a, b Vector2) float64 {
  product := (a.x * b.x) + (a.y * b.y)
  return product
}
