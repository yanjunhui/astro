package star

import (
	"fmt"
	"github.com/yanjunhui/astro/tools"
	"testing"
	"time"
)

func TestStar(t *testing.T) {
	err := InitStarDatabase()
	if err != nil {
		t.Fatal(err)
	}
	sirius, err := StarDataByHR(2491)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%+v\n", sirius)
	now := time.Now()
	ra, dec := sirius.RaDecByDate(now)
	fmt.Println(tools.Format(ra/15, 1), tools.Format(dec, 0))
	fmt.Println(RiseTime(now, ra, dec, 115, 40, 0, true))
	fmt.Println(CulminationTime(now, ra, 115))
	fmt.Println(DownTime(now, ra, dec, 115, 40, 0, true))
}
