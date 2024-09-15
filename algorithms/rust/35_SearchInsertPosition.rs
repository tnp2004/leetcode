pub fn search_insert(nums: Vec<i32>, target: i32) -> i32 {
    let mut l = 0;
    let mut h = nums.len();
    
     while l < h {
         let mid = l + (h - l) / 2;
 
         if target == nums[mid] {
             return mid as i32
         }
 
         if target < nums[mid] {
             h = mid;
         }else if target > nums[mid] {
             l = mid + 1;
         }
     }
 
     (l + (h - l) / 2) as i32
 }