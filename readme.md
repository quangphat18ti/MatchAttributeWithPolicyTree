## Bài toán

### Cấu trúc Policy Tree
- Policy tree là biểu thức mệnh đề (**boolean formulas**). Bao gồm:   
  1. OR
  2. AND
  3. Các phép so sánh: `<, >, <=, >=, =`,   
    a.  Date Comparison: Chuyển date về thành days since the beginning of unix epoch time (Jan 1, 1970).  
       - Example:  `Date > March 1, 2015`
  4. DATE: `[Prefix] = [Month] [Day], [Year]`
  5. RANGE:   
    a. Integer range: `[Attribute] in ([int] - int)`.    (Attribute > int1 and Attribute < int2)
    b. Date Range:    `[Attribute] = [Month] [Day] - [Day], [Year]`  (Attribute >= day1 and Attribute <= day2)
  6. Dấu ngoặc `()`

### Attribute Lists
- Được viết cách nhau bởi `|`. 
- Example: `|Manager|IT|Experience=5|Date = December 20, 2015|`

### Yêu cầu
- Input:   
  1. `Attribute Lists` dạng String. 
  2. `Policy Tree` dạng String.

- Output: True/ False

## Thuật giải

### Cách tiếp cận Stack với biểu thức mệnh đề 
- Sử dụng 2 Stack: 
  1. Stack chứa các phép toán: `operators`
  2. Stack chứa các chân trị của các biểu thức  : `values`

- Đọc từng phần tử trong `Policy Tree`:
  1. Nếu là `(`: Push vào `operators`
  2. Nếu là `)`: Pop toàn bộ `operators` và `values` để thực hiện phép toán cho tới khi gặp `(`. 
      - Push kết quả tính được vào lại `values`
      - Pop `(` ra khỏi `operators`
  3. Nếu là `AND/and`, `OR/or`: Push vào `operators`
  4. Nếu là `biểu thức mệnh đề`: Giải quyết biểu thức mệnh đề và push vào `values`.

### Cách tiếp cận đệ quy để xây cây