# go-rest-api

Install Packages

```
go mod tidy
```

Run Project
```
go run main.go
```
<br>

Since the main branch exists on the remote repository but not locally, you need to create it and switch to it:

```
git fetch origin
git checkout -b main origin/main
```
<br>
Set Upstream Tracking for main
If the main branch is not already tracking origin/main, set it up explicitly:

```
git branch --set-upstream-to=origin/main main
```

<br>
Alternatively, you can use the -u flag with git push to set the upstream branch:

```
git push -u origin main
```