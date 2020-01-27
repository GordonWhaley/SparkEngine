package main

import "math"
//main structs

type Vector2 struct {
  x float64
  y float64
}

type Vector3 struct {
  x float64
  y float64
  z float64
}

type Vector4 struct {
  x float64
  y float64
  z float64
  w float64
}

type Matrix2 struct {
  mtx [4]float64
}

type Matrix3 struct {
  mtx [9]float64
}

type Matrix4 struct {
  mtx [12]float64
}

type Quaternion struct {
  w float64 //scalar value
  vec Vector3
}

//2D Vector Functions

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

func Vector2Magnitude(a Vector2) float64 {
  magnitude := math.Sqrt((a.x * a.x) + (a.y * a.y))
  return magnitude
}

func Vector2Normalize(a Vector2) Vector2 {
  newvec := Vector2 {
    a.x / Vector2Magnitude(a),
    a.y / Vector2Magnitude(a),
   }
   return newvec
}

//3D Vector Functions

func Vector3Create(x, y, z float64) Vector3 {
  newvec := Vector3 {x, y, z}
  return newvec
}

func Vector3Add(a, b Vector3) Vector3 {
  newvec := Vector3 {a.x + b.x, a.y + b.y, a.z + b.z}
  return newvec
}

func Vector3Subtract(a, b Vector3) Vector3 {
  newvec := Vector3 {a.x - b.x, a.y - b.y, a.z - b.z}
  return newvec
}

func Vector3Multiply(a Vector3, b float64) Vector3 {
  newvec := Vector3 {a.x * b, a.y * b, a.z * b}
  return newvec
}

func Vector3Divide(a Vector3, b float64) Vector3 {
  newvec := Vector3 {a.x / b, a.y / b, a.z / b}
  return newvec
}

func Vector3Dot(a, b Vector3) float64 {
  product := (a.x * b.x) + (a.y * b.y) + (a.z * b.z)
  return product
}

func Vector3Cross(a, b Vector3) Vector3 {
   newvec := Vector3 {
        (a.y * b.z) - (a.z * b.y),
        (a.z * b.x) - (a.x * b.x),
        (a.x * b.y) - (a.y * b.z),
     }
  return newvec
}

func Vector3Magnitude(a Vector3) float64 {
  magnitude := math.Sqrt((a.x * a.x) + (a.y * a.y) + (a.z * a.z))
  return magnitude
}

func Vector3Normalize(a Vector3) Vector3 {
  newvec := Vector3 {
    a.x / Vector3Magnitude(a),
    a.y / Vector3Magnitude(a),
    a.z / Vector3Magnitude(a),
   }
   return newvec
}

//Vector4 Functions (used by shaders)

func Vector4Create(x, y, z, w float64) Vector4 {
  newvec := Vector4 {x, y, z, w}
  return newvec
}

func Vector4Add(a, b Vector4) Vector4 {
  newvec := Vector4 {a.x + b.x, a.y + b.y, a.z + b.z, a.w + b.w}
  return newvec
}

func Vector4Subtract(a, b Vector4) Vector4 {
  newvec := Vector4 {a.x - b.x, a.y - b.y, a.z - b.z, a.w - b.w}
  return newvec
}

func Vector4Multiply(a Vector4, b float64) Vector4 {
  newvec := Vector4 {a.x * b, a.y * b, a.z * b, a.w * b}
  return newvec
}

func Vector4Divide(a Vector4, b float64) Vector4 {
  newvec := Vector4 {a.x / b, a.y / b, a.z / b, a.w / b}
  return newvec
}

func Vector4Dot(a, b Vector4) float64 {
  product := (a.x * b.x) + (a.y * b.y) + (a.z * b.z) + (a.w * b.w)
  return product
}

func Vector4Magnitude(a Vector4) float64 {
  magnitude := math.Sqrt((a.x * a.x) + (a.y * a.y) + (a.z * a.z) + (a.w * a.w))
  return magnitude
}

func Vector4Normalize(a Vector4) Vector4 {
  newvec := Vector4 {
    a.x / Vector4Magnitude(a),
    a.y / Vector4Magnitude(a),
    a.z / Vector4Magnitude(a),
    a.w / Vector4Magnitude(a),
   }
   return newvec
}

