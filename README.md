# ALERT


# DATASOURCE
```
Datasource evaluates a configured script,and return some results to rule for judgment
And now we support mysql,prometheus...
More to be implemented
```

- mysql
    - a `SELECT` script return any `METRICS` will be a `[]*DatasourceResult`
- prometheus
    - a `Counter` metrics saves as a script,and will be converted to `sum(sample_metrics{}) by (instance)`
    - when evaluated it,we will use the absolute `Current day(from 00:00:00 to now)` range query
    - so the result `[]*DatasourceResult` will just has one element,and the meaning of it is `the total count of sample_data of current day`
- and so on......

# RULE
```
Rule Intermittent evaluates a script to related Datasource 
and judge the alert state.
```
- a rule defined the frequency  
- configure the rule evaluate INTERVAL and a FOR predicate




# channel

```
Channel is something receive an alert message,and send it  somewhere else...
And now we just support webhook channel.
``` 
- webhook
    - when use webhook alert channel,an alert should POST a message like  `{"title":"xxx","message":"xxxxx"}` to your configured webhook URL
- so on......