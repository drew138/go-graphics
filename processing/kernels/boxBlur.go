package kernels

var BoxBlur = Kernel{
	Determinant: 1 / 16,
	Kernel: [][]float32{
		{1, 1, 1},
		{1, 1, 1},
		{1, 1, 1},
	},
}
