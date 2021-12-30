package kernels

type Kernel struct {
	Kernel [3][3]float32
}

var GaussianBlur = Kernel{

	Kernel: [3][3]float32{
		{1.0 / 16.0, 1.0 / 8.0, 1.0 / 16.0},
		{1.0 / 8.0, 1.0 / 4.0, 1.0 / 8.0},
		{1.0 / 16.0, 1.0 / 8.0, 1.0 / 16.0},
	},
}

var EdgeDetection = Kernel{

	Kernel: [3][3]float32{
		{-1, -1, -1},
		{-1, 8, -1},
		{-1, -1, -1},
	},
}

var BoxBlur = Kernel{

	Kernel: [3][3]float32{
		{1.0 / 9.0, 1.0 / 9.0, 1.0 / 9.0},
		{1.0 / 9.0, 1.0 / 9.0, 1.0 / 9.0},
		{1.0 / 9.0, 1.0 / 9.0, 1.0 / 9.0},
	},
}

var Sharpen = Kernel{

	Kernel: [3][3]float32{
		{0, -1, 0},
		{-1, 5, -1},
		{0, -1, 0},
	},
}