//2x2 Matrix functions

func Matrix2Create(m0, m2, m1, m3 float64) Matrix2 {
  var mtxParams [4]float64
  //written as column major to comply with OpenGL standards
  mtxParams[0] = m0
  mtxParams[2] = m2

  mtxParams[1] = m1
  mtxParams[3] = m3

  newmtx := Matrix2 {
    mtxParams,
  }
  return newmtx
}

func Matrix2Add(a, b Matrix2) Matrix2 {
  var mtxParams [4]float64
  //written as column major to comply with OpenGL standards
  mtxParams[0] = a.mtx[0] + b.mtx[0]
  mtxParams[2] = a.mtx[2] + b.mtx[2]

  mtxParams[1] = a.mtx[1] + b.mtx[1]
  mtxParams[3] = a.mtx[3] + b.mtx[3]

  newmtx := Matrix2 {
    mtxParams,
  }
  return newmtx
}

func Matrix2Subtract(a, b Matrix2) Matrix2 {
  var mtxParams [4]float64
  //written as column major to comply with OpenGL standards
  mtxParams[0] = a.mtx[0] - b.mtx[0]
  mtxParams[2] = a.mtx[2] - b.mtx[2]

  mtxParams[1] = a.mtx[1] - b.mtx[1]
  mtxParams[3] = a.mtx[3] - b.mtx[3]

  newmtx := Matrix2 {
    mtxParams,
  }
  return newmtx
}

func Matrix2Multiply(a Matrix2, b float64) Matrix2 {
  var mtxParams [4]float64
  //written as column major to comply with OpenGL standards
  mtxParams[0] = a.mtx[0] * b
  mtxParams[2] = a.mtx[2] * b

  mtxParams[1] = a.mtx[1] * b
  mtxParams[3] = a.mtx[3] * b

  newmtx := Matrix2 {
    mtxParams,
  }
  return newmtx
}

func Matrix2Divide(a Matrix2, b float64) Matrix2 {
  var mtxParams [4]float64
  //written as column major to comply with OpenGL standards
  mtxParams[0] = a.mtx[0] / b
  mtxParams[2] = a.mtx[2] / b

  mtxParams[1] = a.mtx[1] / b
  mtxParams[3] = a.mtx[3] / b

  newmtx := Matrix2 {
    mtxParams,
  }
  return newmtx
}

func Matrix2SetToIdentity(a Matrix2) Matrix2 {
  for i := 0; i < 4; i++ {
    a.mtx[i] = 0
  }
  a.mtx[0] = 1
  a.mtx[3] = 1
  return a
}

//3x3 Matrix Functions

func Matrix3Create(m0, m3, m6, m1, m4, m7, m2, m5, m8 float64) Matrix3 {
  var mtxParams [9]float64
  //written as column major to comply with OpenGL standards
  mtxParams[0] = m0
  mtxParams[3] = m3
  mtxParams[6] = m6

  mtxParams[1] = m1
  mtxParams[4] = m4
  mtxParams[7] = m7

  mtxParams[2] = m2
  mtxParams[5] = m5
  mtxParams[8] = m8
  newmtx := Matrix3 {
    mtxParams,
  }
  return newmtx
}

func Matrix3Add(a, b Matrix3) Matrix3 {
  var mtxParams [9]float64
  //written as column major to comply with OpenGL standards
  mtxParams[0] = a.mtx[0] + b.mtx[0]
  mtxParams[3] = a.mtx[3] + b.mtx[3]
  mtxParams[6] = a.mtx[6] + b.mtx[6]

  mtxParams[1] = a.mtx[1] + b.mtx[2]
  mtxParams[4] = a.mtx[4] + b.mtx[4]
  mtxParams[7] = a.mtx[7] + b.mtx[7]

  mtxParams[2] = a.mtx[2] + b.mtx[2]
  mtxParams[5] = a.mtx[5] + b.mtx[5]
  mtxParams[8] = a.mtx[8] + b.mtx[8]
  newmtx := Matrix3 {
    mtxParams,
  }
  return newmtx
}

