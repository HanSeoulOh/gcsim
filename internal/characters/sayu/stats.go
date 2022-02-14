package sayu

var (
	attack = [][]float64{
		{
			72.24 / 100,
			78.12 / 100,
			84 / 100,
			92.4 / 100,
			98.28 / 100,
			105 / 100,
			114.24 / 100,
			123.48 / 100,
			132.72 / 100,
			142.8 / 100,
			154.35 / 100,
			167.93 / 100,
			181.52 / 100,
			195.1 / 100,
			209.92 / 100,
		},
		{
			71.38 / 100,
			77.19 / 100,
			83.0 / 100,
			91.3 / 100,
			97.11 / 100,
			103.75 / 100,
			112.88 / 100,
			122.01 / 100,
			131.14 / 100,
			141.1 / 100,
			152.51 / 100,
			165.93 / 100,
			179.35 / 100,
			192.78 / 100,
			207.42 / 100,
		},
		{
			43.43 / 100,
			46.97 / 100,
			50.5 / 100,
			55.55 / 100,
			59.09 / 100,
			63.13 / 100,
			68.68 / 100,
			74.23 / 100,
			79.79 / 100,
			85.85 / 100,
			92.79 / 100,
			100.96 / 100,
			109.13 / 100,
			117.29 / 100,
			126.2 / 100,
		},
		{
			98.13 / 100,
			106.11 / 100,
			114.1 / 100,
			125.51 / 100,
			133.5 / 100,
			142.63 / 100,
			155.18 / 100,
			167.73 / 100,
			180.28 / 100,
			193.97 / 100,
			209.66 / 100,
			228.11 / 100,
			246.56 / 100,
			265.01 / 100,
			285.14 / 100,
		},
	}
	skillpress = []float64{
		36.0 / 100,
		38.7 / 100,
		41.4 / 100,
		45.0 / 100,
		47.7 / 100,
		50.4 / 100,
		54.0 / 100,
		57.6 / 100,
		61.2 / 100,
		64.8 / 100,
		68.4 / 100,
		72.0 / 100,
		76.5 / 100,
		81.0 / 100,
		85.5 / 100,
	}
	skillholdend = []float64{
		217.6 / 100,
		233.92 / 100,
		250.24 / 100,
		272 / 100,
		288.32 / 100,
		304.64 / 100,
		326.4 / 100,
		348.16 / 100,
		369.92 / 100,
		391.68 / 100,
		413.44 / 100,
		435.2 / 100,
		462.4 / 100,
		489.6 / 100,
		516.8 / 100,
	}
	skillpressend = []float64{
		158.4 / 100,
		170.28 / 100,
		182.16 / 100,
		198 / 100,
		209.88 / 100,
		221.76 / 100,
		237.6 / 100,
		253.44 / 100,
		269.28 / 100,
		285.12 / 100,
		300.96 / 100,
		316.8 / 100,
		336.6 / 100,
		356.4 / 100,
		376.2 / 100,
	}
	skillabsorb = []float64{
		16.8 / 100,
		18.06 / 100,
		19.32 / 100,
		21.0 / 100,
		22.26 / 100,
		23.52 / 100,
		25.2 / 100,
		26.88 / 100,
		28.56 / 100,
		30.24 / 100,
		31.92 / 100,
		33.6 / 100,
		35.7 / 100,
		37.8 / 100,
		39.9 / 100,
	}
	skillabsorbend = []float64{
		76.16 / 100,
		81.87 / 100,
		87.58 / 100,
		95.2 / 100,
		100.91 / 100,
		106.62 / 100,
		114.24 / 100,
		121.86 / 100,
		129.47 / 100,
		137.09 / 100,
		144.7 / 100,
		152.32 / 100,
		161.84 / 100,
		171.36 / 100,
		180.88 / 100,
	}

	burst = []float64{
		116.8 / 100,
		125.56 / 100,
		134.32 / 100,
		146.0 / 100,
		154.76 / 100,
		163.52 / 100,
		175.2 / 100,
		186.88 / 100,
		198.56 / 100,
		210.24 / 100,
		221.92 / 100,
		233.6 / 100,
		248.2 / 100,
		262.8 / 100,
		277.4 / 100,
	}
	burstskill = []float64{
		52.0 / 100,
		55.9 / 100,
		59.8 / 100,
		65.0 / 100,
		68.9 / 100,
		72.8 / 100,
		78.0 / 100,
		83.2 / 100,
		88.4 / 100,
		93.6 / 100,
		98.8 / 100,
		104.0 / 100,
		110.5 / 100,
		117.0 / 100,
		123.5 / 100,
	}
	inithealflat = []float64{
		577,
		635,
		698,
		765,
		837,
		914,
		996,
		1083,
		1174,
		1270,
		1371,
		1477,
		1588,
		1703,
		1824,
	}
	inithealpp = []float64{
		92.16 / 100,
		99.07 / 100,
		105.98 / 100,
		115.2 / 100,
		122.11 / 100,
		129.02 / 100,
		138.24 / 100,
		147.46 / 100,
		156.67 / 100,
		165.89 / 100,
		175.1 / 100,
		184.32 / 100,
		195.84 / 100,
		207.36 / 100,
		218.88 / 100,
	}
	bursthealflat = []float64{
		500,
		550,
		605,
		663,
		726,
		792,
		863,
		938,
		1017,
		1101,
		1188,
		1280,
		1376,
		1476,
		1580,
	}
	bursthealpp = []float64{
		79.87 / 100,
		85.86 / 100,
		91.85 / 100,
		99.84 / 100,
		105.83 / 100,
		111.82 / 100,
		119.81 / 100,
		127.8 / 100,
		135.78 / 100,
		143.77 / 100,
		151.76 / 100,
		159.74 / 100,
		169.73 / 100,
		179.71 / 100,
		189.7 / 100,
	}
)
