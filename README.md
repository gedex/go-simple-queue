go-simple-queue
========================

> Implementing simple queue, in Go, for creating background jobs.

Consider a use case where you need to send email after a user is suscessfully
registered. Sending email takes time and it will block request handler to finish
the response. However, you can delay sending email by pushing it into queue and
let the background worker picking up the job later.

## Run

```
$ git clone $THIS_GITHUB_URL
$ cd go-simple-queue
$ go build
$ ./go-simple-queue
Running 'SendEmail' queue with 4 worker(s)
Starting worker 'Worker#SendEmail#1'
Starting worker 'Worker#SendEmail#2'
Starting worker 'Worker#SendEmail#3'
Starting worker 'Worker#SendEmail#4'
Running 'GenerateThumbnail' queue with 3 worker(s)
Starting worker 'Worker#GenerateThumbnail#1'
Starting worker 'Worker#GenerateThumbnail#2'
Starting worker 'Worker#GenerateThumbnail#3'
Worker#SendEmail#1: taking Job#SendEmail#1
Worker#GenerateThumbnail#1: taking Job#GenerateThumbnail#1
Worker#SendEmail#2: taking Job#SendEmail#2
Job#SendEmail#1 done by Worker#SendEmail#1 [SUCCESS]
...
```

Once you built and ran the program, it will publish jobs periodically. There are
4 `SendEmail` workers and 3 `GenerateThumbnail` workers fetching jobs from
queue. Worker will take [0,3) seconds to finish its job. The result may `SUCCESS`
or `FAIL`.

## License

MIT License - see [LICENSE](./LICENSE) file.
