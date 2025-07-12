
package main

func maxFreeTime(eventTime int, startEventTime []int, endEventTime []int) int {
    numberOfEvents := len(startEventTime)

    forwardMaxFreeTime := make([]int, numberOfEvents + 1)
    forwardMaxFreeTime[0] = startEventTime[0]
    for i := 1; i < numberOfEvents; i++ {
        forwardMaxFreeTime[i] = max(forwardMaxFreeTime[i - 1], startEventTime[i] - endEventTime[i - 1])
    }

    backwardMaxFreeTime := make([]int, numberOfEvents + 1)
    backwardMaxFreeTime[len(backwardMaxFreeTime) - 1] = eventTime - endEventTime[numberOfEvents - 1]
    for i := numberOfEvents - 1; i > 0; i-- {
        backwardMaxFreeTime[i] = max(backwardMaxFreeTime[i + 1], startEventTime[i] - endEventTime[i - 1])
    }

    var maxContinuousFreeTime = 0
    for i := range numberOfEvents {

        startAdjacentFreeTimeBeforeEvent := 0
        if (i - 1) >= 0 {
            startAdjacentFreeTimeBeforeEvent = endEventTime[i - 1]
        }

        endAdjacentFreeTimeAfterEvent := eventTime
        if (i + 1) < numberOfEvents {
            endAdjacentFreeTimeAfterEvent = startEventTime[i + 1]
        }

        maxPreceedingNonadjacentFreeTime := 0
        if i - 1 >= 0 {
            maxPreceedingNonadjacentFreeTime = forwardMaxFreeTime[i - 1]
        }

        maxFollowingNonadjacentFreeTime := 0
        if i + 2 <= len(backwardMaxFreeTime) - 1 {
            maxFollowingNonadjacentFreeTime = backwardMaxFreeTime[i + 2]
        }

        maxNonadjacentFreeTime := max(maxPreceedingNonadjacentFreeTime, maxFollowingNonadjacentFreeTime)
        if maxNonadjacentFreeTime < endEventTime[i] - startEventTime[i] {
            maxContinuousFreeTime = max(maxContinuousFreeTime,
                                        endAdjacentFreeTimeAfterEvent -
                                        startAdjacentFreeTimeBeforeEvent -
                                        (endEventTime[i] - startEventTime[i]))

            continue
        }

        maxContinuousFreeTime = max(maxContinuousFreeTime,
                                    endAdjacentFreeTimeAfterEvent -
                                    startAdjacentFreeTimeBeforeEvent)
    }

    return maxContinuousFreeTime
}
