package main

import (
	"log"
	"os"

	"github.com/MicahParks/go-ad"
)

func main() {
	// Create a logger.
	logger := log.New(os.Stdout, "", 0)

	// Iterate through the rest of the periods' data and calculate the A/D line's point for the given period.
	var a *ad.AD
	var result float64
	for i := range open {
		input := ad.Input{
			Close:  closePrices[i],
			Low:    low[i],
			High:   high[i],
			Volume: volume[i],
		}

		if a == nil {
			a, result = ad.New(input)
		} else {
			result = a.Calculate(input)
		}
		logger.Printf("Index: %d AD: %.4f", i, result)
	}
}

var (
	open = []float64{
		48.43,
		48.2,
		48.11,
		48.52,
		48.66,
		49.01,
		49.12,
		48.53,
		48.43,
		48.43,
		48.08,
		48.06,
		48.62,
		48.88,
		48.96,
		48.62,
		47.51,
		47.43,
		47.07,
		47.62,
		46.68,
		46.73,
		45.97,
		46.22,
		46.58,
		47.23,
		47.34,
		48.02,
		48.58,
		48.14,
		48.53,
		49.34,
		49.3,
		49.73,
		50.02,
		50,
		50.73,
		50.62,
		50.38,
		50.27,
		49.39,
		49.28,
		48.9,
		48.63,
		48.53,
		48.22,
		48.41,
		48.1,
		48.37,
		47.52,
		47.47,
		47.47,
		47.55,
		46.42,
		46.48,
		46.79,
		46.52,
		47.6,
		47.18,
		47.9,
		48.4,
		48.31,
		48.62,
		48.41,
		48.67,
		48.45,
		48.28,
		48.56,
		48.71,
		47.83,
		47.16,
		47.82,
		48.26,
		48.59,
		48.94,
		49.14,
		48.85,
		49.17,
		49.89,
		49.27,
		49.98,
		50.67,
		50.56,
		50.69,
		50.67,
		51.01,
		51.08,
		51.36,
		50.74,
		50.88,
		50.57,
		50.22,
		50.7,
		51.02,
		51.33,
		50.97,
		51.4,
		51.18,
		51.16,
		50.98,
		50.69,
		50.76,
		50.59,
		50.76,
		50.5,
		50.07,
		50.36,
		50.48,
		50.87,
		50.77,
		50.75,
		51.05,
		51.9,
		51.59,
		51.38,
		51.98,
		52.9,
		52.82,
		53,
		53.45,
		52.72,
		52.59,
		52.62,
		52.58,
		52.9,
	}
	high = []float64{
		48.44,
		48.29,
		48.24,
		48.65,
		48.82,
		49.19,
		49.2,
		48.88,
		48.44,
		48.48,
		48.12,
		48.47,
		48.69,
		49.16,
		48.99,
		48.64,
		47.65,
		47.59,
		47.48,
		47.62,
		47.23,
		46.87,
		46.59,
		46.34,
		47.36,
		47.65,
		47.5,
		48.74,
		48.6,
		48.58,
		48.96,
		49.63,
		49.7,
		49.96,
		50.2,
		50.59,
		51.02,
		50.72,
		50.55,
		50.89,
		49.87,
		49.44,
		49.17,
		48.95,
		48.59,
		48.45,
		48.65,
		48.32,
		48.38,
		47.81,
		47.87,
		47.72,
		47.78,
		46.68,
		46.93,
		46.85,
		47.14,
		47.66,
		47.63,
		48.48,
		48.43,
		48.49,
		48.78,
		48.45,
		48.86,
		48.48,
		48.41,
		48.83,
		48.94,
		48.01,
		47.8,
		47.97,
		48.39,
		48.97,
		49.13,
		49.21,
		48.96,
		49.5,
		50.02,
		49.35,
		50.73,
		50.77,
		50.57,
		50.92,
		50.9,
		51.23,
		51.25,
		51.38,
		50.75,
		50.94,
		50.62,
		50.72,
		50.99,
		51.43,
		51.45,
		51.2,
		51.56,
		51.42,
		51.21,
		51.04,
		50.9,
		50.78,
		50.63,
		50.86,
		50.56,
		50.12,
		50.48,
		50.88,
		50.94,
		50.99,
		50.98,
		51.59,
		52.05,
		51.83,
		51.65,
		52.72,
		52.98,
		53.09,
		53.14,
		53.48,
		52.85,
		52.84,
		52.71,
		52.95,
		53.1,
	}
	low = []float64{
		47.89,
		47.85,
		47.9,
		48.32,
		48.53,
		48.77,
		48.99,
		48.1,
		48.12,
		48.07,
		47.8,
		48.01,
		48.35,
		48.68,
		48.43,
		47.73,
		47.06,
		47.32,
		46.96,
		47.13,
		46.62,
		46.01,
		45.93,
		45.83,
		46.5,
		47.21,
		47.24,
		47.94,
		48.22,
		48.08,
		48.47,
		49.26,
		49.2,
		49.61,
		49.76,
		49.96,
		50.43,
		50.27,
		50.13,
		50.27,
		48.89,
		49.03,
		48.65,
		48.62,
		48.15,
		48.13,
		48.11,
		48.03,
		47.52,
		47.52,
		47.33,
		47.38,
		47.16,
		46.29,
		46.36,
		46.35,
		46.52,
		46.97,
		47.08,
		47.83,
		47.97,
		48.27,
		48.42,
		48.18,
		48.4,
		48.1,
		47.88,
		48.24,
		48.24,
		47.45,
		47.14,
		47.42,
		48.1,
		48.51,
		48.84,
		48.96,
		48.72,
		48.96,
		49.67,
		48.9,
		49.94,
		50.39,
		50.25,
		50.52,
		50.58,
		50.92,
		50.78,
		50.77,
		50.39,
		50.63,
		50.41,
		50.21,
		50.34,
		50.94,
		51.23,
		50.82,
		51.27,
		51.16,
		50.97,
		50.67,
		50.65,
		50.59,
		50.34,
		50.49,
		50.34,
		49.87,
		50.14,
		50.45,
		50.7,
		50.67,
		50.71,
		51.04,
		51.55,
		51.5,
		51.29,
		51.94,
		52.58,
		52.75,
		52.91,
		52.78,
		52.41,
		52.17,
		52.4,
		52.47,
		52.81,
	}
	closePrices = []float64{
		47.98,
		47.87,
		48,
		48.55,
		48.73,
		48.85,
		49.2,
		48.15,
		48.34,
		48.1,
		47.91,
		48.2,
		48.53,
		49.04,
		48.63,
		47.75,
		47.25,
		47.48,
		47.27,
		47.41,
		46.89,
		46.11,
		46.27,
		46.04,
		47.32,
		47.61,
		47.5,
		48.61,
		48.38,
		48.45,
		48.81,
		49.41,
		49.33,
		49.86,
		49.81,
		50.27,
		50.67,
		50.42,
		50.41,
		50.61,
		49,
		49.34,
		48.82,
		48.79,
		48.43,
		48.44,
		48.22,
		48.21,
		47.59,
		47.69,
		47.66,
		47.56,
		47.32,
		46.52,
		46.51,
		46.7,
		47.01,
		46.97,
		47.57,
		48.37,
		48.1,
		48.44,
		48.44,
		48.32,
		48.45,
		48.11,
		47.96,
		48.78,
		48.28,
		47.45,
		47.54,
		47.77,
		48.39,
		48.88,
		49.03,
		49.04,
		48.81,
		49.39,
		49.72,
		49.04,
		50.29,
		50.62,
		50.4,
		50.87,
		50.81,
		51.12,
		51.02,
		50.78,
		50.56,
		50.7,
		50.51,
		50.43,
		50.9,
		51.3,
		51.36,
		50.91,
		51.47,
		51.21,
		51.1,
		50.7,
		50.67,
		50.65,
		50.45,
		50.58,
		50.52,
		50.07,
		50.44,
		50.81,
		50.84,
		50.95,
		50.86,
		51.34,
		51.68,
		51.74,
		51.64,
		52.67,
		52.67,
		52.78,
		53.07,
		52.84,
		52.62,
		52.75,
		52.6,
		52.75,
		53,
	}
	volume = []float64{
		4737500,
		5531200,
		5735900,
		7513500,
		6801800,
		7556500,
		5310100,
		8277700,
		6072200,
		6437200,
		4549800,
		5872700,
		5103100,
		8903400,
		5949900,
		8983500,
		4921600,
		2641700,
		4831600,
		3760700,
		4796700,
		5033900,
		5315100,
		3780900,
		5952100,
		2833100,
		2432700,
		4794000,
		4615200,
		3546000,
		2597700,
		10061800,
		5479300,
		8097900,
		6282700,
		8709300,
		8166400,
		4941500,
		5585800,
		5542000,
		9857500,
		6511700,
		5747400,
		4632100,
		5904400,
		2540200,
		2690700,
		2249000,
		3666000,
		3900700,
		5008000,
		5596800,
		4149400,
		17857900,
		15397700,
		7705700,
		4947800,
		5063000,
		3782300,
		7463100,
		4276800,
		3353300,
		6207900,
		2797400,
		3776300,
		2371700,
		3730500,
		3605500,
		3194000,
		6383300,
		4853900,
		5041700,
		3101200,
		3574700,
		2639600,
		2638600,
		3413600,
		4960100,
		4727200,
		4998100,
		7620800,
		7198000,
		3212600,
		3317900,
		3186900,
		5644400,
		6850400,
		6879300,
		3553000,
		4856400,
		4211800,
		2435300,
		7527300,
		5618400,
		3837800,
		2959800,
		3933300,
		2246600,
		3507300,
		3507000,
		2628300,
		2476900,
		2837200,
		4517800,
		2972600,
		4315300,
		1825200,
		3224600,
		2139300,
		2995500,
		1724500,
		4212300,
		5407800,
		3497400,
		2100000,
		4270500,
		3203600,
		2660300,
		2234100,
		5946000,
		4973700,
		2525100,
		3292500,
		2175500,
		3435700,
	}
)
