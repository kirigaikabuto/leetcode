package main
//
//import (
//	"bufio"
//	"fmt"
//	"io"
//	"os"
//	"strconv"
//	"strings"
//)
//
//package main
//
//import (
//"bufio"
//"fmt"
//"io"
//"os"
//"strconv"
//"strings"
//)
//
///*
// * Complete the 'climbingLeaderboard' function below.
// *
// * The function is expected to return an INTEGER_ARRAY.
// * The function accepts following parameters:
// *  1. INTEGER_ARRAY ranked
// *  2. INTEGER_ARRAY player
// */
//
//func climbingLeaderboard(ranked []int32, player []int32) []int32 {
//	// Write your code here
//	k := make(map[int32][]int32)
//	var previous int32
//	var rank int32 = 1
//	for _,v := range ranked{
//		if previous == v{
//			predArra := k[rank]
//			predArra = append(predArra, v)
//			k[rank] = predArra
//		} else {
//			k[rank] = []int32{v}
//			rank += 1
//		}
//		previous = v
//	}
//	n := rank-1
//	var i int32
//	result := []int32{}
//	for _, v := range player{
//		isFind := false
//		var res int32 = 0
//		for i=1; i<=n;i++ {
//			if v >= k[i][0]{
//				isFind = true
//				res = i
//				break
//			}
//		}
//		if isFind{
//			result = append(result, res)
//		} else{
//			result = append(result, n+1)
//		}
//	}
//	return result
//}
//
//func main() {
//	reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)
//
//	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
//	checkError(err)
//
//	defer stdout.Close()
//
//	writer := bufio.NewWriterSize(stdout, 16 * 1024 * 1024)
//
//	rankedCount, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
//	checkError(err)
//
//	rankedTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")
//
//	var ranked []int32
//
//	for i := 0; i < int(rankedCount); i++ {
//		rankedItemTemp, err := strconv.ParseInt(rankedTemp[i], 10, 64)
//		checkError(err)
//		rankedItem := int32(rankedItemTemp)
//		ranked = append(ranked, rankedItem)
//	}
//
//	playerCount, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
//	checkError(err)
//
//	playerTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")
//
//	var player []int32
//
//	for i := 0; i < int(playerCount); i++ {
//		playerItemTemp, err := strconv.ParseInt(playerTemp[i], 10, 64)
//		checkError(err)
//		playerItem := int32(playerItemTemp)
//		player = append(player, playerItem)
//	}
//
//	result := climbingLeaderboard(ranked, player)
//
//	for i, resultItem := range result {
//		fmt.Fprintf(writer, "%d", resultItem)
//
//		if i != len(result) - 1 {
//			fmt.Fprintf(writer, "\n")
//		}
//	}
//
//	fmt.Fprintf(writer, "\n")
//
//	writer.Flush()
//}
//
//func readLine(reader *bufio.Reader) string {
//	str, _, err := reader.ReadLine()
//	if err == io.EOF {
//		return ""
//	}
//
//	return strings.TrimRight(string(str), "\r\n")
//}
//
//func checkError(err error) {
//	if err != nil {
//		panic(err)
//	}
//}
//
//
//
//res := []int32{}
//for _, v := range player{
//isFind := false
//var prev int32
//var counter int32 = 0
//if v < ranked[0]{
//for _, value := range ranked{
//if prev != value{
//counter+=1
//}
//if v >= value && prev != value{
//isFind = true
//res = append(res, counter)
//break
//}
//prev = value
//}
//}
//if !isFind{
//res = append(res, counter+1)
//}
//}
//return res
