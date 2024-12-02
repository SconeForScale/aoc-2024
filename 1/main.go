package main

import (
	"bufio"
	"errors"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func mergeSortedLists(list1 []int, list2 []int) []int {
    leftIdx := 0
    rightIdx := 0
    sorted_list := make([]int, 0)
    for {
        if len(sorted_list) == (len(list1) + len(list2)) {
            break 
        }
        rightElt := 0
        leftElt := 0
        if leftIdx < len(list1) {
            leftElt = list1[leftIdx]
        } else {
            leftElt = math.MaxInt32
        }
        if rightIdx < len(list2) {
            rightElt = list2[rightIdx]
        } else {
            rightElt = math.MaxInt32
        }
        if leftElt < rightElt {
            sorted_list = append(sorted_list, leftElt)
            leftIdx += 1
        } else {
            sorted_list = append(sorted_list, rightElt)
            rightIdx += 1 
        }
    }
    return sorted_list
}

func computeDistance(list1 []int, list2 []int) (int, error) {
    if len(list1) != len(list2) {
        return -1, errors.New("lists must be equal length")
    }
    distance := 0
    for idx, val := range list1 {
        if val > list2[idx] {
            distance += val - list2[idx]
        } else {
            distance += list2[idx] - val 
        }
    }
    return distance, nil
}

func mergeSort(list []int) []int {
    if len(list) <= 1 {
        return list
    }
    midpoint := len(list)/2 
    l1 := list[0:midpoint]
    l2 := list[midpoint:]

   
    return mergeSortedLists(mergeSort(l1), mergeSort(l2))
}

func computeSimilarity(list1 []int, list2[]int) (int, error) {
    simMap := map[int]int{}
    countsMap := map[int]int{}
    for _, val := range list1 {
        simMap[val] = 0
    }
    for _, val := range list2 {
        _, valExists := countsMap[val]
        if valExists {
            countsMap[val] += 1 
        } else {
            countsMap[val] = 1 
        }
    }
    similarity := 0
    for _, val := range list1 {
        count, valExists := countsMap[val]
        if valExists {
            similarity += count*val
        }
    }

    return similarity, nil

}

func main() {
    input, inputErr := os.Open("input")
    if inputErr != nil{
        panic(inputErr)
    }
    defer input.Close()
    input_scanner := bufio.NewScanner(input)

    inputA := make([]int, 0)
    inputB := make([]int, 0)

    for input_scanner.Scan() {
       lineVals := strings.Fields(input_scanner.Text())

        lv_a, err_a := strconv.Atoi(lineVals[0])
        lv_b, err_b := strconv.Atoi(lineVals[1])

        if err_a != nil {
            panic(err_a)
        }
        if err_b != nil {
            panic(err_b)
        }

        inputA = append(inputA, lv_a)
        inputB = append(inputB, lv_b)
    }
    distance, distErr := computeDistance(mergeSort(inputA), mergeSort(inputB))
    if distErr != nil {
        panic(distErr)
    }
    fmt.Println(distance)
    similarity, simErr :=  computeSimilarity(inputA, inputB)
    if simErr != nil {
        panic(simErr)
    }
    fmt.Println(similarity)

}
