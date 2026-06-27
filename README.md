# Mnet Auto Reg

Automatically register Mnet Plus accounts. Choose the count, the tool does the rest.

## Install

### Pre-built binary (recommended)

Download from [releases](https://github.com/yngpiu/mnet-auto-reg/releases):

| File | Platform |
|---|---|
| `mnet-auto-x64.exe` | Windows |
| `mnet-auto-x64.bin` | Linux |
| `mnet-auto-macos-x64` | macOS |

### Build from source

```
go build -o mnet-auto-x64.exe .
.\mnet-auto-x64.exe
```

## Usage

```
Mnet Plus Account Auto Creator

  Current password: nmixx0222-
  Current mode: Auto

  1. Create Account
  2. Switch Mode
  3. Change Password
  4. Exit
```

- **Auto** — auto temp email + API registration, retries 3 times with domain rotation
- **Manual** — enter your own email + API registration

Change the default password before creating accounts (requirements: 8-20 chars, uppercase + lowercase + number + special char).

## Notes

- Random 10-30s delay between accounts to avoid rate limiting
- Accounts saved to `accounts.txt`
- Emails are temporary — save them if you need them later
