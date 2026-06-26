# Mnet Auto

Tự động tạo tài khoản Mnet Plus. Chọn số lượng, tool làm phần còn lại.

## Cài đặt

### Pre-built binary (khuyên dùng)

Tải file từ [releases](https://github.com/yngpiu/mnet-auto/releases) và chạy:

| File | Nền tảng |
|---|---|
| `mnet-auto-x64.exe` | Windows |
| `mnet-auto-x64.bin` | Linux |
| `mnet-auto-macos-x64` | macOS |

### Build từ source

```
go build -o mnet-auto-x64.exe .
.\mnet-auto-x64.exe
```

## Sử dụng

```
Mnet Plus Account Auto Creator

  Current password: nmixx0222-
  Current mode: Auto

  1. Create Account
  2. Switch Mode
  3. Change Password
  4. Exit
```

- **Auto** — tạo email tạm + đăng ký qua API, retry 3 lần + domain rotation
- **Manual** — nhập email thủ công + đăng ký qua API

Đổi mật khẩu mặc định trước khi tạo (yêu cầu: 8-20 ký tự, hoa + thường + số + ký tự đặc biệt).

## Lưu ý

- Tool chờ ngẫu nhiên 10-30s giữa các tài khoản để tránh bị chặn
- Tài khoản lưu trong `accounts.txt`
- Email tạm thời, lưu lại nếu cần dùng sau
