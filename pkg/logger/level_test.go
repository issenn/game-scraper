package logger_test


import (
	"fmt"
	// "time"
	"testing"
	// "bytes"
	// "strconv"

	// "go.uber.org/zap"
	// _ "go.uber.org/zap/zapcore"

	// zapAdapter "github.com/issenn/game-scraper/pkg/logger/adapter/zap"

	. "github.com/issenn/game-scraper/pkg/logger"
)


func TestLevel_setlevel(t *testing.T) {
	fmt.Printf("111 %v\n", NewAtomicLevelAt(ErrorLevel).Enabled(FatalLevel))
	// logger, _ := zap.NewProduction()
	// defer logger.Sync() // flushes buffer, if any
	// logger.Info("failed to fetch URL",
	// 	// Structured context as strongly typed Field values.
	// 	zap.String("url", url),
	// 	zap.Int("attempt", 3),
	// 	zap.Duration("backoff", time.Second),
	// )
}

func TestLevel_ComparableFunc(t *testing.T) {
	// var unmarshaller ComparableFunc
	// data := "eq"
	// if err := unmarshaller.UnmarshalText([]byte(data.(string))); err != nil {
	// 	return nil, err
	// }
}


func TestOther(t *testing.T) {
	arr := []int{-1, -1, -1, -1, -1, 0}
	result := pivotIndex(arr)
	fmt.Println(result)
}

func pivotIndex(nums []int) int {
	leftSum := 0
	// rightSum := addArray(nums...)
	rightSum := reduce(nums, func (x int, y int) int {
		return x + y
	})
	for i, v := range nums {
		rightSum -= v
		if leftSum == rightSum {
			return i
		}
		leftSum += v
	}
	return -1
}

func reduce[T, M any](iterable []T, fn func(M, T) M, initValue ...M) M {
	var value M
	if len(initValue) > 0 {
		value = initValue[0]
	}
	for _, element := range iterable {
		value = fn(value, element)
	}
	return value
}

func addArray(nums ...int) int {
	result := 0
	for _, v := range nums {
		result += v
	}
	return result
}