func Matrix3Subtract(a, b Matrix3) Matrix3 {
  var mtxParams [9]float64
  //written as column major to comply with OpenGL standards
  mtxParams[0] = a.mtx[0] - b.mtx[0]
  mtxParams[3] = a.mtx[3] - b.mtx[3]
  mtxParams[6] = a.mtx[6] - b.mtx[6]

  mtxParams[1] = a.mtx[1] - b.mtx[2]
  mtxParams[4] = a.mtx[4] - b.mtx[4]
  mtxParams[7] = a.mtx[7] - b.mtx[7]

  mtxParams[2] = a.mtx[2] - b.mtx[2]
  mtxParams[5] = a.mtx[5] - b.mtx[5]
  mtxParams[8] = a.mtx[8] - b.mtx[8]
  newmtx := Matrix3 {
    mtxParams,
  }
  return newmtx
}

func Matrix3Multiply(a Matrix3, b float64) Matrix3 {
  var mtxParams [9]float64
  //written as column major to comply with OpenGL standards
  mtxParams[0] = a.mtx[0] * b
  mtxParams[3] = a.mtx[3] * b
  mtxParams[6] = a.mtx[6] * b

  mtxParams[1] = a.mtx[1] * b
  mtxParams[4] = a.mtx[4] * b
  mtxParams[7] = a.mtx[7] * b

  mtxParams[2] = a.mtx[2] * b
  mtxParams[5] = a.mtx[5] * b
  mtxParams[8] = a.mtx[8] * b
  newmtx := Matrix3 {
    mtxParams,
  }
  return newmtx
}

func Matrix3Divide(a Matrix3, b float64) Matrix3 {
  var mtxParams [9]float64
  //written as column major to comply with OpenGL standards
  mtxParams[0] = a.mtx[0] / b
  mtxParams[3] = a.mtx[3] / b
  mtxParams[6] = a.mtx[6] / b

  mtxParams[1] = a.mtx[1] / b
  mtxParams[4] = a.mtx[4] / b
  mtxParams[7] = a.mtx[7] / b

  mtxParams[2] = a.mtx[2] / b
  mtxParams[5] = a.mtx[5] / b
  mtxParams[8] = a.mtx[8] / b
  newmtx := Matrix3 {
    mtxParams,
  }
  return newmtx
}

func Matrix3Combine(a, b Matrix3) Matrix3 {
  var mtxParams [9]float64

  mtxParams[0] = (a.mtx[0] * b.mtx[0]) + (a.mtx[1] * b.mtx[3]) + (a.mtx[2] * b.mtx[6])
  mtxParams[3] = (a.mtx[3] * b.mtx[0]) + (a.mtx[4] * b.mtx[3]) + (a.mtx[5] * b.mtx[6])
  mtxParams[6] = (a.mtx[6] * b.mtx[0]) + (a.mtx[7] * b.mtx[3]) + (a.mtx[8] * b.mtx[6])

  mtxParams[1] = (a.mtx[0] * b.mtx[1]) + (a.mtx[1] * b.mtx[4]) + (a.mtx[2] * b.mtx[7])
  mtxParams[4] = (a.mtx[3] * b.mtx[1]) + (a.mtx[4] * b.mtx[4]) + (a.mtx[5] * b.mtx[7])
  mtxParams[7] = (a.mtx[6] * b.mtx[1]) + (a.mtx[7] * b.mtx[4]) + (a.mtx[8] * b.mtx[7])

  mtxParams[2] = (a.mtx[0] * b.mtx[2]) + (a.mtx[1] * b.mtx[5]) + (a.mtx[2] * b.mtx[8])
  mtxParams[5] = (a.mtx[3] * b.mtx[2]) + (a.mtx[4] * b.mtx[5]) + (a.mtx[5] * b.mtx[8])
  mtxParams[8] = (a.mtx[6] * b.mtx[2]) + (a.mtx[7] * b.mtx[5]) + (a.mtx[8] * b.mtx[8])
  newmtx := Matrix3 {
    mtxParams,
  }
  return newmtx
}

func Matrix3SetToIdentity(a Matrix3) Matrix3 {
  for i := 0; i < 9; i++ {
    a.mtx[i] = 0
  }
  a.mtx[0] = 1
  a.mtx[4] = 1
  a.mtx[8] = 1
  return a
}
