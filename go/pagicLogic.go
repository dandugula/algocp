package main

import (
	"github.com/google/uuid"
	"math/rand/v2"
	//"fmt"
)

type StatsRecord struct {
	txFrames   int64
	rxFrames   int64
	minLatency float64
	maxLatency float64
	avgLatency float64
}

type ItemId = int64
type ViewId = string

type ItemStatsMap map[ItemId]*StatsRecord
type ViewIdItemIdMap map[ViewId][]ItemId
type ViewIdResultsMap map[ViewId][]*StatsRecord

func generateViewIds() []ViewId {
	var ret []ViewId
	for range 10 {
		ret = append(ret, uuid.New().String())
	}
	return ret
}

func setUpViewIdItemIdMap(viewIds []ViewId) (ViewIdItemIdMap,
	map[ItemId]struct{}) {
	ret := make(ViewIdItemIdMap)
	itemIds := make(map[ItemId]struct{})
	for _, i := range viewIds {
		for range 1000000 {
			n := rand.Int64()
			ret[i] = append(ret[i], n)
			itemIds[n] = struct{}{}
		}
	}
	return ret, itemIds
}

func setUpItemStatsMap(items map[ItemId]struct{}) ItemStatsMap {
	ret := make(ItemStatsMap)
	for i := range items {
		ret[i] = &StatsRecord{
			txFrames:   rand.Int64(),
			rxFrames:   rand.Int64(),
			minLatency: rand.Float64(),
			maxLatency: rand.Float64(),
			avgLatency: rand.Float64(),
		}
	}
	return ret
}

func copyStat(s *StatsRecord) *StatsRecord {
	return &StatsRecord{
		txFrames:   s.txFrames,
		rxFrames:   s.rxFrames,
		minLatency: s.minLatency,
		maxLatency: s.maxLatency,
		avgLatency: s.avgLatency,
	}
}

func TransformAndProduce(itemStats ItemStatsMap,
	viewIdItemIdMap ViewIdItemIdMap) ViewIdResultsMap {

	// Transform
	liveItemIds := make(map[ItemId]struct{})
	//totalItems, uniqueItems := 0, 0
	for _, v := range viewIdItemIdMap {
		pageSize := 50
		pageNum := 2
		startIdx := (pageNum - 1) * pageSize
		endIdx := startIdx + pageSize

		//totalItems += len(v)
		for i := startIdx; i < len(v) && i < endIdx; i++ {
			liveItemIds[v[i]] = struct{}{}
		}
	}
	//uniqueItems = len(liveItemIds)
	//fmt.Println("Total items: ", totalItems)
	//fmt.Println("Unique items: ", uniqueItems)

	resMap := make(ViewIdResultsMap)
	// Produce
	//fmt.Println("Result Items: ", len(itemStats))
	for currViewId, currItems := range viewIdItemIdMap {
		for i := 0; i < len(currItems); i++ {
			if v, ok := itemStats[currItems[i]]; ok {
				resMap[currViewId] = append(resMap[currViewId], copyStat(v))
			}
		}
	}
	return resMap
}
