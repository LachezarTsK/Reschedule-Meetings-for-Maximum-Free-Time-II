
#include <vector>
#include <algorithm>
using namespace std;

class Solution {

public:
    int maxFreeTime(int eventTime, const vector<int>& startEventTime, const vector<int>& endEventTime) const {
        int numberOfEvents = startEventTime.size();

        vector<int> forwardMaxFreeTime(numberOfEvents + 1);
        forwardMaxFreeTime[0] = startEventTime[0];
        for (int i = 1; i < numberOfEvents; ++i) {
                forwardMaxFreeTime[i] = max(forwardMaxFreeTime[i - 1], startEventTime[i] - endEventTime[i - 1]);
        }

        vector<int> backwardMaxFreeTime(numberOfEvents + 1);
        backwardMaxFreeTime[backwardMaxFreeTime.size() - 1] = eventTime - endEventTime[numberOfEvents - 1];
        for (int i = numberOfEvents - 1; i > 0; --i) {
                backwardMaxFreeTime[i] = max(backwardMaxFreeTime[i + 1], startEventTime[i] - endEventTime[i - 1]);
        }

        int maxContinuousFreeTime = 0;
        for (int i = 0; i < numberOfEvents; ++i) {
            int startAdjacentFreeTimeBeforeEvent = (i - 1) >= 0 ? endEventTime[i - 1] : 0;
            int endAdjacentFreeTimeAfterEvent = (i + 1) < numberOfEvents ? startEventTime[i + 1] : eventTime;

            int maxPreceedingNonadjacentFreeTime = (i - 1 >= 0) ? forwardMaxFreeTime[i - 1] : 0;
            int maxFollowingNonadjacentFreeTime = (i + 2 <= backwardMaxFreeTime.size() - 1) ? backwardMaxFreeTime[i + 2] : 0;
            int maxNonadjacentFreeTime = max(maxPreceedingNonadjacentFreeTime, maxFollowingNonadjacentFreeTime);
           
            if (maxNonadjacentFreeTime < endEventTime[i] - startEventTime[i]) {
                maxContinuousFreeTime = max(maxContinuousFreeTime,
                                            endAdjacentFreeTimeAfterEvent
                                          - startAdjacentFreeTimeBeforeEvent
                                          - (endEventTime[i] - startEventTime[i]));
                continue;
            }

            maxContinuousFreeTime = max(maxContinuousFreeTime,
                                        endAdjacentFreeTimeAfterEvent
                                      - startAdjacentFreeTimeBeforeEvent);
        }

            return maxContinuousFreeTime;
    }
};
