// Copyright 2018 The Energi Core Authors
// Copyright 2018 The go-ethereum Authors
// This file is part of the Energi Core library.
//
// The Energi Core library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The Energi Core library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the Energi Core library. If not, see <http://www.gnu.org/licenses/>.

package metrics

import "testing"

func BenchmarkHistogram(b *testing.B) {
	h := NewHistogram(NewUniformSample(100))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		h.Update(int64(i))
	}
}

func TestGetOrRegisterHistogram(t *testing.T) {
	r := NewRegistry()
	s := NewUniformSample(100)
	NewRegisteredHistogram("foo", r, s).Update(47)
	if h := GetOrRegisterHistogram("foo", r, s); 1 != h.Count() {
		t.Fatal(h)
	}
}

func TestHistogram10000(t *testing.T) {
	h := NewHistogram(NewUniformSample(100000))
	for i := 1; i <= 10000; i++ {
		h.Update(int64(i))
	}
	testHistogram10000(t, h)
}

func TestHistogramEmpty(t *testing.T) {
	h := NewHistogram(NewUniformSample(100))
	if count := h.Count(); 0 != count {
		t.Errorf("h.Count(): 0 != %v\n", count)
	}
	if min := h.Min(); 0 != min {
		t.Errorf("h.Min(): 0 != %v\n", min)
	}
	if max := h.Max(); 0 != max {
		t.Errorf("h.Max(): 0 != %v\n", max)
	}
	if mean := h.Mean(); 0.0 != mean {
		t.Errorf("h.Mean(): 0.0 != %v\n", mean)
	}
	if stdDev := h.StdDev(); 0.0 != stdDev {
		t.Errorf("h.StdDev(): 0.0 != %v\n", stdDev)
	}
	ps := h.Percentiles([]float64{0.5, 0.75, 0.99})
	if 0.0 != ps[0] {
		t.Errorf("median: 0.0 != %v\n", ps[0])
	}
	if 0.0 != ps[1] {
		t.Errorf("75th percentile: 0.0 != %v\n", ps[1])
	}
	if 0.0 != ps[2] {
		t.Errorf("99th percentile: 0.0 != %v\n", ps[2])
	}
}

func TestHistogramSnapshot(t *testing.T) {
	h := NewHistogram(NewUniformSample(100000))
	for i := 1; i <= 10000; i++ {
		h.Update(int64(i))
	}
	snapshot := h.Snapshot()
	h.Update(0)
	testHistogram10000(t, snapshot)
}

func testHistogram10000(t *testing.T, h Histogram) {
	if count := h.Count(); 10000 != count {
		t.Errorf("h.Count(): 10000 != %v\n", count)
	}
	if min := h.Min(); 1 != min {
		t.Errorf("h.Min(): 1 != %v\n", min)
	}
	if max := h.Max(); 10000 != max {
		t.Errorf("h.Max(): 10000 != %v\n", max)
	}
	if mean := h.Mean(); 5000.5 != mean {
		t.Errorf("h.Mean(): 5000.5 != %v\n", mean)
	}
	if stdDev := h.StdDev(); 2886.751331514372 != stdDev {
		t.Errorf("h.StdDev(): 2886.751331514372 != %v\n", stdDev)
	}
	ps := h.Percentiles([]float64{0.5, 0.75, 0.99})
	if 5000.5 != ps[0] {
		t.Errorf("median: 5000.5 != %v\n", ps[0])
	}
	if 7500.75 != ps[1] {
		t.Errorf("75th percentile: 7500.75 != %v\n", ps[1])
	}
	if 9900.99 != ps[2] {
		t.Errorf("99th percentile: 9900.99 != %v\n", ps[2])
	}
}
