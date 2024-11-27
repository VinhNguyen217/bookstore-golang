# Book-Store

Book-Store là một hệ thống bán sách trực tuyến.
## Điều kiện tiên quyết
- Yêu cầu phải cài đặt docker.

## Cài đặt

 - Bước 1: Tải dự án về máy bằng lệnh git như sau :
```bash
 git clone https://github.com/VinhNguyen217/bookstore-golang.git
```
- Bước 2: Truy cập vào thư mục dự án, rồi chạy lệnh :
```bash 
 docker compose up -d
```
- Bước 3: Khởi tạo cơ sở dữ liệu :
```bash 
 go run main.go migrate
```
