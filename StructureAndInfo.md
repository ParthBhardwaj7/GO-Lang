Normal structure of GO Files
project-name/
    go.mod      <-- ONLY ONE
    go.sum
    main.go
    /internal
        utils.go
    /handlers
        user.go
    /db
        connect.go

---What Package Do:
Code organize karne ke liye

Code reuse karne ke liye

Logic ko alag-alag blocks me todne ke liye (clean project structure)