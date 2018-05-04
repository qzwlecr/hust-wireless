# HUST-WIRELESS-HACK

## Intro

This tool was used for brute-force attacking HUST Campus Network in 2017. However, it was prohibited by the network administrator. So I package it for future usage.

## Usage

import "hust_wireless".

```go
    u:= new(hust_wireless.User)
    u.Init("Your Username","Your Passwd")
    status, str:= u.Login()
    //u.Logout()
```

The Status is a digit.
- 0: Can not connect to the website. Maybe you have logged in.
- 1: Login successfully.
- 2: Wrong password.
- 3: Wrong username.
- 4: Failed to log in with other reasons.

## Help

Please log in before you want to log out. Because a user index is needed.
