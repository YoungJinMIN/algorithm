package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"strconv"
)

func main() {
	// 첫번째 인풋값 받을 변수
	var firstInput string
	// 두번째 인풋값 받을 변수
	var secondInput string
	// 최대 원소 개수
	var maxElement int
	// 뽑을 원소 위치 리스트
	var elementIdxList []int
	// 이미 뽑은 원소 위치 리스트
	var deleteElementIdxList []int
	// 이동 최솟값
	var countMinMove int
	// 현재 위치 변수
	nowIdx := 1
	// 오른쪽인지, 왼쪽인지 구분 변수
	var right bool
	// 오른쪽, 왼쪽 이동값 변수
	var moveRight int
	var moveLeft int

	scn := bufio.NewScanner(os.Stdin)
	// 최대 원소 개수, 뽑을 원소 개수 입력 받기
	scn.Scan()
	firstInput = scn.Text()
	// 뽑을 원소 위치들 입력 받기
	scn.Scan()
	secondInput = scn.Text()

	// 첫번째 Input 형변환해서 각 변수에 저장
	firstInfo := strings.Split(firstInput, " ")
	maxElement, _ = strconv.Atoi(firstInfo[0])

	// 두번째 Input 형변환해서 각 변수에 저장
	secondInfo := strings.Split(secondInput, " ")
	for _, value := range secondInfo {
		valueInt, _ := strconv.Atoi(value)
		elementIdxList = append(elementIdxList, valueInt)
	}

	// 위치값 슬라이스로 루프돌기
	for _, numIdx := range elementIdxList {
		// 만약 현재 위치와 numIdx가 다를때만
		if numIdx != nowIdx {
			// 오른쪽, 왼쪽 이동 값 구하기
			// 여기서 먼저 현재위치와 이동하려는 위치 값 비교해서 이동 위치가 클경우 
				// 이동 위치 - 현위치 = moveRight(right 값 true)
				// 총 원소 개수 - moveRight = moveLeft
			// 아닌경우 
				// 현위치 - 이동위치 = moveLeft(right 값 false)
				// 총 원소 개수 - moveLeft = moveRight
			if numIdx > nowIdx {
				moveRight = numIdx - nowIdx
				moveLeft = maxElement - moveRight
				right = true
			} else {
				moveLeft = nowIdx - numIdx
				moveRight = maxElement - moveLeft
				right = false
			}
			
			// 뺀 원소값 loop 돌면서 right, left 값 감소 시키기
			for _, deleteIdx := range deleteElementIdxList {
				// 만약 right 값이 true인경우 
					// 뺀 원소가 현재위치보다는 크고, 이동위치보다 작으면 moveRight 값을 감소
					// 아니면 moveLeft 값을 감소
				// 아니면 right 값이 false 경우
					// 뺀 원소가 이동위치보다 크고, 현재위치보다 작으면 moveLeft 값을 감소
					// 아니면 moveRight 값을 감소

				if right == true {
					if deleteIdx > nowIdx && deleteIdx < numIdx {
						moveRight = moveRight - 1
					} else {
						moveLeft = moveLeft - 1 
					}
				} else {
					if deleteIdx > numIdx && deleteIdx < nowIdx {
						moveLeft = moveLeft - 1
					} else {
						moveRight = moveRight - 1
					}
				}
			}

			// 만약 오른쪽 값이, 왼쪽 값보다 작거나 같으면 오른쪽 값을 count 값으로 증가
			// 아니면 왼쪽 값이 작으면 왼쪽 값으로 count 값 증가
			if moveRight <= moveLeft {
				countMinMove = countMinMove + moveRight
			} else {
				countMinMove = countMinMove + moveLeft
			}
		} 
		// 뺀 원소 값 저장
		deleteElementIdxList = append(deleteElementIdxList, numIdx)
		// fmt.Println("deleteElementIdxList", deleteElementIdxList)
		// 현재 위치값 변경
		// 현재 위치값 + 1하면서 확인, 언제까지? max까지
			// 확인했떠니 있어? 그러면 그값으로 변경
		// max까지 확인했는데 없어? 그러면 현재값 - (현재값 - 1)하면서 확인, 언제까지? 현재값 전까지
			// 있으면 그값으로변경, 밑에 애들도 확인했는데 없어? 그럼 끝난거지 말이안되고?

		// 변경 위치값 찾았는지 여부
		findIdx := false
		// 현재 위치부터 맥스값까지 확인하면서 
		for i:=numIdx+1;i<=maxElement;i++ {
			changeIdx := i
			for _, delNum := range deleteElementIdxList {
				if changeIdx == delNum {
					findIdx = false
					break
				} else {

					findIdx = true
					continue
				}
			}
			if findIdx == true {
				nowIdx = changeIdx
				break
			}
		}
		// 큰값에서 못찾음 => 낮은 위치에서 찾아야함
		if findIdx == false {
			// 처음위치부터 
			for i:=1;i<numIdx;i++ {
				changeIdx := i

				for _, delNum := range deleteElementIdxList {
					if changeIdx == delNum {
						findIdx = false
						break
					} else {
	
						findIdx = true
						continue
					}
				}
				if findIdx == true {
					nowIdx = changeIdx
					break
				}
			}
		}
	}
	fmt.Println(countMinMove)
}