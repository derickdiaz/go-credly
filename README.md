# go-credly

This Golang Module is used to get a user's badges from credly

## Installtion
```bash
go get github.com/derickdiaz/go-credly@v0.0.1
```

## Usage
```golang
svc := credly.CredlyService{}
badges, err := svc.GetBadges("username")
if err != nil {
    panic()
}
for _, badge := range badges.Data {
    fmt.Println(badge.GetName())
    fmt.Println(badge.GetIssueDate())
    fmt.Println(badge.GetExpiredDate())
}
```