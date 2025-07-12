
import kotlin.math.max

class Solution {

    fun maxFreeTime(eventTime: Int, startEventTime: IntArray, endEventTime: IntArray): Int {
        val numberOfEvents = startEventTime.size

        val forwardMaxFreeTime = IntArray(numberOfEvents + 1)
        forwardMaxFreeTime[0] = startEventTime[0]
        for (i in 1..<numberOfEvents) {
            forwardMaxFreeTime[i] = max(forwardMaxFreeTime[i - 1], startEventTime[i] - endEventTime[i - 1])
        }

        val backwardMaxFreeTime = IntArray(numberOfEvents + 1)
        backwardMaxFreeTime[backwardMaxFreeTime.size - 1] = eventTime - endEventTime[numberOfEvents - 1]
        for (i in numberOfEvents - 1 downTo 1) {
            backwardMaxFreeTime[i] = max(backwardMaxFreeTime[i + 1], startEventTime[i] - endEventTime[i - 1])
        }

        var maxContinuousFreeTime = 0
        for (i in 0..<numberOfEvents) {
            val startAdjacentFreeTimeBeforeEvent = if ((i - 1) >= 0) endEventTime[i - 1] else 0
            val endAdjacentFreeTimeAfterEvent = if ((i + 1) < numberOfEvents) startEventTime[i + 1] else eventTime

            val maxPreceedingNonadjacentFreeTime = if (i - 1 >= 0) forwardMaxFreeTime[i - 1] else 0
            val maxFollowingNonadjacentFreeTime = if (i + 2 <= backwardMaxFreeTime.size - 1) backwardMaxFreeTime[i + 2] else 0
            val maxNonadjacentFreeTime = max(maxPreceedingNonadjacentFreeTime, maxFollowingNonadjacentFreeTime)
          
            if (maxNonadjacentFreeTime < endEventTime[i] - startEventTime[i]) {
                maxContinuousFreeTime = max(maxContinuousFreeTime,
                                            endAdjacentFreeTimeAfterEvent
                                          - startAdjacentFreeTimeBeforeEvent
                                          - (endEventTime[i] - startEventTime[i]))
                continue
            }

            maxContinuousFreeTime = max(maxContinuousFreeTime,
                                        endAdjacentFreeTimeAfterEvent
                                      - startAdjacentFreeTimeBeforeEvent
            )
        }

        return maxContinuousFreeTime
    }
}
