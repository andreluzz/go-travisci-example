@ECHO OFF

for /F %%p in ('go list ./...') do (
    go test -race -coverprofile=profile.out -covermode=atomic %%p
    type profile.out >> profile.out
)