+ #### 回溯法：适用于排列，组合，切割字符串，子集，棋盘等问题，本质是递归函数，理解为n叉树。
### 1. 组合总和
##### [力扣39](https://leetcode-cn.com/problems/combination-sum/)
##### 切片的切片，切片的append，回溯，递归.
```go
func combinationSum(candidates []int, target int) [][]int {
    var trcak []int
    var res [][]int
    back(0,0,target,candidates,trcak,&res)
    return res
}
func back(startIndex,sum,target int,candidates,trcak []int,res *[][]int){
    // 递归出口
    if sum==target{
        tmp:=make([]int,len(trcak))
        copy(tmp,trcak)         // 切片的切片，拷贝
        *res=append(*res,tmp)   // 放入结果集
        return
    }

    if sum > target{return}

    // 多叉树，回溯，i++->i从之前index的下一个开始
    for i:=startIndex;i<len(candidates);i++{
        // 更新路径集合和sum
        trcak=append(trcak,candidates[i])
        sum+=candidates[i]
        // 递归，注意：i能重复使用
        back(i,sum,target,candidates,trcak,res)
        // 回溯，上一条路走完了，回去换路
        trcak=trcak[:len(trcak)-1]
        sum-=candidates[i]
    }

}
```
### 2. 组合
##### [力扣77](https://leetcode-cn.com/problems/combinations/)
#####
```go
func combine(n int, k int) [][]int {
    ans := [][]int{}
    path := []int{}
    back(0, n, k, path, &ans)
    return ans
}

func back(startindex, n, k int, path []int, ans *[][]int) {
    // 递归出口
    if len(path) == k {
        temp := make([]int, len(path))
        copy(temp, path)
        *ans = append(*ans, temp)
        return
    }
    // 回溯
    for i := startindex; i < n; i++ {
        // 更新路径
        path = append(path, i+1)
        // 递归
        back(i+1, n, k, path, ans)
        // 回溯,上一条路走完了,回去换路
        path = path[:len(path)-1]
    }
}
```

### 2. 全排列
##### [力扣46](https://leetcode-cn.com/problems/permutations/)
##### 回溯，递归
```go
func permute(nums []int) [][]int {
    ans := [][]int{}
    path := []int{}
    back(nums, path, &ans)
    return ans
}

// 哈希表记录使用过的元素
var used map[int]bool=map[int]bool{}

func back(nums, path []int, ans *[][]int) {
    // 递归出口
    if len(path) == len(nums) {
        temp := make([]int, len(path))
        copy(temp, path)                // 拷贝
        *ans = append(*ans, temp)       // 加入结果
        return
    }
    // 多叉树，回溯.注意：与组合不同的是，使用哈希表记录了使用过的，所以i每次从0开始
    for i := 0; i < len(nums); i++ {
        if used[nums[i]] == true {
            continue
        }
        used[nums[i]] = true
        // 更新路径
        path = append(path, nums[i])
        // 递归
        back(nums, path, ans)
        // 回溯,上一条路走完了,回去换路
        path = path[:len(path)-1]
        used[nums[i]] = false
    }
}
```
### 2. 电话号码的字母匹配
+ ##### [力扣17](https://leetcode-cn.com/problems/letter-combinations-of-a-phone-number/)
+ ##### 字符串数组，切片，递归，深度优先遍历DFS。
```go
func letterCombinations(digits string) []string {
    if len(digits) == 0 {
        return []string{}
    }
    mp := map[string]string {
        "2": "abc",
        "3": "def",
        "4": "ghi",
        "5": "jkl",
        "6": "mno",
        "7": "pqrs",
        "8": "tuv",
        "9": "wxyz",
    }
    ans := make([]string, 0)
    var dfs func(int, string)

    // 定义深度优先遍历函数
    dfs = func(i int, path string) {
        if i >= len(digits) {
            ans = append(ans, path)
            return
        }

        // 不同数字有不同的多叉树,dfs
        for _, c := range mp[string(digits[i])] {
            // 递归
            dfs(i + 1, path + string(c))
        }
    }
    // 深度优先遍历
    dfs(0, "")
    return ans
}
```

### 3.括号生成
##### [力扣22](https://leetcode-cn.com/problems/generate-parentheses/)
##### 画二叉树，深度优先遍历DFS，递归，指针。
```go
func generateParenthesis(n int) []string {
    ans  := make([]string, 0)
    dfs(n, n, n, "", &ans)
    return ans
}

// 深度优先遍历DFS
func dfs(n, left, right int, path string, ans *[]string) {
    if 2*n == len(path) {
        // 递归出口
        *ans = append(*ans, path)
        return
    }
    if left > 0 {
        dfs(n, left - 1, right, path + "(", ans)    // 这里ans已经是地址了，不用加&取地址了
    }
    // 剩下的")"数 比"("数 多才会遍历右子树
    if left < right {
        dfs(n, left, right - 1, path + ")", ans)
    }
}
```

### 4. 目标和的不同表达式数
##### [力扣494](https://leetcode-cn.com/problems/target-sum/submissions/)
#####
```go
func findTargetSumWays(nums []int, target int) int {
    n := len(nums)
    if n == 0 {
        return 0
    }
    ans := 0
    var path func(int, int)
    path = func(cur, sum int) {
        if cur == len(nums) {
            if sum == target {
                ans ++
            }
            return
        }
        path(cur+1, sum+nums[cur])
        path(cur+1, sum-nums[cur])
    }
    path(0,0)
    return ans
}
```