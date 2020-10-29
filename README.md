# alert
This repo is a simple guage metrics alert framework.

# datasource
```
Datasource evaluates a configured script,and return some results to rule for judgment
And now we support mysql,prometheus...
More to be implemented
```

- mysql
- prometheus
- and so on......

# rules
```
Rule Intermittent evaluates a script to related Datasource 
and judge the alert state.
```
- like grafana alert
- configure the rule evaluate INTERVAL and a FOR predicate




# channel

```
Channel is something receive an alert message,and send it  somewhere else...
And now we just support webhook channel.
``` 
- webhook
    - when use webhook alert channel,an alert should POST a message like  `{"title":"xxx","message":"xxxxx"}` to your configured webhook URL
- and so on......