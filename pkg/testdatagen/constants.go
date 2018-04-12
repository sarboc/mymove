package testdatagen

import (
	"time"
)

// TestYear is the default year for testing.
var TestYear = 2019

// DefaultSrcGBLOC is the default GBLOC for testing.
var DefaultSrcGBLOC = "OHAI"

// DefaultMarket is the default market for testing.
var DefaultMarket = "dHHG"

// DefaultSrcRateArea is a default rate area (California) for testing.
var DefaultSrcRateArea = "US87"

// DefaultDstRegion is a default region (US South West) for testing.
var DefaultDstRegion = "6"

// PeakRateCycleStart is the first instant that the peak rate cycle starts
var PeakRateCycleStart = time.Date(TestYear, time.May, 15, 0, 0, 0, 0, time.UTC)

// PeakRateCycleEnd is the first instant that the peak rate cycle ends
var PeakRateCycleEnd = time.Date(TestYear, time.October, 1, 0, 0, 0, 0, time.UTC)

// DateInsidePeakRateCycle is available as a convenient test date inside the Peak Rate Cycle
var DateInsidePeakRateCycle = time.Date(TestYear, time.May, 16, 0, 0, 0, 0, time.UTC)

// DateOutsidePeakRateCycle is available as a convenient test date outside the Peak Rate Cycle
var DateOutsidePeakRateCycle = time.Date(TestYear, time.October, 10, 0, 0, 0, 0, time.UTC)

// NonPeakRateCycleStart is the first instant that the peak rate cycle starts
var NonPeakRateCycleStart = time.Date(TestYear, time.October, 1, 0, 0, 0, 0, time.UTC)

// NonPeakRateCycleEnd is the first instant that the peak rate cycle ends
var NonPeakRateCycleEnd = time.Date(TestYear+1, time.May, 15, 0, 0, 0, 0, time.UTC)

// DateInsideNonPeakRateCycle is available as a convenient test date inside the NonPeak Rate Cycle
var DateInsideNonPeakRateCycle = time.Date(TestYear, time.October, 2, 0, 0, 0, 0, time.UTC)

// DateOutsideNonPeakRateCycle is available as a convenient test date outside the NonPeak Rate Cycle
var DateOutsideNonPeakRateCycle = time.Date(TestYear+1, time.May, 16, 0, 0, 0, 0, time.UTC)

// PerformancePeriodStart is the first day of the first performance period
var PerformancePeriodStart = time.Date(TestYear, time.May, 15, 0, 0, 0, 0, time.UTC)

// PerformancePeriodEnd is the last day of the first performance period
var PerformancePeriodEnd = time.Date(TestYear, time.July, 31, 0, 0, 0, 0, time.UTC)

// DateInsidePerformancePeriod is within the performance period defined by
// PerformancePeriodStart and PerformancePeriodEnd.
var DateInsidePerformancePeriod = time.Date(TestYear, time.May, 16, 0, 0, 0, 0, time.UTC)

// DateOutsidePerformancePeriod is after the performance period defined by
// PerformancePeriodStart and PerformancePeriodEnd.
var DateOutsidePerformancePeriod = time.Date(TestYear, time.August, 1, 0, 0, 0, 0, time.UTC)

// RateEngineDate is a date for the rate engine to use on generation for tests.
var RateEngineDate = time.Date(TestYear, time.May, 18, 0, 0, 0, 0, time.UTC)
