package main

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func TestSearchDataFromES(t *testing.T) {
	InitLoaclCache()
	start := time.Now().Add(-time.Minute).UnixMilli()
	end := time.Now().UnixMilli()
	resp, err := SearchDataFromES(context.Background(), start, end)
	if err != nil {
		panic(err)
	}
	fmt.Println(len(resp))
}
