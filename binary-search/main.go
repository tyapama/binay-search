package main

import (
	"fmt"
	"log"
	"sort"
)

func main() {
	array := []string{"あおい", "あい", "あおぎり", "あかね", "あし", "あし", "あずさ", "あんず", "いばら", "いばら", "いも", "うめ", "えごま", "おぎ", "おだまき", "かえで", "かえで", "かし", "かし", "かしわ", "かしわ", "かじ", "かつら", "かば", "かば", "かや", "からむし", "がま", "ききょう", "ききょう", "きく", "きり", "くすのき", "くすのき", "くず", "くぬぎ", "くり", "くわ", "けし", "こけ", "こなら", "ごんずい", "さかき", "さくら", "さつき", "さねかずら", "さんかよう", "しい", "しば", "しょうぶ", "すぎ", "すぎ", "すげ", "すすき", "せり", "せんだん", "たけ", "たちばな", "たまごたけ", "ちゃ", "つが", "つげ", "つげ", "つた", "つばき", "でいご", "でいご", "とち", "とち", "とちゅう", "なし", "なのはな", "なら", "はぎ", "はしばみ", "はぜ", "ばしょう", "ばしょう", "ばんれいし", "ひいらぎ", "ひし", "ひのき", "ひる", "びわ", "びわ", "ふき", "ふじ", "ふよう", "ふよう", "ぶな", "ほお", "まさき", "まつ", "まゆみ", "みの", "むく", "も", "もみじ", "やえ", "やくしそう", "やなぎ", "やなぎ", "やぶらん", "やますげ", "よもき", "よもぎ", "らん", "わた"}
	sort.Strings(array)

	targetName := "かえで"

	times, err := findIndexByBinarySearch(array, targetName)
	if err != nil {
		log.Fatal("err occurred")
	}
	fmt.Println("binary search loop times is ", times)
}

func findIndexByBinarySearch(array []string, targetName string) (int, error) {

	startArraryLen := len(array)
	middleLen := startArraryLen / 2
	loopLimit := startArraryLen
	tmp := array
	loopTimes := 0

	for loopTimes < loopLimit {
		if tmp[middleLen] == targetName {
			return loopTimes, nil
		}

		if tmp[middleLen] < targetName {
			tmp = tmp[middleLen+1:]
		} else {
			tmp = tmp[:middleLen]
		}
		middleLen = len(tmp) / 2
		loopTimes = loopTimes + 1
	}

	// ここに来ることはない
	return 0, nil
}
