from collections import deque

def sliding_window_maximum(nums, k):
    dq = deque()  # stores indices
    result = []
    
    for i in range(len(nums)):
    # Remove indices outside window
        while dq and dq[0] <= i - k:
            dq.popleft()
        
    # Remove smaller elements
        while dq and nums[dq[-1]] <= nums[i]:
            dq.pop()
        
        dq.append(i)
        
        if i >= k - 1:
            result.append(nums[dq[0]])
    
    return result

def main():
    nums = [4, 0, -1, 3, 5, 3, 6, 8]
    res = sliding_window_maximum(nums, 3)
    print(res)


if __name__ == "__main__":
    main()
