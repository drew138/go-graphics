package kernels

type Kernel struct {
	Kernel [][]float32
}

var GaussianBlur = Kernel{

	Kernel: [][]float32{
		{1.0 / 16.0, 1.0 / 8.0, 1.0 / 16.0},
		{1.0 / 8.0, 1.0 / 4.0, 1.0 / 8.0},
		{1.0 / 16.0, 1.0 / 8.0, 1.0 / 16.0},
	},
}

var EdgeDetection = Kernel{

	Kernel: [][]float32{
		{-1, -1, -1},
		{-1, 8, -1},
		{-1, -1, -1},
	},
}

var BoxBlur = Kernel{

	Kernel: [][]float32{
		{1.0 / 9.0, 1.0 / 9.0, 1.0 / 9.0},
		{1.0 / 9.0, 1.0 / 9.0, 1.0 / 9.0},
		{1.0 / 9.0, 1.0 / 9.0, 1.0 / 9.0},
	},
}

var Sharpen = Kernel{

	Kernel: [][]float32{
		{0, -1, 0},
		{-1, 5, -1},
		{0, -1, 0},
	},
}
