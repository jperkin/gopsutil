// +build solaris

package load

import (
	"context"
	"fmt"

	"github.com/siebenmann/go-kstat"
)

func Avg() (*AvgStat, error) {
	return AvgWithContext(context.Background())
}

func AvgWithContext(ctx context.Context) (*AvgStat, error) {
	tok, err := kstat.Open()
	if err != nil {
		return nil, fmt.Errorf("Open failure: %s", err)
	}

	ks, err := tok.Lookup("unix", -1, "system_misc")
	if err != nil {
		return nil, fmt.Errorf("lookup failure on unix:-1:system_misc: %s", err)
	}

	load1, err := ks.GetNamed("avenrun_1min")
	if err != nil {
		return nil, err
	}

	load5, err := ks.GetNamed("avenrun_5min")
	if err != nil {
		return nil, err
	}

	load15, err := ks.GetNamed("avenrun_15min")
	if err != nil {
		return nil, err
	}

	ret := &AvgStat{
		Load1: float64(load1.UintVal) / 256,
		Load5: float64(load5.UintVal) / 256,
		Load15: float64(load15.UintVal) / 256,
	}

	return ret, nil
}
