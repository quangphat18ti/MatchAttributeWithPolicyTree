## Bài toán

### Cấu trúc Policy Tree
- Policy tree là biểu thức mệnh đề (**boolean formulas**). Bao gồm:   
  1. OR
  2. AND
  3. Các phép so sánh: `<, >, <=, >=`  
    a.  Date Comparison: Chuyển date về thành days since the beginning of unix epoch time (Jan 1, 1970).  
       - Example:  `Date > March 1, 2015`
  4. DATE: `[Prefix] = [Month] [Day], [Year]`
  5. RANGE:   
    a. Integer range: `[Attribute] in ([int], int)`.  
    b. Date Range:    `[Attribute] = [Month] [Day] - [Day], [Year]`

### Attribute Lists
- Được viết cách nhau bởi `|`. 
- Example: `|Manager|IT|Experience=5|Date = December 20, 2015|`

### Yêu cầu
- Input:   
  1. `Attribute Lists` dạng String. 
  2. `Policy Tree` dạng String.

- Output: True/ False