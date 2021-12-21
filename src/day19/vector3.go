package day19

import "math"

type (
	Vector3 struct {
		x, y, z int
	}
)

func (v Vector3) Rotate(id int) Vector3 {
	switch id {
	case 0:
		return Vector3{v.x, v.y, v.z}
	case 1:
		return Vector3{v.x, -v.z, v.y}
	case 2:
		return Vector3{v.x, -v.y, -v.z}
	case 3:
		return Vector3{v.x, v.z, -v.y}
	case 4:
		return Vector3{-v.x, -v.y, v.z}
	case 5:
		return Vector3{-v.x, -v.z, -v.y}
	case 6:
		return Vector3{-v.x, v.y, -v.z}
	case 7:
		return Vector3{-v.x, v.z, v.y}
	case 8:
		return Vector3{v.y, v.x, -v.z}
	case 9:
		return Vector3{v.y, -v.x, v.z}
	case 10:
		return Vector3{v.y, v.z, v.x}
	case 11:
		return Vector3{v.y, -v.z, -v.x}
	case 12:
		return Vector3{-v.y, v.x, v.z}
	case 13:
		return Vector3{-v.y, -v.x, -v.z}
	case 14:
		return Vector3{-v.y, -v.z, v.x}
	case 15:
		return Vector3{-v.y, v.z, -v.x}
	case 16:
		return Vector3{v.z, v.x, v.y}
	case 17:
		return Vector3{v.z, -v.x, -v.y}
	case 18:
		return Vector3{v.z, -v.y, v.x}
	case 19:
		return Vector3{v.z, v.y, -v.x}
	case 20:
		return Vector3{-v.z, v.x, -v.y}
	case 21:
		return Vector3{-v.z, -v.x, v.y}
	case 22:
		return Vector3{-v.z, v.y, v.x}
	case 23:
		return Vector3{-v.z, -v.y, -v.x}
	default:
		panic(id)

	}
}

func (v Vector3) Inverse() Vector3 {
	return Vector3{-v.x, -v.y, -v.z}
}

func (v Vector3) Add(w Vector3) Vector3 {
	return Vector3{
		x: v.x + w.x,
		y: v.y + w.y,
		z: v.z + w.z,
	}
}

func (v Vector3) Sub(w Vector3) Vector3 {
	return Vector3{
		x: v.x - w.x,
		y: v.y - w.y,
		z: v.z - w.z,
	}
}

func (v Vector3) Dist(w Vector3) int {
	return int(math.Abs(float64(v.x-w.x)) +
		math.Abs(float64(v.y-w.y)) +
		math.Abs(float64(v.z-w.z)))
}
