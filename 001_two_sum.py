nums = [2,7,11,15]
target = 13

seenHash = {}

for i in range(len(nums)):
    if ((target - nums[i]) in seenHash):
        print("[ "+str(seenHash[target - nums[i]])+", "+str(i), "]")
    seenHash[nums[i]] = i