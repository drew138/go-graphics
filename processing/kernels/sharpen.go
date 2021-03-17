package kernels

var Sharpen = Kernel{
	Determinant: 1,
	Kernel: [][]float32{
		{0, -1, 0},
		{-1, 5, -1},
		{0, -1, 0},
	},
}
