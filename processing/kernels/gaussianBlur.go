package kernels

var GaussianBlur = Kernel{
	Determinant: 1 / 16,
	Kernel: [][]float32{
		{1, 2, 1},
		{2, 4, 2},
		{1, 2, 1},
	},
}
