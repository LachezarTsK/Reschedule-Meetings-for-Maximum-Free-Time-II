
using System;

public class Solution
{
    public int MaxFreeTime(int eventTime, int[] startEventTime, int[] endEventTime)
    {
        int numberOfEvents = startEventTime.Length;

        int[] forwardMaxFreeTime = new int[numberOfEvents + 1];
        forwardMaxFreeTime[0] = startEventTime[0];
        for (int i = 1; i < numberOfEvents; ++i)
        {
            forwardMaxFreeTime[i] = Math.Max(forwardMaxFreeTime[i - 1], startEventTime[i] - endEventTime[i - 1]);
        }

        int[] backwardMaxFreeTime = new int[numberOfEvents + 1];
        backwardMaxFreeTime[backwardMaxFreeTime.Length - 1] = eventTime - endEventTime[numberOfEvents - 1];
        for (int i = numberOfEvents - 1; i > 0; --i)
        {
            backwardMaxFreeTime[i] = Math.Max(backwardMaxFreeTime[i + 1], startEventTime[i] - endEventTime[i - 1]);
        }

        int maxContinuousFreeTime = 0;
        for (int i = 0; i < numberOfEvents; ++i)
        {
            int startAdjacentFreeTimeBeforeEvent = (i - 1) >= 0 ? endEventTime[i - 1] : 0;
            int endAdjacentFreeTimeAfterEvent = (i + 1) < numberOfEvents ? startEventTime[i + 1] : eventTime;

            int maxPreceedingNonadjacentFreeTime = (i - 1 >= 0) ? forwardMaxFreeTime[i - 1] : 0;
            int maxFollowingNonadjacentFreeTime = (i + 2 <= backwardMaxFreeTime.Length - 1) ? backwardMaxFreeTime[i + 2] : 0;
            int maxNonadjacentFreeTime = Math.Max(maxPreceedingNonadjacentFreeTime, maxFollowingNonadjacentFreeTime);
           
            if (maxNonadjacentFreeTime < endEventTime[i] - startEventTime[i])
            {
                maxContinuousFreeTime = Math.Max(maxContinuousFreeTime,
                                                 endAdjacentFreeTimeAfterEvent
                                               - startAdjacentFreeTimeBeforeEvent
                                               - (endEventTime[i] - startEventTime[i]));
                continue;
            }

            maxContinuousFreeTime = Math.Max(maxContinuousFreeTime,
                                             endAdjacentFreeTimeAfterEvent
                                           - startAdjacentFreeTimeBeforeEvent);
        }

        return maxContinuousFreeTime;
    }
}
