package runningvariance

import "testing"

func TestMeanSD(t *testing.T) {
	s := NewRunningStat()
	for i := 1; i <= 5; i++ {
		s.Push(float64(i))
	}

	if actual, expected := s.Mean(), 3.0; expected != actual {
		t.Errorf("Expected mean %f, got %f", expected, actual)
	}

	if actual, expected := s.StandardDeviation(), 1.58113883008418976139353162579937; expected != actual {
		t.Errorf("Expected standard deviation %f, got %f", expected, actual)
	}
}

func TestMean(t *testing.T) {
	s := NewRunningStat()
	s.Push(1)
	s.Push(1)
	s.Push(1)
	s.Push(0)
	s.Push(0)
	s.Push(0)

	if actual, expected := s.Mean(), 0.5; expected != actual {
		t.Errorf("Expected %f, got %f", expected, actual)
	}
}

func TestStandardDeviation(t *testing.T) {
	s := NewRunningStat()

	if actual, expected := s.StandardDeviation(), 0.0; expected != actual {
		t.Errorf("Expected %f, got %f", expected, actual)
	}

	s.Push(1)
	s.Push(1)
	s.Push(1)

	if actual, expected := s.StandardDeviation(), 0.0; expected != actual {
		t.Errorf("Expected %f, got %f", expected, actual)
	}
}

func TestNumDataValues(t *testing.T) {
	s := NewRunningStat()
	s.Push(1)

	if actual, expected := s.NumDataValues(), uint(1); expected != actual {
		t.Errorf("Expected %d, got %d", expected, actual)
	}
}
