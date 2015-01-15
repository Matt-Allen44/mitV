#MitV
Lightweight WebApp used to view Commits to a Github Repository in Realtime!
Developed using GO and the GITHUB API

#Contributing
Send a Pull Request if you wish to commit!

#Building
No external packages required to build!
``` 
$ go build mitV.go
```

#Running
```
$ mitV.exe <GithubOauthToken> <repo in format (author/repo)> <host (ip:port>> <update interval (seconds)>

eg.
$ mitV.exe <token> Matt-Allen44/mitV 127.0.0.1:8080 10
  This would launch a server watching commits to this repo, and host it on localhost:8080, updating commits every 10 seconds!
```

```
MitV will host a page as shown in the screenshot below, this is a static webpage produced by 
MitV, the page is static - to update your view, refresh the page!

Automatic refreshing can be done using the Easy-Auto-Refresh plugin for Google Chrome 
```


#Screenshots
![MITV Web View as of 11/01/2015](https://raw.githubusercontent.com/Matt-Allen44/mitV/master/mitv.png)

[Easy-Auto-Refresh plugin for Google Chrome}:https://chrome.google.com/webstore/detail/easy-auto-refresh/aabcgdmkeabbnleenpncegpcngjpnjkc

