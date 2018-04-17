package utils

type RateLimitInfo struct{
    LimitID         int         `json:"limitID,omitempty"`
    IntervalType    int         `json:"intervalType"`
    ViewType        int         `json:"viewType"`
    LimitCount      int         `json:"limitCount"`
    PriorityLevel   int         `json:"priorityLevel,omitempty"`
    StrategyID      int         `json:"strategyID,omitempty"`
    StartTime       string      `json:"startTime,omitempty"`
    EndTime         string      `json:"endTime,omitempty"`
}