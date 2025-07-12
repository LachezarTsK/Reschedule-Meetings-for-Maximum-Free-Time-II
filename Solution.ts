
function maxFreeTime(eventTime: number, startEventTime: number[], endEventTime: number[]): number {
    const numberOfEvents = startEventTime.length;

    const forwardMaxFreeTime: number[] = new Array(numberOfEvents + 1);
    forwardMaxFreeTime[0] = startEventTime[0];
    for (let i = 1; i < numberOfEvents; ++i) {
        forwardMaxFreeTime[i] = Math.max(forwardMaxFreeTime[i - 1], startEventTime[i] - endEventTime[i - 1]);
    }

    const backwardMaxFreeTime: number[] = new Array(numberOfEvents + 1);
    backwardMaxFreeTime[backwardMaxFreeTime.length - 1] = eventTime - endEventTime[numberOfEvents - 1];
    for (let i = numberOfEvents - 1; i > 0; --i) {
        backwardMaxFreeTime[i] = Math.max(backwardMaxFreeTime[i + 1], startEventTime[i] - endEventTime[i - 1]);
    }

    let maxContinuousFreeTime = 0;
    for (let i = 0; i < numberOfEvents; ++i) {
        const startAdjacentFreeTimeBeforeEvent = (i - 1) >= 0 ? endEventTime[i - 1] : 0;
        const endAdjacentFreeTimeAfterEvent = (i + 1) < numberOfEvents ? startEventTime[i + 1] : eventTime;

        const maxPreceedingNonadjacentFreeTime = (i - 1 >= 0) ? forwardMaxFreeTime[i - 1] : 0;
        const maxFollowingNonadjacentFreeTime = (i + 2 <= backwardMaxFreeTime.length - 1) ? backwardMaxFreeTime[i + 2] : 0;
        const maxNonadjacentFreeTime = Math.max(maxPreceedingNonadjacentFreeTime, maxFollowingNonadjacentFreeTime);

        if (maxNonadjacentFreeTime < endEventTime[i] - startEventTime[i]) {
            maxContinuousFreeTime = Math.max(maxContinuousFreeTime,
                                             endAdjacentFreeTimeAfterEvent
                                           - startAdjacentFreeTimeBeforeEvent
                                           - (endEventTime[i] - startEventTime[i]));
            continue;
        }

        maxContinuousFreeTime = Math.max(maxContinuousFreeTime,
                                         endAdjacentFreeTimeAfterEvent
                                       - startAdjacentFreeTimeBeforeEvent);
    }

    return maxContinuousFreeTime;
};
