package service

type IndicatorService interface {
	Calculate(prev, value float64, period int) float64
}

type indicatorServiceImpl struct{}

func NewIndicator() IndicatorService {
	return &indicatorServiceImpl{}
}

func (i *indicatorServiceImpl) Calculate(prev, value float64, period int) float64 {
	if prev == 0.0 {
		return value
	}

	multiplier := i.getMultiplier(period)

	return multiplier*value + (1-multiplier)*prev
}

func (i *indicatorServiceImpl) getMultiplier(period int) float64 {
	return float64(2) / float64(period+1)
}
