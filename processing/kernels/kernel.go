package kernels

type Kernel [3][3]float32

var GaussianBlur = [3][3]float32{
	{1.0 / 16.0, 1.0 / 8.0, 1.0 / 16.0},
	{1.0 / 8.0, 1.0 / 4.0, 1.0 / 8.0},
	{1.0 / 16.0, 1.0 / 8.0, 1.0 / 16.0},
}

var EdgeDetection = [3][3]float32{
	{-1, -1, -1},
	{-1, 8, -1},
	{-1, -1, -1},
}

var BoxBlur = [3][3]float32{
	{1.0 / 9.0, 1.0 / 9.0, 1.0 / 9.0},
	{1.0 / 9.0, 1.0 / 9.0, 1.0 / 9.0},
	{1.0 / 9.0, 1.0 / 9.0, 1.0 / 9.0},
}

var Sharpen = [3][3]float32{
	{0, -1, 0},
	{-1, 5, -1},
	{0, -1, 0},
}

// CreKernelFromFloats constructs a custom 3x3
// kernel from a list of 9 floats
func CreateKernelFromFloats(first, second, third, fourth, fifth, sixth, seventh, eigth, ninth float32) Kernel {
	return [3][3]float32{
		{first, second, third},
		{fourth, fifth, sixth},
		{seventh, eigth, ninth},
	}
}
