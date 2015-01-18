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
MitV will host a page as shown in the screenshot below, this page will automatically refresh at the same rate as the server, eg. if the update interval is set to 10 seconds, the page will update every 10 seconds!
```


#Screenshots
![MITV Web View as of 11/01/2015](https://raw.githubusercontent.com/Matt-Allen44/mitV/master/mitv.png)

